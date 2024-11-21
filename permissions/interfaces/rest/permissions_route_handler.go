/*
 * File: permissions_route_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the handler layer where the routes are located.
 *
 * Last Modified: 2023-11-15
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	permissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
)

type permissionsHandler struct {
	permissionsUseCase permissionsDomain.PermissionUseCase
	authMiddleware     authMiddleware.AuthMiddleware
	err                *errDomain.SmartError
}

func NewPermissionsHandler(
	permissions permissionsDomain.PermissionUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &permissionsHandler{
		permissionsUseCase: permissions,
		authMiddleware:     authMiddleware,
		err:                errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/modules/:moduleId/permissions", handler.GetPermissions)
	api.POST("/modules/:moduleId/permissions", handler.CreatePermission)
	api.PUT("/modules/:moduleId/permissions/:permissionId", handler.UpdatePermission)
	api.DELETE("/modules/:moduleId/permissions/:permissionId", handler.DeletePermission)
}
