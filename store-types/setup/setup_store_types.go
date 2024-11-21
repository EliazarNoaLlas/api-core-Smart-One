/*
 * File: setup_store_types.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the store types.
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

	storeTypesRepository "gitlab.smartcitiesperu.com/smartone/api-core/store-types/infrastructure/persistence/mysql"
	storeTypesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/store-types/interfaces/rest"
	storeTypesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/store-types/usecase"
)

func LoadStoreTypes(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	permissionRepository := storeTypesRepository.NewStoreTypesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	permissionsUCase := storeTypesUseCase.NewStoreTypesUseCase(
		permissionRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	storeTypesHttpDelivery.NewStoreTypesHandler(permissionsUCase, router, authMiddleware)
}
