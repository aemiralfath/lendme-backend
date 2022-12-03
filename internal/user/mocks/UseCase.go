// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	body "final-project-backend/internal/user/delivery/body"

	mock "github.com/stretchr/testify/mock"

	models "final-project-backend/internal/models"

	utils "final-project-backend/pkg/utils"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// ConfirmContract provides a mock function with given fields: ctx, userID
func (_m *UseCase) ConfirmContract(ctx context.Context, userID string) (*models.Debtor, error) {
	ret := _m.Called(ctx, userID)

	var r0 *models.Debtor
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Debtor); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Debtor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLoan provides a mock function with given fields: ctx, userID, _a2
func (_m *UseCase) CreateLoan(ctx context.Context, userID string, _a2 body.CreateLoan) (*models.Lending, error) {
	ret := _m.Called(ctx, userID, _a2)

	var r0 *models.Lending
	if rf, ok := ret.Get(0).(func(context.Context, string, body.CreateLoan) *models.Lending); ok {
		r0 = rf(ctx, userID, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Lending)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, body.CreateLoan) error); ok {
		r1 = rf(ctx, userID, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePayment provides a mock function with given fields: ctx, userID, installmentID, _a3
func (_m *UseCase) CreatePayment(ctx context.Context, userID string, installmentID string, _a3 body.CreatePayment) (*models.Payment, error) {
	ret := _m.Called(ctx, userID, installmentID, _a3)

	var r0 *models.Payment
	if rf, ok := ret.Get(0).(func(context.Context, string, string, body.CreatePayment) *models.Payment); ok {
		r0 = rf(ctx, userID, installmentID, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, body.CreatePayment) error); ok {
		r1 = rf(ctx, userID, installmentID, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDebtorDetails provides a mock function with given fields: ctx, userID
func (_m *UseCase) GetDebtorDetails(ctx context.Context, userID string) (*models.Debtor, error) {
	ret := _m.Called(ctx, userID)

	var r0 *models.Debtor
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Debtor); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Debtor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstallmentByID provides a mock function with given fields: ctx, installmentID
func (_m *UseCase) GetInstallmentByID(ctx context.Context, installmentID string) (*models.Installment, error) {
	ret := _m.Called(ctx, installmentID)

	var r0 *models.Installment
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Installment); ok {
		r0 = rf(ctx, installmentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Installment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, installmentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoanByID provides a mock function with given fields: ctx, lendingID
func (_m *UseCase) GetLoanByID(ctx context.Context, lendingID string) (*models.Lending, error) {
	ret := _m.Called(ctx, lendingID)

	var r0 *models.Lending
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Lending); ok {
		r0 = rf(ctx, lendingID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Lending)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, lendingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoans provides a mock function with given fields: ctx, userID, name, status, pagination
func (_m *UseCase) GetLoans(ctx context.Context, userID string, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error) {
	ret := _m.Called(ctx, userID, name, status, pagination)

	var r0 *utils.Pagination
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []int, *utils.Pagination) *utils.Pagination); ok {
		r0 = rf(ctx, userID, name, status, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.Pagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, []int, *utils.Pagination) error); ok {
		r1 = rf(ctx, userID, name, status, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPayments provides a mock function with given fields: ctx, userID, name, pagination
func (_m *UseCase) GetPayments(ctx context.Context, userID string, name string, pagination *utils.Pagination) (*utils.Pagination, error) {
	ret := _m.Called(ctx, userID, name, pagination)

	var r0 *utils.Pagination
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *utils.Pagination) *utils.Pagination); ok {
		r0 = rf(ctx, userID, name, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.Pagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *utils.Pagination) error); ok {
		r1 = rf(ctx, userID, name, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVouchers provides a mock function with given fields: ctx, name, pagination
func (_m *UseCase) GetVouchers(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error) {
	ret := _m.Called(ctx, name, pagination)

	var r0 *utils.Pagination
	if rf, ok := ret.Get(0).(func(context.Context, string, *utils.Pagination) *utils.Pagination); ok {
		r0 = rf(ctx, name, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.Pagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *utils.Pagination) error); ok {
		r1 = rf(ctx, name, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserByID provides a mock function with given fields: ctx, userID, _a2
func (_m *UseCase) UpdateUserByID(ctx context.Context, userID string, _a2 body.UpdateUserRequest) (*models.User, error) {
	ret := _m.Called(ctx, userID, _a2)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string, body.UpdateUserRequest) *models.User); ok {
		r0 = rf(ctx, userID, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, body.UpdateUserRequest) error); ok {
		r1 = rf(ctx, userID, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}