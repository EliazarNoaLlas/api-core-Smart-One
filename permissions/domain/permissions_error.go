package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrPermissionNotFoundCode         = "ERR_PERMISSION_NOT_FOUND"
	ErrPermissionNameAlreadyExistCode = "ERR_PERMISSION_NAME_ALREADY_EXIST"
	ErrPermissionIdHasBeenDeletedCode = "ERR_PERMISSION_ID_HAS_BEEN_DELETED"
)

var (
	ErrPermissionNotFound = errDomain.NewErr().
				SetCode(ErrPermissionNotFoundCode).
				SetDescription("PERMISSION NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdatePermission")

	ErrPermissionNameAlreadyExist = errDomain.NewErr().
					SetCode(ErrPermissionNameAlreadyExistCode).
					SetDescription("PERMISSION NAME ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreatePermission")

	ErrPermissionIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrPermissionIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeletePermission")
)
