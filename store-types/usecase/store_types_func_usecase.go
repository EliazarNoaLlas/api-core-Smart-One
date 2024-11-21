/*
 * File: store_types_func_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the use case layer where the functions for store types are located.
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

	storeTypeDomain "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain"
)

func (u StoreTypeUseCase) GetStoreTypes(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
) (
	res []storeTypeDomain.StoreType,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetStoreTypes, errGetTotalStoreTypes error
	var total *int
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		res, errGetStoreTypes = u.storeTypesRepository.GetStoreTypes(ctx, pagination)
		wg.Done()
	}()

	go func() {
		total, errGetTotalStoreTypes = u.storeTypesRepository.GetTotalStoreTypes(ctx, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetStoreTypes != nil {
		return nil, nil, errGetStoreTypes
	}
	if errGetTotalStoreTypes != nil {
		return nil, nil, errGetTotalStoreTypes
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u StoreTypeUseCase) CreateStoreType(
	ctx context.Context,
	body storeTypeDomain.CreateStoreTypeBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_store_types",
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
		return nil, storeTypeDomain.ErrStoreTypeDescriptionAlreadyExist
	}

	storeTypeId := uuid.New().String()
	id, err = u.storeTypesRepository.CreateStoreType(ctx, storeTypeId, body)
	return

}

func (u StoreTypeUseCase) UpdateStoreType(
	ctx context.Context,
	body storeTypeDomain.UpdateStoreTypeBody,
	storeTypeId string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_store_types",
		IdColumnName:     "id",
		IdValue:          storeTypeId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(storeTypeDomain.ErrStoreTypeNotFound).SetFunction("UpdateStoreType")
	}

	err = u.storeTypesRepository.UpdateStoreType(ctx, storeTypeId, body)
	return
}

func (u StoreTypeUseCase) DeleteStoreType(
	ctx context.Context,
	storeTypeId string,
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
		Table:            "core_store_types",
		IdColumnName:     "id",
		IdValue:          storeTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, storeTypeDomain.ErrStoreTypeIdHasBeenDeleted
	}

	res, err := u.storeTypesRepository.DeleteStoreType(ctx, storeTypeId)
	return res, err
}
