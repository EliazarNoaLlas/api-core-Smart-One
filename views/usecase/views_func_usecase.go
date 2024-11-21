/*
 * File: views_func_usecase.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to views.
 *
 * Last Modified: 2023-11-24
 */

package usecases

import (
	"context"
	"sync"

	"github.com/google/uuid"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	viewDomain "gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
)

func (u viewsUseCase) GetViews(
	ctx context.Context,
	moduleId string,
	searchParams viewDomain.GetViewsParams,
	pagination paramsDomain.PaginationParams,
) (
	views []viewDomain.View,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetViews, errGetTotalViews error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		views, errGetViews = u.viewsRepository.GetViews(ctx, moduleId, searchParams, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalViews = u.viewsRepository.GetTotalViews(ctx, moduleId, searchParams, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetViews != nil {
		return nil, nil, errGetViews
	}
	if errGetTotalViews != nil {
		return nil, nil, errGetTotalViews
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return views, &paginationRes, nil
}

func (u viewsUseCase) CreateView(
	ctx context.Context,
	moduleId string,
	body viewDomain.CreateViewBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	viewId := uuid.New().String()
	id, err = u.viewsRepository.CreateView(ctx, moduleId, viewId, body)
	return
}

func (u viewsUseCase) UpdateView(
	ctx context.Context,
	moduleId string,
	viewId string,
	body viewDomain.UpdateViewBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:        "core_views",
		IdColumnName: "id",
		IdValue:      viewId,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(viewDomain.ErrViewNotFound).SetFunction(
			"UpdateView").SetLayer(errDomain.UseCase)
	}
	err = u.viewsRepository.UpdateView(ctx, moduleId, viewId, body)
	return
}

func (u viewsUseCase) DeleteView(
	ctx context.Context,
	moduleId string,
	viewId string,
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
		Table:            "core_views",
		IdColumnName:     "id",
		IdValue:          viewId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, viewDomain.ErrViewsIdHasBeenDeleted
	}
	res, err := u.viewsRepository.DeleteView(ctx, moduleId, viewId)
	return res, err
}
