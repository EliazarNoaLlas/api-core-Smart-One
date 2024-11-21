/*
 * File: economic_activities_route_handler.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-04
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	economicActivitiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
)

type economicActivitiesHandler struct {
	economicActivitiesUseCase economicActivitiesDomain.EconomicActivityUseCase
	authMiddleware            authMiddleware.AuthMiddleware
	err                       *errDomain.SmartError
}

func NewEconomicActivitiesHandler(
	economicActivities economicActivitiesDomain.EconomicActivityUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &economicActivitiesHandler{
		economicActivitiesUseCase: economicActivities,
		authMiddleware:            authMiddleware,
		err:                       errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/economic_activities", handler.GetEconomicActivities)
}
