/*
 * File: view_permissions_error.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the errors of the viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrViewPermissionNotFoundCode         = "ERR_VIEW_PERMISSION_NOT_FOUND"
	ErrViewNotFoundCode                   = "ERR_VIEW_NOT_FOUND"
	ErrPermissionNotFoundCode             = "ERR_PERMISSION_NOT_FOUND"
	ErrViewPermissionIdHasBeenDeletedCode = "ERR_VIEW_PERMISSION_ID_HAS_BEEN_DELETED"
	ErrUnknownCode                        = "ERR_UNKNOWN"
)

var (
	ErrViewPermissionNotFound = errDomain.NewErr().
					SetCode(ErrViewPermissionNotFoundCode).
					SetDescription("CLASSIFICATIONS VIEW PERMISSIONS NOT FOUND").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("UpdateViewPermission")

	ErrViewNotFound = errDomain.NewErr().
			SetCode(ErrViewNotFoundCode).
			SetDescription("VIEW NOT FOUND").
			SetLevel(errDomain.LevelError).
			SetHttpStatus(http.StatusNotFound).
			SetLayer(errDomain.UseCase).
			SetFunction("CreateViewPermission")

	ErrPermissionNotFound = errDomain.NewErr().
				SetCode(ErrPermissionNotFoundCode).
				SetDescription("PERMISSION NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("CreateViewPermission")

	ErrViewPermissionIdHasBeenDeleted = errDomain.NewErr().
						SetCode(ErrViewPermissionIdHasBeenDeletedCode).
						SetDescription("ID HAS BEEN DELETED").
						SetLevel(errDomain.LevelError).
						SetHttpStatus(http.StatusConflict).
						SetLayer(errDomain.UseCase).
						SetFunction("DeleteViewPermission")
)
