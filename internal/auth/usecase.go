package auth

import (
	"context"
	"final-project-backend/internal/models"
	"final-project-backend/internal/models/auth"
)

type UseCase interface {
	Register(ctx context.Context, body auth.RegisterRequest) (*models.User, error)
	Login(ctx context.Context, body auth.LoginRequest) (*models.UserWithToken, error)
	GetUserDetails(ctx context.Context, userID string) (*models.User, error)
}
