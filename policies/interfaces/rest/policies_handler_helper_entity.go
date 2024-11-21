/*
 * File: policies_handler_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to handler for policies.
 *
 * Last Modified: 2023-11-14
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	policiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain"
)

type policiesResult struct {
	Data       []policiesDomain.Policy            `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type deletePoliciesResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
