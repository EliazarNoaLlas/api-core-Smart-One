/*
 * File: user_roles_handler_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to handler for userRoles.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	userRolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
)

type deleteUserRolesResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}

type userRolesResult struct {
	Data       []userRolesDomain.UserRole         `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}
