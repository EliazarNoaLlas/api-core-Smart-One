/*
 * File: document_types_error.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-06
 */

package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrDocumentTypeNotFoundCode                = "ERR_DOCUMENT_TYPE_NOT_FOUND"
	ErrDocumentTypeDescriptionAlreadyExistCode = "ERR_DOCUMENT_TYPE_DESCRIPTION_ALREADY_EXIST"
	ErrDocumentTypeIdHasBeenDeletedCode        = "ERR_DOCUMENT_TYPE_ID_HAS_BEEN_DELETED"
)

var (
	ErrDocumentTypeNotFound = errDomain.NewErr().
				SetCode(ErrDocumentTypeNotFoundCode).
				SetDescription("DOCUMENT TYPE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateDocumentType")

	ErrDocumentTypeDescriptionAlreadyExist = errDomain.NewErr().
						SetCode(ErrDocumentTypeDescriptionAlreadyExistCode).
						SetDescription("DESCRIPTION ALREADY EXIST").
						SetLevel(errDomain.LevelError).
						SetHttpStatus(http.StatusConflict).
						SetLayer(errDomain.UseCase).
						SetFunction("CreateDocumentType")

	ErrDocumentTypeIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrDocumentTypeIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteDocumentType")
)
