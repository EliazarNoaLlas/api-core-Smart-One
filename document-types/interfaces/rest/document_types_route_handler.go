/*
 * File: document_types_route_handler.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-07
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain"
)

type documentTypesHandler struct {
	documentTypesUseCase domain.DocumentTypeUseCase
	authMiddleware       authMiddleware.AuthMiddleware
	err                  *errDomain.SmartError
}

func NewDocumentTypesHandler(
	documentTypes domain.DocumentTypeUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &documentTypesHandler{
		documentTypesUseCase: documentTypes,
		authMiddleware:       authMiddleware,
		err:                  errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)

	api.GET("/document_types/", handler.GetDocumentTypes)
	api.POST("/document_types/create_document_types/:documentTypeId", handler.CreateDocumentType)
	api.PUT("/document_types/update_document_types/:documentTypeId", handler.UpdateDocumentType)
	api.DELETE("/document_types/delete_document_types/:documentTypeId", handler.DeleteDocumentType)
}
