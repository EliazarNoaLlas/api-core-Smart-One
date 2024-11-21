// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package userTypes

import (
	context "context"
	domain "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain"

	mock "github.com/stretchr/testify/mock"

	paramsdomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

// UserTypeUseCase is an autogenerated mock type for the UserTypeUseCase type
type UserTypeUseCase struct {
	mock.Mock
}

// CreateUserType provides a mock function with given fields: ctx, body
func (_m *UserTypeUseCase) CreateUserType(ctx context.Context, body domain.CreateUserTypeBody) (*string, error) {
	ret := _m.Called(ctx, body)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateUserTypeBody) (*string, error)); ok {
		return rf(ctx, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateUserTypeBody) *string); ok {
		r0 = rf(ctx, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CreateUserTypeBody) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserType provides a mock function with given fields: ctx, userTypeId
func (_m *UserTypeUseCase) DeleteUserType(ctx context.Context, userTypeId string) (bool, error) {
	ret := _m.Called(ctx, userTypeId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, userTypeId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, userTypeId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userTypeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserTypes provides a mock function with given fields: ctx, pagination
func (_m *UserTypeUseCase) GetUserTypes(ctx context.Context, pagination paramsdomain.PaginationParams) ([]domain.UserType, *paramsdomain.PaginationResults, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []domain.UserType
	var r1 *paramsdomain.PaginationResults
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, paramsdomain.PaginationParams) ([]domain.UserType, *paramsdomain.PaginationResults, error)); ok {
		return rf(ctx, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, paramsdomain.PaginationParams) []domain.UserType); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UserType)
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

// UpdateUserType provides a mock function with given fields: ctx, userTypeId, body
func (_m *UserTypeUseCase) UpdateUserType(ctx context.Context, userTypeId string, body domain.UpdateUserTypeBody) error {
	ret := _m.Called(ctx, userTypeId, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.UpdateUserTypeBody) error); ok {
		r0 = rf(ctx, userTypeId, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserTypeUseCase creates a new instance of UserTypeUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserTypeUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserTypeUseCase {
	mock := &UserTypeUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
