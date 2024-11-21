/*
 * File: merchant_economic_activities_func_mysql_repository.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the function repository of the merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package infrastructure

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	"gitlab.smartcitiesperu.com/smartone/api-shared/db"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	merchantEconomicActivitiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
)

//go:embed sql/get_merchant_economic_activities.sql
var QueryGetMerchantEconomicActivities string

//go:embed sql/get_total_merchant_economic_activities.sql
var QueryGetTotalMerchantEconomicActivities string

//go:embed sql/create_merchant_economic_activities.sql
var QueryCreateMerchantEconomicActivity string

//go:embed sql/update_merchant_economic_activities.sql
var QueryUpdateMerchantEconomicActivity string

//go:embed sql/delete_merchant_economic_activities.sql
var QueryDeleteMerchantEconomicActivity string

func (r merchantEconomicActivitiesMySQLRepo) GetMerchantEconomicActivities(
	ctx context.Context,
	merchantId string,
	pagination paramsDomain.PaginationParams,
) (
	merchantEconomicActivitiesRows []merchantEconomicActivitiesDomain.MerchantEconomicActivity,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMerchantEconomicActivities").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetMerchantEconomicActivities,
			merchantId,
			sizePage,
			offset,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMerchantEconomicActivities").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	merchantEconomicActivitiesTmp := make([]MerchantActivity, 0)
	err = carta.Map(results, &merchantEconomicActivitiesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction(
			"GetMerchantEconomicActivities").SetRaw(err)
	}
	automapper.Map(merchantEconomicActivitiesTmp, &merchantEconomicActivitiesRows)
	return merchantEconomicActivitiesRows, nil
}

func (r merchantEconomicActivitiesMySQLRepo) GetTotalEconomicActivities(
	ctx context.Context,
	merchantId string,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalEconomicActivities").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalMerchantEconomicActivities,
			merchantId,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalEconomicActivities").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r merchantEconomicActivitiesMySQLRepo) CreateEconomicActivity(
	ctx context.Context,
	merchantEconomicActivityId string,
	body merchantEconomicActivitiesDomain.CreateMerchantEconomicActivityBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateEconomicActivity").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateMerchantEconomicActivity,
		merchantEconomicActivityId,
		body.MerchantId,
		body.EconomicActivityId,
		body.Sequence,
		now,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateEconomicActivity").SetRaw(err)
	}
	lastId = &merchantEconomicActivityId
	return
}

func (r merchantEconomicActivitiesMySQLRepo) UpdateEconomicActivity(
	ctx context.Context,
	merchantEconomicActivityId string,
	body merchantEconomicActivitiesDomain.UpdateMerchantEconomicActivityBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateEconomicActivity").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateMerchantEconomicActivity,
		body.MerchantId,
		body.EconomicActivityId,
		body.Sequence,
		merchantEconomicActivityId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateEconomicActivity").SetRaw(err)
	}
	return
}

func (r merchantEconomicActivitiesMySQLRepo) DeleteEconomicActivity(
	ctx context.Context,
	merchantEconomicActivityId string,
) (
	update bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteEconomicActivity").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteMerchantEconomicActivity,
		now,
		merchantEconomicActivityId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteEconomicActivity").SetRaw(err)
	}
	return true, nil
}
