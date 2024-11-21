package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrStoreNotFoundCode         = "ERR_STORE_NOT_FOUND"
	ErrStoreNameAlreadyExistCode = "ERR_STORE_NAME_ALREADY_EXIST"
	ErrStoreIdHasBeenDeletedCode = "ERR_STORE_ID_HAS_BEEN_DELETED"
)

var (
	ErrStoreNotFound = errDomain.NewErr().
				SetCode(ErrStoreNotFoundCode).
				SetDescription("STORE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateStore")

	ErrStoreNameAlreadyExist = errDomain.NewErr().
					SetCode(ErrStoreNameAlreadyExistCode).
					SetDescription("NAME ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateStore")
	ErrStoreIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrStoreIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteStore")
)
