package repository

import (
	"context"
	"final-project-backend/internal/models"
	"final-project-backend/internal/user"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepo{db: db}
}

func (r *userRepo) GetDebtorDetailsByID(ctx context.Context, userID string) (*models.Debtor, error) {
	userDebtor := &models.Debtor{}
	if err := r.db.Preload("CreditHealth").Preload("ContractTracking").WithContext(ctx).
		Where("user_id = ?", userID).First(userDebtor).Error; err != nil {
		return userDebtor, err
	}

	return userDebtor, nil
}
