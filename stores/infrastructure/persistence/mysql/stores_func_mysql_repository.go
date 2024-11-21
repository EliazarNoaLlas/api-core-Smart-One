/*
 * File: stores_func_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for stores
 *
 * Last Modified: 2023-11-14
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

	storesDomain "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain"
)

//go:embed sql/get_stores.sql
var QueryGetStores string

//go:embed sql/get_total_stores.sql
var QueryGetTotalStores string

//go:embed sql/update_store.sql
var QueryUpdateStore string

//go:embed sql/delete_store.sql
var QueryDeleteStore string

//go:embed sql/create_store.sql
var QueryCreateStore string

func (r storesMySQLRepo) GetStores(
	ctx context.Context,
	merchantId string,
	pagination paramsDomain.PaginationParams,
) (
	storesRows []storesDomain.Store,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStores").SetRaw(err)
	}
	results, err := client.QueryContext(
		ctx,
		QueryGetStores,
		merchantId,
		sizePage,
		offset)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStores").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	storesTmp := make([]Store, 0)
	err = carta.Map(results, &storesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStores").SetRaw(err)
	}
	automapper.Map(storesTmp, &storesRows)
	return storesRows, nil
}
func (r storesMySQLRepo) GetTotalStores(
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
		return nil, r.err.Clone().SetFunction("GetTotalStores").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalStores,
			merchantId,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalStores").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r storesMySQLRepo) CreateStore(
	ctx context.Context,
	merchantId string,
	storeId string,
	body storesDomain.CreateStoreBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateStore").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateStore,
		storeId,
		body.Name,
		body.Shortname,
		merchantId,
		body.StoreTypeId,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateStore").SetRaw(err)
	}
	lastId = &storeId
	return
}

func (r storesMySQLRepo) UpdateStore(
	ctx context.Context,
	merchantId string,
	storeId string,
	body storesDomain.CreateStoreBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateStore").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateStore,
		body.Name,
		body.Shortname,
		merchantId,
		body.StoreTypeId,
		storeId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateStore").SetRaw(err)
	}
	return
}

func (r storesMySQLRepo) DeleteStore(
	ctx context.Context,
	merchantId string,
	storeId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteStore").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteStore,
		now,
		storeId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteStore").SetRaw(err)
	}
	return true, nil
}
