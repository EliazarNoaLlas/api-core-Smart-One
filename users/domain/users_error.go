/*
 * File: users_error.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * errors of users.
 *
 * Last Modified: 2023-12-11
 */

package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrUserNotFoundCode                 = "ERR_USER_NOT_FOUND"
	ErrUserUsernameAlreadyExistCode     = "ERR_USER_USERNAME_ALREADY_EXIST"
	ErrUserIdHasBeenDeletedCode         = "ERR_USER_ID_HAS_BEEN_DELETED"
	ErrPersonIdNotExistCode             = "ERR_PERSON_ID_NOT_EXIST"
	ErrDocumentOfPersonAlreadyExistCode = "ERR_DOCUMENT_OF_PERSON_ALREADY_EXIST"
	ErrUserIdAppearsMoreThanOnceCode    = "ERR_USER_ID_APPEARS_MORE_THAN_ONCE"
	ErrUserNotExistCode                 = "ERR_USER_NOT_EXIST"
	ErrUserIdAlreadyExistCode           = "ERR_USER_ID_ALREADY_EXIST"
	ErrStoreIdEmptyCode                 = "ERR_STORE_ID_EMPTY"
	ErrInvalidCodeModuleCode            = "ENTER A VALID CODE OF MODULE"
)

var (
	ErrUserNotFound = errDomain.NewErr().
			SetCode(ErrUserNotFoundCode).
			SetDescription("USER NOT FOUND").
			SetLevel(errDomain.LevelError).
			SetHttpStatus(http.StatusNotFound).
			SetLayer(errDomain.Infra).
			SetFunction("GetUser")

	ErrUserUsernameAlreadyExist = errDomain.NewErr().
					SetCode(ErrUserUsernameAlreadyExistCode).
					SetDescription("USERNAME ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateUser")
	ErrUserIdHasBeenDeleted = errDomain.NewErr().
				SetCode(ErrUserIdHasBeenDeletedCode).
				SetDescription("ID HAS BEEN DELETED").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusConflict).
				SetLayer(errDomain.UseCase).
				SetFunction("DeleteUser")

	ErrPersonIdNotExist = errDomain.NewErr().
				SetCode(ErrPersonIdNotExistCode).
				SetDescription("THE PERSON ID NOT EXIST").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusConflict).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateUser")

	ErrUserIdAlreadyExist = errDomain.NewErr().
				SetCode(ErrUserIdAlreadyExistCode).
				SetDescription("THE USER ID ALREADY EXIST").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusConflict).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateUser")

	ErrDocumentOfPersonAlreadyExist = errDomain.NewErr().
					SetCode(ErrDocumentOfPersonAlreadyExistCode).
					SetDescription("THE DOCUMENT OF THE PERSON ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateUserMain")

	ErrUserIdAppearsMoreThanOnce = errDomain.NewErr().
					SetCode(ErrUserIdAppearsMoreThanOnceCode).
					SetDescription("THE USER ID APPEARS MORE THAN ONCE").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.Infra).
					SetFunction("ValidateUniqueUserExistence")

	ErrUserNotExist = errDomain.NewErr().
			SetCode(ErrUserNotExistCode).
			SetDescription("THE USER NOT EXIST").
			SetLevel(errDomain.LevelError).
			SetHttpStatus(http.StatusConflict).
			SetLayer(errDomain.Infra).
			SetFunction("GetUserById")

	ErrStoreIdEmpty = errDomain.NewErr().
			SetCode(ErrStoreIdEmptyCode).
			SetDescription("ENTER A VALID ID FOR A STORE").
			SetLevel(errDomain.LevelError).
			SetHttpStatus(http.StatusConflict).
			SetLayer(errDomain.UseCase).
			SetFunction("VerifyPermissionsByUser")

	ErrInvalidCodeModule = errDomain.NewErr().
				SetCode(ErrInvalidCodeModuleCode).
				SetDescription("ENTER A VALID CODE OF MODULE").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusConflict).
				SetLayer(errDomain.UseCase).
				SetFunction("GetModulePermissions")
)
