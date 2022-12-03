package admin

import (
	"context"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/utils"
)

type Repository interface {
	GetLoans(ctx context.Context, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error)
	GetPayments(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetVouchers(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error)
	GetDebtors(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetDebtorByID(ctx context.Context, debtorID string) (*models.Debtor, error)
	GetLendingByID(ctx context.Context, lendingID string) (*models.Lending, error)
	GetLendingWithInstallmentByID(ctx context.Context, lendingID string) (*models.Lending, error)
	GetContractStatusByID(ctx context.Context, contractID int) (*models.ContractTrackingType, error)
	GetCreditHealthByID(ctx context.Context, healthID int) (*models.CreditHealthType, error)
	GetInstallmentByID(ctx context.Context, installmentID string) (*models.Installment, error)
	CreateInstallments(ctx context.Context, lendingID string, installments []*models.Installment) (*models.Lending, error)
	UpdateDebtorByID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
	UpdateLendingByID(ctx context.Context, lending *models.Lending) (*models.Lending, error)
	UpdateInstallmentByID(ctx context.Context, installment *models.Installment) (*models.Installment, error)
	CreateVoucher(ctx context.Context, voucher *models.Voucher) (*models.Voucher, error)
	GetVoucherByID(ctx context.Context, voucherID string) (*models.Voucher, error)
	UpdateVoucherByID(ctx context.Context, voucher *models.Voucher) error
	DeleteVoucher(ctx context.Context, voucher *models.Voucher) error
	GetUserTotal(ctx context.Context) (int64, error)
	GetLendingTotal(ctx context.Context) (int64, error)
	GetLendingAmount(ctx context.Context) (float64, error)
	GetReturnAmount(ctx context.Context) (float64, error)
	GetLendingAction(ctx context.Context) ([]*models.Lending, error)
	GetUserAction(ctx context.Context) ([]*models.Debtor, error)
}
