/*
 * File: merchants_handler_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines helper entities used in the handler layer for managing merchants related operations.
 *
 * Last Modified: 2023-11-10
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	merchantsDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
)

type merchantsResult struct {
	Data       []merchantsDomain.Merchant         `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type deleteMerchantsResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
