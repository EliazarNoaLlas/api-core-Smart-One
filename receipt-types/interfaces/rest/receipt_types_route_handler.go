/*
 * File: receipt_types_route_handler.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the handler for receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	ReceiptTypes "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
)

type ReceiptTypesHandler struct {
	ReceiptTypesUseCase ReceiptTypes.ReceiptTypesUseCase
	authMiddleware      authMiddleware.AuthMiddleware
	err                 *errDomain.SmartError
}

func NewReceiptTypesHandler(
	ReceiptTypes ReceiptTypes.ReceiptTypesUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &ReceiptTypesHandler{
		ReceiptTypesUseCase: ReceiptTypes,
		authMiddleware:      authMiddleware,
		err:                 errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/receipt_types", handler.GetReceiptTypes)
	api.POST("/receipt_types", handler.CreateReceiptType)
	api.PUT("/receipt_types/:receiptTypeId", handler.UpdateReceiptType)
	api.DELETE("/receipt_types/:receiptTypeId", handler.DeleteReceiptType)
}
