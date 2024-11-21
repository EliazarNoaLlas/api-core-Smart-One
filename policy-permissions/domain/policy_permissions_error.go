/*
 * File: policy_permissions_error.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the policy permissions error.
 *
 * Last Modified: 2023-11-30
 */

package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrPolicyPermissionsNotFoundCode        = "ERR_POLICY_PERMISSION_NOT_FOUND"
	ErrPolicyHasNotPermissionCode           = "ERR_POLICY_PERMISSION_ID_ALREADY_EXIST"
	ErrPolicyPermissionIdHasBeenDeletedCode = "ERR_POLICY_PERMISSION_ID_HAS_BEEN_DELETED"
	ErrPolicyHasPermissionAlreadyExistCode  = "ERR_POLICY_PERMISSION_ALREADY_EXIST"
)

var (
	ErrPolicyPermissionNotFound = errDomain.NewErr().
					SetCode(ErrPolicyPermissionsNotFoundCode).
					SetDescription("POLICY PERMISSION NOT FOUND").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("UpdatePolicyPermission")
	ErrPolicyHasPermissionAlreadyExist = errDomain.NewErr().
						SetCode(ErrPolicyHasPermissionAlreadyExistCode).
						SetDescription("POLICY HAS PERMISSION ALREADY EXIST").
						SetLevel(errDomain.LevelError).
						SetHttpStatus(http.StatusConflict).
						SetLayer(errDomain.UseCase).
						SetFunction("CreatePolicyPermission")
	ErrPolicyHasNotPermission = errDomain.NewErr().
					SetCode(ErrPolicyHasNotPermissionCode).
					SetDescription("POLICY HAS NOT THIS PERMISSION").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreatePolicyPermission")

	ErrPolicyPermissionIdHasBeenDeleted = errDomain.NewErr().
						SetCode(ErrPolicyPermissionIdHasBeenDeletedCode).
						SetDescription("ID HAS BEEN DELETED").
						SetLevel(errDomain.LevelError).
						SetHttpStatus(http.StatusConflict).
						SetLayer(errDomain.UseCase).
						SetFunction("DeletePolicyPermission")
)
