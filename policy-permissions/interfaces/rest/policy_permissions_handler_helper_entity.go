/*
 * File: policyPermissions_handler_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to handler for policyPermissions.
 *
 * Last Modified: 2023-11-20
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	policyPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
)

type policyPermissionsResult struct {
	Data       []policyPermissionsDomain.PolicyPermission `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults         `json:"pagination" binding:"required"`
	Status     int                                        `json:"status" binding:"required"`
}

type deletePolicyPermissionsResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
