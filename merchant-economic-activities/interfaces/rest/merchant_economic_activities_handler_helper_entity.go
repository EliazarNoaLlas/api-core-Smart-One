/*
 * File: merchant_economic_activities_handler_helper_entity.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the handler function of the merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package interfaces

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
)

type merchantEconomicActivitiesResult struct {
	Data       []domain.MerchantEconomicActivity  `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}
