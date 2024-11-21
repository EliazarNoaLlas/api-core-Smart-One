/*
 * File: receipt_types_error.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the errors of the receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrReceiptTypeNotFoundCode         = "ERR_RECEIPT_TYPE_NOT_FOUND"
	ErrSunatCodeAlreadyExistCode       = "ERR_SUNAT_CODE_ALREADY_EXIST"
	ErrReceiptTypeIdHasBeenDeletedCode = "ERR_RECEIPT_TYPE_ID_HAS_BEEN_DELETED"
)

var (
	ErrReceiptTypeNotFound = errDomain.NewErr().
				SetCode(ErrReceiptTypeNotFoundCode).
				SetDescription("RECEIPT TYPES NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateReceiptType")

	ErrSunatCodeAlreadyExist = errDomain.NewErr().
					SetCode(ErrSunatCodeAlreadyExistCode).
					SetDescription("SUNAT CODE ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateReceiptType")

	ErrReceiptTypeIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrReceiptTypeIdHasBeenDeletedCode).
					SetDescription("RECEIPT TYPE ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteReceiptType")
)
