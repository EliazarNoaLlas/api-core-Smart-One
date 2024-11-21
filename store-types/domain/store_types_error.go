package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrStoreTypeNotFoundCode                = "ERR_STORE_TYPE_NOT_FOUND"
	ErrStoreTypeDescriptionAlreadyExistCode = "ERR_STORE_TYPE_DESCRIPTION_ALREADY_EXIST"
	ErrStoreTypeIdHasBeenDeletedCode        = "ERR_STORE_TYPE_ID_HAS_BEEN_DELETED"
)

var (
	ErrStoreTypeNotFound = errDomain.NewErr().
				SetCode(ErrStoreTypeNotFoundCode).
				SetDescription("STORE TYPE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.Infra).
				SetFunction("GetStoreTypes")

	ErrStoreTypeDescriptionAlreadyExist = errDomain.NewErr().
						SetCode(ErrStoreTypeDescriptionAlreadyExistCode).
						SetDescription("STORE TYPE DESCRIPTION ALREADY EXIST").
						SetLevel(errDomain.LevelError).
						SetHttpStatus(http.StatusConflict).
						SetLayer(errDomain.UseCase).
						SetFunction("CreateStoreType")
	ErrStoreTypeIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrStoreTypeIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteStoreType")
)
