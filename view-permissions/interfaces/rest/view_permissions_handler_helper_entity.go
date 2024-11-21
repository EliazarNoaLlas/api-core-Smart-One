/*
 * File: view_permissions_handler_helper_entity.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities helper to handler for viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	ViewPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
)

type viewPermissionsResult struct {
	Data       []ViewPermissionsDomain.ViewPermission `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults     `json:"pagination" binding:"required"`
	Status     int                                    `json:"status" binding:"required"`
}

type deleteViewPermissionsResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
