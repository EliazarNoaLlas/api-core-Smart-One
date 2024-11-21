/*
 * File: roles_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the RolesRepository interface roles data operations.
 *
 * Last Modified: 2023-11-14
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type RoleRepository interface {
	GetRoles(ctx context.Context, pagination paramsDomain.PaginationParams) ([]Role, error)
	GetTotalRoles(ctx context.Context, pagination paramsDomain.PaginationParams) (*int, error)
	CreateRole(ctx context.Context, roleId string, body CreateRoleBody) (*string, error)
	UpdateRole(ctx context.Context, roleId string, body CreateRoleBody) error
	DeleteRole(ctx context.Context, roleId string) (bool, error)
}
