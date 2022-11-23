package usecase

import (
	"context"
	"final-project-backend/config"
	"final-project-backend/internal/admin"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"gorm.io/gorm"
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
	installmentAmount := lending.Amount / float64(lending.LoanPeriod.Duration)
	for i := 0; i < lending.LoanPeriod.Duration; i++ {
		installment := &models.Installment{}
		installmentDate := time.Now()
		installmentDate = installmentDate.AddDate(0, i+1, 0)

		installment.LendingID = lending.LendingID
		installment.Amount = installmentAmount
		installment.DueDate = time.Date(installmentDate.Year(), installmentDate.Month(), 25, 0, 0, 0, installmentDate.Nanosecond(), installmentDate.Location())
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
