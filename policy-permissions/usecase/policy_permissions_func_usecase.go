/*
 * File: policyPermissions_func_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to policyPermissions.
 *
 * Last Modified: 2023-11-20
 */

package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	policyPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
)

func (u policyPermissionsUseCase) GetPolicyPermissionsByPolicy(
	ctx context.Context,
	policyId string,
	pagination paramsDomain.PaginationParams,
) (
	policyPermission []policyPermissionsDomain.PolicyPermission,
	resultPagination *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetPolicyPermissions, errGetTotalPolicyPermissions error
	var total *int
	var wg sync.WaitGroup

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_policies",
		IdColumnName:     "id",
		IdValue:          policyId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, nil, err
	}
	if !exist {
		return nil, nil, u.err.Clone().CopyCodeDescription(policyPermissionsDomain.
			ErrPolicyPermissionNotFound).SetFunction("GetPolicyPermissionByPolicy")
	}

	wg.Add(2)

	go func() {
		policyPermission, errGetPolicyPermissions = u.policyPermissionsRepository.GetPolicyPermissionsByPolicy(
			ctx, policyId, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalPolicyPermissions = u.policyPermissionsRepository.GetTotalPolicyPermissionsByPolicy(
			ctx, policyId, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetPolicyPermissions != nil {
		return nil, nil, errGetPolicyPermissions
	}
	if errGetTotalPolicyPermissions != nil {
		return nil, nil, errGetTotalPolicyPermissions
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return policyPermission, &paginationRes, nil
}

func (u policyPermissionsUseCase) CreatePolicyPermission(
	ctx context.Context,
	policyId string,
	body policyPermissionsDomain.CreatePolicyPermissionBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	policyPermissionId := uuid.New().String()
	// verify if exist the policy has permission
	policyHasPermission, err := u.policyPermissionsRepository.VerifyPolicyHasPermission(
		ctx,
		policyId,
		body.PermissionId)
	if err != nil {
		return nil, err
	}
	if policyHasPermission {
		return nil, policyPermissionsDomain.ErrPolicyHasPermissionAlreadyExist
	}
	id, err = u.policyPermissionsRepository.CreatePolicyPermission(ctx, policyId, policyPermissionId, body)
	return
}

func (u policyPermissionsUseCase) CreatePolicyPermissions(
	ctx context.Context,
	policyId string,
	body []policyPermissionsDomain.CreatePolicyPermissionBody,
) (
	ids []string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	policyPermissions := make([]policyPermissionsDomain.CreatePolicyPermissionMultipleBody, 0)
	policyPermissionIds := make([]string, 0)

	// verify if exist the policy has permission

	for _, policyPermission := range body {
		policyHasPermission, err := u.policyPermissionsRepository.VerifyPolicyHasPermission(
			ctx,
			policyId,
			policyPermission.PermissionId,
		)
		if err != nil {
			return nil, err
		}
		if policyHasPermission {
			return nil, policyPermissionsDomain.ErrPolicyHasNotPermission
		}
		policyPermissionId := uuid.New().String()
		policyPermissionIds = append(policyPermissionIds, policyPermissionId)
		policyPermissions = append(policyPermissions, policyPermissionsDomain.CreatePolicyPermissionMultipleBody{
			Id:                         policyPermissionId,
			CreatePolicyPermissionBody: policyPermission,
		})
	}
	err = u.policyPermissionsRepository.CreatePolicyPermissions(ctx, policyId, policyPermissions)
	ids = policyPermissionIds
	return
}

func (u policyPermissionsUseCase) UpdatePolicyPermission(
	ctx context.Context,
	policyId string,
	policyPermissionId string,
	body policyPermissionsDomain.CreatePolicyPermissionBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:        "core_policy_permissions",
		IdColumnName: "id",
		IdValue:      policyPermissionId,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(policyPermissionsDomain.ErrPolicyPermissionNotFound).SetFunction("UpdateUser")
	}

	err = u.policyPermissionsRepository.UpdatePolicyPermission(ctx, policyId, policyPermissionId, body)
	return
}

func (u policyPermissionsUseCase) DeletePolicyPermission(
	ctx context.Context,
	policyId string,
	policyPermissionId string,
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
		Table:            "core_policy_permissions",
		IdColumnName:     "id",
		IdValue:          policyPermissionId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, policyPermissionsDomain.ErrPolicyPermissionIdHasBeenDeleted
	}

	res, err := u.policyPermissionsRepository.DeletePolicyPermission(ctx, policyId, policyPermissionId)
	return res, err
}

func (u policyPermissionsUseCase) DeletePolicyPermissions(
	ctx context.Context,
	policyId string,
	policyPermissionIds []string,
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
	for _, policyPermissionId := range policyPermissionIds {
		recordExistsParams := validationsDomain.RecordExistsParams{
			Table:            "core_policy_permissions",
			IdColumnName:     "id",
			IdValue:          policyPermissionId,
			StatusColumnName: &deleted,
			StatusValue:      nil,
		}
		var exist bool
		exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
		if err != nil {
			return err
		}
		if !exist {
			return policyPermissionsDomain.ErrPolicyPermissionIdHasBeenDeleted
		}
	}

	err = u.policyPermissionsRepository.DeletePolicyPermissions(ctx, policyId, policyPermissionIds)
	return err
}
