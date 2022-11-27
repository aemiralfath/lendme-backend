package repository

import (
	"context"
	"final-project-backend/internal/models"
	"final-project-backend/internal/user"
	"final-project-backend/pkg/utils"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
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

	lending, err := r.GetLoanByID(ctx, lending.LendingID.String())
	if err != nil {
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

	if err := r.db.Preload(clause.Associations).WithContext(ctx).Where("debtor_id = ?", debtor.DebtorID).First(debtor).Error; err != nil {
		return debtor, err
	}

	return debtor, nil
}

func (r *userRepo) GetDebtorDetailsByID(ctx context.Context, userID string) (*models.Debtor, error) {
	userDebtor := &models.Debtor{}
	if err := r.db.Preload(clause.Associations).WithContext(ctx).
		Where("user_id = ?", userID).First(userDebtor).Error; err != nil {
		return userDebtor, err
	}

	return userDebtor, nil
}

func (r *userRepo) GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error) {
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

func (r *userRepo) GetLoans(ctx context.Context, debtorID, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error) {
	var loans []*models.Lending

	var totalRows int64
	r.db.Model(loans).
		Where("debtor_id = ? AND name ILIKE ? AND lending_status_id in ?", debtorID, fmt.Sprintf("%%%s%%", name), status).
		Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	if err := r.db.WithContext(ctx).
		Preload("Debtor."+clause.Associations).
		Preload("Installments."+clause.Associations).
		Preload(clause.Associations).
		Where("debtor_id = ? AND name ILIKE ? AND lending_status_id in ?", debtorID, fmt.Sprintf("%%%s%%", name), status).
		Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).
		Find(&loans).Error; err != nil {
		return pagination, err
	}

	pagination.Rows = loans
	return pagination, nil
}

func (r *userRepo) GetInstallmentByID(ctx context.Context, installmentID string) (*models.Installment, error) {
	installment := &models.Installment{}
	if err := r.db.Preload(clause.Associations).WithContext(ctx).
		Where("installment_id = ?", installmentID).First(installment).Error; err != nil {
		return installment, err
	}

	return installment, nil
}

func (r *userRepo) GetVoucherByID(ctx context.Context, voucherID string) (*models.Voucher, error) {
	voucher := &models.Voucher{}
	if err := r.db.WithContext(ctx).
		Where("voucher_id = ?", voucherID).First(voucher).Error; err != nil {
		return voucher, err
	}

	return voucher, nil
}

func (r *userRepo) CreatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	if err := r.db.WithContext(ctx).Create(payment).Error; err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *userRepo) UpdateInstallment(ctx context.Context, installment *models.Installment) (*models.Installment, error) {
	if err := r.db.Omit("InstallmentStatus", "Lending").WithContext(ctx).Where("installment_id = ?", installment.InstallmentID).Save(installment).Error; err != nil {
		return installment, err
	}

	installment, err := r.GetInstallmentByID(ctx, installment.InstallmentID.String())
	if err != nil {
		return installment, nil
	}

	return installment, nil
}

func (r *userRepo) UpdateLending(ctx context.Context, lending *models.Lending) (*models.Lending, error) {
	if err := r.db.Omit("Debtor", "LoanPeriod", "LendingStatus", "Installments").WithContext(ctx).Where("lending_id = ?", lending.LendingID).Save(lending).Error; err != nil {
		return lending, err
	}

	return lending, nil
}
