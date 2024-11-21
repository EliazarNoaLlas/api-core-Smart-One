/*
 * File: merchants_error.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the modules errors.
 *
 * Last Modified: 2023-12-04
 */

package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrMerchantNotFoundCode             = "ERR_MERCHANT_NOT_FOUND"
	ErrMerchantDocumentAlreadyExistCode = "ERR_MERCHANT_DOCUMENT_ALREADY_EXIST"
	ErrMerchantIdHasBeenDeletedCode     = "ERR_MERCHANT_ID_HAS_BEEN_DELETED"
)

var (
	ErrMerchantNotFound = errDomain.NewErr().
				SetCode(ErrMerchantNotFoundCode).
				SetDescription("MERCHANT NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateMerchant")

	ErrMerchantDocumentAlreadyExist = errDomain.NewErr().
					SetCode(ErrMerchantDocumentAlreadyExistCode).
					SetDescription("DOCUMENT ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateMerchant")

	ErrMerchantIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrMerchantIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteAccount")
)
