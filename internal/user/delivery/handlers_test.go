package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"final-project-backend/config"
	"final-project-backend/internal/models"
	"final-project-backend/internal/user/delivery/body"
	"final-project-backend/internal/user/mocks"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/logger"
	"final-project-backend/pkg/pagination"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func TestUserHandlers_CreatePayment(t *testing.T) {
	invalidRequestBody := struct {
		LendingID int `json:"lending_id"`
	}{123}

	testCase := []struct {
		name         string
		body         interface{}
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success create payment",
			body: body.CreatePayment{
				LendingID: "1",
				VoucherID: "1",
			},
			mock: func(s *mocks.UseCase) {
				s.On("CreatePayment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&models.Payment{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "internal server error",
			body: body.CreatePayment{
				LendingID: "1",
				VoucherID: "1",
			},
			mock: func(s *mocks.UseCase) {
				s.On("CreatePayment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&models.Payment{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "custome error",
			body: body.CreatePayment{
				LendingID: "1",
				VoucherID: "1",
			},
			mock: func(s *mocks.UseCase) {
				s.On("CreatePayment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&models.Payment{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name: "field not valid",
			body: body.CreatePayment{
				LendingID: "",
				VoucherID: "",
			},
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnprocessableEntity,
			unauthorized: true,
		},
		{
			name: "unauthorized",
			body: body.CreatePayment{
				LendingID: "1",
				VoucherID: "1",
			},
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnauthorized,
			unauthorized: false,
		},
		{
			name:         "invalid request body",
			body:         invalidRequestBody,
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			jsonValue, err := json.Marshal(tc.body)
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodPost, "/loans/installments/1", bytes.NewBuffer(jsonValue))
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")
			MockJsonPost(c, tc.body)

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.CreatePayment(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestUserHandlers_CreateLoan(t *testing.T) {
	invalidRequestBody := struct {
		LoadPeriodID string `json:"loan_period_id"`
	}{"1"}

	testCase := []struct {
		name         string
		body         interface{}
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success create loan",
			body: body.CreateLoan{
				LoadPeriodID: 1,
				Name:         "test",
				Amount:       1000000,
			},
			mock: func(s *mocks.UseCase) {
				s.On("CreateLoan", mock.Anything, mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "internal server error",
			body: body.CreateLoan{
				LoadPeriodID: 1,
				Name:         "test",
				Amount:       1000000,
			},
			mock: func(s *mocks.UseCase) {
				s.On("CreateLoan", mock.Anything, mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "custom error",
			body: body.CreateLoan{
				LoadPeriodID: 1,
				Name:         "test",
				Amount:       1000000,
			},
			mock: func(s *mocks.UseCase) {
				s.On("CreateLoan", mock.Anything, mock.Anything, mock.Anything).Return(&models.Lending{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name: "Field not valid",
			body: body.CreateLoan{
				LoadPeriodID: 0,
				Name:         "",
				Amount:       1,
			},
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnprocessableEntity,
			unauthorized: true,
		},
		{
			name: "unauthorized",
			body: body.CreateLoan{
				LoadPeriodID: 1,
				Name:         "test",
				Amount:       1000000,
			},
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnauthorized,
			unauthorized: false,
		},
		{
			name:         "invalid request",
			body:         invalidRequestBody,
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			jsonValue, err := json.Marshal(tc.body)
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodPost, "/loans", bytes.NewBuffer(jsonValue))
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")
			MockJsonPost(c, tc.body)

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.CreateLoan(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestUserHandlers_ContractConfirm(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success confirm contract",
			mock: func(s *mocks.UseCase) {
				s.On("ConfirmContract", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("ConfirmContract", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("ConfirmContract", mock.Anything, mock.Anything).Return(&models.Debtor{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name:         "unauthorized",
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnauthorized,
			unauthorized: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodPatch, "/details", nil)
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.ContractConfirm(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestUserHandlers_UpdateUser(t *testing.T) {
	invalidRequestBody := struct {
		Name int `json:"name"`
	}{1}

	testCase := []struct {
		name         string
		body         interface{}
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success update user",
			body: body.UpdateUserRequest{
				Name:        "Emir",
				PhoneNumber: "080808",
				Address:     "Jakarta",
				Email:       "emir@gmail.com",
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateUserByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.User{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "custom error",
			body: body.UpdateUserRequest{
				Name:        "Emir",
				PhoneNumber: "080808",
				Address:     "Jakarta",
				Email:       "emir@gmail.com",
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateUserByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.User{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name: "internal server error",
			body: body.UpdateUserRequest{
				Name:        "Emir",
				PhoneNumber: "080808",
				Address:     "Jakarta",
				Email:       "emir@gmail.com",
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateUserByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.User{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "Field not valid",
			body: body.UpdateUserRequest{
				Name:        "",
				PhoneNumber: "",
				Address:     "",
				Email:       "",
			},
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnprocessableEntity,
			unauthorized: true,
		},
		{
			name: "unauthorized",
			body: body.UpdateUserRequest{
				Name:        "Emir",
				PhoneNumber: "080808",
				Address:     "Jakarta",
				Email:       "emir@gmail.com",
			},
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnauthorized,
			unauthorized: false,
		},
		{
			name:         "invalid request",
			body:         invalidRequestBody,
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			jsonValue, err := json.Marshal(tc.body)
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodPut, "/details", bytes.NewBuffer(jsonValue))
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")
			MockJsonPost(c, tc.body)

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.UpdateUser(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestUserHandlers_DebtorDetails(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success get debtor",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtorDetails", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtorDetails", mock.Anything, mock.Anything).Return(&models.Debtor{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtorDetails", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name:         "unauthorized",
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnauthorized,
			unauthorized: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/details", nil)
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.DebtorDetails(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestUserHandlers_GetLoanByID(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success get loan",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/loans/1", nil)
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.GetLoanByID(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestUserHandlers_GetInstallmentByID(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success get installment",
			mock: func(s *mocks.UseCase) {
				s.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/loans/installments/1", nil)
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.GetInstallmentByID(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestUserHandlers_GetLoans(t *testing.T) {
	testCase := []struct {
		name         string
		filter       string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name:   "success",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=amount&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "custom error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name:   "internal server error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name:         "unauthorized",
			filter:       "",
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnauthorized,
			unauthorized: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/loans"+tc.filter, nil)
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.GetLoans(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestUserHandlers_GetVouchers(t *testing.T) {
	testCase := []struct {
		name         string
		filter       string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name:   "success",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=amount&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=discount_quota&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=discount_payment&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=expire_date&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "custom error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name:   "internal server error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/vouchers"+tc.filter, nil)
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.GetVouchers(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestUserHandlers_GetPayments(t *testing.T) {
	testCase := []struct {
		name         string
		filter       string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name:   "success",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetPayments", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=payment_amount&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetPayments", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "custom error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetPayments", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name:   "internal server error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetPayments", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&pagination.Pagination{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name:         "unauthorized",
			filter:       "",
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnauthorized,
			unauthorized: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/payments"+tc.filter, nil)
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")

			if tc.unauthorized {
				c.Set("userID", "1")
			}

			s := mocks.NewUseCase(t)

			cfg := &config.Config{
				Logger: config.Logger{
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
					Level:             "info",
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			h := NewUserHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.GetPayments(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}
