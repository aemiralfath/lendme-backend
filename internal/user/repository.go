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
}
