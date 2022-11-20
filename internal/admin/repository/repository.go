package repository

import (
	"context"
	"final-project-backend/internal/admin"
	"final-project-backend/internal/models"
	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) admin.Repository {
	return &adminRepo{db: db}
}

func (r *adminRepo) GetDebtorByID(ctx context.Context, debtorID string) (*models.Debtor, error) {
	debtor := &models.Debtor{}
	if err := r.db.Preload("ContractTracking").Preload("CreditHealth").WithContext(ctx).Where("debtor_id = ?", debtorID).First(debtor).Error; err != nil {
		return debtor, err
	}

	return debtor, nil
}
func (r *adminRepo) GetContractStatusByID(ctx context.Context, contractID int) (*models.ContractTrackingType, error) {
	contract := &models.ContractTrackingType{}
	if err := r.db.WithContext(ctx).Where("contract_tracking_id = ?", contractID).First(contract).Error; err != nil {
		return contract, err
	}

	return contract, nil
}
func (r *adminRepo) UpdateDebtorStatusByID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error) {
	if err := r.db.Omit("ContractTracking", "CreditHealth").WithContext(ctx).Where("debtor_id = ?", debtor.DebtorID).Save(debtor).Error; err != nil {
		return debtor, err
	}

	if err := r.db.Preload("ContractTracking").Preload("CreditHealth").WithContext(ctx).Where("debtor_id = ?", debtor.DebtorID).First(debtor).Error; err != nil {
		return debtor, err
	}

	return debtor, nil
}
