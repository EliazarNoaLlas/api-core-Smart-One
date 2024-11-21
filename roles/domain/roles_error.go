package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrRoleNotFoundCode             = "ERR_ROLE_NOT_FOUND"
	ErrRoleRoleNameAlreadyExistCode = "ERR_ROLE_NAME_ALREADY_EXIST"
	ErrRoleIdHasBeenDeletedCode     = "ERR_ROLE_ID_HAS_BEEN_DELETED"
)

var (
	ErrRoleNotFound = errDomain.NewErr().
			SetCode(ErrRoleNotFoundCode).
			SetDescription("ROLE NOT FOUND").
			SetLevel(errDomain.LevelError).
			SetHttpStatus(http.StatusNotFound).
			SetLayer(errDomain.UseCase).
			SetFunction("UpdateRole")
	ErrRoleNameAlreadyExist = errDomain.NewErr().
				SetCode(ErrRoleRoleNameAlreadyExistCode).
				SetDescription("ROLE NAME ALREADY EXIST").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusConflict).
				SetLayer(errDomain.UseCase).
				SetFunction("CreateRole")
	ErrRoleIdHasBeenDeleted = errDomain.NewErr().
				SetCode(ErrRoleIdHasBeenDeletedCode).
				SetDescription("ID HAS BEEN DELETED").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusConflict).
				SetLayer(errDomain.UseCase).
				SetFunction("DeleteRole")
)
