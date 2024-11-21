/*
 * File: economic_activities_handler_helper_entity.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the helper entity.
 *
 * Last Modified: 2023-12-04
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	economicActivityDomain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
)

type economicActivitiesResult struct {
	Data       []economicActivityDomain.EconomicActivity `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults        `json:"pagination" binding:"required"`
	Status     int                                       `json:"status" binding:"required"`
}
