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

func (r *userRepo) CreateLending(ctx context.Context, lending *models.Lending) (*models.Lending, error) {
	if err := r.db.WithContext(ctx).Create(lending).Error; err != nil {
		return lending, err
	}

	if err := r.db.WithContext(ctx).Preload("LoanPeriod").Preload("LendingStatus").Where("lending_id = ?", lending.LendingID).First(lending).Error; err != nil {
		return lending, err
	}

	return lending, nil
}

func (r *userRepo) GetLoanPeriodByID(ctx context.Context, periodID int) (*models.LoanPeriod, error) {
	loanPeriod := &models.LoanPeriod{}
	if err := r.db.WithContext(ctx).Where("loan_period_id = ?", periodID).First(loanPeriod).Error; err != nil {
		return loanPeriod, err
	}

	return loanPeriod, nil
}

func (r *userRepo) UpdateDebtorByID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error) {
	if err := r.db.Omit("ContractTracking", "CreditHealth", "User").WithContext(ctx).Where("debtor_id = ?", debtor.DebtorID).Save(debtor).Error; err != nil {
		return debtor, err
	}

	if err := r.db.Preload("User").Preload("ContractTracking").Preload("CreditHealth").WithContext(ctx).Where("debtor_id = ?", debtor.DebtorID).First(debtor).Error; err != nil {
		return debtor, err
	}

	return debtor, nil
}

func (r *userRepo) GetDebtorDetailsByID(ctx context.Context, userID string) (*models.Debtor, error) {
	userDebtor := &models.Debtor{}
	if err := r.db.Preload("User").Preload("CreditHealth").Preload("ContractTracking").WithContext(ctx).
		Where("user_id = ?", userID).First(userDebtor).Error; err != nil {
		return userDebtor, err
	}

	return userDebtor, nil
}