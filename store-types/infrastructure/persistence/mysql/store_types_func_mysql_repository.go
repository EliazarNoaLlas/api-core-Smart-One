/*
 * File: store_types_func_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains functions for interacting with the repository layer related to store_types.
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

	storeTypeDomain "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain"
)

//go:embed sql/get_store_types.sql
var QueryGetStoreTypes string

//go:embed sql/get_total_store_types.sql
var QueryGetTotalStoreTypes string

//go:embed sql/update_store_type.sql
var QueryUpdateStoreType string

//go:embed sql/delete_store_type.sql
var QueryDeleteStoreType string

//go:embed sql/create_store_type.sql
var QueryCreateStoreType string

func (r storeTypeMySQLRepo) GetStoreTypes(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
) (
	storeTypesRows []storeTypeDomain.StoreType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStoreTypes").SetRaw(err)
	}
	results, err := client.QueryContext(
		ctx,
		QueryGetStoreTypes,
		sizePage,
		offset,
	)

	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStoreTypes").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	storeTypeTmp := make([]storeType, 0)
	err = carta.Map(results, &storeTypeTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStoreTypes").SetRaw(err)
	}
	automapper.Map(storeTypeTmp, &storeTypesRows)
	return storeTypesRows, nil
}

func (r storeTypeMySQLRepo) GetTotalStoreTypes(
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
		return nil, r.err.Clone().SetFunction("GetTotalStoreTypes").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalStoreTypes,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalStoreTypes").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r storeTypeMySQLRepo) CreateStoreType(
	ctx context.Context,
	storeTypeId string,
	body storeTypeDomain.CreateStoreTypeBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateStoreType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateStoreType,
		storeTypeId,
		body.Description,
		body.Abbreviation,
		now,
	)

	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateStoreType").SetRaw(err)
	}
	lastId = &storeTypeId
	return
}

func (r storeTypeMySQLRepo) UpdateStoreType(
	ctx context.Context,
	storeTypeId string,
	body storeTypeDomain.UpdateStoreTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateStoreType").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateStoreType,
		body.Description,
		body.Abbreviation,
		storeTypeId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateStoreType").SetRaw(err)
	}
	return
}

func (r storeTypeMySQLRepo) DeleteStoreType(
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
		return false, r.err.Clone().SetFunction("DeleteStoreType").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteStoreType,
		now,
		id)

	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteStoreType").SetRaw(err)
	}
	return true, nil
}
