/*
 * File: permissions_func_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains functions for interacting with the repository layer related to permissions.
 *
 * Last Modified: 2023-11-15
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

	permissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
)

//go:embed sql/get_total_permissions.sql
var QueryGetTotalPermissions string

//go:embed sql/get_permissions.sql
var QueryGetPermissions string

//go:embed sql/update_permission.sql
var QueryUpdatePermission string

//go:embed sql/delete_permission.sql
var QueryDeletePermission string

//go:embed sql/create_permission.sql
var QueryCreatePermission string

func (r permissionMySQLRepo) GetPermissions(
	ctx context.Context,
	moduleId string,
	searchParams permissionsDomain.GetPermissionsParams,
	pagination paramsDomain.PaginationParams,
) (
	permissionsRows []permissionsDomain.Permission,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPermissions").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetPermissions,
			moduleId,
			searchParams.Code,
			searchParams.Code,
			searchParams.Name,
			searchParams.Name,
			sizePage,
			offset)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPermissions").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	permissionsTmp := make([]Permission, 0)
	err = carta.Map(results, &permissionsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPermissions").SetRaw(err)
	}
	automapper.Map(permissionsTmp, &permissionsRows)
	return permissionsRows, nil

}

func (r permissionMySQLRepo) GetTotalPermissions(
	ctx context.Context,
	moduleId string,
	searchParams permissionsDomain.GetPermissionsParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalPermissions").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalPermissions,
			moduleId,
			searchParams.Code,
			searchParams.Code,
			searchParams.Name,
			searchParams.Name).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalPermissions").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r permissionMySQLRepo) CreatePermission(
	ctx context.Context,
	moduleId string,
	permissionId string,
	body permissionsDomain.CreatePermissionBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreatePermission").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreatePermission,
		permissionId,
		body.Code,
		body.Name,
		body.Description,
		moduleId,
		now,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreatePermission").SetRaw(err)
	}
	lastId = &permissionId
	return
}

func (r permissionMySQLRepo) UpdatePermission(
	ctx context.Context,
	moduleId string,
	permissionId string,
	body permissionsDomain.UpdatePermissionBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdatePermission").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdatePermission,
		body.Code,
		body.Name,
		body.Description,
		moduleId,
		permissionId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdatePermission").SetRaw(err)
	}
	return
}

func (r permissionMySQLRepo) DeletePermission(
	ctx context.Context,
	moduleId string,
	permissionId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeletePermission").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeletePermission,
		now,
		permissionId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeletePermission").SetRaw(err)
	}
	return true, nil
}
