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

func (r *adminRepo) GetDebtors(ctx context.Context) ([]*models.Debtor, error) {
	var debtors []*models.Debtor
	if err := r.db.Preload("ContractTracking").Preload("CreditHealth").Preload("User").WithContext(ctx).Find(&debtors).Error; err != nil {
		return debtors, nil
	}

	return debtors, nil
}

func (r *adminRepo) GetDebtorByID(ctx context.Context, debtorID string) (*models.Debtor, error) {
	debtor := &models.Debtor{}
	if err := r.db.Preload("User").Preload("ContractTracking").Preload("CreditHealth").WithContext(ctx).Where("debtor_id = ?", debtorID).First(debtor).Error; err != nil {
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

func (r *adminRepo) GetCreditHealthByID(ctx context.Context, healthID int) (*models.CreditHealthType, error) {
	health := &models.CreditHealthType{}
	if err := r.db.WithContext(ctx).Where("credit_health_id = ?", healthID).First(health).Error; err != nil {
		return health, err
	}

	return health, nil
}

func (r *adminRepo) UpdateDebtorByID(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error) {
	if err := r.db.Omit("ContractTracking", "CreditHealth", "User").WithContext(ctx).Where("debtor_id = ?", debtor.DebtorID).Save(debtor).Error; err != nil {
		return debtor, err
	}

	if err := r.db.Preload("User").Preload("ContractTracking").Preload("CreditHealth").WithContext(ctx).Where("debtor_id = ?", debtor.DebtorID).First(debtor).Error; err != nil {
		return debtor, err
	}

	return debtor, nil
}
