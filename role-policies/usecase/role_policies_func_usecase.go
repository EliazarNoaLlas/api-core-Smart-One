/*
 * File: role_policies_func_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to rolePolicies.
 *
 * Last Modified: 2023-11-23
 */

package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	rolePoliciesDomain "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain"
)

func (u rolePoliciesUseCase) GetPolicies(
	ctx context.Context,
	searchParams rolePoliciesDomain.GetRolePoliciesParams,
	pagination paramsDomain.PaginationParams,
) (
	policies []rolePoliciesDomain.RolePolicy,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetPolicies, errGetTotalPolicies error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		policies, errGetPolicies = u.rolePoliciesRepository.GetPolicies(ctx, searchParams, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalPolicies = u.rolePoliciesRepository.GetTotalPolicies(ctx, searchParams, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetPolicies != nil {
		return nil, nil, errGetPolicies
	}
	if errGetTotalPolicies != nil {
		return nil, nil, errGetTotalPolicies
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return policies, &paginationRes, nil
}

func (u rolePoliciesUseCase) CreateRolePolicy(
	ctx context.Context,
	roleId string,
	body rolePoliciesDomain.CreateRolePolicyBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	rolePolicyID := uuid.New().String()
	// verify if exist the policy has permission
	existRolePolicy, err := u.rolePoliciesRepository.VerifyRoleHasPolicy(ctx, roleId, body.PolicyId)
	if err != nil {
		return nil, err
	}
	if existRolePolicy {
		return nil, rolePoliciesDomain.ErrRoleAlreadyHasThePolicy
	}
	id, err = u.rolePoliciesRepository.CreateRolePolicy(ctx, rolePolicyID, roleId, body)
	return
}

func (u rolePoliciesUseCase) CreateRolePolicies(
	ctx context.Context,
	roleId string,
	body []rolePoliciesDomain.CreateRolePolicyBody,
) (
	ids []string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	rolePolicies := make([]rolePoliciesDomain.CreateMultipleRolePolicyBody, len(body))
	rolePolicyIds := make([]string, len(body))
	// verify if exist the role has policy
	for iRolePolicy, rolePolicy := range body {
		existRolePolicy, err := u.rolePoliciesRepository.VerifyRoleHasPolicy(ctx, roleId, rolePolicy.PolicyId)
		if err != nil {
			return nil, err
		}
		if existRolePolicy {
			return nil, rolePoliciesDomain.ErrRoleAlreadyHasThePolicy
		}
		rolePolicyId := uuid.New().String()
		rolePolicyIds[iRolePolicy] = rolePolicyId
		rolePolicies[iRolePolicy] = rolePoliciesDomain.CreateMultipleRolePolicyBody{
			Id:                   rolePolicyId,
			CreateRolePolicyBody: rolePolicy,
		}
	}
	err = u.rolePoliciesRepository.CreateRolePolicies(ctx, roleId, rolePolicies)
	ids = rolePolicyIds
	return
}

func (u rolePoliciesUseCase) UpdateRolePolicy(
	ctx context.Context,
	roleId string,
	rolePolicyId string,
	body rolePoliciesDomain.UpdateRolePolicyBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:        "core_role_policies",
		IdColumnName: "id",
		IdValue:      rolePolicyId,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(rolePoliciesDomain.ErrRolePolicyNotFound).SetFunction("UpdateUser")
	}

	err = u.rolePoliciesRepository.UpdateRolePolicy(ctx, roleId, rolePolicyId, body)
	return
}

func (u rolePoliciesUseCase) DeleteRolePolicy(
	ctx context.Context,
	roleId string,
	rolePolicyId string,
) (
	update bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	var deleted string
	deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_role_policies",
		IdColumnName:     "id",
		IdValue:          rolePolicyId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, rolePoliciesDomain.ErrRolePolicyIdHasBeenDeleted
	}

	res, err := u.rolePoliciesRepository.DeleteRolePolicy(ctx, roleId, rolePolicyId)
	return res, err
}

func (u rolePoliciesUseCase) DeleteRolePolicies(
	ctx context.Context,
	roleId string,
	rolePolicyIds []string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	var deleted string
	deleted = "deleted_at"
	// ensure that something cannot be deleted if it's already deleted or doesn't exist.
	for _, rolePolicyId := range rolePolicyIds {
		recordExistsParams := validationsDomain.RecordExistsParams{
			Table:            "core_role_policies",
			IdColumnName:     "id",
			IdValue:          rolePolicyId,
			StatusColumnName: &deleted,
			StatusValue:      nil,
		}
		var exist bool
		exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
		if err != nil {
			return err
		}
		if !exist {
			return rolePoliciesDomain.ErrRolePolicyIdHasBeenDeleted
		}
	}

	err = u.rolePoliciesRepository.DeleteRolePolicies(ctx, roleId, rolePolicyIds)
	return err
}
