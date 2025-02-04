// Code generated by mockery v2.20.0. DO NOT EDIT.

package users

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "gitlab.smartcitiesperu.com/smartone/api-core/users/domain"

	paramsdomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, body
func (_m *UserUseCase) CreateUser(ctx context.Context, body domain.CreateUserBody) (*string, error) {
	ret := _m.Called(ctx, body)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateUserBody) (*string, error)); ok {
		return rf(ctx, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateUserBody) *string); ok {
		r0 = rf(ctx, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CreateUserBody) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, userId
func (_m *UserUseCase) DeleteUser(ctx context.Context, userId string) (bool, error) {
	ret := _m.Called(ctx, userId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMeByUser provides a mock function with given fields: ctx, userId
func (_m *UserUseCase) GetMeByUser(ctx context.Context, userId string) (*domain.UserMe, error) {
	ret := _m.Called(ctx, userId)

	var r0 *domain.UserMe
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.UserMe, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.UserMe); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserMe)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMenuByUser provides a mock function with given fields: ctx, userId
func (_m *UserUseCase) GetMenuByUser(ctx context.Context, userId string) ([]domain.MenuModule, error) {
	ret := _m.Called(ctx, userId)

	var r0 []domain.MenuModule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.MenuModule, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.MenuModule); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.MenuModule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModulePermissions provides a mock function with given fields: ctx, userId, codeModule
func (_m *UserUseCase) GetModulePermissions(ctx context.Context, userId string, codeModule string) ([]domain.Permissions, error) {
	ret := _m.Called(ctx, userId, codeModule)

	var r0 []domain.Permissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]domain.Permissions, error)); ok {
		return rf(ctx, userId, codeModule)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []domain.Permissions); ok {
		r0 = rf(ctx, userId, codeModule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Permissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userId, codeModule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, userId
func (_m *UserUseCase) GetUser(ctx context.Context, userId string) (*domain.User, error) {
	ret := _m.Called(ctx, userId)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx, searchParams, pagination
func (_m *UserUseCase) GetUsers(ctx context.Context, searchParams domain.GetUsersParams, pagination paramsdomain.PaginationParams) ([]domain.UserMultiple, *paramsdomain.PaginationResults, error) {
	ret := _m.Called(ctx, searchParams, pagination)

	var r0 []domain.UserMultiple
	var r1 *paramsdomain.PaginationResults
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetUsersParams, paramsdomain.PaginationParams) ([]domain.UserMultiple, *paramsdomain.PaginationResults, error)); ok {
		return rf(ctx, searchParams, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetUsersParams, paramsdomain.PaginationParams) []domain.UserMultiple); ok {
		r0 = rf(ctx, searchParams, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UserMultiple)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.GetUsersParams, paramsdomain.PaginationParams) *paramsdomain.PaginationResults); ok {
		r1 = rf(ctx, searchParams, pagination)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*paramsdomain.PaginationResults)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.GetUsersParams, paramsdomain.PaginationParams) error); ok {
		r2 = rf(ctx, searchParams, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LoginUser provides a mock function with given fields: ctx, body
func (_m *UserUseCase) LoginUser(ctx context.Context, body domain.LoginUserBody) (*string, *string, error) {
	ret := _m.Called(ctx, body)

	var r0 *string
	var r1 *string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.LoginUserBody) (*string, *string, error)); ok {
		return rf(ctx, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.LoginUserBody) *string); ok {
		r0 = rf(ctx, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.LoginUserBody) *string); ok {
		r1 = rf(ctx, body)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.LoginUserBody) error); ok {
		r2 = rf(ctx, body)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ResetPasswordUser provides a mock function with given fields: ctx, userId, body
func (_m *UserUseCase) ResetPasswordUser(ctx context.Context, userId string, body domain.ResetUserPasswordBody) (bool, error) {
	ret := _m.Called(ctx, userId, body)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.ResetUserPasswordBody) (bool, error)); ok {
		return rf(ctx, userId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.ResetUserPasswordBody) bool); ok {
		r0 = rf(ctx, userId, body)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.ResetUserPasswordBody) error); ok {
		r1 = rf(ctx, userId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, userId, body
func (_m *UserUseCase) UpdateUser(ctx context.Context, userId string, body domain.UpdateUserBody) error {
	ret := _m.Called(ctx, userId, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.UpdateUserBody) error); ok {
		r0 = rf(ctx, userId, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyPermissionsByUser provides a mock function with given fields: ctx, userId, storeId, codePermission
func (_m *UserUseCase) VerifyPermissionsByUser(ctx context.Context, userId string, storeId string, codePermission string) (bool, error) {
	ret := _m.Called(ctx, userId, storeId, codePermission)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (bool, error)); ok {
		return rf(ctx, userId, storeId, codePermission)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) bool); ok {
		r0 = rf(ctx, userId, storeId, codePermission)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, userId, storeId, codePermission)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUseCase(t mockConstructorTestingTNewUserUseCase) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
