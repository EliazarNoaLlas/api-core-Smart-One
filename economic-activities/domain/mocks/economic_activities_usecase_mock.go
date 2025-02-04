// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package economic_activities

import (
	context "context"
	domain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"

	mock "github.com/stretchr/testify/mock"

	paramsdomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

// EconomicActivityUseCase is an autogenerated mock type for the EconomicActivityUseCase type
type EconomicActivityUseCase struct {
	mock.Mock
}

// GetEconomicActivities provides a mock function with given fields: ctx, searchParams, pagination
func (_m *EconomicActivityUseCase) GetEconomicActivities(ctx context.Context, searchParams domain.GetEconomicActivitiesParams, pagination paramsdomain.PaginationParams) ([]domain.EconomicActivity, *paramsdomain.PaginationResults, error) {
	ret := _m.Called(ctx, searchParams, pagination)

	var r0 []domain.EconomicActivity
	var r1 *paramsdomain.PaginationResults
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetEconomicActivitiesParams, paramsdomain.PaginationParams) ([]domain.EconomicActivity, *paramsdomain.PaginationResults, error)); ok {
		return rf(ctx, searchParams, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetEconomicActivitiesParams, paramsdomain.PaginationParams) []domain.EconomicActivity); ok {
		r0 = rf(ctx, searchParams, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.EconomicActivity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.GetEconomicActivitiesParams, paramsdomain.PaginationParams) *paramsdomain.PaginationResults); ok {
		r1 = rf(ctx, searchParams, pagination)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*paramsdomain.PaginationResults)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.GetEconomicActivitiesParams, paramsdomain.PaginationParams) error); ok {
		r2 = rf(ctx, searchParams, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewEconomicActivityUseCase creates a new instance of EconomicActivityUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEconomicActivityUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *EconomicActivityUseCase {
	mock := &EconomicActivityUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
