// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package policyPermissions

import (
	context "context"
	domain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"

	mock "github.com/stretchr/testify/mock"

	paramsdomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

// PolicyPermissionUseCase is an autogenerated mock type for the PolicyPermissionUseCase type
type PolicyPermissionUseCase struct {
	mock.Mock
}

// CreatePolicyPermission provides a mock function with given fields: ctx, policyId, body
func (_m *PolicyPermissionUseCase) CreatePolicyPermission(ctx context.Context, policyId string, body domain.CreatePolicyPermissionBody) (*string, error) {
	ret := _m.Called(ctx, policyId, body)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.CreatePolicyPermissionBody) (*string, error)); ok {
		return rf(ctx, policyId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.CreatePolicyPermissionBody) *string); ok {
		r0 = rf(ctx, policyId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.CreatePolicyPermissionBody) error); ok {
		r1 = rf(ctx, policyId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePolicyPermissions provides a mock function with given fields: ctx, policyId, body
func (_m *PolicyPermissionUseCase) CreatePolicyPermissions(ctx context.Context, policyId string, body []domain.CreatePolicyPermissionBody) ([]string, error) {
	ret := _m.Called(ctx, policyId, body)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []domain.CreatePolicyPermissionBody) ([]string, error)); ok {
		return rf(ctx, policyId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []domain.CreatePolicyPermissionBody) []string); ok {
		r0 = rf(ctx, policyId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []domain.CreatePolicyPermissionBody) error); ok {
		r1 = rf(ctx, policyId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePolicyPermission provides a mock function with given fields: ctx, policyId, policyPermissionId
func (_m *PolicyPermissionUseCase) DeletePolicyPermission(ctx context.Context, policyId string, policyPermissionId string) (bool, error) {
	ret := _m.Called(ctx, policyId, policyPermissionId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, policyId, policyPermissionId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, policyId, policyPermissionId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, policyId, policyPermissionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePolicyPermissions provides a mock function with given fields: ctx, policyId, policyPermissionIds
func (_m *PolicyPermissionUseCase) DeletePolicyPermissions(ctx context.Context, policyId string, policyPermissionIds []string) error {
	ret := _m.Called(ctx, policyId, policyPermissionIds)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, policyId, policyPermissionIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPolicyPermissionsByPolicy provides a mock function with given fields: ctx, policyId, pagination
func (_m *PolicyPermissionUseCase) GetPolicyPermissionsByPolicy(ctx context.Context, policyId string, pagination paramsdomain.PaginationParams) ([]domain.PolicyPermission, *paramsdomain.PaginationResults, error) {
	ret := _m.Called(ctx, policyId, pagination)

	var r0 []domain.PolicyPermission
	var r1 *paramsdomain.PaginationResults
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, paramsdomain.PaginationParams) ([]domain.PolicyPermission, *paramsdomain.PaginationResults, error)); ok {
		return rf(ctx, policyId, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, paramsdomain.PaginationParams) []domain.PolicyPermission); ok {
		r0 = rf(ctx, policyId, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PolicyPermission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, paramsdomain.PaginationParams) *paramsdomain.PaginationResults); ok {
		r1 = rf(ctx, policyId, pagination)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*paramsdomain.PaginationResults)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, paramsdomain.PaginationParams) error); ok {
		r2 = rf(ctx, policyId, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdatePolicyPermission provides a mock function with given fields: ctx, policyId, policyPermissionId, body
func (_m *PolicyPermissionUseCase) UpdatePolicyPermission(ctx context.Context, policyId string, policyPermissionId string, body domain.CreatePolicyPermissionBody) error {
	ret := _m.Called(ctx, policyId, policyPermissionId, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, domain.CreatePolicyPermissionBody) error); ok {
		r0 = rf(ctx, policyId, policyPermissionId, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPolicyPermissionUseCase creates a new instance of PolicyPermissionUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPolicyPermissionUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *PolicyPermissionUseCase {
	mock := &PolicyPermissionUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
