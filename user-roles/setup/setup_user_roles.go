/*
 * File: setup_user_roles.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the user roles.
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

	userRolesRepository "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/infrastructure/persistence/mysql"
	userRolesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/interfaces/rest"
	userRolesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/usecase"
)

func LoadUserRoles(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	userRoleRepository := userRolesRepository.NewUserRolesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	userRolesUCase := userRolesUseCase.NewUserRolesUseCase(
		userRoleRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	userRolesHttpDelivery.NewUserRolesHandler(userRolesUCase, router, authMiddleware)
}
