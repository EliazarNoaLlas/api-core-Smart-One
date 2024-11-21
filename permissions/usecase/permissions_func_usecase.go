/*
 * File: permissions_func_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the use case layer where the functions for permissions are located.
 *
 * Last Modified: 2023-11-15
 */

package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	permissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
)

func (u PermissionUseCase) GetPermissions(
	ctx context.Context,
	moduleId string,
	searchParams permissionsDomain.GetPermissionsParams,
	pagination paramsDomain.PaginationParams,
) (
	res []permissionsDomain.Permission,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetPermissions, errGetTotalPermissions error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		res, errGetPermissions = u.permissionsRepository.GetPermissions(ctx, moduleId, searchParams, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalPermissions = u.permissionsRepository.GetTotalPermissions(ctx, moduleId, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetPermissions != nil {
		return nil, nil, errGetPermissions
	}
	if errGetTotalPermissions != nil {
		return nil, nil, errGetTotalPermissions
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u PermissionUseCase) CreatePermission(
	ctx context.Context,
	moduleId string,
	body permissionsDomain.CreatePermissionBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	permissionId := uuid.New().String()
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_permissions",
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
		return nil, permissionsDomain.ErrPermissionNameAlreadyExist
	}
	id, err = u.permissionsRepository.CreatePermission(ctx, moduleId, permissionId, body)
	return

}

func (u PermissionUseCase) UpdatePermission(
	ctx context.Context,
	moduleId string,
	permissionId string,
	body permissionsDomain.UpdatePermissionBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:        "core_permissions",
		IdColumnName: "id",
		IdValue:      permissionId,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(permissionsDomain.ErrPermissionNotFound).
			SetFunction("UpdatePermission")
	}

	err = u.permissionsRepository.UpdatePermission(ctx, moduleId, permissionId, body)
	return
}

func (u PermissionUseCase) DeletePermission(
	ctx context.Context,
	moduleId string,
	permissionId string,
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
		Table:            "core_permissions",
		IdColumnName:     "id",
		IdValue:          permissionId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, permissionsDomain.ErrPermissionIdHasBeenDeleted
	}
	update, err = u.permissionsRepository.DeletePermission(ctx, moduleId, permissionId)
	return
}
