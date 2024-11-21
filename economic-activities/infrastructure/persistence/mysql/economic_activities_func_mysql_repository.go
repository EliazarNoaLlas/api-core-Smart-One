/*
 * File: economic_activities_func_mysql_repository.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the function repository.
 *
 * Last Modified: 2023-12-04
 */

package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	"gitlab.smartcitiesperu.com/smartone/api-shared/db"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	economicActivityDomain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
)

//go:embed sql/get_economic_activities.sql
var QueryGetEconomicActivities string

//go:embed sql/get_total_economic_activities.sql
var QueryGetTotalEconomicActivities string

func (r economicActivitiesMySQLRepo) GetEconomicActivities(
	ctx context.Context,
	searchParams economicActivityDomain.GetEconomicActivitiesParams,
	pagination paramsDomain.PaginationParams,
) (
	economicActivitiesRows []economicActivityDomain.EconomicActivity,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEconomicActivities").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetEconomicActivities,
			searchParams.CuuiId,
			searchParams.CuuiId,
			searchParams.Description,
			searchParams.Description,
			sizePage,
			offset,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEconomicActivities").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	economicActivitiesTmp := make([]economicActivityDomain.EconomicActivity, 0)
	err = carta.Map(results, &economicActivitiesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEconomicActivities").SetRaw(err)
	}
	automapper.Map(economicActivitiesTmp, &economicActivitiesRows)
	return economicActivitiesRows, nil
}

func (r economicActivitiesMySQLRepo) GetTotalGetEconomicActivities(ctx context.Context,
	searchParams economicActivityDomain.GetEconomicActivitiesParams,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalGetEconomicActivities").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalEconomicActivities,
			searchParams.CuuiId,
			searchParams.CuuiId,
			searchParams.Description,
			searchParams.Description,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalGetEconomicActivities").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}
