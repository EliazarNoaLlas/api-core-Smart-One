/*
 * File: merchants_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the use case layer where the layer is initialized.
 *
 * Last Modified: 2023-11-10
 */

package usecase

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
)

type merchantsUseCase struct {
	merchantsRepository  domain.MerchantRepository
	validationRepository validationsDomain.ValidationRepository
	authRepository       authDomain.AuthRepository
	contextTimeout       time.Duration
	err                  *errDomain.SmartError
}

func NewMerchantsUseCase(
	ur domain.MerchantRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.MerchantUseCase {
	return &merchantsUseCase{
		merchantsRepository:  ur,
		validationRepository: validation,
		authRepository:       authRepository,
		contextTimeout:       timeout,
		err:                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
