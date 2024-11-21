/*
 * File: roles_handler_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines helper entities used in the handler layer for managing roles related operations.
 *
 * Last Modified: 2023-11-14
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	rolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/roles/domain"
)

type rolesResult struct {
	Data       []rolesDomain.Role                 `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type deleteRoleResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
