/*
 * File: stores_func_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to stores.
 *
 * Last Modified: 2023-11-14
 */

package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	storesDomain "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain"
)

func (u storesUseCase) GetStores(
	ctx context.Context,
	merchantId string,
	pagination paramsDomain.PaginationParams,
) (
	stores []storesDomain.Store,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetStores, errGetTotalStores error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		stores, errGetStores = u.storesRepository.GetStores(ctx, merchantId, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalStores = u.storesRepository.GetTotalStores(ctx, merchantId, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetStores != nil {
		return nil, nil, errGetStores
	}
	if errGetTotalStores != nil {
		return nil, nil, errGetTotalStores
	}
	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return stores, &paginationRes, nil
}

func (u storesUseCase) CreateStore(
	ctx context.Context,
	merchantId string,
	body storesDomain.CreateStoreBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_stores",
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
		return nil, storesDomain.ErrStoreNameAlreadyExist
	}
	storeId := uuid.New().String()
	// REVIEW exist the same store in merchant by name

	id, err = u.storesRepository.CreateStore(ctx, merchantId, storeId, body)
	return
}

func (u storesUseCase) UpdateStore(
	ctx context.Context,
	merchantId string,
	storeId string,
	body storesDomain.CreateStoreBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_stores",
		IdColumnName:     "id",
		IdValue:          storeId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(storesDomain.ErrStoreNotFound).SetFunction("UpdateStore")
	}

	err = u.storesRepository.UpdateStore(ctx, merchantId, storeId, body)
	return
}

func (u storesUseCase) DeleteStore(
	ctx context.Context,
	merchantId string,
	storeId string,
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
		Table:            "core_stores",
		IdColumnName:     "id",
		IdValue:          storeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, storesDomain.ErrStoreIdHasBeenDeleted
	}

	res, err := u.storesRepository.DeleteStore(ctx, merchantId, storeId)
	return res, err
}
