/*
 * File: policies_func_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to policies.
 *
 * Last Modified: 2023-11-14
 */

package usecase

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	policiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain"
)

func (u policiesUseCase) GetPolicies(
	ctx context.Context,
	searchParams policiesDomain.GetPoliciesParams,
	pagination paramsDomain.PaginationParams,
) (
	policies []policiesDomain.Policy,
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
		policies, errGetPolicies = u.policiesRepository.GetPolicies(ctx, searchParams, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalPolicies = u.policiesRepository.GetTotalPolicies(ctx, searchParams, pagination)
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

func (u policiesUseCase) CreatePolicy(
	ctx context.Context,
	body policiesDomain.CreatePolicyBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_policies",
		IdColumnName:     "name",
		IdValue:          body.Name,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, policiesDomain.ErrPolicyNameAlreadyExist
	}
	// validate merchant_id is required when store_id is not null
	if body.StoreId != nil && body.MerchantId == nil {
		err = u.err.Clone().SetFunction("CreatePolicy").SetRaw(errors.New("merchant_id is required"))
		return nil, err
	}
	policyId := uuid.New().String()
	id, err = u.policiesRepository.CreatePolicy(ctx, body, policyId)
	return
}

func (u policiesUseCase) UpdatePolicy(
	ctx context.Context,
	body policiesDomain.UpdatePolicyBody,
	policyId string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_policies",
		IdColumnName:     "id",
		IdValue:          policyId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(policiesDomain.ErrPolicyNotFound).SetFunction("UpdatePolicy")
	}
	// validate merchant_id is required when store_id is not null
	if body.StoreId != nil && body.MerchantId == nil {
		err = u.err.Clone().SetFunction("UpdatePolicy").SetRaw(errors.New("merchant_id is required"))
		return err
	}

	err = u.policiesRepository.UpdatePolicy(ctx, body, policyId)
	return
}

func (u policiesUseCase) DeletePolicy(
	ctx context.Context,
	policyId string,
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
		Table:            "core_policies",
		IdColumnName:     "id",
		IdValue:          policyId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, policiesDomain.ErrPolicyIdAlreadyExist
	}

	res, err := u.policiesRepository.DeletePolicy(ctx, policyId)
	return res, err
}
