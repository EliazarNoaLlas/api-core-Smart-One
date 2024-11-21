package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrUserRoleNotFoundCode         = "ERR_USER_ROLE_NOT_FOUND"
	ErrUserHasRoleAlreadyExistCode  = "ERR_USER_ROLE_HAS_ALREADY_EXIST"
	ErrUserRoleIdHasBeenDeletedCode = "ERR_USER_ROLE_ID_HAS_BEEN_DELETED"
)

var (
	ErrUserRoleNotFound = errDomain.NewErr().
				SetCode(ErrUserRoleNotFoundCode).
				SetDescription("USER ROLE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateUserRole")

	ErrUserHasRoleAlreadyExist = errDomain.NewErr().
					SetCode(ErrUserHasRoleAlreadyExistCode).
					SetDescription("USER HAS ROLE ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateUserRole")
	ErrUserRoleIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrUserRoleIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteUserRole")
)
