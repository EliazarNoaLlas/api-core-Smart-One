/*
 * File: merchants_usecase_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the test of the merchants UseCase
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

	merchantsDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
	mockMerchants "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain/mocks"
)

func TestGet_Merchant(t *testing.T) {
	t.Run("When get merchants successfully", func(t *testing.T) {
		merchantsRepository := &mockMerchants.MerchantRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		merchantsRepository.
			On("GetMerchants",
				mock.Anything,
				mock.Anything).
			Return([]merchantsDomain.Merchant{}, nil)
		merchantsRepository.
			On("GetTotalMerchants",
				mock.Anything,
				mock.Anything).
			Return(&total, nil)
		merchantsUCase := NewMerchantsUseCase(
			merchantsRepository,
			validationRepository,
			authRepository,
			60,
		)
		pagination := paramsDomain.NewPaginationParams(nil)
		merchants, _, err := merchantsUCase.GetMerchants(context.Background(), pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, merchants, []merchantsDomain.Merchant{})
	})

	t.Run("When get merchants error", func(t *testing.T) {
		merchantsRepository := &mockMerchants.MerchantRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		merchantsRepository.
			On("GetMerchants",
				mock.Anything,
				mock.Anything).
			Return(nil, errors.New("random error"))
		merchantsRepository.
			On("GetTotalMerchants",
				mock.Anything,
				mock.Anything).
			Return(&total, nil)
		merchantsUCase := NewMerchantsUseCase(
			merchantsRepository,
			validationRepository,
			authRepository,
			60,
		)
		pagination := paramsDomain.NewPaginationParams(nil)
		ok, _, err := merchantsUCase.GetMerchants(context.Background(), pagination)
		assert.Error(t, err)
		assert.EqualValues(t, ok, []merchantsDomain.Merchant(nil))
	})
}

func TestMerchant_CreateMerchant(t *testing.T) {
	t.Run("When create merchant successfully", func(t *testing.T) {
		merchantsRepository := &mockMerchants.MerchantRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		createMerchantBody := merchantsDomain.CreateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}

		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		merchantsRepository.
			On("CreateMerchant",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(&merchantId, nil)
		validationRepository.
			On("ValidateExistence",
				mock.Anything,
				mock.Anything,
				mock.Anything).Return(
			false, nil)
		merchantsUCase := NewMerchantsUseCase(
			merchantsRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := merchantsUCase.CreateMerchant(
			context.Background(), createMerchantBody,
		)
		assert.NoError(t, err)
	})

	t.Run("When creating a user fails and that merchant already exists", func(t *testing.T) {
		merchantsRepository := &mockMerchants.MerchantRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		createMerchantBody := merchantsDomain.CreateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		merchantsRepository.
			On("CreateMerchant",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, errors.New("random error"))
		validationRepository.
			On("ValidateExistence",
				mock.Anything,
				mock.Anything,
				mock.Anything).Return(
			true, nil)
		merchantsUCase := NewMerchantsUseCase(
			merchantsRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := merchantsUCase.CreateMerchant(
			context.Background(), createMerchantBody)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, merchantsDomain.ErrMerchantDocumentAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateMerchant")
	})

	t.Run("When create merchant error", func(t *testing.T) {
		merchantsRepository := &mockMerchants.MerchantRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		errCreate := errDomain.NewErr().SetFunction("CreateMerchant").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		createMerchantBody := merchantsDomain.CreateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		merchantsRepository.
			On("CreateMerchant",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, errCreate)
		validationRepository.
			On("ValidateExistence",
				mock.Anything,
				mock.Anything,
				mock.Anything).Return(
			false, nil)
		merchantsUCase := NewMerchantsUseCase(
			merchantsRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := merchantsUCase.CreateMerchant(
			context.Background(), createMerchantBody)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateMerchant")
	})
}

func TestMerchant_UpdateMerchant(t *testing.T) {
	t.Run("When update merchant successfully", func(t *testing.T) {
		merchantsRepository := &mockMerchants.MerchantRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		merchantsRepository.
			On("UpdateMerchant",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(nil)
		merchantsUCase := NewMerchantsUseCase(
			merchantsRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := merchantsUCase.UpdateMerchant(
			context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
			merchantsDomain.UpdateMerchantBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When update merchant error", func(t *testing.T) {
		merchantsRepository := &mockMerchants.MerchantRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).Return(true, nil)
		merchantsRepository.
			On("UpdateMerchant",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(errors.New("random error"))
		merchantsUCase := NewMerchantsUseCase(
			merchantsRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := merchantsUCase.UpdateMerchant(
			context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
			merchantsDomain.UpdateMerchantBody{},
		)
		assert.Error(t, err)
	})
}

func TestMerchant_DeleteMerchant(t *testing.T) {
	t.Run("When delete merchant by id successfully", func(t *testing.T) {
		merchantsRepository := &mockMerchants.MerchantRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).Return(true, nil)
		merchantsRepository.
			On("DeleteMerchant", mock.Anything, mock.Anything).
			Return(true, nil)
		merchantsUCase := NewMerchantsUseCase(
			merchantsRepository,
			validationRepository,
			authRepository,
			60,
		)
		res, err := merchantsUCase.DeleteMerchant(context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016")
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete merchant by id error", func(t *testing.T) {
		merchantsRepository := &mockMerchants.MerchantRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantsError := errors.New("random error")
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).Return(true, nil)
		merchantsRepository.
			On("DeleteMerchant", mock.Anything, mock.Anything).
			Return(false, merchantsError)
		merchantsUCase := NewMerchantsUseCase(
			merchantsRepository,
			validationRepository,
			authRepository,
			60,
		)
		res, err := merchantsUCase.DeleteMerchant(context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016")
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
