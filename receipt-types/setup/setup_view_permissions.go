/*
 * File: setup_receipt_types.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the setup for receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.smartcitiesperu.com/smartone/api-shared/auth"
	authRepository "gitlab.smartcitiesperu.com/smartone/api-shared/auth/infrastructure/jwt"
	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"
	validationsRepository "gitlab.smartcitiesperu.com/smartone/api-shared/validations/infrastructure/persistence/mysql"

	ReceiptTypesRepository "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/infrastructure/persistence/mysql"
	ReceiptTypesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/interfaces/rest"
	ReceiptTypesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/usecase"
)

func LoadReceiptTypes(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	receiptTypesRepository := ReceiptTypesRepository.NewReceiptTypesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	ReceiptTypesUCase := ReceiptTypesUseCase.NewReceiptTypesUseCase(
		receiptTypesRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	ReceiptTypesHttpDelivery.NewReceiptTypesHandler(ReceiptTypesUCase, router, authMiddleware)
}
