/*
 * File: view_permissions_usecase.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the use case for viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package usecase

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	viewPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
)

type ViewPermissionsUseCase struct {
	ViewPermissionsRepository viewPermissionsDomain.ViewPermissionsRepository
	validationRepository      validationsDomain.ValidationRepository
	authRepository            authDomain.AuthRepository
	contextTimeout            time.Duration
	err                       *errDomain.SmartError
}

func NewViewPermissionsUseCase(
	ur viewPermissionsDomain.ViewPermissionsRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) viewPermissionsDomain.ViewPermissionsUseCase {
	return &ViewPermissionsUseCase{
		ViewPermissionsRepository: ur,
		validationRepository:      validation,
		authRepository:            authRepository,
		contextTimeout:            timeout,
		err:                       errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
