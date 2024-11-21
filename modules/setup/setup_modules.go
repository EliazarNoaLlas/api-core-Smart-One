/*
 * File: setup_modules.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the modules.
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

	modulesRepository "gitlab.smartcitiesperu.com/smartone/api-core/modules/infrastructure/persistence/mysql"
	modulesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/modules/interfaces/rest"
	modulesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/modules/usecase"
)

func LoadModules(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	moduleRepository := modulesRepository.NewModulesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	modulesUCase := modulesUseCase.NewModulesUseCase(
		moduleRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	modulesHttpDelivery.NewModulesHandler(modulesUCase, router, authMiddleware)
}
