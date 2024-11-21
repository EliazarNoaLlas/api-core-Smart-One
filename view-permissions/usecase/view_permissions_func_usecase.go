/*
 * File: view_permissions_func_usecase.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the use case for viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package usecase

import (
	"context"

	"github.com/google/uuid"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	ViewPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
)

func (u ViewPermissionsUseCase) GetViewPermissions(
	ctx context.Context,
	viewId string,
) (
	viewPermissions []ViewPermissionsDomain.ViewPermission,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	viewPermissions, err = u.ViewPermissionsRepository.GetViewPermissions(ctx, viewId)
	if err != nil {
		return nil, err
	}
	return viewPermissions, nil
}

func (u ViewPermissionsUseCase) CreateViewPermission(
	ctx context.Context,
	viewId string,
	userId string,
	body ViewPermissionsDomain.CreateViewPermissionBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParamsView := validationsDomain.RecordExistsParams{
		Table:            "core_views",
		IdColumnName:     "id",
		IdValue:          viewId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}

	var existView bool
	existView, err = u.validationRepository.RecordExists(ctx, recordExistsParamsView)
	if err != nil {
		return nil, err
	}
	if !existView {
		return nil, ViewPermissionsDomain.ErrViewNotFound
	}

	recordExistsParamsPermission := validationsDomain.RecordExistsParams{
		Table:            "core_permissions",
		IdColumnName:     "id",
		IdValue:          body.PermissionId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}

	var existPermission bool
	existPermission, err = u.validationRepository.RecordExists(ctx, recordExistsParamsPermission)
	if err != nil {
		return nil, err
	}
	if !existPermission {
		return nil, ViewPermissionsDomain.ErrPermissionNotFound
	}

	viewPermissionId := uuid.New().String()
	id, err = u.ViewPermissionsRepository.CreateViewPermission(ctx, viewId, userId, viewPermissionId, body)
	id = &viewPermissionId
	return
}

func (u ViewPermissionsUseCase) UpdateViewPermission(
	ctx context.Context,
	viewId string,
	ViewPermissionsId string,
	body ViewPermissionsDomain.UpdateViewPermissionBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted string
	deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_view_permissions",
		IdColumnName:     "id",
		IdValue:          ViewPermissionsId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(ViewPermissionsDomain.ErrViewPermissionNotFound).
			SetFunction("UpdateViewPermission")
	}
	err = u.ViewPermissionsRepository.UpdateViewPermission(ctx, viewId, ViewPermissionsId, body)
	return
}

func (u ViewPermissionsUseCase) DeleteViewPermission(
	ctx context.Context,
	viewId string,
	ViewPermissionsId string,
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
		Table:            "core_view_permissions",
		IdColumnName:     "id",
		IdValue:          ViewPermissionsId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, ViewPermissionsDomain.ErrViewPermissionIdHasBeenDeleted
	}
	update, err = u.ViewPermissionsRepository.DeleteViewPermission(ctx, viewId, ViewPermissionsId)
	return
}
