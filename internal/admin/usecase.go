package admin

import (
	"context"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/internal/models"
)

type UseCase interface {
	GetDebtors(ctx context.Context) ([]*models.Debtor, error)
	GetDebtorByID(ctx context.Context, id string) (*models.Debtor, error)
	UpdateDebtorByID(ctx context.Context, body body.UpdateContractRequest) (*models.Debtor, error)
	ApproveLoan(ctx context.Context, lendingID string) (*models.Lending, error)
}
