/*
 * File: store_types_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the StoreTypeRepository interface for store_types data operations.
 *
 * Last Modified: 2023-11-10
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type StoreTypeRepository interface {
	GetStoreTypes(ctx context.Context, pagination paramsDomain.PaginationParams) ([]StoreType, error)
	GetTotalStoreTypes(ctx context.Context, pagination paramsDomain.PaginationParams) (*int, error)
	CreateStoreType(ctx context.Context, storeTypeId string, body CreateStoreTypeBody) (*string, error)
	UpdateStoreType(ctx context.Context, storeTypeId string, body UpdateStoreTypeBody) error
	DeleteStoreType(ctx context.Context, storeTypeId string) (bool, error)
}
