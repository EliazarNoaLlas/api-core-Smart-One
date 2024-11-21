/*
 * File: setup_policy_permissions.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the policy permissions.
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

	policyPermissionsRepository "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/infrastructure/persistence/mysql"
	policyPermissionsHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/interfaces/rest"
	policyPermissionsUseCase "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/usecase"
)

func LoadPolicyPermissions(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	permissionRepository := policyPermissionsRepository.NewPolicyPermissionsRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	permissionsUCase := policyPermissionsUseCase.NewPolicyPermissionsUseCase(
		permissionRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	policyPermissionsHttpDelivery.NewPolicyPermissionsHandler(permissionsUCase, router, authMiddleware)
}
