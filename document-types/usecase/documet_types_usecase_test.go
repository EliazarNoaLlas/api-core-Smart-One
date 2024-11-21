/*
 * File: documentTypes_usecase_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of documentTypes.
 *
 * Last Modified: 2023-11-10
 */

package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	mockValidation "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain/mocks"

	documentTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain"
	mockDocumentTypes "gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain/mocks"
)

func TestUseCaseDocumentTypes_GetDocumentType(t *testing.T) {
	t.Run("When get document types successfully", func(t *testing.T) {
		documentTypesRepository := &mockDocumentTypes.DocumentTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		documentTypesRepository.
			On("GetDocumentTypes", mock.Anything, mock.Anything, mock.Anything).
			Return([]documentTypesDomain.DocumentType{}, nil)
		documentTypesRepository.
			On("GetTotalDocumentTypes", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		documentTypesUCase := NewDocumentTypesUseCase(
			documentTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		searchParams := documentTypesDomain.GetDocumentTypeParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		documentTypes, _, err := documentTypesUCase.GetDocumentTypes(context.Background(), searchParams, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, documentTypes, []documentTypesDomain.DocumentType{})
	})

	t.Run("When get document types error", func(t *testing.T) {
		documentTypesRepository := &mockDocumentTypes.DocumentTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		searchParams := documentTypesDomain.GetDocumentTypeParams{}
		documentTypesRepository.
			On("GetDocumentTypes", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		documentTypesRepository.
			On("GetTotalDocumentTypes", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, errors.New("random error"))
		documentTypesUCase := NewDocumentTypesUseCase(
			documentTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		pagination := paramsDomain.NewPaginationParams(nil)
		ok, _, err := documentTypesUCase.GetDocumentTypes(context.Background(), searchParams, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, ok, []documentTypesDomain.DocumentType(nil))
	})
}

func TestUseCaseDocumentTypes_CreateDocumentType(t *testing.T) {
	t.Run("When create document type successfully", func(t *testing.T) {
		documentTypesRepository := &mockDocumentTypes.DocumentTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		documentTypesRepository.
			On("CreateDocumentType", mock.Anything, mock.Anything, mock.Anything).
			Return(&documentTypeId, nil)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything, mock.Anything).
			Return(false, nil)
		documentTypesUCase := NewDocumentTypesUseCase(
			documentTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := documentTypesUCase.CreateDocumentType(
			context.Background(),
			documentTypeId,
			documentTypesDomain.CreateDocumentTypeBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When create document type error", func(t *testing.T) {
		documentTypesRepository := &mockDocumentTypes.DocumentTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		documentTypesRepository.
			On("CreateDocumentType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		documentTypesUCase := NewDocumentTypesUseCase(
			documentTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := documentTypesUCase.CreateDocumentType(
			context.Background(),
			documentTypeId,
			documentTypesDomain.CreateDocumentTypeBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, documentTypesDomain.ErrDocumentTypeDescriptionAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateDocumentType")
	})

	t.Run("When create document type error", func(t *testing.T) {
		documentTypesRepository := &mockDocumentTypes.DocumentTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		errCreate := errDomain.NewErr().SetFunction("CreateDocumentType").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		documentTypesRepository.
			On("CreateDocumentType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything, mock.Anything).
			Return(false, nil)
		documentTypesUCase := NewDocumentTypesUseCase(
			documentTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		_, err := documentTypesUCase.CreateDocumentType(
			context.Background(),
			documentTypeId,
			documentTypesDomain.CreateDocumentTypeBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateDocumentType")
	})
}

func TestUseCaseDocumentTypes_UpdateDocumentType(t *testing.T) {
	t.Run("When update document type successfully", func(t *testing.T) {
		documentTypesRepository := &mockDocumentTypes.DocumentTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		documentTypesRepository.
			On("UpdateDocumentType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		documentTypesUCase := NewDocumentTypesUseCase(
			documentTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := documentTypesUCase.UpdateDocumentType(
			context.Background(),
			documentTypeId,
			documentTypesDomain.UpdateDocumentTypeBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When update document type error", func(t *testing.T) {
		documentTypesRepository := &mockDocumentTypes.DocumentTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		documentTypesRepository.
			On("UpdateDocumentType",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(errors.New("random error"))
		documentTypesUCase := NewDocumentTypesUseCase(
			documentTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := documentTypesUCase.UpdateDocumentType(
			context.Background(),
			documentTypeId,
			documentTypesDomain.UpdateDocumentTypeBody{},
		)
		assert.Error(t, err)
	})
}

func TestUseCaseDocumentTypes_DeleteDocumentType(t *testing.T) {
	t.Run("When delete document type by id successfully", func(t *testing.T) {
		documentTypesRepository := &mockDocumentTypes.DocumentTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		documentTypesRepository.
			On("DeleteDocumentType", mock.Anything, mock.Anything).
			Return(true, nil)
		documentTypesUCase := NewDocumentTypesUseCase(
			documentTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := documentTypesUCase.DeleteDocumentType(context.Background(), documentTypeId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete document type by id error", func(t *testing.T) {
		documentTypesRepository := &mockDocumentTypes.DocumentTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		documentTypesError := errors.New("random error")
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		documentTypesRepository.
			On("DeleteDocumentType", mock.Anything, mock.Anything).
			Return(false, documentTypesError)
		documentTypesUCase := NewDocumentTypesUseCase(
			documentTypesRepository,
			validationRepository,
			authRepository, 60)
		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := documentTypesUCase.DeleteDocumentType(context.Background(), documentTypeId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
