/*
 * File: merchants_func_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains functions for interacting with the repository layer related to merchants.
 *
 * Last Modified: 2023-11-10
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

	merchantDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
)

//go:embed sql/get_total_merchants.sql
var QueryGetTotalMerchants string

//go:embed sql/get_merchants.sql
var QueryGetMerchants string

//go:embed sql/update_merchant.sql
var QueryUpdateMerchant string

//go:embed sql/delete_merchant.sql
var QueryDeleteMerchant string

//go:embed sql/create_merchant.sql
var QueryCreateMerchant string

func (r merchantsMySQLRepo) GetMerchants(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
) (
	merchantsRows []merchantDomain.Merchant,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMerchants").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetMerchants, sizePage, offset)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMerchants").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	merchantTmp := make([]merchantModel, 0)
	err = carta.Map(results, &merchantTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMerchants").SetRaw(err)
	}
	automapper.Map(merchantTmp, &merchantsRows)
	return merchantsRows, nil
}

func (r merchantsMySQLRepo) GetTotalMerchants(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalMerchants").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalMerchants).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalMerchants").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r merchantsMySQLRepo) CreateMerchant(
	ctx context.Context,
	merchantId string,
	body merchantDomain.CreateMerchantBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateMerchant").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateMerchant,
		merchantId,
		body.Name,
		body.Description,
		body.Phone,
		body.Document,
		body.Address,
		body.Industry,
		body.ImagePath,
		now)

	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateMerchant").SetRaw(err)
	}
	lastId = &merchantId
	return
}

func (r merchantsMySQLRepo) UpdateMerchant(
	ctx context.Context,
	merchantId string,
	body merchantDomain.UpdateMerchantBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateMerchant").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateMerchant,
		body.Name,
		body.Description,
		body.Phone,
		body.Document,
		body.Address,
		body.Industry,
		body.ImagePath,
		merchantId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateMerchant").SetRaw(err)
	}
	return
}

func (r merchantsMySQLRepo) DeleteMerchant(
	ctx context.Context,
	id string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteMerchant").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteMerchant,
		now,
		id)

	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteMerchant").SetRaw(err)
	}
	return true, nil
}
