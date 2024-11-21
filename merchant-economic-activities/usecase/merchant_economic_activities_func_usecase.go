/*
 * File: merchant_economic_activities_func_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the use case function .
 *
 * Last Modified: 2023-12-05
 */

package usecase

import (
	"context"
	"sync"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
)

func (u merchantEconomicActivitiesUseCase) GetMerchantEconomicActivities(
	ctx context.Context,
	merchantId string,
	pagination paramsDomain.PaginationParams,
) (
	res []domain.MerchantEconomicActivity,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetMerchantEconomicActivities, errGetTotalMerchantEconomicActivities error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		res, errGetMerchantEconomicActivities = u.merchantEconomicActivitiesRepository.GetMerchantEconomicActivities(
			ctx, merchantId, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalMerchantEconomicActivities = u.merchantEconomicActivitiesRepository.
			GetTotalEconomicActivities(ctx, merchantId, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetMerchantEconomicActivities != nil {
		return nil, nil, errGetMerchantEconomicActivities
	}
	if errGetTotalMerchantEconomicActivities != nil {
		return nil, nil, errGetTotalMerchantEconomicActivities
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u merchantEconomicActivitiesUseCase) CreateEconomicActivity(
	ctx context.Context,
	merchantEconomicActivityId string,
	body domain.CreateMerchantEconomicActivityBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_merchant_economic_activities",
		IdColumnName:     "id",
		IdValue:          merchantEconomicActivityId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, domain.ErrMerchantEconomicActivityIdAlreadyExist
	}
	id, err = u.merchantEconomicActivitiesRepository.CreateEconomicActivity(ctx, merchantEconomicActivityId, body)
	return
}

func (u merchantEconomicActivitiesUseCase) UpdateEconomicActivity(
	ctx context.Context,
	merchantEconomicActivityId string,
	body domain.UpdateMerchantEconomicActivityBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_merchant_economic_activities",
		IdColumnName:     "id",
		IdValue:          merchantEconomicActivityId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(domain.ErrMerchantEconomicActivityNotFound).SetFunction(
			"UpdateEconomicActivity")
	}

	err = u.merchantEconomicActivitiesRepository.UpdateEconomicActivity(ctx, merchantEconomicActivityId, body)
	return
}

func (u merchantEconomicActivitiesUseCase) DeleteEconomicActivity(
	ctx context.Context,
	merchantEconomicActivityId string,
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
		Table:            "core_merchant_economic_activities",
		IdColumnName:     "id",
		IdValue:          merchantEconomicActivityId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, domain.ErrMerchantEconomicActivityIdHasBeenDeleted
	}

	res, err := u.merchantEconomicActivitiesRepository.DeleteEconomicActivity(ctx, merchantEconomicActivityId)
	return res, err
}
