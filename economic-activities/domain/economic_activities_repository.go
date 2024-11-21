/*
 * File: economic_activities_repository.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository economic activities.
 *
 * Last Modified: 2023-12-04
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type EconomicActivityRepository interface {
	GetEconomicActivities(ctx context.Context, searchParams GetEconomicActivitiesParams,
		pagination paramsDomain.PaginationParams) ([]EconomicActivity, error)
	GetTotalGetEconomicActivities(ctx context.Context, searchParams GetEconomicActivitiesParams,
		pagination paramsDomain.PaginationParams) (*int, error)
}
