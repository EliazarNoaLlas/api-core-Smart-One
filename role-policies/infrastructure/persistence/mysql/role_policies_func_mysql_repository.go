/*
 * File: role_policies_func_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for rolePolicies
 *
 * Last Modified: 2023-11-22
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

	rolePolicyDomain "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain"
)

//go:embed sql/get_role_policies.sql
var QueryGetRolePolicies string

//go:embed sql/get_total_role_policies.sql
var QueryGetTotalRolePolicies string

//go:embed sql/update_role_policy.sql
var QueryUpdateRolePolicy string

//go:embed sql/verify_role_has_policy.sql
var QueryVerifyRoleHasPolicy string

//go:embed sql/delete_role_policy.sql
var QueryDeleteRolePolicy string

//go:embed sql/create_role_policy.sql
var QueryCreateRolePolicy string

func (r rolePoliciesMySQLRepo) GetPolicies(
	ctx context.Context,
	searchParams rolePolicyDomain.GetRolePoliciesParams,
	pagination paramsDomain.PaginationParams,
) (
	rolePolicies []rolePolicyDomain.RolePolicy,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPolicies").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetRolePolicies,
			searchParams.RoleId,
			sizePage,
			offset,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPolicies").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	rolePoliciesTmp := make([]RolePolicy, 0)
	err = carta.Map(results, &rolePoliciesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPolicies").SetRaw(err)
	}
	automapper.Map(rolePoliciesTmp, &rolePolicies)
	return rolePolicies, nil
}

func (r rolePoliciesMySQLRepo) GetTotalPolicies(
	ctx context.Context,
	searchParams rolePolicyDomain.GetRolePoliciesParams,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalPolicies").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalRolePolicies,
			searchParams.RoleId,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalPolicies").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r rolePoliciesMySQLRepo) CreateRolePolicy(
	ctx context.Context,
	rolePolicyId string,
	roleId string,
	body rolePolicyDomain.CreateRolePolicyBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateRolePolicy").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateRolePolicy,
		rolePolicyId,
		body.PolicyId,
		roleId,
		body.Enable,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateRolePolicy").SetRaw(err)
	}
	lastId = &rolePolicyId
	return
}

func (r rolePoliciesMySQLRepo) CreateRolePolicies(
	ctx context.Context,
	roleId string,
	body []rolePolicyDomain.CreateMultipleRolePolicyBody,
) (
	err error,
) {
	var tx *sql.Tx
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateRolePolicies").SetRaw(err)
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
	// REVIEW
	//client, _, err := db.ClientDB(ctx)
	//if err != nil {
	//	return r.err.Clone().SetFunction("CreateRolePolicies").SetRaw(err)
	//}
	for _, rolePolicy := range body {
		_, err = client.ExecContext(ctx,
			QueryCreateRolePolicy,
			rolePolicy.Id,
			rolePolicy.PolicyId,
			roleId,
			rolePolicy.Enable,
			now)
		if err != nil {
			return r.err.Clone().SetFunction("CreateRolePolicies").SetRaw(err)
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return
}

func (r rolePoliciesMySQLRepo) VerifyRoleHasPolicy(
	ctx context.Context,
	roleId string,
	policyId string,
) (
	has bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("VerifyRoleHasPolicy").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryVerifyRoleHasPolicy,
		roleId,
		policyId,
	).Scan(&totalTmp)
	if err != nil {
		return false, r.err.Clone().SetFunction("VerifyRoleHasPolicy").SetRaw(err)
	}
	if totalTmp > 0 {
		has = true
	}
	return has, nil
}

func (r rolePoliciesMySQLRepo) UpdateRolePolicy(
	ctx context.Context,
	roleId string,
	rolePolicyId string,
	body rolePolicyDomain.UpdateRolePolicyBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRolePolicy").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateRolePolicy,
		roleId,
		body.PolicyId,
		body.Enable,
		rolePolicyId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRolePolicy").SetRaw(err)
	}
	return
}

func (r rolePoliciesMySQLRepo) DeleteRolePolicy(
	ctx context.Context,
	roleId string,
	rolePolicyId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteRolePolicy").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteRolePolicy,
		now,
		rolePolicyId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteRolePolicy").SetRaw(err)
	}
	return true, nil
}

func (r rolePoliciesMySQLRepo) DeleteRolePolicies(
	ctx context.Context,
	roleId string,
	rolePolicyIds []string,
) (
	err error,
) {
	var tx *sql.Tx
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteRolePolicies").SetRaw(err)
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
	for _, rolePolicyId := range rolePolicyIds {
		_, err = tx.ExecContext(
			ctx,
			QueryDeleteRolePolicy,
			now,
			rolePolicyId)
		if err != nil {
			return r.err.Clone().SetFunction("DeleteRolePolicies").SetRaw(err)
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
