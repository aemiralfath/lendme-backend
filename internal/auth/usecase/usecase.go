package usecase

import (
	"context"
	"final-project-backend/config"
	"final-project-backend/internal/auth"
	"final-project-backend/internal/auth/delivery/body"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"final-project-backend/pkg/utils"
	"net/http"
)

type authUC struct {
	cfg      *config.Config
	authRepo auth.Repository
}

func NewAuthUseCase(cfg *config.Config, authRepo auth.Repository) auth.UseCase {
	return &authUC{cfg: cfg, authRepo: authRepo}
}

func (u *authUC) Register(ctx context.Context, body body.RegisterRequest) (*models.User, error) {
	user := &models.User{}
	user.Name = body.Name
	user.Email = body.Email
	user.Password = body.Password

	existsUser, err := u.authRepo.CheckEmailExist(ctx, user)
	if existsUser.Email != "" {
		return nil, httperror.New(http.StatusBadRequest, response.AuthEmailAlreadyExistMessage)
	}

	if err = user.PrepareCreate(2); err != nil {
		return nil, err
	}

	createdUser, err := u.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	createdUser.SanitizePassword()

	debtor := &models.Debtor{}
	if err = debtor.PrepareCreate(createdUser.UserID, 1, 1); err != nil {
		return nil, err
	}

	if _, err := u.authRepo.CreateDebtor(ctx, debtor); err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *authUC) Login(ctx context.Context, body body.LoginRequest) (*models.UserWithToken, error) {
	user := &models.User{}
	user.Email = body.Email
	user.Password = body.Password

	foundUser, err := u.authRepo.FindByEmail(ctx, user)
	if err != nil {
		return nil, httperror.New(http.StatusUnauthorized, response.UnauthorizedMessage)
	}

	if err = foundUser.ComparePasswords(user.Password); err != nil {
		return nil, httperror.New(http.StatusUnauthorized, response.UnauthorizedMessage)
	}

	foundUser.SanitizePassword()

	token, err := utils.GenerateJWTToken(foundUser.UserID.String(), foundUser.RoleID, u.cfg)
	if err != nil {
		return nil, err
	}

	return &models.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}

func (u *authUC) GetUserDetails(ctx context.Context, userID string) (*models.User, error) {
	user, err := u.authRepo.GetUserDetailsByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
