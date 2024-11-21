/*
 * File: receipt_types_usecase.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the use case for receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package usecase

import (
	"time"

	authDomain "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	receiptTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
)

type ReceiptTypesUseCase struct {
	ReceiptTypesRepository receiptTypesDomain.ReceiptTypesRepository
	validationRepository   validationsDomain.ValidationRepository
	authRepository         authDomain.AuthRepository
	contextTimeout         time.Duration
	err                    *errDomain.SmartError
}

func NewReceiptTypesUseCase(
	ur receiptTypesDomain.ReceiptTypesRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) receiptTypesDomain.ReceiptTypesUseCase {
	return &ReceiptTypesUseCase{
		ReceiptTypesRepository: ur,
		validationRepository:   validation,
		authRepository:         authRepository,
		contextTimeout:         timeout,
		err:                    errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
