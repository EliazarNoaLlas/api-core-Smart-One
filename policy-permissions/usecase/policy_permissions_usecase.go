/*
 * File: policyPermissions_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Initializing use cases for policyPermissions
 *
 * Last Modified: 2023-11-20
 */

package usecase

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
)

type policyPermissionsUseCase struct {
	policyPermissionsRepository domain.PolicyPermissionRepository
	validationRepository        validationsDomain.ValidationRepository
	authRepository              authDomain.AuthRepository
	contextTimeout              time.Duration
	err                         *errDomain.SmartError
}

func NewPolicyPermissionsUseCase(
	ur domain.PolicyPermissionRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.PolicyPermissionUseCase {
	return &policyPermissionsUseCase{
		policyPermissionsRepository: ur,
		validationRepository:        validation,
		authRepository:              authRepository,
		contextTimeout:              timeout,
		err:                         errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
