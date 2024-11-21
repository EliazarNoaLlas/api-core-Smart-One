/*
 * File: merchant_economic_activities_usecase_test.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the use case test of merchant economic activities.
 *
 * Last Modified: 2023-12-05
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

	"gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
	mockMerchantEconomicActivities "gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain/mocks"
)

func TestMerchantEconomicActivities_GetMerchantEconomicActivities(t *testing.T) {
	t.Run("When merchant economic activities are successfully listed", func(t *testing.T) {
		merchantEconomicActivitiesRepository := &mockMerchantEconomicActivities.MerchantEconomicActivityRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		merchantEconomicActivitiesRepository.
			On("GetMerchantEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return([]domain.MerchantEconomicActivity{}, nil)
		merchantEconomicActivitiesRepository.
			On("GetTotalEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		merchantEconomicActivitiesUCase := NewMerchantEconomicActivitiesUseCase(merchantEconomicActivitiesRepository, validationRepository, authRepository, 60)
		pagination := paramsDomain.NewPaginationParams(nil)

		merchantId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		users, _, err := merchantEconomicActivitiesUCase.GetMerchantEconomicActivities(context.Background(),
			merchantId, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, users, []domain.MerchantEconomicActivity{})
	})

	t.Run("When merchant economic activities are successfully listed", func(t *testing.T) {
		merchantEconomicActivitiesRepository := &mockMerchantEconomicActivities.MerchantEconomicActivityRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantEconomicActivitiesRepository.
			On("GetMerchantEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return([]domain.MerchantEconomicActivity{}, errors.New("random error"))
		merchantEconomicActivitiesRepository.
			On("GetTotalEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		merchantEconomicActivitiesUCase := NewMerchantEconomicActivitiesUseCase(merchantEconomicActivitiesRepository, validationRepository, authRepository, 60)
		pagination := paramsDomain.NewPaginationParams(nil)

		merchantId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		activities, _, err := merchantEconomicActivitiesUCase.GetMerchantEconomicActivities(context.Background(),
			merchantId, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, activities, []domain.MerchantEconomicActivity(nil))
	})
}

func TestMerchantEconomicActivities_CreateEconomicActivities(t *testing.T) {
	t.Run("When attempting to create a merchant economic activities, the operation is successful.", func(t *testing.T) {
		merchantEconomicActivitiesRepository := &mockMerchantEconomicActivities.MerchantEconomicActivityRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantEconomicActivityId := "cf6e4017-f918-4ef0-b641-236d89901a5c"
		merchantEconomicActivitiesRepository.
			On("CreateEconomicActivity",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(&merchantEconomicActivityId, nil)
		validationRepository.
			On("ValidateExistence",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(false, nil)
		merchantEconomicActivitiesUCase := NewMerchantEconomicActivitiesUseCase(merchantEconomicActivitiesRepository,
			validationRepository, authRepository, 60)
		_, err := merchantEconomicActivitiesUCase.CreateEconomicActivity(
			context.Background(),
			merchantEconomicActivityId,
			domain.CreateMerchantEconomicActivityBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("Encountered an error during merchant economic activities creation.", func(t *testing.T) {
		merchantEconomicActivitiesRepository := &mockMerchantEconomicActivities.MerchantEconomicActivityRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantEconomicActivityBody := domain.CreateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		merchantEconomicActivityId := "cf6e4017-f918-4ef0-b641-236d89901a5c"
		merchantEconomicActivitiesRepository.
			On(merchantEconomicActivityId, "CreateEconomicActivity",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, errors.New("random error"))
		validationRepository.
			On("ValidateExistence",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		merchantEconomicActivitiesUCase := NewMerchantEconomicActivitiesUseCase(merchantEconomicActivitiesRepository,
			validationRepository, authRepository, 60)
		_, err := merchantEconomicActivitiesUCase.CreateEconomicActivity(
			context.Background(),
			merchantEconomicActivityId,
			merchantEconomicActivityBody,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, domain.ErrMerchantEconomicActivityIdAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateEconomicActivities")
	})

	t.Run("Encountered an error during merchant economic activities creation.", func(t *testing.T) {
		merchantEconomicActivitiesRepository := &mockMerchantEconomicActivities.MerchantEconomicActivityRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		errCreate := errDomain.NewErr().SetFunction("CreateEconomicActivity").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		merchantEconomicActivityUCaseBody := domain.CreateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		merchantEconomicActivityId := "cf6e4017-f918-4ef0-b641-236d89901a5c"
		merchantEconomicActivitiesRepository.
			On("CreateEconomicActivity",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, errCreate)
		validationRepository.
			On("ValidateExistence",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(false, nil)
		merchantEconomicActivitiesUCase := NewMerchantEconomicActivitiesUseCase(merchantEconomicActivitiesRepository,
			validationRepository, authRepository, 60)
		_, err := merchantEconomicActivitiesUCase.CreateEconomicActivity(
			context.Background(),
			merchantEconomicActivityId,
			merchantEconomicActivityUCaseBody,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateEconomicActivity")
	})
}

func TestMerchantEconomicActivities_UpdateEconomicActivity(t *testing.T) {
	t.Run("When attempting to update a merchant economic activities, the operation is successful.", func(t *testing.T) {
		merchantEconomicActivitiesRepository := &mockMerchantEconomicActivities.MerchantEconomicActivityRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantEconomicActivityBody := domain.UpdateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		merchantEconomicActivityId := "cf6e4017-f918-4ef0-b641-236d89901a5c"
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		merchantEconomicActivitiesRepository.
			On("UpdateEconomicActivity", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		merchantEconomicActivitiesUCase := NewMerchantEconomicActivitiesUseCase(merchantEconomicActivitiesRepository, validationRepository, authRepository, 60)
		err := merchantEconomicActivitiesUCase.UpdateEconomicActivity(
			context.Background(),
			merchantEconomicActivityId,
			merchantEconomicActivityBody,
		)
		assert.NoError(t, err)
	})

	t.Run("Encountered an error during merchant economic activities update.", func(t *testing.T) {
		merchantEconomicActivitiesRepository := &mockMerchantEconomicActivities.MerchantEconomicActivityRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantEconomicActivityId := "cf6e4017-f918-4ef0-b641-236d89901a5c"
		merchantEconomicActivityBody := domain.UpdateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything, mock.Anything).
			Return(false, nil)
		merchantEconomicActivitiesRepository.
			On("UpdateEconomicActivity", mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		merchantEconomicActivitiesUCase := NewMerchantEconomicActivitiesUseCase(merchantEconomicActivitiesRepository, validationRepository, authRepository, 60)
		err := merchantEconomicActivitiesUCase.UpdateEconomicActivity(
			context.Background(),
			merchantEconomicActivityId,
			merchantEconomicActivityBody,
		)
		assert.Error(t, err)
	})
}

func TestMerchantEconomicActivities_DeleteEconomicActivity(t *testing.T) {
	t.Run("When attempting to delete a merchant economic activities by id, the operation is successful.", func(t *testing.T) {
		merchantEconomicActivitiesRepository := &mockMerchantEconomicActivities.MerchantEconomicActivityRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantEconomicActivityId := "cf6e4017-f918-4ef0-b641-236d89901a5c"
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		merchantEconomicActivitiesRepository.
			On("DeleteEconomicActivity", mock.Anything, mock.Anything).
			Return(true, nil)
		merchantEconomicActivitiesUCase := NewMerchantEconomicActivitiesUseCase(merchantEconomicActivitiesRepository,
			validationRepository, authRepository, 60)
		res, err := merchantEconomicActivitiesUCase.DeleteEconomicActivity(context.Background(),
			merchantEconomicActivityId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("Encountered an error during merchant economic activities deletion by id.", func(t *testing.T) {
		merchantEconomicActivitiesRepository := &mockMerchantEconomicActivities.MerchantEconomicActivityRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		merchantEconomicActivitiesError := errors.New("random error")
		merchantEconomicActivityId := "cf6e4017-f918-4ef0-b641-236d89901a5c"
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		merchantEconomicActivitiesRepository.
			On("DeleteEconomicActivity", mock.Anything, mock.Anything).
			Return(false, merchantEconomicActivitiesError)
		merchantEconomicActivitiesUCase := NewMerchantEconomicActivitiesUseCase(merchantEconomicActivitiesRepository,
			validationRepository, authRepository, 60)
		res, err := merchantEconomicActivitiesUCase.DeleteEconomicActivity(context.Background(),
			merchantEconomicActivityId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
