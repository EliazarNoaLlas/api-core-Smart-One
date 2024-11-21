/*
 * File: setup_server.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the setup of the server.
 *
 * Last Modified: 2024-04-09
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gitlab.smartcitiesperu.com/smartone/api-shared/auth"

	serverHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/server/interfaces/rest"
	serverUseCase "gitlab.smartcitiesperu.com/smartone/api-core/server/usecase"
)

func LoadServer(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	authMiddleware := auth.LoadAuthMiddleware()
	serverUCase := serverUseCase.NewServerUseCase(timeoutContext)
	serverHttpDelivery.NewServerHandler(serverUCase, router, authMiddleware)
}
