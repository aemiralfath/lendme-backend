package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"final-project-backend/config"
	"final-project-backend/internal/auth/delivery/body"
	"final-project-backend/internal/auth/mocks"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/logger"
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

func TestAuthHandlers_Register(t *testing.T) {
	invalidRequestBody := struct {
		Email int `json:"email"`
	}{123}

	testCase := []struct {
		name     string
		body     interface{}
		mock     func(s *mocks.UseCase)
		expected int
	}{
		{
			name: "success register",
			body: body.RegisterRequest{
				Name:        "emir",
				PhoneNumber: "080808",
				Address:     "Jakarta",
				Email:       "emir@gmail.com",
				Password:    "Tested8*",
			},
			mock: func(s *mocks.UseCase) {
				s.On("Register", mock.Anything, mock.Anything).Return(&models.User{}, nil)
			},
			expected: http.StatusCreated,
		},
		{
			name:     "invalid request",
			body:     invalidRequestBody,
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusBadRequest,
		},
		{
			name: "invalid request",
			body: body.RegisterRequest{
				Name:        "emir",
				PhoneNumber: "080808",
				Address:     "Jakarta",
				Email:       "emir@gmail.com",
				Password:    "Tested8",
			},
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusUnprocessableEntity,
		},
		{
			name: "register error internal",
			body: body.RegisterRequest{
				Name:        "emir",
				PhoneNumber: "080808",
				Address:     "Jakarta",
				Email:       "emir@gmail.com",
				Password:    "Tested8*",
			},
			mock: func(s *mocks.UseCase) {
				s.On("Register", mock.Anything, mock.Anything).Return(nil, errors.New("test"))
			},
			expected: http.StatusInternalServerError,
		},
		{
			name: "register error custom",
			body: body.RegisterRequest{
				Name:        "emir",
				PhoneNumber: "080808",
				Address:     "Jakarta",
				Email:       "emir@gmail.com",
				Password:    "Tested8*",
			},
			mock: func(s *mocks.UseCase) {
				s.On("Register", mock.Anything, mock.Anything).Return(nil, httperror.New(http.StatusBadRequest, "test"))
			},
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

			r := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(jsonValue))
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

			h := NewAuthHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.Register(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAuthHandlers_Login(t *testing.T) {
	invalidRequestBody := struct {
		Email int `json:"email"`
	}{123}

	testCase := []struct {
		name     string
		body     interface{}
		mock     func(s *mocks.UseCase)
		expected int
	}{
		{
			name: "success login",
			body: body.LoginRequest{
				Email:    "emir@gmail.com",
				Password: "Tested8*",
			},
			mock: func(s *mocks.UseCase) {
				s.On("Login", mock.Anything, mock.Anything).Return(&models.UserWithToken{}, nil)
			},
			expected: http.StatusOK,
		},
		{
			name:     "invalid request",
			body:     invalidRequestBody,
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusBadRequest,
		},
		{
			name: "invalid request entity",
			body: body.LoginRequest{
				Email:    "",
				Password: "",
			},
			mock:     func(s *mocks.UseCase) {},
			expected: http.StatusUnprocessableEntity,
		},
		{
			name: "register error internal",
			body: body.LoginRequest{
				Email:    "emir@gmail.com",
				Password: "Tested8*",
			},
			mock: func(s *mocks.UseCase) {
				s.On("Login", mock.Anything, mock.Anything).Return(nil, errors.New("test"))
			},
			expected: http.StatusInternalServerError,
		},
		{
			name: "register error custom",
			body: body.LoginRequest{
				Email:    "emir@gmail.com",
				Password: "Tested8*",
			},
			mock: func(s *mocks.UseCase) {
				s.On("Login", mock.Anything, mock.Anything).Return(nil, httperror.New(http.StatusBadRequest, "test"))
			},
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

			r := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(jsonValue))
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

			h := NewAuthHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.Login(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}

func TestAuthHandlers_UserDetails(t *testing.T) {
	testCase := []struct {
		name         string
		userID       string
		mock         func(s *mocks.UseCase)
		expected     int
		unauthorized bool
	}{
		{
			name:   "success user details",
			userID: "94222718-b84a-49a0-b403-5bf9173c3b55",
			mock: func(s *mocks.UseCase) {
				s.On("GetUserDetails", mock.Anything, mock.Anything).Return(&models.User{}, nil)
			},
			expected:     http.StatusOK,
			unauthorized: true,
		},
		{
			name:         "error unauthorized verify",
			userID:       "94222718-b84a-49a0-b403-5bf9173c3b55",
			mock:         func(s *mocks.UseCase) {},
			expected:     http.StatusUnauthorized,
			unauthorized: false,
		},
		{
			name:   "reset verify error internal",
			userID: "94222718-b84a-49a0-b403-5bf9173c3b55",
			mock: func(s *mocks.UseCase) {
				s.On("GetUserDetails", mock.Anything, mock.Anything).Return(nil, errors.New("test"))
			},
			expected:     http.StatusInternalServerError,
			unauthorized: true,
		},
		{
			name:   "reset verify error custom",
			userID: "94222718-b84a-49a0-b403-5bf9173c3b55",
			mock: func(s *mocks.UseCase) {
				s.On("GetUserDetails", mock.Anything, mock.Anything).Return(nil, httperror.New(http.StatusBadRequest, "test"))
			},
			expected:     http.StatusBadRequest,
			unauthorized: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			r := httptest.NewRequest(http.MethodGet, "/api/v1/auth/reset", nil)
			r.Header = make(http.Header)

			c.Request = r
			c.Request.Header.Set("Content-Type", "application/json")

			if tc.unauthorized {
				c.Set("userID", tc.userID)
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

			h := NewAuthHandlers(cfg, s, appLogger)

			tc.mock(s)
			h.UserDetails(c)

			assert.Equal(t, rr.Code, tc.expected)
		})
	}
}
