package usecase

import (
	"context"
	"final-project-backend/config"
	"final-project-backend/internal/admin"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"final-project-backend/pkg/utils"
	"gorm.io/gorm"
	"math"
	"net/http"
	"time"
)

type adminUC struct {
	cfg       *config.Config
	adminRepo admin.Repository
}

func NewAdminUseCase(cfg *config.Config, adminRepo admin.Repository) admin.UseCase {
	return &adminUC{cfg: cfg, adminRepo: adminRepo}
}

func (u *adminUC) GetDebtors(ctx context.Context) ([]*models.Debtor, error) {
	debtors, err := u.adminRepo.GetDebtors(ctx)
	if err != nil {
		return debtors, err
	}

	return debtors, nil
}

func (u *adminUC) GetDebtorByID(ctx context.Context, id string) (*models.Debtor, error) {
	debtor, err := u.adminRepo.GetDebtorByID(ctx, id)
	if err != nil {
		return debtor, err
	}

	return debtor, nil
}

func (u *adminUC) ApproveLoan(ctx context.Context, lendingID string) (*models.Lending, error) {
	lending, err := u.adminRepo.GetLendingByID(ctx, lendingID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return lending, httperror.New(http.StatusBadRequest, response.DebtorIDNotExist)
		}
		return lending, err
	}

	lending.LendingStatusID = 2
	lending, err = u.adminRepo.UpdateLendingByID(ctx, lending)
	if err != nil {
		return lending, err
	}

	var installments []*models.Installment
	installmentAmount := math.Round(lending.Amount / float64(lending.LoanPeriod.Duration))

	loc, _ := time.LoadLocation("Asia/Jakarta")
	for i := 0; i < lending.LoanPeriod.Duration; i++ {
		installment := &models.Installment{}
		installmentDate := time.Now().In(loc)
		installmentDate = installmentDate.AddDate(0, i+1, 0)

		installment.LendingID = lending.LendingID
		installment.Amount = installmentAmount
		installment.DueDate = time.Date(installmentDate.Year(), installmentDate.Month(), 25, 23, 59, 59, installmentDate.Nanosecond(), installmentDate.Location())
		installments = append(installments, installment)

		if err := installment.PrepareCreate(); err != nil {
			return lending, err
		}
	}

	lending, err = u.adminRepo.CreateInstallments(ctx, lendingID, installments)
	if err != nil {
		return lending, err
	}

	return lending, nil
}

func (u *adminUC) UpdateDebtorByID(ctx context.Context, body body.UpdateContractRequest) (*models.Debtor, error) {
	debtor, err := u.adminRepo.GetDebtorByID(ctx, body.DebtorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return debtor, httperror.New(http.StatusBadRequest, response.DebtorIDNotExist)
		}
		return debtor, err
	}

	health, err := u.adminRepo.GetCreditHealthByID(ctx, body.CreditHealthID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return debtor, httperror.New(http.StatusBadRequest, response.ContractIDNotExist)
		}
		return debtor, err
	}

	contract, err := u.adminRepo.GetContractStatusByID(ctx, body.ContractStatusID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return debtor, httperror.New(http.StatusBadRequest, response.ContractIDNotExist)
		}
		return debtor, err
	}

	debtor.CreditLimit = body.CreditLimit
	debtor.CreditHealthID = health.CreditHealthID
	debtor.ContractTrackingID = contract.ContractTrackingID
	debtor, err = u.adminRepo.UpdateDebtorByID(ctx, debtor)
	if err != nil {
		return debtor, err
	}

	return debtor, nil
}

func (u *adminUC) UpdateInstallmentByID(ctx context.Context, body body.UpdateInstallmentRequest) (*models.Installment, error) {
	installment, err := u.adminRepo.GetInstallmentByID(ctx, body.InstallmentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return installment, httperror.New(http.StatusBadRequest, response.InstallmentNotExist)
		}
		return installment, err
	}

	installment.DueDate = body.DueDateTime
	installment, err = u.adminRepo.UpdateInstallmentByID(ctx, installment)
	if err != nil {
		return installment, err
	}

	return installment, nil
}

func (u *adminUC) GetLoans(ctx context.Context, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error) {
	loans, err := u.adminRepo.GetLoans(ctx, name, status, pagination)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (u *adminUC) GetPayments(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error) {
	payments, err := u.adminRepo.GetPayments(ctx, name, pagination)
	if err != nil {
		return payments, err
	}

	return payments, nil
}

func (u *adminUC) GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error) {
	lending, err := u.adminRepo.GetLoanByID(ctx, lendingID)
	if err != nil {
		return lending, err
	}

	return lending, nil
}
