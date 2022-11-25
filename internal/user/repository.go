package user

import (
	"context"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/utils"
)

type Repository interface {
	GetLoans(ctx context.Context, debtorID, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error)
	GetLoanPeriodByID(ctx context.Context, periodID int) (*models.LoanPeriod, error)
	GetDebtorDetailsByID(ctx context.Context, userID string) (*models.Debtor, error)
	UpdateDebtorByID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
	CreateLending(ctx context.Context, lending *models.Lending) (*models.Lending, error)
	GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error)
	GetInstallmentByID(ctx context.Context, installmentID string) (*models.Installment, error)
	GetVoucherByID(ctx context.Context, voucherID string) (*models.Voucher, error)
	CreatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	UpdateInstallment(ctx context.Context, installment *models.Installment) (*models.Installment, error)
	UpdateLending(ctx context.Context, lending *models.Lending) (*models.Lending, error)
}
