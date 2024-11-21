/*
 * File: economic_activities_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-04
 */

package usecase

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
)

type economicActivitiesUseCase struct {
	economicActivitiesRepository domain.EconomicActivityRepository
	authRepository               authDomain.AuthRepository
	contextTimeout               time.Duration
	err                          *errDomain.SmartError
}

func NewEconomicActivitiesUseCase(
	ur domain.EconomicActivityRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.EconomicActivityUseCase {
	return &economicActivitiesUseCase{
		economicActivitiesRepository: ur,
		authRepository:               authRepository,
		contextTimeout:               timeout,
		err:                          errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
