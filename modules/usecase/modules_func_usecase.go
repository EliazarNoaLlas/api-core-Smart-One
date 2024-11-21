/*
 * File: modules_func_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to modules.
 *
 * Last Modified: 2023-11-10
 */

package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	modulesDomain "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"
)

func (u modulesUseCase) GetModules(
	ctx context.Context,
	searchParams modulesDomain.GetModulesParams,
	pagination paramsDomain.PaginationParams,
) (
	res []modulesDomain.Module,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetModules, errGetTotalModules error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		res, errGetModules = u.modulesRepository.GetModules(ctx, searchParams, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalModules = u.modulesRepository.GetTotalModules(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetModules != nil {
		return nil, nil, errGetModules
	}
	if errGetTotalModules != nil {
		return nil, nil, errGetTotalModules
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u modulesUseCase) CreateModule(
	ctx context.Context,
	body modulesDomain.CreateModuleBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	moduleId := uuid.New().String()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_modules",
		IdColumnName:     "code",
		IdValue:          body.Code,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, modulesDomain.ErrModuleCodeAlreadyExist
	}
	id, err = u.modulesRepository.CreateModule(ctx, moduleId, body)
	return
}

func (u modulesUseCase) UpdateModule(
	ctx context.Context,
	moduleId string,
	body modulesDomain.UpdateModuleBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_modules",
		IdColumnName:     "id",
		IdValue:          moduleId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(modulesDomain.ErrModuleNotFound).SetFunction("UpdateModule")
	}

	err = u.modulesRepository.UpdateModule(ctx, moduleId, body)
	return
}

func (u modulesUseCase) DeleteModule(
	ctx context.Context,
	moduleId string,
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
		Table:            "core_modules",
		IdColumnName:     "id",
		IdValue:          moduleId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, modulesDomain.ErrModuleIdHasBeenDeleted
	}

	res, err := u.modulesRepository.DeleteModule(ctx, moduleId)
	return res, err
}
