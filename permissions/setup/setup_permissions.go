/*
 * File: setup_permissions.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the permissions.
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

	permissionsRepository "gitlab.smartcitiesperu.com/smartone/api-core/permissions/infrastructure/persistence/mysql"
	permissionsHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/permissions/interfaces/rest"
	permissionsUseCase "gitlab.smartcitiesperu.com/smartone/api-core/permissions/usecase"
)

func LoadPermissions(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	permissionRepository := permissionsRepository.NewPermissionsRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	permissionsUCase := permissionsUseCase.NewPermissionsUseCase(
		permissionRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	permissionsHttpDelivery.NewPermissionsHandler(permissionsUCase, router, authMiddleware)
}
