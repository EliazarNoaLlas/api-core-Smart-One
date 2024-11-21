/*
 * File: `roles_usecase.go`
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the RoleUseCase interface, which declares methods for interacting with roles entities.
 * It includes methods for retrieving, creating, updating, and deleting roles data.
 *
 * Last Modified: 2023-11-14
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type RoleUseCase interface {
	GetRoles(ctx context.Context, pagination paramsDomain.PaginationParams) ([]Role,
		*paramsDomain.PaginationResults, error)
	CreateRole(ctx context.Context, body CreateRoleBody) (*string, error)
	UpdateRole(ctx context.Context, roleId string, body CreateRoleBody) error
	DeleteRole(ctx context.Context, roleId string) (bool, error)
}
