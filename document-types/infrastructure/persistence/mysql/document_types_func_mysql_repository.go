/*
 * File: document_types_func_mysql_repository.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-06
 */

package mysql

import (
	"context"
	"database/sql"

	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	_ "embed"
	"gitlab.smartcitiesperu.com/smartone/api-shared/db"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain"
)

//go:embed sql/get_document_types.sql
var QueryGetDocumentTypes string

//go:embed sql/get_total_document_types.sql
var QueryGetTotalDocumentTypes string

//go:embed sql/create_document_type.sql
var QueryCreateDocumentType string

//go:embed sql/update_document_type.sql
var QueryUpdateDocumentType string

//go:embed sql/delete_document_type.sql
var QueryDeleteDocumentType string

func (r documentTypesMySQLRepo) GetDocumentTypes(
	ctx context.Context,
	searchParams domain.GetDocumentTypeParams,
	pagination paramsDomain.PaginationParams,
) (
	documentTypesRows []domain.DocumentType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetDocumentTypes").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetDocumentTypes,
			searchParams.SearchDescription,
			searchParams.SearchDescription,
			searchParams.SearchDescription,
			searchParams.SearchDescription,
			sizePage,
			offset,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetDocumentTypes").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	documentTypesTmp := make([]DocumentType, 0)
	err = carta.Map(results, &documentTypesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetDocumentTypes").SetRaw(err)
	}
	automapper.Map(documentTypesTmp, &documentTypesRows)
	return documentTypesRows, nil
}

func (r documentTypesMySQLRepo) GetTotalDocumentTypes(
	ctx context.Context,
	searchParams domain.GetDocumentTypeParams,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalDocumentTypes").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalDocumentTypes,
			searchParams.SearchDescription,
			searchParams.SearchDescription,
			searchParams.SearchDescription,
			searchParams.SearchDescription,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalDocumentTypes").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r documentTypesMySQLRepo) CreateDocumentType(
	ctx context.Context,
	documentTypeId string,
	body domain.CreateDocumentTypeBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateDocumentType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateDocumentType,
		documentTypeId,
		body.Number,
		body.Description,
		body.AbbreviatedDescription,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateDocumentType").SetRaw(err)
	}
	lastId = &documentTypeId
	return
}

func (r documentTypesMySQLRepo) UpdateDocumentType(
	ctx context.Context,
	documentTypeId string,
	body domain.UpdateDocumentTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateDocumentType").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateDocumentType,
		body.Number,
		body.Description,
		body.AbbreviatedDescription,
		body.Enable,
		documentTypeId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateDocumentType").SetRaw(err)
	}
	return
}

func (r documentTypesMySQLRepo) DeleteDocumentType(
	ctx context.Context,
	documentTypeId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteDocumentType").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteDocumentType,
		now,
		documentTypeId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteDocumentType").SetRaw(err)
	}
	return true, nil
}
