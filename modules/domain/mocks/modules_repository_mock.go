// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package modules

import (
	context "context"
	domain "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"

	mock "github.com/stretchr/testify/mock"

	paramsdomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

// ModuleRepository is an autogenerated mock type for the ModuleRepository type
type ModuleRepository struct {
	mock.Mock
}

// CreateModule provides a mock function with given fields: ctx, moduleId, body
func (_m *ModuleRepository) CreateModule(ctx context.Context, moduleId string, body domain.CreateModuleBody) (*string, error) {
	ret := _m.Called(ctx, moduleId, body)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.CreateModuleBody) (*string, error)); ok {
		return rf(ctx, moduleId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.CreateModuleBody) *string); ok {
		r0 = rf(ctx, moduleId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.CreateModuleBody) error); ok {
		r1 = rf(ctx, moduleId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteModule provides a mock function with given fields: ctx, moduleId
func (_m *ModuleRepository) DeleteModule(ctx context.Context, moduleId string) (bool, error) {
	ret := _m.Called(ctx, moduleId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, moduleId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, moduleId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, moduleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModules provides a mock function with given fields: ctx, searchParams, pagination
func (_m *ModuleRepository) GetModules(ctx context.Context, searchParams domain.GetModulesParams, pagination paramsdomain.PaginationParams) ([]domain.Module, error) {
	ret := _m.Called(ctx, searchParams, pagination)

	var r0 []domain.Module
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetModulesParams, paramsdomain.PaginationParams) ([]domain.Module, error)); ok {
		return rf(ctx, searchParams, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetModulesParams, paramsdomain.PaginationParams) []domain.Module); ok {
		r0 = rf(ctx, searchParams, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Module)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.GetModulesParams, paramsdomain.PaginationParams) error); ok {
		r1 = rf(ctx, searchParams, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalModules provides a mock function with given fields: ctx, searchParams
func (_m *ModuleRepository) GetTotalModules(ctx context.Context, searchParams domain.GetModulesParams) (*int, error) {
	ret := _m.Called(ctx, searchParams)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetModulesParams) (*int, error)); ok {
		return rf(ctx, searchParams)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetModulesParams) *int); ok {
		r0 = rf(ctx, searchParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.GetModulesParams) error); ok {
		r1 = rf(ctx, searchParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateModule provides a mock function with given fields: ctx, moduleId, body
func (_m *ModuleRepository) UpdateModule(ctx context.Context, moduleId string, body domain.UpdateModuleBody) error {
	ret := _m.Called(ctx, moduleId, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.UpdateModuleBody) error); ok {
		r0 = rf(ctx, moduleId, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewModuleRepository creates a new instance of ModuleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewModuleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ModuleRepository {
	mock := &ModuleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
