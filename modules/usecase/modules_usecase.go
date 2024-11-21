/*
 * File: modules_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Initializing use cases for modules
 *
 * Last Modified: 2023-11-10
 */

package usecase

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"
)

type modulesUseCase struct {
	modulesRepository    domain.ModuleRepository
	validationRepository validationsDomain.ValidationRepository
	authRepository       authDomain.AuthRepository
	contextTimeout       time.Duration
	err                  *errDomain.SmartError
}

func NewModulesUseCase(
	ur domain.ModuleRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.ModuleUseCase {
	return &modulesUseCase{
		modulesRepository:    ur,
		validationRepository: validation,
		authRepository:       authRepository,
		contextTimeout:       timeout,
		err:                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
