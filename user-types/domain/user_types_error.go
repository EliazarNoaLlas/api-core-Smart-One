package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrUserTypeNotFoundCode                = "ERR_USER_TYPE_NOT_FOUND"
	ErrUserTypeDescriptionAlreadyExistCode = "ERR_USER_TYPE_DESCRIPTION_ALREADY_EXIST"
	ErrUserTypeIdHasBeenDeletedCode        = "ERR_USER_TYPE_ID_HAS_BEEN_DELETED"
)

var (
	ErrUserTypeNotFound = errDomain.NewErr().
				SetCode(ErrUserTypeNotFoundCode).
				SetDescription("USER TYPE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateUserType")

	ErrUserTypeDescriptionAlreadyExist = errDomain.NewErr().
						SetCode(ErrUserTypeDescriptionAlreadyExistCode).
						SetDescription("DESCRIPTION ALREADY EXIST").
						SetLevel(errDomain.LevelError).
						SetHttpStatus(http.StatusConflict).
						SetLayer(errDomain.UseCase).
						SetFunction("CreateUserType")
	ErrUserTypeIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrUserTypeIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteUserType")
)
