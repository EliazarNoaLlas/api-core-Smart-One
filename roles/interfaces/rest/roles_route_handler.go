/*
 * File: roles_route_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the handler layer where the routes are located.
 *
 * Last Modified: 2023-11-14
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	rolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/roles/domain"
)

type rolesHandler struct {
	rolesUseCase   rolesDomain.RoleUseCase
	authMiddleware authMiddleware.AuthMiddleware
	err            *errDomain.SmartError
}

func NewRolesHandler(
	roles rolesDomain.RoleUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {

	handler := &rolesHandler{
		rolesUseCase:   roles,
		authMiddleware: authMiddleware,
		err:            errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/roles", handler.GetRoles)
	api.POST("/roles", handler.CreateRole)
	api.PUT("/roles/:roleId", handler.UpdateRole)
	api.DELETE("/roles/:roleId", handler.DeleteRole)
}
