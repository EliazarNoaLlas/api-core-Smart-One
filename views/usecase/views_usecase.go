/*
 * File: views_usecase.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Initializing use cases for views
 *
 * Last Modified: 2023-11-24
 */

package usecases

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
)

type viewsUseCase struct {
	viewsRepository      domain.ViewRepository
	validationRepository validationsDomain.ValidationRepository
	authRepository       authDomain.AuthRepository
	contextTimeout       time.Duration
	err                  *errDomain.SmartError
}

func NewViewUseCase(
	ur domain.ViewRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.ViewUseCase {
	return &viewsUseCase{
		viewsRepository:      ur,
		validationRepository: validation,
		authRepository:       authRepository,
		contextTimeout:       timeout,
		err:                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
