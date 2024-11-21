/*
 * File: economic_activities_func_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-04
 */

package usecase

import (
	"context"
	"sync"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	economicActivityDomain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
)

func (u economicActivitiesUseCase) GetEconomicActivities(
	ctx context.Context,
	searchParams economicActivityDomain.GetEconomicActivitiesParams,
	pagination paramsDomain.PaginationParams,
) (
	res []economicActivityDomain.EconomicActivity,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetEconomicActivities, errGetTotalEconomicActivities error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		res, errGetEconomicActivities = u.economicActivitiesRepository.GetEconomicActivities(ctx, searchParams,
			pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalEconomicActivities = u.economicActivitiesRepository.GetTotalGetEconomicActivities(ctx,
			searchParams, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetEconomicActivities != nil {
		return nil, nil, errGetEconomicActivities
	}
	if errGetTotalEconomicActivities != nil {
		return nil, nil, errGetTotalEconomicActivities
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}
