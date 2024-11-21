/*
 * File: setup_merchants.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the merchants.
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

	merchantsRepository "gitlab.smartcitiesperu.com/smartone/api-core/merchants/infrastructure/persistence/mysql"
	merchantsHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/merchants/interfaces/rest"
	merchantsUseCase "gitlab.smartcitiesperu.com/smartone/api-core/merchants/usecase"
)

func LoadMerchants(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	merchantRepository := merchantsRepository.NewMerchantsRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	merchantUCase := merchantsUseCase.NewMerchantsUseCase(
		merchantRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	merchantsHttpDelivery.NewMerchantsHandler(merchantUCase, router, authMiddleware)
}
