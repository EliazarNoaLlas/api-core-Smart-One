/*
 * File: view_permissions_func_mysql_repository.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the func mysql repository of the viewPermissions.
 *
 * Last Modified: 2024-02-26
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

	ViewPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
)

//go:embed sql/get_view_permissions.sql
var QueryGetViewPermissions string

//go:embed sql/update_view_permission.sql
var QueryUpdateViewPermission string

//go:embed sql/create_view_permission.sql
var QueryCreateViewPermission string

//go:embed sql/delete_view_permission.sql
var QueryDeleteViewPermission string

func (r ViewPermissionsMySQLRepo) GetViewPermissions(
	ctx context.Context,
	viewId string,
) (
	ViewPermissionRows []ViewPermissionsDomain.ViewPermission,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetViewPermissions").SetRaw(err)
	}
	results, err := client.QueryContext(
		ctx,
		QueryGetViewPermissions,
		viewId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetViewPermissions").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	ViewPermissionTmp := make([]ViewPermission, 0)
	err = carta.Map(results, &ViewPermissionTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetViewPermissions").SetRaw(err)
	}
	automapper.Map(ViewPermissionTmp, &ViewPermissionRows)
	return ViewPermissionRows, nil
}

func (r ViewPermissionsMySQLRepo) CreateViewPermission(
	ctx context.Context,
	viewId string,
	userId string,
	ViewPermissionId string,
	body ViewPermissionsDomain.CreateViewPermissionBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateViewPermission").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateViewPermission,
		ViewPermissionId,
		viewId,
		body.PermissionId,
		userId,
		now,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateViewPermission").SetRaw(err)
	}
	lastId = &ViewPermissionId
	return
}

func (r ViewPermissionsMySQLRepo) UpdateViewPermission(
	ctx context.Context,
	viewId string,
	ViewPermissionId string,
	body ViewPermissionsDomain.UpdateViewPermissionBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateViewPermission").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateViewPermission,
		body.PermissionId,
		ViewPermissionId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateViewPermission").SetRaw(err)
	}
	return
}

func (r ViewPermissionsMySQLRepo) DeleteViewPermission(
	ctx context.Context,
	viewId string,
	ViewPermissionId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteViewPermission").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteViewPermission,
		now,
		ViewPermissionId,
		viewId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteViewPermission").SetRaw(err)
	}
	return true, nil
}
