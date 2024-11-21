/*
 * File: view_permissions_usecase.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the use case of the viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package domain

import (
	"context"
)

type ViewPermissionsUseCase interface {
	GetViewPermissions(ctx context.Context, viewId string) ([]ViewPermission, error)
	CreateViewPermission(ctx context.Context, viewId string, userId string, body CreateViewPermissionBody) (*string, error)
	UpdateViewPermission(ctx context.Context, viewId string, viewPermissionId string, body UpdateViewPermissionBody) error
	DeleteViewPermission(ctx context.Context, viewId string, viewPermissionId string) (bool, error)
}
