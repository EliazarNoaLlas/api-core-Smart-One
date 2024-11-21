/*
 * File: store_types_route_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the handler layer where the routes are located.
 *
 * Last Modified: 2023-08-02
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	storeTypeDomain "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain"
)

type storeTypesHandler struct {
	storeTypesUseCase storeTypeDomain.StoreTypeUseCase
	authMiddleware    authMiddleware.AuthMiddleware
	err               *errDomain.SmartError
}

func NewStoreTypesHandler(
	storeTypes storeTypeDomain.StoreTypeUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {

	handler := &storeTypesHandler{
		storeTypesUseCase: storeTypes,
		authMiddleware:    authMiddleware,
		err:               errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Cors)
	api.Use(handler.authMiddleware.Auth)
	api.GET("/store_types", handler.GetStoreTypes)
	api.POST("/store_types", handler.CreateStoreType)
	api.PUT("/store_types/:id", handler.UpdateStoreType)
	api.DELETE("/store_types/:storeTypeId", handler.DeleteStoreType)
}
