/*
 * File: view_permissions_route_handler.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the handler for viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	ViewPermissions "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
)

type ViewPermissionsHandler struct {
	ViewPermissionsUseCase ViewPermissions.ViewPermissionsUseCase
	authMiddleware         authMiddleware.AuthMiddleware
	err                    *errDomain.SmartError
}

func NewViewPermissionsHandler(
	ViewPermissions ViewPermissions.ViewPermissionsUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &ViewPermissionsHandler{
		ViewPermissionsUseCase: ViewPermissions,
		authMiddleware:         authMiddleware,
		err:                    errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/views/:viewId/permissions", handler.GetViewPermissions)
	api.POST("/views/:viewId/permissions", handler.CreateViewPermission)
	api.PUT("/views/:viewId/permissions/:viewPermissionId", handler.UpdateViewPermission)
	api.DELETE("/views/:viewId/permissions/:viewPermissionId", handler.DeleteViewPermission)
}
