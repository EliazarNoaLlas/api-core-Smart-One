/*
 * File: role_policies_route_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for rolePolicies.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	rolePoliciesDomain "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain"
)

type rolePoliciesHandler struct {
	rolePoliciesUseCase rolePoliciesDomain.RolePolicyUseCase
	authMiddleware      authMiddleware.AuthMiddleware
	err                 *errDomain.SmartError
}

func NewRolePoliciesHandler(
	rolePolicies rolePoliciesDomain.RolePolicyUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &rolePoliciesHandler{
		rolePoliciesUseCase: rolePolicies,
		authMiddleware:      authMiddleware,
		err:                 errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)

	api.GET("/roles/:roleId/policies", handler.GetPolicies)
	api.POST("/roles/:roleId/policies", handler.CreateRolePolicy)
	api.POST("/roles/:roleId/policies/batch", handler.CreateMultipleRolePolicies)
	api.PUT("/roles/:roleId/policies/:rolePolicyId", handler.UpdateRolePolicy)
	api.DELETE("/roles/:roleId/policies/batch", handler.DeleteMultipleRolePolicies)
	api.DELETE("/roles/:roleId/policies/:rolePolicyId", handler.DeleteRolePolicy)
}
