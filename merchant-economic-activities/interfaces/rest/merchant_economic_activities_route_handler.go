/*
 * File: merchant_economic_activities_route_handler.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the route handler of the merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package interfaces

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
)

type merchantEconomicActivitiesHandler struct {
	merchantEconomicActivitiesUseCase domain.MerchantEconomicActivityUseCase
	authMiddleware                    authMiddleware.AuthMiddleware
	err                               *errDomain.SmartError
}

func NewMerchantEconomicActivitiesHandler(
	merchantEconomicActivities domain.MerchantEconomicActivityUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &merchantEconomicActivitiesHandler{
		merchantEconomicActivitiesUseCase: merchantEconomicActivities,
		authMiddleware:                    authMiddleware,
		err:                               errDomain.NewErr().SetLayer(errDomain.Interface),
	}

	apiAuth := router.Group("/api/v1/auth")
	api := router.Group("/api/v1/core")
	apiAuth.Use(handler.authMiddleware.Cors)
	api.Use(handler.authMiddleware.Auth)
	api.GET("/merchant_economic_activities/:merchantId", handler.GetMerchantEconomicActivities)
	api.POST("/merchant_economic_activities/url/:merchantEconomicActivityId", handler.CreateEconomicActivity)
	api.PUT("/merchant_economic_activities/:merchantEconomicActivityId", handler.UpdateEconomicActivity)
	api.DELETE("/merchant_economic_activities/:merchantEconomicActivityId", handler.DeleteEconomicActivity)
}
