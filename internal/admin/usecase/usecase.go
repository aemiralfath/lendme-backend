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
)

type adminUC struct {
	cfg       *config.Config
	adminRepo admin.Repository
}

func NewAdminUseCase(cfg *config.Config, adminRepo admin.Repository) admin.UseCase {
	return &adminUC{cfg: cfg, adminRepo: adminRepo}
}

func (u *adminUC) UpdateContractStatus(ctx context.Context, body body.UpdateContractRequest) (*models.Debtor, error) {
	debtor, err := u.adminRepo.GetDebtorByID(ctx, body.DebtorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return debtor, httperror.New(http.StatusBadRequest, response.DebtorIDNotExist)
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

	debtor.ContractTrackingID = contract.ContractTrackingID
	debtor, err = u.adminRepo.UpdateDebtorStatusByID(ctx, debtor)
	if err != nil {
		return debtor, err
	}

	return debtor, nil
}
