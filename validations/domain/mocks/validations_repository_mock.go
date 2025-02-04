// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package validations

import (
	"context"

	validations "gitlab.smartcitiesperu.com/smartone/api-core/validations/domain"

	mock "github.com/stretchr/testify/mock"
)

// ValidationRepository is an autogenerated mock type for the ValidationRepository type
type ValidationRepository struct {
	mock.Mock
}

// RecordExists provides a mock function with given fields: ctx, params
func (_m *ValidationRepository) RecordExists(ctx context.Context, params validations.RecordExistsParams) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, validations.RecordExistsParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateExistence provides a mock function with given fields: ctx, params
func (_m *ValidationRepository) ValidateExistence(ctx context.Context, params validations.RecordExistsParams) (bool, error) {
	ret := _m.Called(ctx, params)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, validations.RecordExistsParams) (bool, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, validations.RecordExistsParams) bool); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, validations.RecordExistsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewValidationRepository creates a new instance of ValidationRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValidationRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ValidationRepository {
	mock := &ValidationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
