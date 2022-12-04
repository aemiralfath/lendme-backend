package usecase

import (
	"context"
	"errors"
	"final-project-backend/config"
	"final-project-backend/internal/models"
	"final-project-backend/internal/user/delivery/body"
	"final-project-backend/internal/user/mocks"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/pagination"
	"final-project-backend/pkg/response"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

func TestUserUC_GetLoanByID(t *testing.T) {
	testCase := []struct {
		name        string
		input       string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name:  "success",
			input: "lending_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name:  "error not found",
			input: "lending_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.LendingIDNotExist),
		},
		{
			name:  "error",
			input: "lending_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetLoanByID(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestUserUC_GetInstallmentByID(t *testing.T) {
	testCase := []struct {
		name        string
		input       string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name:  "success",
			input: "installment_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
			},
			expectedErr: nil,
		},
		{
			name:  "error not found",
			input: "installment_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.InstallmentNotExist),
		},
		{
			name:  "error",
			input: "installment_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetInstallmentByID(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestUserUC_GetDebtorDetails(t *testing.T) {
	testCase := []struct {
		name        string
		input       string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name:  "success",
			input: "user_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expectedErr: nil,
		},
		{
			name:  "error not found",
			input: "user_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name:  "error",
			input: "user_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetDebtorDetails(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestUserUC_GetLoans(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "loan error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "error not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name: "error debtor id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetLoans(context.Background(), "id", "name", []int{}, &pagination.Pagination{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestUserUC_GetPayments(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetPayments", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "payment error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("GetPayments", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "error not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name: "error debtor id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetPayments(context.Background(), "id", "name", &pagination.Pagination{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestUserUC_GetVouchers(t *testing.T) {
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
			name: "payment error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetVouchers(context.Background(), "name", &pagination.Pagination{})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestUserUC_ConfirmContract(t *testing.T) {
	testCase := []struct {
		name        string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 4}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "error update",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 4}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "already accepted",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 5}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.ContractAlreadyAccepted),
		},
		{
			name: "not accepted",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 1}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.ContractNotAccepted),
		},
		{
			name: "error not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name: "error debtor id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.ConfirmContract(context.Background(), "id")
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestUserUC_UpdateUserByID(t *testing.T) {
	testCase := []struct {
		name        string
		email       string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name:  "success with same email",
			email: "email@email.com",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, nil)
				r.On("UpdateUser", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, nil)
			},
			expectedErr: nil,
		},
		{
			name:  "error update",
			email: "email@email.com",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, nil)
				r.On("UpdateUser", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name:  "success with diff email",
			email: "email2@email.com",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, nil)
				r.On("CheckEmailExist", mock.Anything, mock.Anything).Return(&models.User{Email: ""}, nil)
				r.On("UpdateUser", mock.Anything, mock.Anything).Return(&models.User{}, nil)
			},
			expectedErr: nil,
		},
		{
			name:  "error email exist",
			email: "email2@email.com",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, nil)
				r.On("CheckEmailExist", mock.Anything, mock.Anything).Return(&models.User{Email: "email2@email.com"}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.EmailAlreadyExistMessage),
		},
		{
			name:  "error email exist",
			email: "email2@email.com",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, nil)
				r.On("CheckEmailExist", mock.Anything, mock.Anything).Return(&models.User{Email: "email2@email.com"}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.EmailAlreadyExistMessage),
		},
		{
			name:  "error email",
			email: "email2@email.com",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, nil)
				r.On("CheckEmailExist", mock.Anything, mock.Anything).Return(&models.User{Email: ""}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name:  "error user not found",
			email: "email@email.com",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.UserIDNotExist),
		},
		{
			name:  "error user",
			email: "email@email.com",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).Return(&models.User{Email: "email@email.com"}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.UpdateUserByID(context.Background(), "useID", body.UpdateUserRequest{Email: tc.email})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestUserUC_CreateLoan(t *testing.T) {
	testCase := []struct {
		name        string
		amount      float64
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name:   "success",
			amount: 0,
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 5}, nil)
				r.On("GetLoanPeriodByID", mock.Anything, mock.Anything).Return(&models.LoanPeriod{}, nil)
				r.On("CreateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expectedErr: nil,
		},
		{
			name:   "error update credit used",
			amount: 0,
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 5}, nil)
				r.On("GetLoanPeriodByID", mock.Anything, mock.Anything).Return(&models.LoanPeriod{}, nil)
				r.On("CreateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name:   "error create lending",
			amount: 0,
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 5}, nil)
				r.On("GetLoanPeriodByID", mock.Anything, mock.Anything).Return(&models.LoanPeriod{}, nil)
				r.On("CreateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name:   "exceed credit limit",
			amount: 1000000,
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 5, CreditHealthID: 1, CreditLimit: 0}, nil)
				r.On("GetLoanPeriodByID", mock.Anything, mock.Anything).Return(&models.LoanPeriod{Percentage: 100}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.LoanAmountExceedCreditLimit),
		},
		{
			name:   "exceed credit limit warning",
			amount: 1000000,
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 5, CreditHealthID: 2, CreditLimit: 0}, nil)
				r.On("GetLoanPeriodByID", mock.Anything, mock.Anything).Return(&models.LoanPeriod{Percentage: 100}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.LoanAmountExceedCreditLimitWarning),
		},
		{
			name:   "credit blocked",
			amount: 1000000,
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 5, CreditHealthID: 3, CreditLimit: 0}, nil)
				r.On("GetLoanPeriodByID", mock.Anything, mock.Anything).Return(&models.LoanPeriod{Percentage: 100}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.CreditHealthStatusBlocked),
		},
		{
			name: "loan not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 5}, nil)
				r.On("GetLoanPeriodByID", mock.Anything, mock.Anything).Return(&models.LoanPeriod{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.LoanPeriodNotExist),
		},
		{
			name: "loan err",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 5}, nil)
				r.On("GetLoanPeriodByID", mock.Anything, mock.Anything).Return(&models.LoanPeriod{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "contract not confirmed",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{ContractTrackingID: 4}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.ContractNotConfirmed),
		},
		{
			name: "error not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name: "error debtor id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.CreateLoan(context.Background(), "id", body.CreateLoan{Amount: tc.amount})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestUserUC_CreatePayment(t *testing.T) {
	idA, _ := uuid.Parse("3aae9268-722e-11ed-94c2-aa665a52e827")
	idB, _ := uuid.Parse("3aae9268-722e-11ed-94c2-aa665a52e826")
	testCase := []struct {
		name        string
		voucher     string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateInstallment", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "update lending error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateInstallment", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "success totol paid increment",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{{InstallmentStatusID: 1}, {InstallmentStatusID: 2}}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateInstallment", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "update installment error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateInstallment", mock.Anything, mock.Anything).Return(&models.Installment{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "update debtor failed",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "success delay warning",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA, TotalDelay: 11}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1, DueDate: time.Now()}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateInstallment", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "success credit use minus",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA, CreditUsed: -1}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateInstallment", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "success minus",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1, DueDate: time.Now().Add(time.Hour)}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateInstallment", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "success not minus",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA, TotalDelay: 10}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1, DueDate: time.Now().Add(time.Hour)}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateInstallment", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "create payment error",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name:    "success voucher",
			voucher: "1",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{DiscountQuota: 1, ActiveDate: time.Now(), ExpireDate: time.Now().Add(time.Hour)}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateVoucher", mock.Anything, mock.Anything).Return(nil)
				r.On("DeleteVoucher", mock.Anything, mock.Anything).Return(nil)
				r.On("UpdateDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
				r.On("UpdateInstallment", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
				r.On("UpdateLending", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expectedErr: nil,
		},
		{
			name:    "error delete voucher",
			voucher: "1",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{DiscountQuota: 1, ActiveDate: time.Now(), ExpireDate: time.Now().Add(time.Hour)}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateVoucher", mock.Anything, mock.Anything).Return(nil)
				r.On("DeleteVoucher", mock.Anything, mock.Anything).Return(errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name:    "update voucher error",
			voucher: "1",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{DiscountQuota: 1, ActiveDate: time.Now(), ExpireDate: time.Now().Add(time.Hour)}, nil)
				r.On("CreatePayment", mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
				r.On("UpdateVoucher", mock.Anything, mock.Anything).Return(errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name:    "voucher not found",
			voucher: "1",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{DiscountQuota: 1, ActiveDate: time.Now(), ExpireDate: time.Now().Add(time.Hour)}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.VoucherNotExist),
		},
		{
			name:    "voucher err",
			voucher: "1",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{DiscountQuota: 1, ActiveDate: time.Now(), ExpireDate: time.Now().Add(time.Hour)}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name:    "voucher expired",
			voucher: "1",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
				r.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{DiscountQuota: 0, ActiveDate: time.Now(), ExpireDate: time.Now().Add(time.Hour)}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.VoucherNotExist),
		},
		{
			name: "installment not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.InstallmentNotExist),
		},
		{
			name: "installment err",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "installment is paid",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 2}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.InstallmentAlreadyPaid),
		},
		{
			name: "installment not match",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idB, Installments: &[]models.Installment{}}, nil)
				r.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{InstallmentStatusID: 1}, nil)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.LendingInstallmentNotMatch),
		},
		{
			name: "loan not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.LendingIDNotExist),
		},
		{
			name: "loan err",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{DebtorID: idA}, nil)
				r.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{DebtorID: idA, Installments: &[]models.Installment{}}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "error not found",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, gorm.ErrRecordNotFound)
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.DebtorIDNotExist),
		},
		{
			name: "error debtor id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetDebtorDetailsByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expectedErr: errors.New("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewUserUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.CreatePayment(context.Background(), "id", "id", body.CreatePayment{VoucherID: tc.voucher})
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}
