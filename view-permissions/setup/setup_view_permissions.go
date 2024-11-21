/*
 * File: setup_view_permissions.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the setup for viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.smartcitiesperu.com/smartone/api-shared/auth"
	authRepository "gitlab.smartcitiesperu.com/smartone/api-shared/auth/infrastructure/jwt"
	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"
	validationsRepository "gitlab.smartcitiesperu.com/smartone/api-shared/validations/infrastructure/persistence/mysql"

	ViewPermissionsRepository "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/infrastructure/persistence/mysql"
	ViewPermissionsHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/interfaces/rest"
	ViewPermissionsUseCase "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/usecase"
)

func LoadViewPermissions(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	viewPermissionsRepository := ViewPermissionsRepository.NewViewPermissionsRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	ViewPermissionsUCase := ViewPermissionsUseCase.NewViewPermissionsUseCase(
		viewPermissionsRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	ViewPermissionsHttpDelivery.NewViewPermissionsHandler(ViewPermissionsUCase, router, authMiddleware)
}
