package usecase

import (
	"context"
	"final-project-backend/config"
	"final-project-backend/internal/auth/delivery/body"
	"final-project-backend/internal/auth/mocks"
	"final-project-backend/internal/models"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestAuthUseCase_Register(t *testing.T) {
	testCase := []struct {
		name        string
		body        body.RegisterRequest
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success register",
			body: body.RegisterRequest{Name: "emir", PhoneNumber: "080808", Address: "Jakarta", Email: "a@test.com", Password: "test"},
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("CheckEmailExist", mock.Anything, mock.Anything).Return(&models.User{}, nil)
				r.On("Register", mock.Anything, mock.Anything).Return(&models.User{}, nil)
				r.On("CreateDebtor", mock.Anything, mock.Anything).Return(nil, nil)
			},
			expectedErr: nil,
		},
		{
			name: "email exist",
			body: body.RegisterRequest{Name: "emir", PhoneNumber: "080808", Address: "Jakarta", Email: "a@test.com", Password: "test"},
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("CheckEmailExist", mock.Anything, mock.Anything).Return(&models.User{Email: "a@test.com"}, httperror.New(http.StatusBadRequest, response.EmailAlreadyExistMessage))
			},
			expectedErr: httperror.New(http.StatusBadRequest, response.EmailAlreadyExistMessage),
		},
		{
			name: "register error",
			body: body.RegisterRequest{Name: "emir", PhoneNumber: "080808", Address: "Jakarta", Email: "a@test.com", Password: "test"},
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("CheckEmailExist", mock.Anything, mock.Anything).Return(&models.User{}, nil)
				r.On("Register", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("test"))
			},
			expectedErr: fmt.Errorf("test"),
		},
		{
			name: "create debtor",
			body: body.RegisterRequest{Name: "emir", PhoneNumber: "080808", Address: "Jakarta", Email: "a@test.com", Password: "test"},
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("CheckEmailExist", mock.Anything, mock.Anything).Return(&models.User{}, nil)
				r.On("Register", mock.Anything, mock.Anything).Return(&models.User{}, nil)
				r.On("CreateDebtor", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("test"))
			},
			expectedErr: fmt.Errorf("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAuthUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.Register(context.Background(), tc.body)
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAuthUseCase_Login(t *testing.T) {
	testCase := []struct {
		name        string
		body        body.LoginRequest
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name: "success login",
			body: body.LoginRequest{Email: "a@test.com", Password: "Tested8*"},
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("FindByEmail", mock.Anything, mock.Anything).
					Return(&models.User{Name: "emir", Email: "a@test.com", Password: "$2a$10$WKul/6gjYoYjOXuNVX4XGen1ZkWYb1PKFiI5vlZp5TFerZh6nTujG"}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "email not found",
			body: body.LoginRequest{Email: "a@test.com", Password: "Tested8*"},
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("FindByEmail", mock.Anything, mock.Anything).
					Return(nil, httperror.New(http.StatusUnauthorized, response.UnauthorizedMessage))
			},
			expectedErr: httperror.New(http.StatusUnauthorized, response.UnauthorizedMessage),
		},
		{
			name: "wrong password login",
			body: body.LoginRequest{Email: "a@test.com", Password: "Tested8"},
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("FindByEmail", mock.Anything, mock.Anything).
					Return(&models.User{Name: "emir", Email: "a@test.com", Password: "$2a$10$WKul/6gjYoYjOXuNVX4XGen1ZkWYb1PKFiI5vlZp5TFerZh6nTujG"}, nil)
			},
			expectedErr: httperror.New(http.StatusUnauthorized, response.UnauthorizedMessage),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAuthUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.Login(context.Background(), tc.body)
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}

func TestAuthUseCase_GetUserDetails(t *testing.T) {
	testCase := []struct {
		name        string
		input       string
		mock        func(t *testing.T, r *mocks.Repository)
		expectedErr error
	}{
		{
			name:  "success reset verify",
			input: "user_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).
					Return(&models.User{Name: "emir", Email: "test@gmail.com"}, nil)
			},
			expectedErr: nil,
		},
		{
			name:  "error get user details",
			input: "user_id",
			mock: func(t *testing.T, r *mocks.Repository) {
				r.On("GetUserDetailsByID", mock.Anything, mock.Anything).
					Return(nil, fmt.Errorf("test"))
			},
			expectedErr: fmt.Errorf("test"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			r := mocks.NewRepository(t)
			u := NewAuthUseCase(&config.Config{}, r)

			tc.mock(t, r)
			_, err := u.GetUserDetails(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}
