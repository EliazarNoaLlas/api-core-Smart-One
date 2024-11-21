/*
 * File: permissions_handler_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines helper entities used in the handler layer for managing permissions related operations.
 *
 * Last Modified: 2023-11-15
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	permissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
)

type permissionsResult struct {
	Data       []permissionsDomain.Permission     `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type deletePermissionResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
