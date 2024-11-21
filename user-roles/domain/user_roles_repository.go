/*
 * File: user_roles_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the repository to userRoles.
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type UserRoleRepository interface {
	GetUserRolesByUser(ctx context.Context, userId string, pagination paramsDomain.PaginationParams) (
		[]UserRole, error)
	GetTotalUserRolesByUser(ctx context.Context, userId string, pagination paramsDomain.PaginationParams) (
		*int, error)
	CreateUserRole(ctx context.Context, userRoleId string, userId string, body CreateUserRoleBody) (*string, error)
	VerifyUserHasRole(ctx context.Context, userId string, roleId string) (bool, error)
	UpdateUserRole(ctx context.Context, userId string, userRoleId string, body CreateUserRoleBody) error
	DeleteUserRole(ctx context.Context, userId string, userRoleId string) (bool, error)
}
