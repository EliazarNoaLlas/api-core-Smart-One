/*
 * File: permissions_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the PermissionsRepository interface permissions data operations.
 *
 * Last Modified: 2023-11-15
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type PermissionRepository interface {
	GetPermissions(ctx context.Context, moduleId string, searchParams GetPermissionsParams, pagination paramsDomain.PaginationParams) ([]Permission, error)
	GetTotalPermissions(ctx context.Context, moduleId string, searchParams GetPermissionsParams) (*int, error)
	CreatePermission(ctx context.Context, moduleId string, permissionId string, body CreatePermissionBody) (
		*string, error)
	UpdatePermission(ctx context.Context, moduleId string, permissionId string, body UpdatePermissionBody) error
	DeletePermission(ctx context.Context, moduleId string, permissionId string) (bool, error)
}
