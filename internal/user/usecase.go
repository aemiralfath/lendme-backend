package user

import (
	"context"
	"final-project-backend/internal/models"
	"final-project-backend/internal/user/delivery/body"
)

type UseCase interface {
	GetDebtorDetails(ctx context.Context, userID string) (*models.Debtor, error)
	ConfirmContract(ctx context.Context, userID string) (*models.Debtor, error)
	CreateLoan(ctx context.Context, userID string, body body.CreateLoan) (*models.Lending, error)
}
