package admin

import (
	"context"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/utils"
)

type Repository interface {
	GetLoans(ctx context.Context, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error)
	GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error)
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
