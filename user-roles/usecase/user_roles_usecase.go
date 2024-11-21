/*
 * File: user_roles_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Initializing use cases for userRoles
 *
 * Last Modified: 2023-11-23
 */

package usecase

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
)

type userRolesUseCase struct {
	userRolesRepository  domain.UserRoleRepository
	validationRepository validationsDomain.ValidationRepository
	authRepository       authDomain.AuthRepository
	contextTimeout       time.Duration
	err                  *errDomain.SmartError
}

func NewUserRolesUseCase(
	ur domain.UserRoleRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.UserRoleUseCase {
	return &userRolesUseCase{
		userRolesRepository:  ur,
		validationRepository: validation,
		authRepository:       authRepository,
		contextTimeout:       timeout,
		err:                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
