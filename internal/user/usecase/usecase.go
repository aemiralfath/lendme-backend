package usecase

import (
	"context"
	"final-project-backend/config"
	"final-project-backend/internal/models"
	"final-project-backend/internal/user"
)

type userUC struct {
	cfg      *config.Config
	userRepo user.Repository
}

func NewUserUseCase(cfg *config.Config, userRepo user.Repository) user.UseCase {
	return &userUC{cfg: cfg, userRepo: userRepo}
}

func (u *userUC) GetDebtorDetails(ctx context.Context, userID string) (*models.Debtor, error) {
	debtor, err := u.userRepo.GetDebtorDetailsByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return debtor, nil
}
