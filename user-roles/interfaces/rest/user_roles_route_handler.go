/*
 * File: user_roles_route_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for userRoles.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	userRolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
)

type userRolesHandler struct {
	userRolesUseCase userRolesDomain.UserRoleUseCase
	authMiddleware   authMiddleware.AuthMiddleware
	err              *errDomain.SmartError
}

func NewUserRolesHandler(
	userRoles userRolesDomain.UserRoleUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &userRolesHandler{
		userRolesUseCase: userRoles,
		authMiddleware:   authMiddleware,
		err:              errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Cors)
	api.Use(handler.authMiddleware.Auth)
	api.GET("/users/:userId/roles", handler.GetUserRolesByUser)
	api.POST("/users/:userId/roles", handler.CreateUserRole)
	api.PUT("/users/:userId/roles/:userRoleId", handler.UpdateUserRole)
	api.DELETE("/users/:userId/roles/:userRoleId", handler.DeleteUserRole)
}
