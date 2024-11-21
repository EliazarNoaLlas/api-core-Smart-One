/*
 * File: setup_document_types.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the microservice document types.
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

	documentTypesRepository "gitlab.smartcitiesperu.com/smartone/api-core/document-types/infrastructure/persistence/mysql"
	documnetTypesHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/document-types/interfaces/rest"
	documentTypesUseCase "gitlab.smartcitiesperu.com/smartone/api-core/document-types/usecase"
)

func LoadDocumentTypes(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	validationRepository := validationsRepository.NewValidationsRepository(60)
	clock := smartClock.NewClock()
	documentTypeRepository := documentTypesRepository.NewDocumentTypesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	documentTypesUCase := documentTypesUseCase.NewDocumentTypesUseCase(
		documentTypeRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	documnetTypesHttpDelivery.NewDocumentTypesHandler(documentTypesUCase, router, authMiddleware)
}
