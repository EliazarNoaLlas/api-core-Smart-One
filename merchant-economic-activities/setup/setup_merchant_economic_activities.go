/*
 * File: setup_merchant_economic_activities.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the merchant economic activities types.
 *
 * Last Modified: 2023-12-05
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gitlab.smartcitiesperu.com/smartone/api-shared/auth"
	authRepository "gitlab.smartcitiesperu.com/smartone/api-shared/auth/infrastructure/jwt"
	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"
	validationsRepository "gitlab.smartcitiesperu.com/smartone/api-shared/validations/infrastructure/persistence/mysql"

	merchantEconomicActivityRepository "gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/infrastructure/persistence/mysql"
	merchantEconomicActivitiesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/interfaces/rest"
	merchantEconomicActivitiesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/usecase"
)

func LoadMerchantEconomicActivities(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	validationRepository := validationsRepository.NewValidationsRepository(60)
	clock := smartClock.NewClock()
	merchantActivityRepository := merchantEconomicActivityRepository.NewMerchantEconomicActivitiesRepository(
		clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	merchantEconomicActivitiesUCase := merchantEconomicActivitiesUseCase.NewMerchantEconomicActivitiesUseCase(
		merchantActivityRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	merchantEconomicActivitiesHttpDelivery.NewMerchantEconomicActivitiesHandler(merchantEconomicActivitiesUCase,
		router, authMiddleware)
}
