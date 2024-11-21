/*
 * File: merchants_func_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the use case layer where the merchant functions are located.
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

	merchantsDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
)

func (u merchantsUseCase) GetMerchants(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
) (
	res []merchantsDomain.Merchant,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetMerchants, errGetTotalMerchants error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		res, errGetMerchants = u.merchantsRepository.GetMerchants(ctx, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalMerchants = u.merchantsRepository.GetTotalMerchants(ctx, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetMerchants != nil {
		return nil, nil, errGetMerchants
	}
	if errGetTotalMerchants != nil {
		return nil, nil, errGetTotalMerchants
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u merchantsUseCase) CreateMerchant(
	ctx context.Context,
	body merchantsDomain.CreateMerchantBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	merchantId := uuid.New().String()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_merchants",
		IdColumnName:     "document",
		IdValue:          body.Document,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, merchantsDomain.ErrMerchantDocumentAlreadyExist
	}
	id, err = u.merchantsRepository.CreateMerchant(ctx, merchantId, body)
	return
}

func (u merchantsUseCase) UpdateMerchant(
	ctx context.Context,
	merchantId string,
	body merchantsDomain.UpdateMerchantBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_merchants",
		IdColumnName:     "id",
		IdValue:          merchantId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(merchantsDomain.ErrMerchantNotFound).SetFunction("UpdateMerchant")
	}

	err = u.merchantsRepository.UpdateMerchant(ctx, merchantId, body)
	return
}

func (u merchantsUseCase) DeleteMerchant(
	ctx context.Context,
	merchantId string,
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
		Table:            "core_merchants",
		IdColumnName:     "id",
		IdValue:          merchantId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, merchantsDomain.ErrMerchantIdHasBeenDeleted
	}

	res, err := u.merchantsRepository.DeleteMerchant(ctx, merchantId)
	return res, err
}
