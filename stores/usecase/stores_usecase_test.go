/*
 * File: stores_usecase_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of stores.
 *
 * Last Modified: 2023-11-14
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

	storesDomain "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain"
	mockStores "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain/mocks"
)

func TestUseCaseStores_GetStores(t *testing.T) {
	t.Run("When get stores successfully", func(t *testing.T) {
		storesRepository := &mockStores.StoreRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		storesRepository.
			On("GetStores", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return([]storesDomain.Store{}, nil)
		storesRepository.
			On("GetTotalStores", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		storesUCase := NewStoresUseCase(storesRepository, validationRepository, authRepository, 60)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		pagination := paramsDomain.NewPaginationParams(nil)
		stores, _, err := storesUCase.GetStores(context.Background(), merchantId, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, stores, []storesDomain.Store{})
	})

	t.Run("When get stores error", func(t *testing.T) {
		storesRepository := &mockStores.StoreRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		storesRepository.
			On("GetStores", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		storesRepository.
			On("GetTotalStores", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		storesUCase := NewStoresUseCase(storesRepository, validationRepository, authRepository, 60)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		pagination := paramsDomain.NewPaginationParams(nil)
		ok, _, err := storesUCase.GetStores(context.Background(), merchantId, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, ok, []storesDomain.Store(nil))
	})
}

func TestUseCaseStores_CreateStore(t *testing.T) {
	t.Run("When create store successfully", func(t *testing.T) {
		storesRepository := &mockStores.StoreRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storeID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		storesRepository.
			On("CreateStore", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&storeID, nil)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		storesUCase := NewStoresUseCase(storesRepository, validationRepository, authRepository, 60)
		_, err := storesUCase.CreateStore(
			context.Background(),
			merchantId,
			storesDomain.CreateStoreBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When create store error", func(t *testing.T) {
		storesRepository := &mockStores.StoreRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storesRepository.
			On("CreateStore", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(true, nil)
		storesUCase := NewStoresUseCase(storesRepository, validationRepository, authRepository, 60)
		_, err := storesUCase.CreateStore(
			context.Background(),
			merchantId,
			storesDomain.CreateStoreBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, storesDomain.ErrStoreNameAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateStore")
	})

	t.Run("When create store error", func(t *testing.T) {
		storesRepository := &mockStores.StoreRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		errCreate := errDomain.NewErr().SetFunction("CreateStore").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storesRepository.
			On("CreateStore", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything, mock.Anything).
			Return(false, nil)
		storesUCase := NewStoresUseCase(storesRepository, validationRepository, authRepository, 60)
		_, err := storesUCase.CreateStore(
			context.Background(),
			merchantId,
			storesDomain.CreateStoreBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateStore")
	})
}

func TestUseCaseStores_UpdateStore(t *testing.T) {
	t.Run("When update store successfully", func(t *testing.T) {
		storesRepository := &mockStores.StoreRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		storesRepository.
			On("UpdateStore", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		storesUCase := NewStoresUseCase(storesRepository, validationRepository, authRepository, 60)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		err := storesUCase.UpdateStore(
			context.Background(),
			merchantId,
			storeId,
			storesDomain.CreateStoreBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When update store error", func(t *testing.T) {
		storesRepository := &mockStores.StoreRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		storesRepository.
			On("UpdateStore", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		storesUCase := NewStoresUseCase(storesRepository, validationRepository, authRepository, 60)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		err := storesUCase.UpdateStore(
			context.Background(),
			merchantId,
			storeId,
			storesDomain.CreateStoreBody{},
		)
		assert.Error(t, err)
	})
}

func TestUseCaseStores_DeleteStore(t *testing.T) {
	t.Run("When delete store by id successfully", func(t *testing.T) {
		storesRepository := &mockStores.StoreRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		storesRepository.
			On("DeleteStore", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		storesUCase := NewStoresUseCase(storesRepository, validationRepository, authRepository, 60)
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := storesUCase.DeleteStore(context.Background(), merchantId, storeId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete store by id error", func(t *testing.T) {
		storesRepository := &mockStores.StoreRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		storesError := errors.New("random error")
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		storesRepository.
			On("DeleteStore", mock.Anything, mock.Anything, mock.Anything).
			Return(false, storesError)
		storesUCase := NewStoresUseCase(storesRepository, validationRepository, authRepository, 60)
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := storesUCase.DeleteStore(context.Background(), merchantId, storeId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
