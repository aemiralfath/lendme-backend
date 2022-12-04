package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"final-project-backend/config"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/internal/admin/mocks"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/logger"
	"final-project-backend/pkg/utils"
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

func TestAdminHandlers_CreateVoucher(t *testing.T) {
	invalidRequestBody := struct {
		Name int `json:"name"`
	}{123}

	testCase := []struct {
		name     string
		body     interface{}
		mock     func(s *mocks.UseCase)
		expected int
	}{
		{
			name: "success",
			body: body.CreateVoucherRequest{
				Name:            "end year",
				DiscountPayment: 1,
				DiscountQuota:   1,
				ActiveDate:      "31-12-2022 23:59:59",
				ExpireDate:      "31-12-2022 23:59:59",
			},
			mock: func(s *mocks.UseCase) {
				s.On("CreateVoucher", mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
			},
			expected: http.StatusOK,
		},
		{
			name: "internal server error",
			body: body.CreateVoucherRequest{
				Name:            "end year",
				DiscountPayment: 1,
				DiscountQuota:   1,
				ActiveDate:      "31-12-2022 23:59:59",
				ExpireDate:      "31-12-2022 23:59:59",
			},
			mock: func(s *mocks.UseCase) {
				s.On("CreateVoucher", mock.Anything, mock.Anything).Return(&models.Voucher{}, errors.New("test"))
			},
			expected: http.StatusInternalServerError,
		},
		{
			name: "custom error",
			body: body.CreateVoucherRequest{
				Name:            "end year",
				DiscountPayment: 1,
				DiscountQuota:   1,
				ActiveDate:      "31-12-2022 23:59:59",
				ExpireDate:      "31-12-2022 23:59:59",
			},
			mock: func(s *mocks.UseCase) {
				s.On("CreateVoucher", mock.Anything, mock.Anything).Return(&models.Voucher{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "field not valid",
			body: body.CreateVoucherRequest{
				Name:            "",
				DiscountPayment: 0,
				DiscountQuota:   0,
				ActiveDate:      "",
				ExpireDate:      "3",
			},
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusUnprocessableEntity,
		},
		{
			name:     "invalid request",
			body:     invalidRequestBody,
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusBadRequest,
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

			r := httptest.NewRequest(http.MethodPost, "/vouchers", bytes.NewBuffer(jsonValue))
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")
			MockJsonPost(c, tc.body)

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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.CreateVoucher(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_UpdateDebtorByID(t *testing.T) {
	invalidRequestBody := struct {
		CreditLimit string `json:"credit_limit"`
	}{"123"}

	testCase := []struct {
		name     string
		body     interface{}
		mock     func(s *mocks.UseCase)
		expected int
	}{
		{
			name: "success",
			body: body.UpdateContractRequest{
				CreditLimit:      1000000,
				CreditHealthID:   1,
				ContractStatusID: 5,
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateDebtorByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expected: http.StatusOK,
		},
		{
			name: "internal server error",
			body: body.UpdateContractRequest{
				CreditLimit:      1000000,
				CreditHealthID:   1,
				ContractStatusID: 5,
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateDebtorByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expected: http.StatusInternalServerError,
		},
		{
			name: "custom error",
			body: body.UpdateContractRequest{
				CreditLimit:      1000000,
				CreditHealthID:   1,
				ContractStatusID: 5,
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateDebtorByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.Debtor{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "field not valid",
			body: body.UpdateContractRequest{
				CreditLimit:      0,
				CreditHealthID:   0,
				ContractStatusID: 0,
			},
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusUnprocessableEntity,
		},
		{
			name:     "invalid request",
			body:     invalidRequestBody,
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusBadRequest,
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

			r := httptest.NewRequest(http.MethodPut, "/debtors/1", bytes.NewBuffer(jsonValue))
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")
			MockJsonPost(c, tc.body)

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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.UpdateDebtorByID(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_UpdateVoucher(t *testing.T) {
	invalidRequestBody := struct {
		Name int `json:"name"`
	}{123}

	testCase := []struct {
		name     string
		body     interface{}
		mock     func(s *mocks.UseCase)
		expected int
	}{
		{
			name: "success",
			body: body.UpdateVoucherRequest{
				Name:            "end year",
				DiscountPayment: 1,
				DiscountQuota:   1,
				ActiveDate:      "31-12-2022 23:59:59",
				ExpireDate:      "31-12-2022 23:59:59",
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateVoucherByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
			},
			expected: http.StatusOK,
		},
		{
			name: "internal server error",
			body: body.UpdateVoucherRequest{
				Name:            "end year",
				DiscountPayment: 1,
				DiscountQuota:   1,
				ActiveDate:      "31-12-2022 23:59:59",
				ExpireDate:      "31-12-2022 23:59:59",
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateVoucherByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.Voucher{}, errors.New("test"))
			},
			expected: http.StatusInternalServerError,
		},
		{
			name: "custom error",
			body: body.UpdateVoucherRequest{
				Name:            "end year",
				DiscountPayment: 1,
				DiscountQuota:   1,
				ActiveDate:      "31-12-2022 23:59:59",
				ExpireDate:      "31-12-2022 23:59:59",
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateVoucherByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.Voucher{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "field not valid",
			body: body.UpdateVoucherRequest{
				Name:            "",
				DiscountPayment: 0,
				DiscountQuota:   0,
				ActiveDate:      "",
				ExpireDate:      "",
			},
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusUnprocessableEntity,
		},
		{
			name:     "invalid request",
			body:     invalidRequestBody,
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusBadRequest,
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

			r := httptest.NewRequest(http.MethodPut, "/vouchers/1", bytes.NewBuffer(jsonValue))
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")
			MockJsonPost(c, tc.body)

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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.UpdateVoucher(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_UpdateInstallmentByID(t *testing.T) {
	invalidRequestBody := struct {
		DueDate int `json:"due_date"`
	}{123}

	testCase := []struct {
		name     string
		body     interface{}
		mock     func(s *mocks.UseCase)
		expected int
	}{
		{
			name: "success",
			body: body.UpdateInstallmentRequest{
				DueDate: "31-12-2022 23:59:59",
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateInstallmentByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
			},
			expected: http.StatusOK,
		},
		{
			name: "internal server error",
			body: body.UpdateInstallmentRequest{
				DueDate: "31-12-2022 23:59:59",
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateInstallmentByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.Installment{}, errors.New("test"))
			},
			expected: http.StatusInternalServerError,
		},
		{
			name: "custom error",
			body: body.UpdateInstallmentRequest{
				DueDate: "31-12-2022 23:59:59",
			},
			mock: func(s *mocks.UseCase) {
				s.On("UpdateInstallmentByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.Installment{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "field not valid",
			body: body.UpdateInstallmentRequest{
				DueDate: "",
			},
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusUnprocessableEntity,
		},
		{
			name:     "invalid request",
			body:     invalidRequestBody,
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusBadRequest,
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

			r := httptest.NewRequest(http.MethodPut, "/loans/installments/1", bytes.NewBuffer(jsonValue))
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")
			MockJsonPost(c, tc.body)

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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.UpdateInstallmentByID(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_GetSummary(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success",
			mock: func(s *mocks.UseCase) {
				s.On("GetSummary", mock.Anything).Return(&body.SummaryResponse{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("GetSummary", mock.Anything).Return(&body.SummaryResponse{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("GetSummary", mock.Anything).Return(&body.SummaryResponse{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.GetSummary(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_GetLoanByID(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expected:     http.StatusOK,
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
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoanByID", mock.Anything, mock.Anything).Return(&models.Lending{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.GetLoanByID(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_GetInstallmentByID(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success",
			mock: func(s *mocks.UseCase) {
				s.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, nil)
			},
			expected:     http.StatusOK,
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
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("GetInstallmentByID", mock.Anything, mock.Anything).Return(&models.Installment{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.GetInstallmentByID(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_RejectLoan(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success",
			mock: func(s *mocks.UseCase) {
				s.On("RejectLoan", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("RejectLoan", mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("RejectLoan", mock.Anything, mock.Anything).Return(&models.Lending{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodDelete, "/loans/1", nil)
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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.RejectLoan(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_ApproveLoan(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success",
			mock: func(s *mocks.UseCase) {
				s.On("ApproveLoan", mock.Anything, mock.Anything).Return(&models.Lending{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("ApproveLoan", mock.Anything, mock.Anything).Return(&models.Lending{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("ApproveLoan", mock.Anything, mock.Anything).Return(&models.Lending{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodPut, "/loans/1", nil)
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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.ApproveLoan(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_GetVoucherByID(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success",
			mock: func(s *mocks.UseCase) {
				s.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("GetVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/vouchers/1", nil)
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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.GetVoucherByID(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_GetDebtorByID(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtorByID", mock.Anything, mock.Anything).Return(&models.Debtor{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/debtors/1", nil)
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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.GetDebtorByID(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_DeleteVoucher(t *testing.T) {
	testCase := []struct {
		name         string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name: "success",
			mock: func(s *mocks.UseCase) {
				s.On("DeleteVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name: "internal server error",
			mock: func(s *mocks.UseCase) {
				s.On("DeleteVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name: "custom error",
			mock: func(s *mocks.UseCase) {
				s.On("DeleteVoucherByID", mock.Anything, mock.Anything).Return(&models.Voucher{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodDelete, "/vouchers/1", nil)
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

			h := NewAdminHandlers(cfg, s, appLogger)
			tc.mock(s)
			h.DeleteVoucher(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_GetDebtors(t *testing.T) {
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
				s.On("GetDebtors", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=credit_limit&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtors", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=credit_used&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtors", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=total_delay&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtors", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "custom error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtors", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name:   "internal server error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetDebtors", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/debtors"+tc.filter, nil)
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

			h := NewAdminHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.GetDebtors(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_GetVouchers(t *testing.T) {
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
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=expire_date&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=discount_payment&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=discount_quota&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "custom error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name:   "internal server error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetVouchers", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, errors.New("test"))
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

			h := NewAdminHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.GetVouchers(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_GetPayments(t *testing.T) {
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
				s.On("GetPayments", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=payment_amount&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetPayments", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},

		{
			name:   "custom error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetPayments", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name:   "internal server error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetPayments", mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
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

			h := NewAdminHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.GetPayments(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAdminHandlers_GetLoans(t *testing.T) {
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
				s.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:   "success using filter",
			filter: "?sortBy=amount&sort=asc&status=history",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},

		{
			name:   "custom error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
		{
			name:   "internal server error",
			filter: "",
			mock: func(s *mocks.UseCase) {
				s.On("GetLoans", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&utils.Pagination{}, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
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

			h := NewAdminHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.GetLoans(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}
