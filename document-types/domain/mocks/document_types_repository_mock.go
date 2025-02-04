// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package document_types

import (
	context "context"
	domain "gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain"

	mock "github.com/stretchr/testify/mock"

	paramsdomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

// DocumentTypeRepository is an autogenerated mock type for the DocumentTypeRepository type
type DocumentTypeRepository struct {
	mock.Mock
}

// CreateDocumentType provides a mock function with given fields: ctx, documentTypeId, body
func (_m *DocumentTypeRepository) CreateDocumentType(ctx context.Context, documentTypeId string, body domain.CreateDocumentTypeBody) (*string, error) {
	ret := _m.Called(ctx, documentTypeId, body)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.CreateDocumentTypeBody) (*string, error)); ok {
		return rf(ctx, documentTypeId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.CreateDocumentTypeBody) *string); ok {
		r0 = rf(ctx, documentTypeId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.CreateDocumentTypeBody) error); ok {
		r1 = rf(ctx, documentTypeId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDocumentType provides a mock function with given fields: ctx, documentTypeId
func (_m *DocumentTypeRepository) DeleteDocumentType(ctx context.Context, documentTypeId string) (bool, error) {
	ret := _m.Called(ctx, documentTypeId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, documentTypeId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, documentTypeId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, documentTypeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDocumentTypes provides a mock function with given fields: ctx, searchParams, pagination
func (_m *DocumentTypeRepository) GetDocumentTypes(ctx context.Context, searchParams domain.GetDocumentTypeParams, pagination paramsdomain.PaginationParams) ([]domain.DocumentType, error) {
	ret := _m.Called(ctx, searchParams, pagination)

	var r0 []domain.DocumentType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetDocumentTypeParams, paramsdomain.PaginationParams) ([]domain.DocumentType, error)); ok {
		return rf(ctx, searchParams, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetDocumentTypeParams, paramsdomain.PaginationParams) []domain.DocumentType); ok {
		r0 = rf(ctx, searchParams, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.DocumentType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.GetDocumentTypeParams, paramsdomain.PaginationParams) error); ok {
		r1 = rf(ctx, searchParams, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalDocumentTypes provides a mock function with given fields: ctx, searchParams, pagination
func (_m *DocumentTypeRepository) GetTotalDocumentTypes(ctx context.Context, searchParams domain.GetDocumentTypeParams, pagination paramsdomain.PaginationParams) (*int, error) {
	ret := _m.Called(ctx, searchParams, pagination)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetDocumentTypeParams, paramsdomain.PaginationParams) (*int, error)); ok {
		return rf(ctx, searchParams, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetDocumentTypeParams, paramsdomain.PaginationParams) *int); ok {
		r0 = rf(ctx, searchParams, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.GetDocumentTypeParams, paramsdomain.PaginationParams) error); ok {
		r1 = rf(ctx, searchParams, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDocumentType provides a mock function with given fields: ctx, documentTypeId, body
func (_m *DocumentTypeRepository) UpdateDocumentType(ctx context.Context, documentTypeId string, body domain.UpdateDocumentTypeBody) error {
	ret := _m.Called(ctx, documentTypeId, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.UpdateDocumentTypeBody) error); ok {
		r0 = rf(ctx, documentTypeId, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDocumentTypeRepository creates a new instance of DocumentTypeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDocumentTypeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *DocumentTypeRepository {
	mock := &DocumentTypeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
