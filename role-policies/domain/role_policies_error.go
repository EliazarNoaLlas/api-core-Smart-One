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
	ErrRolePolicyNotFoundCode         = "ERR_ROLE_POLICY_NOT_FOUND"
	ErrRoleAlreadyHasThePolicyCode    = "ERR_ROLE_ALREADY_HAS_THE_POLICY"
	ErrRolePolicyIdHasBeenDeletedCode = "ERR_ROLE_POLICY_ID_HAS_BEEN_DELETED"
)

var (
	ErrRolePolicyNotFound = errDomain.NewErr().
				SetCode(ErrRolePolicyNotFoundCode).
				SetDescription("POLICY PERMISSION NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateRolePolicy")

	ErrRoleAlreadyHasThePolicy = errDomain.NewErr().
					SetCode(ErrRoleAlreadyHasThePolicyCode).
					SetDescription("THE ROLE ALREADY HAS THE POLICY").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateRolePolicy")

	ErrRolePolicyIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrRolePolicyIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteRolePolicy")
)
