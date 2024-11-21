/*
 * File: user_types_func_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to user types.
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

	userTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain"
)

func (u userTypesUseCase) GetUserTypes(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
) (
	res []userTypesDomain.UserType,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetUserTypes, errGetTotalUserTypes error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		res, errGetUserTypes = u.userTypesRepository.GetUserTypes(ctx, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalUserTypes = u.userTypesRepository.GetTotalUserTypes(ctx, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetUserTypes != nil {
		return nil, nil, errGetUserTypes
	}
	if errGetTotalUserTypes != nil {
		return nil, nil, errGetTotalUserTypes
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u userTypesUseCase) CreateUserType(
	ctx context.Context,
	body userTypesDomain.CreateUserTypeBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_user_types",
		IdColumnName:     "description",
		IdValue:          body.Description,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, userTypesDomain.ErrUserTypeDescriptionAlreadyExist
	}
	userTypeId := uuid.New().String()
	id, err = u.userTypesRepository.CreateUserType(ctx, userTypeId, body)
	return
}

func (u userTypesUseCase) UpdateUserType(
	ctx context.Context,
	userTypeId string,
	body userTypesDomain.UpdateUserTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_user_types",
		IdColumnName:     "id",
		IdValue:          userTypeId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(userTypesDomain.ErrUserTypeNotFound).SetFunction("UpdateUserType")
	}

	err = u.userTypesRepository.UpdateUserType(ctx, userTypeId, body)
	return
}

func (u userTypesUseCase) DeleteUserType(
	ctx context.Context,
	userTypeId string,
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
		Table:            "core_user_types",
		IdColumnName:     "id",
		IdValue:          userTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, userTypesDomain.ErrUserTypeIdHasBeenDeleted
	}

	res, err := u.userTypesRepository.DeleteUserType(ctx, userTypeId)
	return res, err
}
