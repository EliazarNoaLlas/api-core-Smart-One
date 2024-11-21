/*
 * File: document_types_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-06
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type DocumentTypeUseCase interface {
	GetDocumentTypes(ctx context.Context, searchParams GetDocumentTypeParams, pagination paramsDomain.PaginationParams) (
		[]DocumentType, *paramsDomain.PaginationResults, error)
	CreateDocumentType(ctx context.Context, documentTypeId string, body CreateDocumentTypeBody) (*string, error)
	UpdateDocumentType(ctx context.Context, documentTypeId string, body UpdateDocumentTypeBody) error
	DeleteDocumentType(ctx context.Context, documentTypeId string) (bool, error)
}
