/*
 * File: views_func_mysql_repository.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for rolePolicies
 *
 * Last Modified: 2023-11-24
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

	viewDomain "gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
)

//go:embed sql/get_views.sql
var QueryGetViews string

//go:embed sql/get_total_views.sql
var QueryGetTotalViews string

//go:embed sql/create_view.sql
var QueryCreateView string

//go:embed sql/update_view.sql
var QueryUpdateView string

//go:embed sql/delete_view.sql
var QueryDeleteView string

func (r viewMySqlRepo) GetViews(
	ctx context.Context,
	moduleId string,
	searchParams viewDomain.GetViewsParams,
	pagination paramsDomain.PaginationParams,
) (
	views []viewDomain.View,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetViews").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetViews,
			moduleId,
			searchParams.Name,
			searchParams.Name,
			searchParams.Name,
			sizePage,
			offset,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetViews").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	viewsTmp := make([]View, 0)
	err = carta.Map(results, &viewsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetViews").SetRaw(err)
	}
	automapper.Map(viewsTmp, &views)
	return views, nil
}

func (r viewMySqlRepo) GetTotalViews(
	ctx context.Context,
	moduleId string,
	searchParams viewDomain.GetViewsParams,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalViews").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalViews,
			moduleId,
			searchParams.Name,
			searchParams.Name,
			searchParams.Name,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalViews").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r viewMySqlRepo) CreateView(
	ctx context.Context,
	moduleId string,
	viewId string,
	body viewDomain.CreateViewBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateView").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryCreateView,
		viewId,
		body.Name,
		body.Description,
		body.Url,
		body.Icon,
		moduleId,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateView").SetRaw(err)
	}
	lastId = &viewId
	return
}

func (r viewMySqlRepo) UpdateView(
	ctx context.Context,
	moduleId string,
	viewId string,
	body viewDomain.UpdateViewBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateView").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateView,
		body.Name,
		body.Description,
		body.Url,
		body.Icon,
		moduleId,
		viewId)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateView").SetRaw(err)
	}
	return
}

func (r viewMySqlRepo) DeleteView(
	ctx context.Context,
	moduleId string,
	viewId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteView").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteView,
		now,
		viewId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteView").SetRaw(err)
	}
	return true, nil
}
