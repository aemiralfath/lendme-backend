package auth

import (
	"context"
	"final-project-backend/internal/models"
)

type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	FindByEmail(ctx context.Context, user *models.User) (*models.User, error)
	CreateDebtor(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
	CheckEmailExist(ctx context.Context, user *models.User) (*models.User, error)
	GetUserDetailsByID(ctx context.Context, userId string) (*models.User, error)
}
