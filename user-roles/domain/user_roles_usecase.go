/*
 * File: user_roles_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the use cases to userRoles.
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type UserRoleUseCase interface {
	GetUserRolesByUser(ctx context.Context, userId string, pagination paramsDomain.PaginationParams) (
		[]UserRole, *paramsDomain.PaginationResults, error)
	CreateUserRole(ctx context.Context, userId string, body CreateUserRoleBody) (*string, error)
	UpdateUserRole(ctx context.Context, userId string, userRoleId string, body CreateUserRoleBody) error
	DeleteUserRole(ctx context.Context, userId string, userRoleId string) (bool, error)
}
