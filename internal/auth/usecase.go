package auth

import (
	"context"
	auth2 "final-project-backend/internal/auth/delivery/body"
	"final-project-backend/internal/models"
)

type UseCase interface {
	Register(ctx context.Context, body auth2.RegisterRequest) (*models.User, error)
	Login(ctx context.Context, body auth2.LoginRequest) (*models.UserWithToken, error)
	GetUserDetails(ctx context.Context, userID string) (*models.User, error)
}
