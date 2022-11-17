package repository

import (
	"context"
	"final-project-backend/internal/auth"
	"final-project-backend/internal/models"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepo{db: db}
}

func (r *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *authRepo) CreateDebtor(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error) {
	if err := r.db.WithContext(ctx).Create(debtor).Error; err != nil {
		return debtor, err
	}

	return debtor, nil
}

func (r *authRepo) CheckEmailExist(ctx context.Context, user *models.User) (*models.User, error) {
	foundUser := &models.User{}
	if err := r.db.WithContext(ctx).Where("email ilike ?", user.Email).First(foundUser).Error; err != nil {
		return foundUser, err
	}
	return foundUser, nil
}

func (r *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	foundUser := &models.User{}
	if err := r.db.WithContext(ctx).Where("email ilike ?", user.Email).First(foundUser).Error; err != nil {
		return foundUser, err
	}
	return foundUser, nil
}

func (r *authRepo) GetUserDetailsByID(ctx context.Context, userId string) (*models.User, error) {
	userWallet := &models.User{}
	if err := r.db.Preload("Role").WithContext(ctx).Where("user_id = ?", userId).First(userWallet).Error; err != nil {
		return userWallet, err
	}

	return userWallet, nil
}
