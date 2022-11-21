package admin

import (
	"context"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/internal/models"
)

type UseCase interface {
	GetDebtors(ctx context.Context) ([]*models.Debtor, error)
	UpdateContractStatus(ctx context.Context, body body.UpdateContractRequest) (*models.Debtor, error)
}
