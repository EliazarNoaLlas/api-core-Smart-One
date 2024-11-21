/*
 * File: policyPermissions_route_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for policyPermissions.
 *
 * Last Modified: 2023-11-20
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	policyPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
)

type policyPermissionsHandler struct {
	policyPermissionsUseCase policyPermissionsDomain.PolicyPermissionUseCase
	authMiddleware           authMiddleware.AuthMiddleware
	err                      *errDomain.SmartError
}

func NewPolicyPermissionsHandler(
	policyPermissions policyPermissionsDomain.PolicyPermissionUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware) {
	handler := &policyPermissionsHandler{
		policyPermissionsUseCase: policyPermissions,
		authMiddleware:           authMiddleware,
		err:                      errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)

	api.GET("/policies/:policyId/permissions", handler.GetPolicyPermissionsByPolicy)
	api.POST("/policies/:policyId/permissions", handler.CreatePolicyPermission)
	api.POST("/policies/:policyId/permissions/batch", handler.CreateMultiplePolicyPermissions)
	api.PUT("/policies/:policyId/permissions/:policyPermissionId", handler.UpdatePolicyPermission)
	api.DELETE("/policies/:policyId/permissions/batch", handler.DeleteMultiplePolicyPermissions)
	api.DELETE("/policies/:policyId/permissions/:policyPermissionId", handler.DeletePolicyPermission)
}
