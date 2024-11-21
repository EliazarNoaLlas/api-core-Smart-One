/*
 * File: policies_func_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for policies
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

	policiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain"
)

//go:embed sql/get_policies.sql
var QueryGetPolicies string

//go:embed sql/get_total_policies.sql
var QueryGetTotalPolicies string

//go:embed sql/update_policy.sql
var QueryUpdatePolicy string

//go:embed sql/delete_policy.sql
var QueryDeletePolicy string

//go:embed sql/create_policy.sql
var QueryCreatePolicy string

func (r policiesMySQLRepo) GetPolicies(
	ctx context.Context,
	searchParams policiesDomain.GetPoliciesParams,
	pagination paramsDomain.PaginationParams,
) (
	policiesRows []policiesDomain.Policy,
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
			QueryGetPolicies,
			searchParams.ModuleId,
			searchParams.ModuleId,
			searchParams.MerchantId,
			searchParams.MerchantId,
			searchParams.StoreId,
			searchParams.StoreId,
			searchParams.Description,
			searchParams.Description,
			searchParams.Description,
			sizePage, offset,
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
	policiesTmp := make([]Policy, 0)
	err = carta.Map(results, &policiesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPolicies").SetRaw(err)
	}
	automapper.Map(policiesTmp, &policiesRows)
	for iPolicy, policy := range policiesRows {
		if policy.Module.Id == nil {
			policiesRows[iPolicy].Module = nil
		}
		if policy.Merchant.Id == nil {
			policiesRows[iPolicy].Merchant = nil
		}
		if policy.Store.Id == nil {
			policiesRows[iPolicy].Store = nil
		}
	}
	return policiesRows, nil
}

func (r policiesMySQLRepo) GetTotalPolicies(
	ctx context.Context,
	searchParams policiesDomain.GetPoliciesParams,
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
			QueryGetTotalPolicies,
			searchParams.ModuleId,
			searchParams.ModuleId,
			searchParams.MerchantId,
			searchParams.MerchantId,
			searchParams.StoreId,
			searchParams.StoreId,
			searchParams.Description,
			searchParams.Description,
			searchParams.Description,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalPolicies").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r policiesMySQLRepo) CreatePolicy(
	ctx context.Context,
	body policiesDomain.CreatePolicyBody,
	policyId string,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreatePolicy").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreatePolicy,
		policyId,
		body.Name,
		body.Description,
		body.ModuleId,
		body.MerchantId,
		body.StoreId,
		body.Level,
		body.Enable,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreatePolicy").SetRaw(err)
	}
	lastId = &policyId
	return
}

func (r policiesMySQLRepo) UpdatePolicy(
	ctx context.Context,
	body policiesDomain.UpdatePolicyBody,
	policyId string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdatePolicy").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdatePolicy,
		body.Name,
		body.Description,
		body.ModuleId,
		body.MerchantId,
		body.StoreId,
		body.Level,
		body.Enable,
		policyId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdatePolicy").SetRaw(err)
	}
	return
}

func (r policiesMySQLRepo) DeletePolicy(
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
		return false, r.err.Clone().SetFunction("DeletePolicy").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeletePolicy,
		now,
		id)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeletePolicy").SetRaw(err)
	}
	return true, nil
}
