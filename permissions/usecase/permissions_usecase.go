/*
 * File: permissions_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the use case layer where the permissions use case is initialized.
 *
 * Last Modified: 2023-11-15
 */

package usecase

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
)

type PermissionUseCase struct {
	permissionsRepository domain.PermissionRepository
	validationRepository  validationsDomain.ValidationRepository
	authRepository        authDomain.AuthRepository
	contextTimeout        time.Duration
	err                   *errDomain.SmartError
}

func NewPermissionsUseCase(
	ur domain.PermissionRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.PermissionUseCase {
	return &PermissionUseCase{
		permissionsRepository: ur,
		validationRepository:  validation,
		authRepository:        authRepository,
		contextTimeout:        timeout,
		err:                   errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
