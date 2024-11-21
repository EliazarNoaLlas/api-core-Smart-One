/*
 * File: policies_route_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for policies.
 *
 * Last Modified: 2023-11-14
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	policiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain"
)

type policiesHandler struct {
	policiesUseCase policiesDomain.PolicyUseCase
	authMiddleware  authMiddleware.AuthMiddleware
	err             *errDomain.SmartError
}

func NewPoliciesHandler(
	policies policiesDomain.PolicyUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &policiesHandler{
		policiesUseCase: policies,
		authMiddleware:  authMiddleware,
		err:             errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Cors)
	api.Use(handler.authMiddleware.Auth)
	api.GET("/policies", handler.GetPolicies)
	api.POST("/policies", handler.CreatePolicy)
	api.PUT("/policies/:policyId", handler.UpdatePolicy)
	api.DELETE("/policies/:policyId", handler.DeletePolicy)
}
