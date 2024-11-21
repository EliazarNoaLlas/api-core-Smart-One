/*
 * File: economic_activities_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the economic activities use case.
 *
 * Last Modified: 2023-12-04
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type EconomicActivityUseCase interface {
	GetEconomicActivities(ctx context.Context, searchParams GetEconomicActivitiesParams,
		pagination paramsDomain.PaginationParams) ([]EconomicActivity, *paramsDomain.PaginationResults, error)
}
