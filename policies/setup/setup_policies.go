/*
 * File: setup_policies.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the policies.
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

	policiesRepository "gitlab.smartcitiesperu.com/smartone/api-core/policies/infrastructure/persistence/mysql"
	policiesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/policies/interfaces/rest"
	policiesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/policies/usecase"
)

func LoadPolicies(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	validationRepository := validationsRepository.NewValidationsRepository(60)
	clock := smartClock.NewClock()
	policyRepository := policiesRepository.NewPoliciesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	policiesUCase := policiesUseCase.NewPoliciesUseCase(
		policyRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	policiesHttpDelivery.NewPoliciesHandler(policiesUCase, router, authMiddleware)
}
