/*
 * File: views_error.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * errors for views.
 *
 * Last Modified: 2023-12-11
 */

package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrViewNotFoundCode          = "ERR_VIEW_NOT_FOUND"
	ErrViewNameAlreadyExistCode  = "ERR_VIEW_NAME_ALREADY_EXIST"
	ErrViewsIdHasBeenDeletedCode = "ERR_VIEWS_ID_HAS_BEEN_DELETED"
)

var (
	ErrViewNotFound = errDomain.NewErr().
			SetCode(ErrViewNotFoundCode).
			SetDescription("VIEW NOT FOUND").
			SetLevel(errDomain.LevelError).
			SetHttpStatus(http.StatusNotFound).
			SetLayer(errDomain.UseCase).
			SetFunction("UpdateView")

	ErrViewNameAlreadyExist = errDomain.NewErr().
				SetCode(ErrViewNameAlreadyExistCode).
				SetDescription("VIEW NAME ALREADY EXIST").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusConflict).
				SetLayer(errDomain.UseCase).
				SetFunction("CreateView")
	ErrViewsIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrViewsIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteView")
)
