/*
 * File: receipt_types_usecase_test.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the test for the receiptTypes use case.
 *
 * Last Modified: 2024-03-06
 */

package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	mockValidation "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain/mocks"

	receiptTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
	mockReceiptTypes "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain/mocks"
)

func TestReceiptTypesUseCase_GetReceiptTypes(t *testing.T) {
	t.Run("When get receipt types  successfully", func(t *testing.T) {
		receiptTypesRepository := &mockReceiptTypes.ReceiptTypesRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		now := time.Now().UTC()
		receiptTypes := []receiptTypesDomain.ReceiptType{
			{
				Id:          "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
				Description: "Recibo por Honorarios",
				SunatCode:   "02",
				Enable:      true,
				CreatedBy:   "91fb86bd-da46-414b-97a1-fcdaa8cd35d1",
				CreatedAt:   &now,
			},
			{
				Id:          "2e8bfdbb-1a58-4b45-9a2c-8ac54a5db723",
				Description: "Recibo por Arrendamiento",
				SunatCode:   "03",
				Enable:      false,
				CreatedBy:   "c3f92a0d-ef58-4e15-a71b-6f0a8b9d147d",
				CreatedAt:   &now,
			},
		}

		receiptTypesRepository.
			On("GetReceiptTypes", mock.Anything, mock.Anything).
			Return(receiptTypes, nil)

		ReceiptTypesUCase := NewReceiptTypesUseCase(
			receiptTypesRepository,
			validationRepository,
			authRepository,
			60,
		)

		res, err := ReceiptTypesUCase.GetReceiptTypes(context.Background())

		assert.NoError(t, err)
		assert.EqualValues(t, res, receiptTypes)
	})

	t.Run("When get receipt types error", func(t *testing.T) {
		receiptTypesRepository := &mockReceiptTypes.ReceiptTypesRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		receiptTypesRepository.
			On("GetReceiptTypes", mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		ReceiptTypesUCase := NewReceiptTypesUseCase(
			receiptTypesRepository,
			validationRepository,
			authRepository,
			60)
		res, err := ReceiptTypesUCase.GetReceiptTypes(context.Background())
		assert.Error(t, err)
		assert.EqualValues(t, res, []receiptTypesDomain.ReceiptType(nil))
	})
}

func TestReceiptTypesUseCase_CreateReceiptType(t *testing.T) {
	t.Run("When create receipt type successfully", func(t *testing.T) {
		receiptTypesRepository := &mockReceiptTypes.ReceiptTypesRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		userId := "91fb86bd-da46-414b-97a1-fcdaa8cd35d1"
		body := receiptTypesDomain.CreateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		receiptTypesRepository.
			On("CreateReceiptType", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		ReceiptTypesUCase := NewReceiptTypesUseCase(
			receiptTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := ReceiptTypesUCase.CreateReceiptType(
			context.Background(),
			userId,
			body,
		)
		assert.NoError(t, err)
	})

	t.Run("When receipt type not found", func(t *testing.T) {
		receiptTypesRepository := &mockReceiptTypes.ReceiptTypesRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userId := "91fb86bd-da46-414b-97a1-fcdaa8cd35d1"
		body := receiptTypesDomain.CreateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		receiptTypesRepository.
			On("CreateReceiptType", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))

		ReceiptTypesUCase := NewReceiptTypesUseCase(
			receiptTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := ReceiptTypesUCase.CreateReceiptType(
			context.Background(),
			userId,
			body,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, receiptTypesDomain.ErrSunatCodeAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateReceiptType")
	})
}

func TestReceiptTypesUseCase_UpdateReceiptType(t *testing.T) {
	t.Run("When update receipt type successfully", func(t *testing.T) {
		receiptTypesRepository := &mockReceiptTypes.ReceiptTypesRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		receiptTypeId := "73900000-7e93-11ee-89fd-0242a500000"
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		receiptTypesRepository.
			On("UpdateReceiptType", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		ReceiptTypesUCase := NewReceiptTypesUseCase(
			receiptTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := ReceiptTypesUCase.UpdateReceiptType(
			context.Background(),
			receiptTypeId,
			receiptTypesDomain.UpdateReceiptTypeBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When update receipt type error", func(t *testing.T) {
		receiptTypesRepository := &mockReceiptTypes.ReceiptTypesRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		receiptTypeId := "73900000-7e93-11ee-89fd-0242a500000"
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		receiptTypesRepository.
			On("UpdateReceiptType", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		ReceiptTypesUCase := NewReceiptTypesUseCase(
			receiptTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := ReceiptTypesUCase.UpdateReceiptType(
			context.Background(),
			receiptTypeId,
			receiptTypesDomain.UpdateReceiptTypeBody{},
		)
		assert.Error(t, err)
	})
}

func TestReceiptTypesUseCase_DeleteReceiptType(t *testing.T) {
	t.Run("When delete receipt type by id successfully", func(t *testing.T) {
		receiptTypesRepository := &mockReceiptTypes.ReceiptTypesRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		receiptTypeId := "73900000-7e93-11ee-89fd-0242a500000"
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		receiptTypesRepository.
			On("DeleteReceiptType", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		ReceiptTypesUCase := NewReceiptTypesUseCase(
			receiptTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		res, err := ReceiptTypesUCase.DeleteReceiptType(
			context.Background(),
			receiptTypeId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete classifications receipt type by id error", func(t *testing.T) {
		receiptTypesRepository := &mockReceiptTypes.ReceiptTypesRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		receiptTypeId := "73900000-7e93-11ee-89fd-0242a500000"
		ReceiptTypesError := errors.New("random error")
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		receiptTypesRepository.
			On("DeleteReceiptType", mock.Anything, mock.Anything, mock.Anything).
			Return(false, ReceiptTypesError)
		ReceiptTypesUCase := NewReceiptTypesUseCase(
			receiptTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		res, err := ReceiptTypesUCase.DeleteReceiptType(
			context.Background(),
			receiptTypeId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
