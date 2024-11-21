/*
 * File: roles_func_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains functions for interacting with the repository layer related to roles.
 *
 * Last Modified: 2023-11-14
 */

package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/google/uuid"
	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	"gitlab.smartcitiesperu.com/smartone/api-shared/db"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	rolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/roles/domain"
)

//go:embed sql/get_roles.sql
var QueryGetRoles string

//go:embed sql/get_total_roles.sql
var QueryGetTotalRoles string

//go:embed sql/update_rol.sql
var QueryUpdateRole string

//go:embed sql/delete_rol.sql
var QueryDeleteRole string

//go:embed sql/create_rol.sql
var QueryCreateRole string

func (r roleMySQLRepo) GetRoles(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
) (
	rolesRows []rolesDomain.Role,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRoles").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetRoles, sizePage, offset)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRoles").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &err)
		}
	}(results)
	roleTmp := make([]RoleModel, 0)
	err = carta.Map(results, &roleTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRoles").SetRaw(err)
	}
	automapper.Map(roleTmp, &rolesRows)
	return rolesRows, nil
}

func (r roleMySQLRepo) GetTotalRoles(
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
		return nil, r.err.Clone().SetFunction("GetTotalRoles").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalRoles,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalRoles").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r roleMySQLRepo) CreateRole(
	ctx context.Context,
	roleId string,
	body rolesDomain.CreateRoleBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	roleID := uuid.New().String()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateRole").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateRole,
		roleId,
		body.Name,
		body.Description,
		body.Enable,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateRole").SetRaw(err)
	}
	lastId = &roleID
	return
}

func (r roleMySQLRepo) UpdateRole(
	ctx context.Context,
	roleId string,
	body rolesDomain.CreateRoleBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRole").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateRole,
		body.Name,
		body.Description,
		body.Enable,
		roleId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRole").SetRaw(err)
	}
	return
}

func (r roleMySQLRepo) DeleteRole(
	ctx context.Context,
	roleId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteRole").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteRole,
		now,
		roleId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteRole").SetRaw(err)
	}
	return true, nil
}
