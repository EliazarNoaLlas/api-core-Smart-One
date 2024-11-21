/*
 * File: user_roles_func_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to userRoles.
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

	userRolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
)

func (u userRolesUseCase) GetUserRolesByUser(
	ctx context.Context,
	userId string,
	pagination paramsDomain.PaginationParams,
) (
	res []userRolesDomain.UserRole,
	resultPagination *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetUserRolesByUser, errGetTotalUserRolesByUser error
	var total *int
	var wg sync.WaitGroup

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_users",
		IdColumnName:     "id",
		IdValue:          userId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, nil, err
	}
	if !exist {
		return nil, nil, u.err.Clone().CopyCodeDescription(userRolesDomain.
			ErrUserRoleNotFound).SetFunction("GetUserRolesByUsers")
	}

	wg.Add(2)

	go func() {
		res, errGetUserRolesByUser = u.userRolesRepository.GetUserRolesByUser(
			ctx, userId, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalUserRolesByUser = u.userRolesRepository.GetTotalUserRolesByUser(
			ctx, userId, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetUserRolesByUser != nil {
		return nil, nil, errGetUserRolesByUser
	}
	if errGetTotalUserRolesByUser != nil {
		return nil, nil, errGetTotalUserRolesByUser
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u userRolesUseCase) CreateUserRole(
	ctx context.Context,
	userId string,
	body userRolesDomain.CreateUserRoleBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	userRoleID := uuid.New().String()
	// verify if already the user has role
	existUserRole, err := u.userRolesRepository.VerifyUserHasRole(ctx, userId, body.RoleId)
	if err != nil {
		return nil, err
	}
	if existUserRole {
		return nil, userRolesDomain.ErrUserHasRoleAlreadyExist
	}
	id, err = u.userRolesRepository.CreateUserRole(ctx, userRoleID, userId, body)
	return
}

func (u userRolesUseCase) UpdateUserRole(
	ctx context.Context,
	userId string,
	userRoleId string,
	body userRolesDomain.CreateUserRoleBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:        "core_user_roles",
		IdColumnName: "id",
		IdValue:      userRoleId,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().
			CopyCodeDescription(userRolesDomain.ErrUserRoleNotFound).
			SetFunction("UpdateUserRole")
	}
	err = u.userRolesRepository.UpdateUserRole(ctx, userId, userRoleId, body)
	return
}

func (u userRolesUseCase) DeleteUserRole(
	ctx context.Context,
	userId string,
	userRoleId string,
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
		Table:            "core_user_roles",
		IdColumnName:     "id",
		IdValue:          userRoleId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, userRolesDomain.ErrUserRoleIdHasBeenDeleted
	}

	res, err := u.userRolesRepository.DeleteUserRole(ctx, userId, userRoleId)
	return res, err
}
