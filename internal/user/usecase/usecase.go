package usecase

import (
	"context"
	"final-project-backend/config"
	"final-project-backend/internal/models"
	"final-project-backend/internal/user"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"net/http"
)

type userUC struct {
	cfg      *config.Config
	userRepo user.Repository
}

func NewUserUseCase(cfg *config.Config, userRepo user.Repository) user.UseCase {
	return &userUC{cfg: cfg, userRepo: userRepo}
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
	debtor, err = u.userRepo.UpdateContractByUserID(ctx, debtor)
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
