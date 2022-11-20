package admin

import (
	"context"
	"final-project-backend/internal/models"
)

type Repository interface {
	GetDebtorByID(ctx context.Context, debtorID string) (*models.Debtor, error)
	UpdateDebtorStatusByID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
}
