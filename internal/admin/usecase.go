package admin

import (
	"context"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/pagination"
)

type UseCase interface {
	GetDebtorByID(ctx context.Context, id string) (*models.Debtor, error)
	GetDebtors(ctx context.Context, name string, pagination *pagination.Pagination) (*pagination.Pagination, error)
	GetLoans(ctx context.Context, name string, status []int, pagination *pagination.Pagination) (*pagination.Pagination, error)
	GetPayments(ctx context.Context, name string, pagination *pagination.Pagination) (*pagination.Pagination, error)
	GetVouchers(ctx context.Context, name string, pagination *pagination.Pagination) (*pagination.Pagination, error)
	GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error)
	GetInstallmentByID(ctx context.Context, installmentID string) (*models.Installment, error)
	UpdateDebtorByID(ctx context.Context, debtorID string, body body.UpdateContractRequest) (*models.Debtor, error)
	UpdateInstallmentByID(ctx context.Context, installmentID string, body body.UpdateInstallmentRequest) (*models.Installment, error)
	ApproveLoan(ctx context.Context, lendingID string) (*models.Lending, error)
	RejectLoan(ctx context.Context, lendingID string) (*models.Lending, error)
	CreateVoucher(ctx context.Context, body body.CreateVoucherRequest) (*models.Voucher, error)
	GetVoucherByID(ctx context.Context, voucherID string) (*models.Voucher, error)
	GetSummary(ctx context.Context) (*body.SummaryResponse, error)
	UpdateVoucherByID(ctx context.Context, voucherID string, body body.UpdateVoucherRequest) (*models.Voucher, error)
	DeleteVoucherByID(ctx context.Context, voucherID string) (*models.Voucher, error)
}
