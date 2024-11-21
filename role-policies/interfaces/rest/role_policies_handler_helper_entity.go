/*
 * File: role_policies_handler_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to handler for rolePolicies.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	rolePoliciesDomain "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain"
)

type deleteRolePoliciesResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}

type deleteMultipleRolePoliciesValidate struct {
	RolePolicyIds []string `json:"role_policy_ids" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
}

type rolePoliciesResult struct {
	Data       []rolePoliciesDomain.RolePolicy    `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}
