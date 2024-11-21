/*
 * File: setup_economic_activities.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the economic activities types.
 *
 * Last Modified: 2023-12-28
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	economicActivitiesRepository "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/infrastructure/persistence/mysql"
	economicActivitiesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/interfaces/rest"
	economicAtivitiesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/usecase"
	"gitlab.smartcitiesperu.com/smartone/api-shared/auth"
	authRepository "gitlab.smartcitiesperu.com/smartone/api-shared/auth/infrastructure/jwt"
	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"
)

func LoadEconomicActivities(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	economicActivityRepository := economicActivitiesRepository.NewEconomicActivitiesRepository(clock,
		60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	economicActivitiesUCase := economicAtivitiesUseCase.NewEconomicActivitiesUseCase(
		economicActivityRepository,
		authJWTRepository,
		timeoutContext)
	economicActivitiesHttpDelivery.NewEconomicActivitiesHandler(economicActivitiesUCase, router, authMiddleware)
}
