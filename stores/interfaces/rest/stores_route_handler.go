/*
 * File: stores_route_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for stores.
 *
 * Last Modified: 2023-11-14
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	storesDomain "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain"
)

type storesHandler struct {
	storesUseCase  storesDomain.StoreUseCase
	authMiddleware authMiddleware.AuthMiddleware
	err            *errDomain.SmartError
}

func NewStoresHandler(
	stores storesDomain.StoreUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &storesHandler{
		storesUseCase:  stores,
		authMiddleware: authMiddleware,
		err:            errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/merchants/:merchantId/stores", handler.GetStores)
	api.POST("/merchants/:merchantId/stores", handler.CreateStore)
	api.PUT("/merchants/:merchantId/stores/:storeId", handler.UpdateStore)
	api.DELETE("/merchants/:merchantId/stores/:storeId", handler.DeleteStore)
}
