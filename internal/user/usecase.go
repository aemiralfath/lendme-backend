package user

import (
	"context"
	"final-project-backend/internal/models"
	"final-project-backend/internal/user/delivery/body"
	"final-project-backend/pkg/utils"
)

type UseCase interface {
	GetDebtorDetails(ctx context.Context, userID string) (*models.Debtor, error)
	ConfirmContract(ctx context.Context, userID string) (*models.Debtor, error)
	CreateLoan(ctx context.Context, userID string, body body.CreateLoan) (*models.Lending, error)
	GetLoans(ctx context.Context, userID, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error)
	GetVouchers(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error)
	GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error)
	GetInstallmentByID(ctx context.Context, installmentID string) (*models.Installment, error)
	CreatePayment(ctx context.Context, userID string, body body.CreatePayment) (*models.Payment, error)
}
