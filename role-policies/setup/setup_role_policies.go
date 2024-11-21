/*
 * File: setup_role_policies.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the role policies.
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

	rolePoliciesRepository "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/infrastructure/persistence/mysql"
	rolePoliciesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/interfaces/rest"
	rolePoliciesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/usecase"
)

func LoadRolePolicies(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	permissionRepository := rolePoliciesRepository.NewRolePoliciesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	permissionsUCase := rolePoliciesUseCase.NewRolePoliciesUseCase(
		permissionRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	rolePoliciesHttpDelivery.NewRolePoliciesHandler(permissionsUCase, router, authMiddleware)
}
