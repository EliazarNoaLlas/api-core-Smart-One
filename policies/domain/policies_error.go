package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrPolicyNotFoundCode         = "ERR_POLICY_NOT_FOUND"
	ErrPolicyNameAlreadyExistCode = "ERR_POLICY_NAME_ALREADY_EXIST"
	ErrPolicyIdHasBeenDeletedCode = "ERR_POLICY_ID_HAS_BEEN_DELETED"
)

var (
	ErrPolicyNotFound = errDomain.NewErr().
				SetCode(ErrPolicyNotFoundCode).
				SetDescription("POLICY NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdatePolicy")

	ErrPolicyNameAlreadyExist = errDomain.NewErr().
					SetCode(ErrPolicyNameAlreadyExistCode).
					SetDescription("POLICY NAME ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreatePolicy")

	ErrPolicyIdAlreadyExist = errDomain.NewErr().
				SetCode(ErrPolicyIdHasBeenDeletedCode).
				SetDescription("ID HAS BEEN DELETED").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusConflict).
				SetLayer(errDomain.UseCase).
				SetFunction("DeletePolicy")
)
