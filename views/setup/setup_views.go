/*
 * File: setup_views.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the views.
 *
 * Last Modified: 2023-12-28
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

	viewsRepository "gitlab.smartcitiesperu.com/smartone/api-core/views/infrastructure/persistence/mysql"
	viewsHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/views/interfaces/rest"
	viewsUseCase "gitlab.smartcitiesperu.com/smartone/api-core/views/usecase"
)

func LoadViews(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	viewRepository := viewsRepository.NewViewRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	viewUseCase := viewsUseCase.NewViewUseCase(
		viewRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	viewsHttpDelivery.NewViewsHandler(
		viewUseCase,
		router,
		authMiddleware,
	)
}
