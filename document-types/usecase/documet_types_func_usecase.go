/*
 * File: documentTypes_func_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to documentTypes.
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

	documentTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain"
)

func (u documentTypesUseCase) GetDocumentTypes(
	ctx context.Context,
	searchParams documentTypesDomain.GetDocumentTypeParams,
	pagination paramsDomain.PaginationParams,
) (
	res []documentTypesDomain.DocumentType,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetDocumentTypes, errGetTotalDocumentTypes error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		res, errGetDocumentTypes = u.documentTypesRepository.GetDocumentTypes(ctx, searchParams, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetDocumentTypes = u.documentTypesRepository.GetTotalDocumentTypes(ctx, searchParams, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetDocumentTypes != nil {
		return nil, nil, errGetDocumentTypes
	}
	if errGetDocumentTypes != nil {
		return nil, nil, errGetTotalDocumentTypes
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u documentTypesUseCase) CreateDocumentType(
	ctx context.Context,
	documentTypeId string,
	body documentTypesDomain.CreateDocumentTypeBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	documentTypeId = uuid.New().String()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_document_types",
		IdColumnName:     "number",
		IdValue:          body.Number,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, documentTypesDomain.ErrDocumentTypeDescriptionAlreadyExist
	}
	id, err = u.documentTypesRepository.CreateDocumentType(ctx, documentTypeId, body)
	return
}

func (u documentTypesUseCase) UpdateDocumentType(
	ctx context.Context,
	documentTypeId string,
	body documentTypesDomain.UpdateDocumentTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_document_types",
		IdColumnName:     "id",
		IdValue:          documentTypeId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(documentTypesDomain.ErrDocumentTypeNotFound).SetFunction("UpdateDocumentType")
	}

	err = u.documentTypesRepository.UpdateDocumentType(ctx, documentTypeId, body)
	return
}

func (u documentTypesUseCase) DeleteDocumentType(
	ctx context.Context,
	documentTypeId string,
) (
	update bool,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	var deleted string
	deleted = "deleted_at"

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_document_types",
		IdColumnName:     "id",
		IdValue:          documentTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}

	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, documentTypesDomain.ErrDocumentTypeIdHasBeenDeleted
	}

	res, err := u.documentTypesRepository.DeleteDocumentType(ctx, documentTypeId)
	return res, err
}
