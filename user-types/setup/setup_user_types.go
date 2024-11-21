/*
 * File: setup_user_types.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the user types.
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

	userTypesRepository "gitlab.smartcitiesperu.com/smartone/api-core/user-types/infrastructure/persistence/mysql"
	userTypesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/user-types/interfaces/rest"
	userTypesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/user-types/usecase"
)

func LoadUserTypes(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	userTypeRepository := userTypesRepository.NewUserTypesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	userTypesUCase := userTypesUseCase.NewUserTypesUseCase(
		userTypeRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	userTypesHttpDelivery.NewUserTypesHandler(userTypesUCase, router, authMiddleware)
}
