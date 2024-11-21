/*
 * File: merchants_error.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the modules errors.
 *
 * Last Modified: 2023-11-29
 */

package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrModuleNotFoundCode         = "ERR_MODULE_NOT_FOUND"
	ErrModuleCodeAlreadyExistCode = "ERR_MODULE_CODE_ALREADY_EXIST"
	ErrModuleIdHasBeenDeletedCode = "ERR_MODULE_ID_HAS_BEEN_DELETED"
)

var (
	ErrModuleNotFound = errDomain.NewErr().
				SetCode(ErrModuleNotFoundCode).
				SetDescription("MODULE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateModule")

	ErrModuleCodeAlreadyExist = errDomain.NewErr().
					SetCode(ErrModuleCodeAlreadyExistCode).
					SetDescription("CODE ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateModule")

	ErrModuleIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrModuleIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteModule")
)
