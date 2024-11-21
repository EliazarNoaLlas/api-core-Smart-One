/*
 * File: permissions_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the PermissionUseCase interface, which declares methods for interacting with permissions entities.
 * It includes methods for retrieving, creating, updating, and deleting permissions data.
 *
 * Last Modified: 2023-11-15
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type PermissionUseCase interface {
	GetPermissions(ctx context.Context, moduleId string, searchParams GetPermissionsParams, pagination paramsDomain.PaginationParams) ([]Permission,
		*paramsDomain.PaginationResults, error)
	CreatePermission(ctx context.Context, moduleId string, body CreatePermissionBody) (*string, error)
	UpdatePermission(ctx context.Context, moduleId string, permissionId string, body UpdatePermissionBody) error
	DeletePermission(ctx context.Context, moduleId string, permissionId string) (bool, error)
}
