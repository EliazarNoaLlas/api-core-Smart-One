/*
 * File: setup_stores.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the stores.
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

	storesReposiotry "gitlab.smartcitiesperu.com/smartone/api-core/stores/infrastructure/persistence/mysql"
	storesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/stores/interfaces/rest"
	storesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/stores/usecase"
)

func LoadStores(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	storeRepository := storesReposiotry.NewStoresRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	storeUCase := storesUseCase.NewStoresUseCase(
		storeRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	storesHttpDelivery.NewStoresHandler(storeUCase, router, authMiddleware)
}
