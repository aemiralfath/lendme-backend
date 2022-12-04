package usecase

import (
	"context"
	"errors"
	"final-project-backend/config"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/internal/admin/mocks"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/pagination"
	"final-project-backend/pkg/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func TestAdminUC_GetDebtors(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtors", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtors", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetDebtors(context.Background(), "name", &pagination.Pagination{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_GetPayments(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetPayments", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetPayments", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetPayments(context.Background(), "name", &pagination.Pagination{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_GetVouchers(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetVouchers(context.Background(), "name", &pagination.Pagination{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_GetLoans(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetLoans(context.Background(), "name", []int{}, &pagination.Pagination{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_GetDebtorByID(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name: "error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetDebtorByID(context.Background(), "id")
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_GetInstallmentByID(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.InstallmentNotExist),
		},
		{
			name: "error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetInstallmentByID(context.Background(), "id")
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_GetVoucherByID(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.VoucherNotExist),
		},
		{
			name: "error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetVoucherByID(context.Background(), "id")
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_GetLoanByID(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.LendingIDNotExist),
		},
		{
			name: "error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetLoanByID(context.Background(), "id")
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_GetSummary(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingAmount", mock.Anything).Return(float64(1), nil)
				r.On("GetReturnAmount", mock.Anything).Return(float64(1), nil)
				r.On("GetLendingAction", mock.Anything).Return([]*models.Lending{}, nil)
				r.On("GetUserAction", mock.Anything).Return([]*models.Debtor{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "user action error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingAmount", mock.Anything).Return(float64(1), nil)
				r.On("GetReturnAmount", mock.Anything).Return(float64(1), nil)
				r.On("GetLendingAction", mock.Anything).Return([]*models.Lending{}, nil)
				r.On("GetUserAction", mock.Anything).Return([]*models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "lending action error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingAmount", mock.Anything).Return(float64(1), nil)
				r.On("GetReturnAmount", mock.Anything).Return(float64(1), nil)
				r.On("GetLendingAction", mock.Anything).Return([]*models.Lending{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "return amount error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingAmount", mock.Anything).Return(float64(1), nil)
				r.On("GetReturnAmount", mock.Anything).Return(float64(1), errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "lending amount error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingAmount", mock.Anything).Return(float64(1), errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "lending total error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserTotal", mock.Anything).Return(int64(1), nil)
				r.On("GetLendingTotal", mock.Anything).Return(int64(1), errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "user total error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserTotal", mock.Anything).Return(int64(1), errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetSummary(context.Background())
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_ApproveLoan(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
				r.On("UpdateLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{Amount: 0, LoanPeriod: &models.LoanPeriod{Duration: 1}}, nil)
				r.On("CreateInstallments", mock.Anything, mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "create error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
				r.On("UpdateLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{Amount: 0, LoanPeriod: &models.LoanPeriod{Duration: 1}}, nil)
				r.On("CreateInstallments", mock.Anything, mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "update error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
				r.On("UpdateLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{Amount: 0, LoanPeriod: &models.LoanPeriod{Duration: 1}}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "lending not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name: "lending not err",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.ApproveLoan(context.Background(), "id")
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_RejectLoan(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
				r.On("UpdateLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{Amount: 0, LoanPeriod: &models.LoanPeriod{Duration: 1}}, nil)
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "update debtor error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
				r.On("UpdateLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{Amount: 0, LoanPeriod: &models.LoanPeriod{Duration: 1}}, nil)
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "debtor error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
				r.On("UpdateLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{Amount: 0, LoanPeriod: &models.LoanPeriod{Duration: 1}}, nil)
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "update error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
				r.On("UpdateLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{Amount: 0, LoanPeriod: &models.LoanPeriod{Duration: 1}}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "lending not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name: "lending not err",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLendingByID", mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.RejectLoan(context.Background(), "id")
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_UpdateDebtorByID(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetCreditHealthByID", mock.Anything, mock.Anything).Return(&models.CreditHealthType{}, nil)
				r.On("GetContractStatusByID", mock.Anything, mock.Anything).Return(&models.ContractTrackingType{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "update error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetCreditHealthByID", mock.Anything, mock.Anything).Return(&models.CreditHealthType{}, nil)
				r.On("GetContractStatusByID", mock.Anything, mock.Anything).Return(&models.ContractTrackingType{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "contract not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetCreditHealthByID", mock.Anything, mock.Anything).Return(&models.CreditHealthType{}, nil)
				r.On("GetContractStatusByID", mock.Anything, mock.Anything).Return(&models.ContractTrackingType{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.ContractIDNotExist),
		},
		{
			name: "contract error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetCreditHealthByID", mock.Anything, mock.Anything).Return(&models.CreditHealthType{}, nil)
				r.On("GetContractStatusByID", mock.Anything, mock.Anything).Return(&models.ContractTrackingType{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "credit not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetCreditHealthByID", mock.Anything, mock.Anything).Return(&models.CreditHealthType{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.CreditIDNotExist),
		},
		{
			name: "credit error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetCreditHealthByID", mock.Anything, mock.Anything).Return(&models.CreditHealthType{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "debtor not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name: "debtor error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.UpdateDebtorByID(context.Background(), "id", body.UpdateContractRequest{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_UpdateVoucherByID(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
				r.On("UpdateVoucherByID", mock.Anything, mock.Anything).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "update error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
				r.On("UpdateVoucherByID", mock.Anything, mock.Anything).Return(errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "voucher not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.VoucherNotExist),
		},
		{
			name: "voucher error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.UpdateVoucherByID(context.Background(), "id", body.UpdateVoucherRequest{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_CreateVoucher(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("CreateVoucher", mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "create error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("CreateVoucher", mock.Anything, mock.Anything).Return(&models.Voucher{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.CreateVoucher(context.Background(), body.CreateVoucherRequest{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_UpdateInstallmentByID(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "update error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "installment not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.InstallmentNotExist),
		},
		{
			name: "installment error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.UpdateInstallmentByID(context.Background(), "id", body.UpdateInstallmentRequest{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAdminUC_DeleteVoucherByID(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
				r.On("DeleteVoucher", mock.Anything, mock.Anything).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "voucher not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.VoucherNotExist),
		},
		{
			name: "voucher error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "delete error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
				r.On("DeleteVoucher", mock.Anything, mock.Anything).Return(errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAdminUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.DeleteVoucherByID(context.Background(), "id")
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}
