// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package merchants

import (
	context "context"
	domain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"

	mock "github.com/stretchr/testify/mock"

	paramsdomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

// MerchantUseCase is an autogenerated mock type for the MerchantUseCase type
type MerchantUseCase struct {
	mock.Mock
}

// CreateMerchant provides a mock function with given fields: ctx, body
func (_m *MerchantUseCase) CreateMerchant(ctx context.Context, body domain.CreateMerchantBody) (*string, error) {
	ret := _m.Called(ctx, body)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateMerchantBody) (*string, error)); ok {
		return rf(ctx, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateMerchantBody) *string); ok {
		r0 = rf(ctx, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CreateMerchantBody) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMerchant provides a mock function with given fields: ctx, merchantId
func (_m *MerchantUseCase) DeleteMerchant(ctx context.Context, merchantId string) (bool, error) {
	ret := _m.Called(ctx, merchantId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, merchantId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, merchantId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, merchantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMerchants provides a mock function with given fields: ctx, pagination
func (_m *MerchantUseCase) GetMerchants(ctx context.Context, pagination paramsdomain.PaginationParams) ([]domain.Merchant, *paramsdomain.PaginationResults, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []domain.Merchant
	var r1 *paramsdomain.PaginationResults
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, paramsdomain.PaginationParams) ([]domain.Merchant, *paramsdomain.PaginationResults, error)); ok {
		return rf(ctx, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, paramsdomain.PaginationParams) []domain.Merchant); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Merchant)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, paramsdomain.PaginationParams) *paramsdomain.PaginationResults); ok {
		r1 = rf(ctx, pagination)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*paramsdomain.PaginationResults)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, paramsdomain.PaginationParams) error); ok {
		r2 = rf(ctx, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateMerchant provides a mock function with given fields: ctx, merchantId, body
func (_m *MerchantUseCase) UpdateMerchant(ctx context.Context, merchantId string, body domain.UpdateMerchantBody) error {
	ret := _m.Called(ctx, merchantId, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.UpdateMerchantBody) error); ok {
		r0 = rf(ctx, merchantId, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMerchantUseCase creates a new instance of MerchantUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMerchantUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MerchantUseCase {
	mock := &MerchantUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
