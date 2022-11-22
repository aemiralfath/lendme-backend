package admin

import (
	"context"
	"final-project-backend/internal/models"
)

type Repository interface {
	GetDebtors(ctx context.Context) ([]*models.Debtor, error)
	GetDebtorByID(ctx context.Context, debtorID string) (*models.Debtor, error)
	GetLendingByID(ctx context.Context, lendingID string) (*models.Lending, error)
	GetLendingWithInstallmentByID(ctx context.Context, lendingID string) (*models.Lending, error)
	GetContractStatusByID(ctx context.Context, contractID int) (*models.ContractTrackingType, error)
	GetCreditHealthByID(ctx context.Context, healthID int) (*models.CreditHealthType, error)
	CreateInstallments(ctx context.Context, lendingID string, installments []*models.Installment) (*models.Lending, error)
	UpdateDebtorByID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
	UpdateLendingByID(ctx context.Context, lending *models.Lending) (*models.Lending, error)
}
