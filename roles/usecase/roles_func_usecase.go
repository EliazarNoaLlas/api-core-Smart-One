/*
 * File: roles_func_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the use case layer where the functions for roles are located.
 *
 * Last Modified: 2023-11-14
 */

package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	rolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/roles/domain"
)

func (u RoleUseCase) GetRoles(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
) (
	res []rolesDomain.Role,
	paginationResult *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetRoles, errGetTotalRoles error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		res, errGetRoles = u.rolesRepository.GetRoles(ctx, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalRoles = u.rolesRepository.GetTotalRoles(ctx, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetRoles != nil {
		return nil, nil, errGetRoles
	}
	if errGetTotalRoles != nil {
		return nil, nil, errGetTotalRoles
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u RoleUseCase) CreateRole(
	ctx context.Context,
	body rolesDomain.CreateRoleBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:        "core_roles",
		IdColumnName: "name",
		IdValue:      body.Name,
	}
	roleId := uuid.New().String()
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, rolesDomain.ErrRoleNameAlreadyExist
	}
	id, err = u.rolesRepository.CreateRole(ctx, roleId, body)
	return

}

func (u RoleUseCase) UpdateRole(
	ctx context.Context,
	roleId string,
	body rolesDomain.CreateRoleBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:        "core_roles",
		IdColumnName: "id",
		IdValue:      roleId,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}

	if !exist {
		return u.err.Clone().CopyCodeDescription(rolesDomain.ErrRoleNotFound).SetFunction(
			"UpdateRole").SetLayer(errDomain.UseCase)
	}
	err = u.rolesRepository.UpdateRole(ctx, roleId, body)
	return
}

func (u RoleUseCase) DeleteRole(
	ctx context.Context,
	roleId string,
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
		Table:            "core_roles",
		IdColumnName:     "id",
		IdValue:          roleId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, rolesDomain.ErrRoleIdHasBeenDeleted
	}

	res, err := u.rolesRepository.DeleteRole(ctx, roleId)
	return res, err
}
