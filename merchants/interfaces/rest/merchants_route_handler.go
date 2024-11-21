/*
 * File: merchants_route_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the handler layer where the routes are located.
 *
 * Last Modified: 2023-11-10
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	merchantsDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
)

type merchantsHandler struct {
	merchantsUseCase merchantsDomain.MerchantUseCase
	authMiddleware   authMiddleware.AuthMiddleware
	err              *errDomain.SmartError
}

func NewMerchantsHandler(
	merchants merchantsDomain.MerchantUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &merchantsHandler{
		merchantsUseCase: merchants,
		authMiddleware:   authMiddleware,
		err:              errDomain.NewErr().SetLayer(errDomain.Interface),
	}

	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/merchants", handler.GetMerchants)
	api.POST("/merchants", handler.CreateMerchant)
	api.PUT("/merchants/:merchantId", handler.UpdateMerchant)
	api.DELETE("/merchants/:merchantId", handler.DeleteMerchant)
}
