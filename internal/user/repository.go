package user

import (
	"context"
	"final-project-backend/internal/models"
)

type Repository interface {
	GetDebtorDetailsByID(ctx context.Context, userID string) (*models.Debtor, error)
	UpdateDebtorByID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
	GetLoanPeriodByID(ctx context.Context, periodID int) (*models.LoanPeriod, error)
	CreateLending(ctx context.Context, lending *models.Lending) (*models.Lending, error)
}
