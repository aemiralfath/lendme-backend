package auth

import (
	"context"
	"final-project-backend/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	FindByEmail(ctx context.Context, user *models.User) (*models.User, error)
	CreateDebtor(ctx context.Context, wallet *models.Debtor) (*models.Debtor, error)
	CheckEmailExist(ctx context.Context, user *models.User) (*models.User, error)
	GetUserDetailsByID(ctx context.Context, userId string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
}
