package admin

import (
	"context"
	"final-project-backend/internal/models"
)

type Repository interface {
	GetDebtors(ctx context.Context) ([]*models.Debtor, error)
	GetDebtorByID(ctx context.Context, debtorID string) (*models.Debtor, error)
	GetContractStatusByID(ctx context.Context, contractID int) (*models.ContractTrackingType, error)
	GetCreditHealthByID(ctx context.Context, healthID int) (*models.CreditHealthType, error)
	UpdateDebtorByID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
}
