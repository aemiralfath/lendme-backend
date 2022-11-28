package repository

import (
	"context"
	"final-project-backend/internal/admin"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/utils"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"time"
)

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) admin.Repository {
	return &adminRepo{db: db}
}

func (r *adminRepo) GetDebtors(ctx context.Context) ([]*models.Debtor, error) {
	var debtors []*models.Debtor
	if err := r.db.Preload(clause.Associations).WithContext(ctx).Find(&debtors).Error; err != nil {
		return debtors, nil
	}

	return debtors, nil
}

func (r *adminRepo) GetLendingByID(ctx context.Context, lendingID string) (*models.Lending, error) {
	lending := &models.Lending{}
	if err := r.db.Preload(clause.Associations).WithContext(ctx).Where("lending_id = ?", lendingID).First(lending).Error; err != nil {
		return lending, err
	}

	return lending, nil
}

func (r *adminRepo) GetInstallmentByID(ctx context.Context, installmentID string) (*models.Installment, error) {
	installment := &models.Installment{}
	if err := r.db.Preload(clause.Associations).WithContext(ctx).Where("installment_id = ?", installmentID).First(installment).Error; err != nil {
		return installment, err
	}

	return installment, nil
}

func (r *adminRepo) GetLendingWithInstallmentByID(ctx context.Context, lendingID string) (*models.Lending, error) {
	lending := &models.Lending{}
	if err := r.db.Preload("Installments.InstallmentStatus").Preload(clause.Associations).WithContext(ctx).Where("lending_id = ?", lendingID).First(lending).Error; err != nil {
		return lending, err
	}

	return lending, nil
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

	debtor, err := r.GetDebtorByID(ctx, debtor.DebtorID.String())
	if err != nil {
		return debtor, err
	}

	return debtor, nil
}

func (r *adminRepo) UpdateLendingByID(ctx context.Context, lending *models.Lending) (*models.Lending, error) {
	if err := r.db.Omit("LoanPeriod", "LendingStatus", "Installments", "Debtor").WithContext(ctx).Where("lending_id = ?", lending.LendingID).Save(lending).Error; err != nil {
		return lending, err
	}

	lending, err := r.GetLendingByID(ctx, lending.LendingID.String())
	if err != nil {
		return lending, err
	}

	return lending, nil
}

func (r *adminRepo) UpdateInstallmentByID(ctx context.Context, installment *models.Installment) (*models.Installment, error) {
	if err := r.db.Omit("InstallmentStatus", "Lending").WithContext(ctx).Where("installment_id = ?", installment.InstallmentID).Save(installment).Error; err != nil {
		return installment, err
	}

	installment, err := r.GetInstallmentByID(ctx, installment.InstallmentID.String())
	if err != nil {
		return installment, err
	}

	return installment, nil
}

func (r *adminRepo) CreateInstallments(ctx context.Context, lendingID string, installments []*models.Installment) (*models.Lending, error) {
	if err := r.db.WithContext(ctx).Create(installments).Error; err != nil {
		return nil, err
	}

	lending, err := r.GetLendingWithInstallmentByID(ctx, lendingID)
	if err != nil {
		return lending, err
	}

	return lending, nil
}

func (r *adminRepo) CreateVoucher(ctx context.Context, voucher *models.Voucher) (*models.Voucher, error) {
	if err := r.db.WithContext(ctx).Create(voucher).Error; err != nil {
		return nil, err
	}

	return voucher, nil
}

func (r *adminRepo) DeleteVoucher(ctx context.Context, voucher *models.Voucher) error {
	if err := r.db.WithContext(ctx).Where("voucher_id = ?", voucher.VoucherID).Delete(voucher).Error; err != nil {
		return err
	}

	return nil
}

func (r *adminRepo) GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error) {
	lending := &models.Lending{}
	if err := r.db.WithContext(ctx).
		Preload("Debtor."+clause.Associations).
		Preload("Installments."+clause.Associations).
		Preload("LendingStatus").
		Preload("LoanPeriod").
		Preload("Debtor").
		Preload("Installments", func(db *gorm.DB) *gorm.DB {
			return db.Order("installments.due_date asc")
		}).Where("lending_id = ?", lendingID).First(lending).Error; err != nil {
		return lending, err
	}

	return lending, nil
}

func (r *adminRepo) GetLoans(ctx context.Context, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error) {
	var loans []*models.Lending

	var totalRows int64
	r.db.Model(loans).WithContext(ctx).
		Where("name ILIKE ? AND lending_status_id in ?", fmt.Sprintf("%%%s%%", name), status).
		Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	if err := r.db.WithContext(ctx).
		Preload("Debtor."+clause.Associations).
		Preload("Installments."+clause.Associations).
		Preload(clause.Associations).
		Where("name ILIKE ? AND lending_status_id in ?", fmt.Sprintf("%%%s%%", name), status).
		Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).
		Find(&loans).Error; err != nil {
		return pagination, err
	}

	pagination.Rows = loans
	return pagination, nil
}

func (r *adminRepo) GetVouchers(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error) {
	var rows []*models.Voucher
	var vouchers []*models.Voucher

	var totalRows int64
	r.db.Model(vouchers).WithContext(ctx).
		Where("name ILIKE ?", fmt.Sprintf("%%%s%%", name)).
		Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	if err := r.db.WithContext(ctx).Where("name ILIKE ?", fmt.Sprintf("%%%s%%", name)).
		Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).
		Find(&vouchers).Error; err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(loc)
	for _, voucher := range vouchers {
		if timeNow.Sub(voucher.ExpireDate) > 0 {
			if err := r.DeleteVoucher(ctx, voucher); err != nil {
				return nil, err
			}
		} else {
			rows = append(rows, voucher)
		}
	}

	pagination.Rows = rows
	return pagination, nil
}

func (r *adminRepo) GetPayments(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error) {
	var payments []*models.Payment

	var totalRows int64
	r.db.WithContext(ctx).Model(payments).
		Joins("inner join installments on installments.installment_id = payments.installment_id").
		Joins("inner join lendings on installments.lending_id = lendings.lending_id").
		Where("lendings.name ILIKE ?", fmt.Sprintf("%%%s%%", name)).
		Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	if err := r.db.WithContext(ctx).
		Joins("inner join installments on installments.installment_id = payments.installment_id").
		Joins("inner join lendings on installments.lending_id = lendings.lending_id").
		Where("lendings.name ILIKE ?", fmt.Sprintf("%%%s%%", name)).
		Preload("Voucher").
		Preload("Installment").
		Preload("Installment.Lending").
		Preload("Installment.InstallmentStatus").
		Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).
		Find(&payments).Error; err != nil {
		return pagination, err
	}

	pagination.Rows = payments
	return pagination, nil
}
