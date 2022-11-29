package usecase

import (
	"context"
	"final-project-backend/config"
	"final-project-backend/internal/models"
	"final-project-backend/internal/user"
	"final-project-backend/internal/user/delivery/body"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"final-project-backend/pkg/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type userUC struct {
	cfg      *config.Config
	userRepo user.Repository
}

func NewUserUseCase(cfg *config.Config, userRepo user.Repository) user.UseCase {
	return &userUC{cfg: cfg, userRepo: userRepo}
}

func (u *userUC) GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error) {
	lending, err := u.userRepo.GetLoanByID(ctx, lendingID)
	if err != nil {
		return lending, err
	}

	return lending, nil
}

func (u *userUC) GetInstallmentByID(ctx context.Context, id string) (*models.Installment, error) {
	installment, err := u.userRepo.GetInstallmentByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return installment, httperror.New(http.StatusBadRequest, response.InstallmentNotExist)
		}
		return installment, err
	}

	return installment, nil
}

func (u *userUC) CreatePayment(ctx context.Context, userID string, body body.CreatePayment) (*models.Payment, error) {
	delay := 0
	payment := &models.Payment{}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(loc)

	debtor, err := u.userRepo.GetDebtorDetailsByID(ctx, userID)
	if err != nil {
		return payment, err
	}

	lending, err := u.userRepo.GetLoanByID(ctx, body.LendingID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return payment, httperror.New(http.StatusBadRequest, response.LendingIDNotExist)
		}
		return payment, err
	}

	if lending.DebtorID != debtor.DebtorID {
		return payment, httperror.New(http.StatusBadRequest, response.LendingInstallmentNotMatch)
	}

	installment, err := u.userRepo.GetInstallmentByID(ctx, body.InstallmentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return payment, httperror.New(http.StatusBadRequest, response.InstallmentNotExist)
		}
		return payment, err
	}

	if installment.InstallmentStatusID != 1 {
		return payment, httperror.New(http.StatusBadRequest, response.InstallmentAlreadyPaid)
	}

	payment.VoucherID = nil
	voucher := &models.Voucher{}
	if body.VoucherID != "" {
		voucher, err = u.userRepo.GetVoucherByID(ctx, body.VoucherID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return payment, httperror.New(http.StatusBadRequest, response.VoucherNotExist)
			}
			return payment, err
		}

		if voucher.DiscountQuota > 0 && (timeNow.Sub(voucher.ActiveDate).Seconds() >= 0 && timeNow.Sub(voucher.ExpireDate).Seconds() <= 0) {
			payment.VoucherID = &voucher.VoucherID
			payment.PaymentDiscount = installment.Amount * float64(voucher.DiscountPayment) / 100.0
		} else {
			return payment, httperror.New(http.StatusBadRequest, response.VoucherNotExist)
		}
	}

	delayTime := timeNow.Sub(installment.DueDate)
	if delayTime.Seconds() > 0 {
		delay = int(delayTime.Hours()/24) + 1
	}

	payment.InstallmentID = installment.InstallmentID
	payment.PaymentDate = timeNow
	payment.PaymentFine = float64(5000 * delay)
	payment.PaymentAmount = installment.Amount - payment.PaymentDiscount + payment.PaymentFine
	if err := payment.PrepareCreate(); err != nil {
		return payment, nil
	}

	payment, err = u.userRepo.CreatePayment(ctx, payment)
	if err != nil {
		return payment, err
	}

	if body.VoucherID != "" {
		voucher.DiscountQuota -= 1
		if err := u.userRepo.UpdateVoucher(ctx, voucher); err != nil {
			return payment, err
		}

		if voucher.DiscountQuota <= 0 {
			if err := u.userRepo.DeleteVoucher(ctx, voucher); err != nil {
				return payment, err
			}
		}
	}

	debtor.CreditUsed = debtor.CreditUsed - installment.Amount
	debtor.TotalDelay = debtor.TotalDelay + delay
	if delay == 0 {
		if debtor.TotalDelay-10 < 0 {
			debtor.TotalDelay = 0
		} else {
			debtor.TotalDelay = debtor.TotalDelay - 10
		}
	}

	switch {
	case debtor.TotalDelay > 20:
		debtor.CreditHealthID = 3
	case debtor.TotalDelay > 10:
		debtor.CreditHealthID = 2
	default:
		debtor.CreditHealthID = 1
	}

	debtor, err = u.userRepo.UpdateDebtorByID(ctx, debtor)
	if err != nil {
		return payment, err
	}

	installment.InstallmentStatusID = 2
	installment, err = u.userRepo.UpdateInstallment(ctx, installment)
	if err != nil {
		return payment, err
	}

	totalPaid := 0
	for _, installment := range *lending.Installments {
		if installment.InstallmentStatusID == 2 {
			totalPaid++
		}
	}

	if totalPaid == len(*lending.Installments)-1 {
		lending.LendingStatusID = 4
	} else {
		lending.LendingStatusID = 3
	}

	lending, err = u.userRepo.UpdateLending(ctx, lending)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (u *userUC) CreateLoan(ctx context.Context, userID string, body body.CreateLoan) (*models.Lending, error) {
	lending := &models.Lending{}
	debtor, err := u.userRepo.GetDebtorDetailsByID(ctx, userID)
	if err != nil {
		return lending, err
	}

	if debtor.ContractTrackingID != 5 {
		return lending, httperror.New(http.StatusBadRequest, response.ContractNotConfirmed)
	}

	period, err := u.userRepo.GetLoanPeriodByID(ctx, body.LoadPeriodID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return lending, httperror.New(http.StatusBadRequest, response.LoanPeriodNotExist)
		}
		return lending, err
	}

	amount := body.Amount * (float64(period.Percentage) / 100)
	switch debtor.CreditHealthID {
	case 1:
		if debtor.CreditLimit-(debtor.CreditUsed+amount) < 0 {
			return lending, httperror.New(http.StatusBadRequest, response.LoanAmountExceedCreditLimit)
		}
	case 2:
		if (debtor.CreditLimit*(float64(80)/100))-(debtor.CreditUsed+amount) < 0 {
			return lending, httperror.New(http.StatusBadRequest, response.LoanAmountExceedCreditLimitWarning)
		}
	case 3:
		return lending, httperror.New(http.StatusBadRequest, response.CreditHealthStatusBlocked)
	}

	lending.DebtorID = debtor.DebtorID
	lending.LoanPeriodID = period.LoanPeriodID
	lending.Name = body.Name
	lending.Amount = amount
	if err = lending.PrepareCreate(); err != nil {
		return lending, err
	}

	createdLending, err := u.userRepo.CreateLending(ctx, lending)
	if err != nil {
		return nil, err
	}

	debtor.CreditUsed = debtor.CreditUsed + amount
	if _, err := u.userRepo.UpdateDebtorByID(ctx, debtor); err != nil {
		return nil, err
	}

	return createdLending, nil
}

func (u *userUC) ConfirmContract(ctx context.Context, userID string) (*models.Debtor, error) {
	debtor, err := u.userRepo.GetDebtorDetailsByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if debtor.ContractTrackingID == 5 {
		return debtor, httperror.New(http.StatusBadRequest, response.ContractAlreadyAccepted)
	}

	if debtor.ContractTrackingID != 4 {
		return debtor, httperror.New(http.StatusBadRequest, response.ContractNotAccepted)
	}

	debtor.ContractTrackingID = 5
	debtor, err = u.userRepo.UpdateDebtorByID(ctx, debtor)
	if err != nil {
		return nil, err
	}

	return debtor, nil
}

func (u *userUC) GetDebtorDetails(ctx context.Context, userID string) (*models.Debtor, error) {
	debtor, err := u.userRepo.GetDebtorDetailsByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return debtor, nil
}

func (u *userUC) GetLoans(ctx context.Context, userID, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error) {
	debtor, err := u.userRepo.GetDebtorDetailsByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	loans, err := u.userRepo.GetLoans(ctx, debtor.DebtorID.String(), name, status, pagination)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (u *userUC) GetVouchers(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error) {
	vouchers, err := u.userRepo.GetVouchers(ctx, name, pagination)
	if err != nil {
		return vouchers, err
	}

	return vouchers, nil
}
