package admin

import (
	"context"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/utils"
)

type UseCase interface {
	GetDebtors(ctx context.Context) ([]*models.Debtor, error)
	GetDebtorByID(ctx context.Context, id string) (*models.Debtor, error)
	GetLoans(ctx context.Context, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error)
	GetPayments(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error)
	UpdateDebtorByID(ctx context.Context, body body.UpdateContractRequest) (*models.Debtor, error)
	UpdateInstallmentByID(ctx context.Context, body body.UpdateInstallmentRequest) (*models.Installment, error)
	ApproveLoan(ctx context.Context, lendingID string) (*models.Lending, error)
}
