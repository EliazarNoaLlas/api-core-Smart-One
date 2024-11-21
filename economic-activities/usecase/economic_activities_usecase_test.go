/*
 * File: economic_activities_usecase_test.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the use case test.
 *
 * Last Modified: 2023-12-04
 */

package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	economicActivitiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	mockEconomicActivities "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain/mocks"
)

func TestUseCaseEconomicActivities_GetEconomicActivities(t *testing.T) {
	t.Run("When economic activities are successfully listed", func(t *testing.T) {
		economicActivityRepository := &mockEconomicActivities.EconomicActivityRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		economicActivityRepository.On("GetEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return([]economicActivitiesDomain.EconomicActivity{}, nil)
		economicActivityRepository.On("GetTotalGetEconomicActivities", mock.Anything, mock.Anything,
			mock.Anything).Return(&total, nil)
		economicActivitiesUCase := NewEconomicActivitiesUseCase(economicActivityRepository, authRepository, 60)
		searchParams := economicActivitiesDomain.GetEconomicActivitiesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		economicActivities, _, err := economicActivitiesUCase.GetEconomicActivities(context.Background(), searchParams,
			pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, economicActivities, []economicActivitiesDomain.EconomicActivity{})
	})

	t.Run("When an error occurs while listing economic activities", func(t *testing.T) {
		economicActivityRepository := &mockEconomicActivities.EconomicActivityRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		economicActivityRepository.On("GetEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		economicActivityRepository.On("GetTotalGetEconomicActivities", mock.Anything, mock.Anything,
			mock.Anything).Return(&total, errors.New("random error"))
		economicActivitiesUCase := NewEconomicActivitiesUseCase(economicActivityRepository, authRepository, 60)
		searchParams := economicActivitiesDomain.GetEconomicActivitiesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		economicActivities, _, err := economicActivitiesUCase.GetEconomicActivities(context.Background(), searchParams,
			pagination)
		assert.Error(t, err)
		assert.EqualValues(t, economicActivities, []economicActivitiesDomain.EconomicActivity(nil))
	})
}
