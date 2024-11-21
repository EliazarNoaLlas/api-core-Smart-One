/*
 * File: setup_roles.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the roles.
 *
 * Last Modified: 2023-11-28
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

	rolesRepository "gitlab.smartcitiesperu.com/smartone/api-core/roles/infrastructure/persistence/mysql"
	rolesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/roles/interfaces/rest"
	rolesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/roles/usecase"
)

func LoadRoles(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	roleRepository := rolesRepository.NewRolesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	roleUseCase := rolesUseCase.NewRolesUseCase(
		roleRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	rolesHttpDelivery.NewRolesHandler(
		roleUseCase,
		router,
		authMiddleware,
	)
}
