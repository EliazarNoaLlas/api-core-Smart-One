/*
 * File: modules_route_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for modules.
 *
 * Last Modified: 2023-11-10
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	modulesDomain "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"
)

type modulesHandler struct {
	modulesUseCase modulesDomain.ModuleUseCase
	authMiddleware authMiddleware.AuthMiddleware
	err            *errDomain.SmartError
}

func NewModulesHandler(
	modules modulesDomain.ModuleUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &modulesHandler{
		modulesUseCase: modules,
		authMiddleware: authMiddleware,
		err:            errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/modules", handler.GetModules)
	api.POST("/modules", handler.CreateModule)
	api.PUT("/modules/:moduleId", handler.UpdateModule)
	api.DELETE("/modules/:moduleId", handler.DeleteModule)
}
