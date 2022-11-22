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
)

type userUC struct {
	cfg      *config.Config
	userRepo user.Repository
}

func NewUserUseCase(cfg *config.Config, userRepo user.Repository) user.UseCase {
	return &userUC{cfg: cfg, userRepo: userRepo}
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
