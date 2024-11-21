/*
 * File: policyPermissions_func_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for policyPermissions
 *
 * Last Modified: 2023-11-20
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

	policyPermissionDomain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
)

//go:embed sql/get_policy_permissions_by_policy.sql
var QueryGetPermissionsByPolicy string

//go:embed sql/get_total_policy_permissions_by_policy.sql
var QueryTotalPermissionByPolicy string

//go:embed sql/update_policy_permission.sql
var QueryUpdatePolicyPermission string

//go:embed sql/verify_policy_has_permission.sql
var QueryVerifyPolicyHasPermission string

//go:embed sql/delete_policy_permission.sql
var QueryDeletePolicyPermission string

//go:embed sql/create_policy_permission.sql
var QueryCreatePolicyPermission string

func (r policyPermissionsMySQLRepo) GetPolicyPermissionsByPolicy(
	ctx context.Context,
	policyId string,
	pagination paramsDomain.PaginationParams,
) (
	policyPermissionRows []policyPermissionDomain.PolicyPermission,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPolicyPermissionsByPolicy").SetRaw(err)
	}
	results, err := client.QueryContext(
		ctx,
		QueryGetPermissionsByPolicy,
		policyId,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPolicyPermissionsByPolicy").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	usersTmp := make([]PermissionPolicy, 0)
	err = carta.Map(results, &usersTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPolicyPermissionsByPolicy").SetRaw(err)
	}
	automapper.Map(usersTmp, &policyPermissionRows)
	return policyPermissionRows, nil
}

func (r policyPermissionsMySQLRepo) GetTotalPolicyPermissionsByPolicy(
	ctx context.Context,
	policyId string,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalPermissionByPolicy").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryTotalPermissionByPolicy,
			policyId,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalPermissionByPolicy").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r policyPermissionsMySQLRepo) CreatePolicyPermission(
	ctx context.Context,
	policyId string,
	policyPermissionId string,
	body policyPermissionDomain.CreatePolicyPermissionBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreatePolicyPermission").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreatePolicyPermission,
		policyPermissionId,
		policyId,
		body.PermissionId,
		body.Enable,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreatePolicyPermission").SetRaw(err)
	}
	lastId = &policyPermissionId
	return
}

func (r policyPermissionsMySQLRepo) CreatePolicyPermissions(
	ctx context.Context,
	policyId string,
	body []policyPermissionDomain.CreatePolicyPermissionMultipleBody,
) (
	err error,
) {
	var tx *sql.Tx
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreatePolicyPermissions").SetRaw(err)
	}
	tx, err = client.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	for _, policyPermission := range body {
		now := r.clock.Now().Format("2006-01-02 15:04:05")
		_, err = tx.ExecContext(ctx,
			QueryCreatePolicyPermission,
			policyPermission.Id,
			policyId,
			policyPermission.PermissionId,
			policyPermission.Enable,
			now)
		if err != nil {
			return r.err.Clone().SetFunction("CreatePolicyPermissions").SetRaw(err)
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return
}

func (r policyPermissionsMySQLRepo) VerifyPolicyHasPermission(
	ctx context.Context,
	policyId string,
	permissionId string,
) (
	has bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("VerifyPolicyHasPermission").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryVerifyPolicyHasPermission,
			policyId,
			permissionId).
		Scan(&totalTmp)
	if err != nil {
		return false, r.err.Clone().SetFunction("VerifyPolicyHasPermission").SetRaw(err)
	}
	if totalTmp > 0 {
		has = true
	}
	return has, nil
}

func (r policyPermissionsMySQLRepo) UpdatePolicyPermission(
	ctx context.Context,
	policyId string,
	policyPermissionId string,
	body policyPermissionDomain.CreatePolicyPermissionBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdatePolicyPermission").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdatePolicyPermission,
		policyId,
		body.PermissionId,
		body.Enable,
		policyPermissionId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdatePolicyPermission").SetRaw(err)
	}
	return
}

func (r policyPermissionsMySQLRepo) DeletePolicyPermission(
	ctx context.Context,
	policyId string,
	policyPermissionId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeletePolicyPermission").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeletePolicyPermission,
		now,
		policyPermissionId,
		policyId,
	)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeletePolicyPermission").SetRaw(err)
	}
	return true, nil
}

func (r policyPermissionsMySQLRepo) DeletePolicyPermissions(
	ctx context.Context,
	policyId string,
	policyPermissionIds []string,
) (
	err error,
) {
	var tx *sql.Tx
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeletePolicyPermissions").SetRaw(err)
	}
	tx, err = client.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	for _, policyPermissionId := range policyPermissionIds {
		_, err = tx.ExecContext(
			ctx,
			QueryDeletePolicyPermission,
			now,
			policyPermissionId,
			policyId,
		)
		if err != nil {
			return r.err.Clone().SetFunction("DeletePolicyPermissions").SetRaw(err)
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return
}
