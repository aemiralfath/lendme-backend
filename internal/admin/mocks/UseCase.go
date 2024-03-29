// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	body "final-project-backend/internal/admin/delivery/body"

	mock "github.com/stretchr/testify/mock"

	models "final-project-backend/internal/models"

	utils "final-project-backend/pkg/utils"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// ApproveLoan provides a mock function with given fields: ctx, lendingID
func (_m *UseCase) ApproveLoan(ctx context.Context, lendingID string) (*models.Lending, error) {
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

// CreateVoucher provides a mock function with given fields: ctx, _a1
func (_m *UseCase) CreateVoucher(ctx context.Context, _a1 body.CreateVoucherRequest) (*models.Voucher, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *models.Voucher
	if rf, ok := ret.Get(0).(func(context.Context, body.CreateVoucherRequest) *models.Voucher); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Voucher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, body.CreateVoucherRequest) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVoucherByID provides a mock function with given fields: ctx, voucherID
func (_m *UseCase) DeleteVoucherByID(ctx context.Context, voucherID string) (*models.Voucher, error) {
	ret := _m.Called(ctx, voucherID)

	var r0 *models.Voucher
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Voucher); ok {
		r0 = rf(ctx, voucherID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Voucher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, voucherID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDebtorByID provides a mock function with given fields: ctx, id
func (_m *UseCase) GetDebtorByID(ctx context.Context, id string) (*models.Debtor, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Debtor
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Debtor); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Debtor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDebtors provides a mock function with given fields: ctx, name, pagination
func (_m *UseCase) GetDebtors(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error) {
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

// GetLoans provides a mock function with given fields: ctx, name, status, pagination
func (_m *UseCase) GetLoans(ctx context.Context, name string, status []int, pagination *utils.Pagination) (*utils.Pagination, error) {
	ret := _m.Called(ctx, name, status, pagination)

	var r0 *utils.Pagination
	if rf, ok := ret.Get(0).(func(context.Context, string, []int, *utils.Pagination) *utils.Pagination); ok {
		r0 = rf(ctx, name, status, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.Pagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []int, *utils.Pagination) error); ok {
		r1 = rf(ctx, name, status, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPayments provides a mock function with given fields: ctx, name, pagination
func (_m *UseCase) GetPayments(ctx context.Context, name string, pagination *utils.Pagination) (*utils.Pagination, error) {
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

// GetSummary provides a mock function with given fields: ctx
func (_m *UseCase) GetSummary(ctx context.Context) (*body.SummaryResponse, error) {
	ret := _m.Called(ctx)

	var r0 *body.SummaryResponse
	if rf, ok := ret.Get(0).(func(context.Context) *body.SummaryResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*body.SummaryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVoucherByID provides a mock function with given fields: ctx, voucherID
func (_m *UseCase) GetVoucherByID(ctx context.Context, voucherID string) (*models.Voucher, error) {
	ret := _m.Called(ctx, voucherID)

	var r0 *models.Voucher
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Voucher); ok {
		r0 = rf(ctx, voucherID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Voucher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, voucherID)
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

// RejectLoan provides a mock function with given fields: ctx, lendingID
func (_m *UseCase) RejectLoan(ctx context.Context, lendingID string) (*models.Lending, error) {
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

// UpdateDebtorByID provides a mock function with given fields: ctx, debtorID, _a2
func (_m *UseCase) UpdateDebtorByID(ctx context.Context, debtorID string, _a2 body.UpdateContractRequest) (*models.Debtor, error) {
	ret := _m.Called(ctx, debtorID, _a2)

	var r0 *models.Debtor
	if rf, ok := ret.Get(0).(func(context.Context, string, body.UpdateContractRequest) *models.Debtor); ok {
		r0 = rf(ctx, debtorID, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Debtor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, body.UpdateContractRequest) error); ok {
		r1 = rf(ctx, debtorID, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateInstallmentByID provides a mock function with given fields: ctx, installmentID, _a2
func (_m *UseCase) UpdateInstallmentByID(ctx context.Context, installmentID string, _a2 body.UpdateInstallmentRequest) (*models.Installment, error) {
	ret := _m.Called(ctx, installmentID, _a2)

	var r0 *models.Installment
	if rf, ok := ret.Get(0).(func(context.Context, string, body.UpdateInstallmentRequest) *models.Installment); ok {
		r0 = rf(ctx, installmentID, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Installment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, body.UpdateInstallmentRequest) error); ok {
		r1 = rf(ctx, installmentID, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateVoucherByID provides a mock function with given fields: ctx, voucherID, _a2
func (_m *UseCase) UpdateVoucherByID(ctx context.Context, voucherID string, _a2 body.UpdateVoucherRequest) (*models.Voucher, error) {
	ret := _m.Called(ctx, voucherID, _a2)

	var r0 *models.Voucher
	if rf, ok := ret.Get(0).(func(context.Context, string, body.UpdateVoucherRequest) *models.Voucher); ok {
		r0 = rf(ctx, voucherID, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Voucher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, body.UpdateVoucherRequest) error); ok {
		r1 = rf(ctx, voucherID, _a2)
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
