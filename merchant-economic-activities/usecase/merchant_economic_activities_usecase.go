/*
 * File: merchant_economic_activities_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the use case of merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package usecase

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
)

type merchantEconomicActivitiesUseCase struct {
	merchantEconomicActivitiesRepository domain.MerchantEconomicActivityRepository
	validationRepository                 validationsDomain.ValidationRepository
	authRepository                       authDomain.AuthRepository
	contextTimeout                       time.Duration
	err                                  *errDomain.SmartError
}

func NewMerchantEconomicActivitiesUseCase(
	ur domain.MerchantEconomicActivityRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.MerchantEconomicActivityUseCase {
	return &merchantEconomicActivitiesUseCase{
		merchantEconomicActivitiesRepository: ur,
		validationRepository:                 validation,
		authRepository:                       authRepository,
		contextTimeout:                       timeout,
		err:                                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
