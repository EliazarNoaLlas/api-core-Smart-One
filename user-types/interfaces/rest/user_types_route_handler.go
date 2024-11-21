/*
 * File: user_types_route_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for user types.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	userTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain"
)

type userTypesHandler struct {
	userTypesUseCase userTypesDomain.UserTypeUseCase
	authMiddleware   authMiddleware.AuthMiddleware
	err              *errDomain.SmartError
}

func NewUserTypesHandler(
	userTypes userTypesDomain.UserTypeUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &userTypesHandler{
		userTypesUseCase: userTypes,
		authMiddleware:   authMiddleware,
		err:              errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/user_types", handler.GetUserTypes)
	api.POST("/user_types", handler.CreateUserType)
	api.PUT("/user_types/:userTypeId", handler.UpdateUserType)
	api.DELETE("/user_types/:userTypeId", handler.DeleteUserType)
}
