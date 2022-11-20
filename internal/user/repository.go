package user

import (
	"context"
	"final-project-backend/internal/models"
)

type Repository interface {
	GetDebtorDetailsByID(ctx context.Context, userID string) (*models.Debtor, error)
	UpdateContractByUserID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
}
