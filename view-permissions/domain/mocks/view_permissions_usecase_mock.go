// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package view_permissions

import (
	context "context"
	domain "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"

	mock "github.com/stretchr/testify/mock"
)

// ViewPermissionsUseCase is an autogenerated mock type for the ViewPermissionsUseCase type
type ViewPermissionsUseCase struct {
	mock.Mock
}

// CreateViewPermission provides a mock function with given fields: ctx, viewId, userId, body
func (_m *ViewPermissionsUseCase) CreateViewPermission(ctx context.Context, viewId string, userId string, body domain.CreateViewPermissionBody) (*string, error) {
	ret := _m.Called(ctx, viewId, userId, body)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, domain.CreateViewPermissionBody) (*string, error)); ok {
		return rf(ctx, viewId, userId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, domain.CreateViewPermissionBody) *string); ok {
		r0 = rf(ctx, viewId, userId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, domain.CreateViewPermissionBody) error); ok {
		r1 = rf(ctx, viewId, userId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteViewPermission provides a mock function with given fields: ctx, viewId, viewPermissionId
func (_m *ViewPermissionsUseCase) DeleteViewPermission(ctx context.Context, viewId string, viewPermissionId string) (bool, error) {
	ret := _m.Called(ctx, viewId, viewPermissionId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, viewId, viewPermissionId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, viewId, viewPermissionId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, viewId, viewPermissionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetViewPermissions provides a mock function with given fields: ctx, viewId
func (_m *ViewPermissionsUseCase) GetViewPermissions(ctx context.Context, viewId string) ([]domain.ViewPermission, error) {
	ret := _m.Called(ctx, viewId)

	var r0 []domain.ViewPermission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.ViewPermission, error)); ok {
		return rf(ctx, viewId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.ViewPermission); ok {
		r0 = rf(ctx, viewId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ViewPermission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, viewId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateViewPermission provides a mock function with given fields: ctx, viewId, viewPermissionId, body
func (_m *ViewPermissionsUseCase) UpdateViewPermission(ctx context.Context, viewId string, viewPermissionId string, body domain.UpdateViewPermissionBody) error {
	ret := _m.Called(ctx, viewId, viewPermissionId, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, domain.UpdateViewPermissionBody) error); ok {
		r0 = rf(ctx, viewId, viewPermissionId, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewViewPermissionsUseCase creates a new instance of ViewPermissionsUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewViewPermissionsUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ViewPermissionsUseCase {
	mock := &ViewPermissionsUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
