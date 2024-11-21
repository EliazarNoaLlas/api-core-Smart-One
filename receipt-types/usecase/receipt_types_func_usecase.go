/*
 * File: receipt_types_func_usecase.go
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
	"context"

	"github.com/google/uuid"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	receiptTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
)

func (u ReceiptTypesUseCase) GetReceiptTypes(
	ctx context.Context,
) (
	receiptTypes []receiptTypesDomain.ReceiptType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	receiptTypes, err = u.ReceiptTypesRepository.GetReceiptTypes(ctx)
	if err != nil {
		return nil, err
	}
	return receiptTypes, nil
}

func (u ReceiptTypesUseCase) CreateReceiptType(
	ctx context.Context,
	userId string,
	body receiptTypesDomain.CreateReceiptTypeBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	deleted := "deleted_at"
	recordExistsParamsPermission := validationsDomain.RecordExistsParams{
		Table:            "core_receipt_types",
		IdColumnName:     "sunat_code",
		IdValue:          body.SunatCode,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}

	var existPermission bool
	existPermission, err = u.validationRepository.RecordExists(ctx, recordExistsParamsPermission)
	if err != nil {
		return nil, err
	}
	if existPermission {
		return nil, receiptTypesDomain.ErrSunatCodeAlreadyExist
	}

	receiptTypeId := uuid.New().String()
	err = u.ReceiptTypesRepository.CreateReceiptType(ctx, receiptTypeId, userId, body)
	id = &receiptTypeId
	return
}

func (u ReceiptTypesUseCase) UpdateReceiptType(
	ctx context.Context,
	receiptTypeId string,
	body receiptTypesDomain.UpdateReceiptTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted string
	deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_receipt_types",
		IdColumnName:     "id",
		IdValue:          receiptTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(receiptTypesDomain.ErrReceiptTypeNotFound).
			SetFunction("UpdateReceiptType")
	}
	err = u.ReceiptTypesRepository.UpdateReceiptType(ctx, receiptTypeId, body)
	return
}

func (u ReceiptTypesUseCase) DeleteReceiptType(
	ctx context.Context,
	receiptTypeId string,
) (
	update bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	var deleted string
	deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_receipt_types",
		IdColumnName:     "id",
		IdValue:          receiptTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, receiptTypesDomain.ErrReceiptTypeIdHasBeenDeleted
	}
	update, err = u.ReceiptTypesRepository.DeleteReceiptType(ctx, receiptTypeId)
	return
}
