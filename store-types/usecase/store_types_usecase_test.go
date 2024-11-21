/*
 * File: store_types_usecase_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the store types use case.
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

	mockStoreTypes "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain/mocks"
	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	mockValidation "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain/mocks"

	storeTypeDomain "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain"
)

func TestUseCaseGetStoreTypes_GetStoreTypes(t *testing.T) {
	t.Run("When get store types  successfully", func(t *testing.T) {
		storeTypesRepository := &mockStoreTypes.StoreTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10

		storeTypesRepository.
			On("GetStoreTypes", mock.Anything, mock.Anything).
			Return([]storeTypeDomain.StoreType{}, nil)
		storeTypesRepository.
			On("GetTotalStoreTypes", mock.Anything, mock.Anything).
			Return(&total, nil)
		storeTypeUCase := NewStoreTypesUseCase(
			storeTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		pagination := paramsDomain.NewPaginationParams(nil)
		storeType, _, err := storeTypeUCase.GetStoreTypes(context.Background(), pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, storeType, []storeTypeDomain.StoreType{})
	})

	t.Run("When get store types  error", func(t *testing.T) {
		storeTypesRepository := &mockStoreTypes.StoreTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		storeTypesRepository.
			On("GetStoreTypes", mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		storeTypesRepository.
			On("GetTotalStoreTypes", mock.Anything, mock.Anything).
			Return(&total, nil)
		storeTypesUCase := NewStoreTypesUseCase(
			storeTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		pagination := paramsDomain.NewPaginationParams(nil)
		ok, _, err := storeTypesUCase.GetStoreTypes(context.Background(), pagination)
		assert.Error(t, err)
		assert.EqualValues(t, ok, []storeTypeDomain.StoreType(nil))
	})
}

func TestUseCaseStoreType_CreateStoreType(t *testing.T) {
	t.Run("When create store type  successfully", func(t *testing.T) {
		storeTypesRepository := &mockStoreTypes.StoreTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		storeTypeID := "73900000-7e93-11ee-89fd-0242a500000"
		storeTypesRepository.
			On("CreateStoreType", mock.Anything, mock.Anything, mock.Anything).
			Return(&storeTypeID, nil)
		validationRepository.On("ValidateExistence", mock.Anything, mock.Anything).Return(
			false, nil)
		storeTypesUCase := NewStoreTypesUseCase(
			storeTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := storeTypesUCase.CreateStoreType(
			context.Background(),
			storeTypeDomain.CreateStoreTypeBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When create store type  error", func(t *testing.T) {
		storeTypesRepository := &mockStoreTypes.StoreTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		storeTypesRepository.
			On("CreateStoreType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		validationRepository.On("ValidateExistence", mock.Anything, mock.Anything).Return(
			true, nil)
		storeTypesUCase := NewStoreTypesUseCase(
			storeTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := storeTypesUCase.CreateStoreType(
			context.Background(),
			storeTypeDomain.CreateStoreTypeBody{},
		)
		assert.Error(t, err)
	})
}

func TestUseCaseStoreType_UpdateStoreType(t *testing.T) {
	t.Run("When update store type  successfully", func(t *testing.T) {
		storeTypesRepository := &mockStoreTypes.StoreTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		storeTypesRepository.
			On("UpdateStoreType",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(nil)
		storeTypesUCase := NewStoreTypesUseCase(
			storeTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := storeTypesUCase.UpdateStoreType(
			context.Background(),
			storeTypeDomain.UpdateStoreTypeBody{},
			"73900000-7e93-11ee-89fd-0242a500000",
		)
		assert.NoError(t, err)
	})

	t.Run("When update store type error", func(t *testing.T) {
		storeTypesRepository := &mockStoreTypes.StoreTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		storeTypesRepository.
			On("UpdateStoreType",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(errors.New("random error"))
		storeTypesUCase := NewStoreTypesUseCase(
			storeTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := storeTypesUCase.UpdateStoreType(
			context.Background(),
			storeTypeDomain.UpdateStoreTypeBody{},
			"73900000-7e93-11ee-89fd-0242a500000",
		)
		assert.Error(t, err)
	})
}

func TestUseCaseStoreType_DeleteStoreType(t *testing.T) {
	t.Run("When delete store type by id successfully", func(t *testing.T) {
		storeTypesRepository := &mockStoreTypes.StoreTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		storeTypesRepository.
			On("DeleteStoreType", mock.Anything, mock.Anything).
			Return(true, nil)
		storeTypesUCase := NewStoreTypesUseCase(
			storeTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		res, err := storeTypesUCase.DeleteStoreType(context.Background(),
			"73900000-7e93-11ee-89fd-0242a500000")
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete store type by id error", func(t *testing.T) {
		storeTypesRepository := &mockStoreTypes.StoreTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		storeTypesError := errors.New("random error")
		storeTypesRepository.
			On("DeleteStoreType", mock.Anything, mock.Anything).
			Return(false, storeTypesError)
		storeTypesUCase := NewStoreTypesUseCase(
			storeTypesRepository,
			validationRepository,
			authRepository,
			60)
		res, err := storeTypesUCase.DeleteStoreType(context.Background(),
			"73900000-7e93-11ee-89fd-0242a500000")
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
