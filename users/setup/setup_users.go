/*
 * File: setup_users.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the users.
 *
 * Last Modified: 2023-12-28
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gitlab.smartcitiesperu.com/smartone/api-shared/auth"
	authRepository "gitlab.smartcitiesperu.com/smartone/api-shared/auth/infrastructure/jwt"
	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"
	validationsRepository "gitlab.smartcitiesperu.com/smartone/api-shared/validations/infrastructure/persistence/mysql"

	usersRepository "gitlab.smartcitiesperu.com/smartone/api-core/users/infrastructure/persistence/mysql"
	usersHttpDelivery "gitlab.smartcitiesperu.com/smartone/api-core/users/interfaces/rest"
	usersUseCase "gitlab.smartcitiesperu.com/smartone/api-core/users/usecase"
)

func LoadUsers(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	validationRepository := validationsRepository.NewValidationsRepository(60)
	clock := smartClock.NewClock()
	userRepository := usersRepository.NewUsersRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	usersUCase := usersUseCase.NewUsersUseCase(
		userRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	usersHttpDelivery.NewUsersHandler(usersUCase, router, authMiddleware)
}
