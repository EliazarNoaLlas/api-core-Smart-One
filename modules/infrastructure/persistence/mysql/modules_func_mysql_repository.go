/*
 * File: modules_func_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for modules
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

	moduleDomain "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"
)

//go:embed sql/get_modules.sql
var QueryGetModules string

//go:embed sql/get_total_modules.sql
var QueryGetTotalModules string

//go:embed sql/update_module.sql
var QueryUpdateModule string

//go:embed sql/delete_module.sql
var QueryDeleteModule string

//go:embed sql/create_module.sql
var QueryCreateModule string

func (r modulesMySQLRepo) GetModules(
	ctx context.Context,
	searchParams moduleDomain.GetModulesParams,
	pagination paramsDomain.PaginationParams,
) (
	modulesRows []moduleDomain.Module,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetModules").SetRaw(err)
	}
	results, err := client.QueryContext(
		ctx,
		QueryGetModules,
		searchParams.Code,
		searchParams.Code,
		searchParams.Name,
		searchParams.Name,
		sizePage,
		offset)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetModules").SetRaw(err)
	}

	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &err)
		}
	}(results)

	modulesTmp := make([]module, 0)
	err = carta.Map(results, &modulesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetModules").SetRaw(err)
	}
	automapper.Map(modulesTmp, &modulesRows)
	return modulesRows, nil
}

func (r modulesMySQLRepo) GetTotalModules(
	ctx context.Context,
	searchParams moduleDomain.GetModulesParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalModules").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalModules,
			searchParams.Code,
			searchParams.Code,
			searchParams.Name,
			searchParams.Name,
		).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalModules").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r modulesMySQLRepo) CreateModule(
	ctx context.Context,
	moduleId string,
	body moduleDomain.CreateModuleBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateModule").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateModule,
		moduleId,
		body.Name,
		body.Description,
		body.Code,
		body.Icon,
		body.Position,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateModule").SetRaw(err)
	}
	lastId = &moduleId
	return
}

func (r modulesMySQLRepo) UpdateModule(
	ctx context.Context,
	moduleId string,
	body moduleDomain.UpdateModuleBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateModule").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateModule,
		body.Name,
		body.Description,
		body.Code,
		body.Icon,
		body.Position,
		moduleId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateModule").SetRaw(err)
	}
	return
}

func (r modulesMySQLRepo) DeleteModule(
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
		return false, r.err.Clone().SetFunction("DeleteModule").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteModule,
		now,
		id)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteModule").SetRaw(err)
	}
	return true, nil
}
