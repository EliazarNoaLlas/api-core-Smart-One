/*
 * File: server_route_handler.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the server route handler.
 *
 * Last Modified: 2024-04-09
 */

package rest

import (
	"github.com/gin-gonic/gin"
	swaggerRest "gitlab.smartcitiesperu.com/smartone/api-shared/swagger/interfaces/rest"

	authMiddleware "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/server/docs"
	"gitlab.smartcitiesperu.com/smartone/api-core/server/domain"
)

type serverHandler struct {
	serverUseCase  domain.ServerUseCase
	authMiddleware authMiddleware.AuthMiddleware
	err            *errDomain.SmartError
}

func NewServerHandler(
	server domain.ServerUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &serverHandler{
		serverUseCase:  server,
		authMiddleware: authMiddleware,
		err:            errDomain.NewErr().SetLayer(errDomain.Interface),
	}

	swaggerRest.Handler(router, docs.SwaggerInfoserver, docs.DocTemplateJson, "core", "server")

	api := router.Group("/api/v1/core")
	api.Use(handler.authMiddleware.Cors)
	api.Use(handler.authMiddleware.Auth)
	api.GET("/server/datetime", handler.GetServerDatetime)
}
