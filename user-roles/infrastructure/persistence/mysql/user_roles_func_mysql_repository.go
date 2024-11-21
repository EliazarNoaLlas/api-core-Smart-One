/*
 * File: user_roles_func_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for userRoles
 *
 * Last Modified: 2023-11-23
 */

package mysql

import (
	"context"
	"database/sql"

	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	_ "embed"
	"gitlab.smartcitiesperu.com/smartone/api-shared/db"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	userRoleDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
)

//go:embed sql/get_user_roles_by_user.sql
var QueryUserRolesbyUser string

//go:embed sql/get_total_user_roles_by_user.sql
var QueryGetTotalRolesbyUser string

//go:embed sql/update_user_role.sql
var QueryUpdateUserRole string

//go:embed sql/verify_user_has_role.sql
var QueryVerifyRoleHasPolicy string

//go:embed sql/delete_user_role.sql
var QueryDeleteUserRole string

//go:embed sql/create_user_role.sql
var QueryCreateUserRole string

func (r userRolesMySQLRepo) GetUserRolesByUser(
	ctx context.Context,
	userId string,
	pagination paramsDomain.PaginationParams,
) (
	userRoleRows []userRoleDomain.UserRole,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUserRolesByUser").SetRaw(err)
	}
	results, err := client.QueryContext(
		ctx,
		QueryUserRolesbyUser,
		userId,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUserRolesByUser").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	usersTmp := make([]UserRole, 0)
	err = carta.Map(results, &usersTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUserRolesByUser").SetRaw(err)
	}
	automapper.Map(usersTmp, &userRoleRows)
	return userRoleRows, nil
}

func (r userRolesMySQLRepo) GetTotalUserRolesByUser(
	ctx context.Context,
	userId string,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalUserRolesByUser").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalRolesbyUser,
			userId,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalUserRolesByUser").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r userRolesMySQLRepo) CreateUserRole(
	ctx context.Context,
	userRoleId string,
	userId string,
	body userRoleDomain.CreateUserRoleBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateUserRole").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateUserRole,
		userRoleId,
		userId,
		body.RoleId,
		body.Enable,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateUserRole").SetRaw(err)
	}
	lastId = &userRoleId
	return
}

func (r userRolesMySQLRepo) VerifyUserHasRole(
	ctx context.Context,
	userId string,
	roleId string,
) (
	has bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("VerifyUserHasRole").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryVerifyRoleHasPolicy,
		userId,
		roleId,
	).Scan(&totalTmp)
	if err != nil {
		return false, r.err.Clone().SetFunction("VerifyUserHasRole").SetRaw(err)
	}
	if totalTmp > 0 {
		has = true
	}
	return has, nil
}

func (r userRolesMySQLRepo) UpdateUserRole(
	ctx context.Context,
	userId string,
	userRoleId string,
	body userRoleDomain.CreateUserRoleBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateUserRole").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateUserRole,
		userId,
		body.RoleId,
		body.Enable,
		userRoleId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateUserRole").SetRaw(err)
	}
	return
}

func (r userRolesMySQLRepo) DeleteUserRole(
	ctx context.Context,
	userId string,
	userRoleId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteUserRole").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteUserRole,
		now,
		userRoleId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteUserRole").SetRaw(err)
	}
	return true, nil
}
