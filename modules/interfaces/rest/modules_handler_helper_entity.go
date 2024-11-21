/*
 * File: modules_handler_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to handler for modules.
 *
 * Last Modified: 2023-11-10
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	modulesDomain "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"
)

type modulesResult struct {
	Data       []modulesDomain.Module             `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type deleteModulesResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
