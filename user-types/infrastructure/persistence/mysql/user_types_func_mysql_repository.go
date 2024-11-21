/*
 * File: user_types_func_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for user types
 *
 * Last Modified: 2023-11-23
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

	userTypeDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain"
)

//go:embed sql/get_user_types.sql
var QueryGetUserTypes string

//go:embed sql/get_total_user_types.sql
var QueryGetTotalUserTypes string

//go:embed sql/update_user_type.sql
var QueryUpdateUserType string

//go:embed sql/delete_user_type.sql
var QueryDeleteUserType string

//go:embed sql/create_user_type.sql
var QueryCreateUserType string

func (r userTypesMySQLRepo) GetUserTypes(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
) (
	userTypesRows []userTypeDomain.UserType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUserTypes").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetUserTypes, sizePage, offset)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUserTypes").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &err)
		}
	}(results)
	userTypesTmp := make([]userType, 0)
	err = carta.Map(results, &userTypesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUserTypes").SetRaw(err)
	}
	automapper.Map(userTypesTmp, &userTypesRows)
	return userTypesRows, nil
}

func (r userTypesMySQLRepo) GetTotalUserTypes(
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
		return nil, r.err.Clone().SetFunction("GetTotalUserTypes").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalUserTypes,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalUserTypes").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r userTypesMySQLRepo) CreateUserType(
	ctx context.Context,
	userTypeId string,
	body userTypeDomain.CreateUserTypeBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateUserType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateUserType,
		userTypeId,
		body.Description,
		body.Code,
		body.Enable,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateUserType").SetRaw(err)
	}
	lastId = &userTypeId
	return
}

func (r userTypesMySQLRepo) UpdateUserType(
	ctx context.Context,
	userTypeId string,
	body userTypeDomain.UpdateUserTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateUserType").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateUserType,
		body.Description,
		body.Code,
		body.Enable,
		userTypeId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateUserType").SetRaw(err)
	}
	return
}

func (r userTypesMySQLRepo) DeleteUserType(
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
		return false, r.err.Clone().SetFunction("DeleteUserType").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteUserType,
		now,
		id)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteUserType").SetRaw(err)
	}
	return true, nil
}
