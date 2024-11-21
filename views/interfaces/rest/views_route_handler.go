/*
 * File: views_route_handler.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for views.
 *
 * Last Modified: 2023-11-24
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	viewsDomain "gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
)

type viewsHandler struct {
	viewsUseCase   viewsDomain.ViewUseCase
	authMiddleware authMiddleware.AuthMiddleware
	err            *errDomain.SmartError
}

func NewViewsHandler(
	views viewsDomain.ViewUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &viewsHandler{
		viewsUseCase:   views,
		authMiddleware: authMiddleware,
		err:            errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)

	api.GET("/modules/:moduleId/views", handler.GetViews)
	api.POST("/modules/:moduleId/views", handler.CreateView)
	api.PUT("/modules/:moduleId/views/:viewId", handler.UpdateView)
	api.DELETE("/modules/:moduleId/views/:viewId", handler.DeleteView)
}
