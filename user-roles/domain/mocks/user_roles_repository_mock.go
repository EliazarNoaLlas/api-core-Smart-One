// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package user_roles

import (
	context "context"
	domain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"

	mock "github.com/stretchr/testify/mock"

	paramsdomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

// UserRoleRepository is an autogenerated mock type for the UserRoleRepository type
type UserRoleRepository struct {
	mock.Mock
}

// CreateUserRole provides a mock function with given fields: ctx, userRoleId, userId, body
func (_m *UserRoleRepository) CreateUserRole(ctx context.Context, userRoleId string, userId string, body domain.CreateUserRoleBody) (*string, error) {
	ret := _m.Called(ctx, userRoleId, userId, body)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, domain.CreateUserRoleBody) (*string, error)); ok {
		return rf(ctx, userRoleId, userId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, domain.CreateUserRoleBody) *string); ok {
		r0 = rf(ctx, userRoleId, userId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, domain.CreateUserRoleBody) error); ok {
		r1 = rf(ctx, userRoleId, userId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserRole provides a mock function with given fields: ctx, userId, userRoleId
func (_m *UserRoleRepository) DeleteUserRole(ctx context.Context, userId string, userRoleId string) (bool, error) {
	ret := _m.Called(ctx, userId, userRoleId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, userId, userRoleId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, userId, userRoleId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userId, userRoleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalUserRolesByUser provides a mock function with given fields: ctx, userId, pagination
func (_m *UserRoleRepository) GetTotalUserRolesByUser(ctx context.Context, userId string, pagination paramsdomain.PaginationParams) (*int, error) {
	ret := _m.Called(ctx, userId, pagination)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, paramsdomain.PaginationParams) (*int, error)); ok {
		return rf(ctx, userId, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, paramsdomain.PaginationParams) *int); ok {
		r0 = rf(ctx, userId, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, paramsdomain.PaginationParams) error); ok {
		r1 = rf(ctx, userId, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserRolesByUser provides a mock function with given fields: ctx, userId, pagination
func (_m *UserRoleRepository) GetUserRolesByUser(ctx context.Context, userId string, pagination paramsdomain.PaginationParams) ([]domain.UserRole, error) {
	ret := _m.Called(ctx, userId, pagination)

	var r0 []domain.UserRole
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, paramsdomain.PaginationParams) ([]domain.UserRole, error)); ok {
		return rf(ctx, userId, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, paramsdomain.PaginationParams) []domain.UserRole); ok {
		r0 = rf(ctx, userId, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UserRole)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, paramsdomain.PaginationParams) error); ok {
		r1 = rf(ctx, userId, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserRole provides a mock function with given fields: ctx, userId, userRoleId, body
func (_m *UserRoleRepository) UpdateUserRole(ctx context.Context, userId string, userRoleId string, body domain.CreateUserRoleBody) error {
	ret := _m.Called(ctx, userId, userRoleId, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, domain.CreateUserRoleBody) error); ok {
		r0 = rf(ctx, userId, userRoleId, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyUserHasRole provides a mock function with given fields: ctx, userId, roleId
func (_m *UserRoleRepository) VerifyUserHasRole(ctx context.Context, userId string, roleId string) (bool, error) {
	ret := _m.Called(ctx, userId, roleId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, userId, roleId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, userId, roleId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userId, roleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRoleRepository creates a new instance of UserRoleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRoleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRoleRepository {
	mock := &UserRoleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
