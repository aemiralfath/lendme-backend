package admin

import (
	"context"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/utils"
)

type UseCase interface {
	GetDebtorByID(ctx context.Context, id string) (*models.Debtor, error)
	GetDebtors(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetLoans(ctx context.Context, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error)
	GetPayments(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetVouchers(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error)
	GetInstallmentByID(ctx context.Context, installmentID string) (*models.Installment, error)
	UpdateDebtorByID(ctx context.Context, debtorID string, body body.UpdateContractRequest) (*models.Debtor, error)
	UpdateInstallmentByID(ctx context.Context, installmentID string, body body.UpdateInstallmentRequest) (*models.Installment, error)
	ApproveLoan(ctx context.Context, lendingID string) (*models.Lending, error)
	CreateVoucher(ctx context.Context, body body.CreateVoucherRequest) (*models.Voucher, error)
}
