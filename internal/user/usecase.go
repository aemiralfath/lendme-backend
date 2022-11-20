package user

import (
	"context"
	"final-project-backend/internal/models"
)

type UseCase interface {
	GetDebtorDetails(ctx context.Context, userID string) (*models.Debtor, error)
	ConfirmContract(ctx context.Context, userID string) (*models.Debtor, error)
}
