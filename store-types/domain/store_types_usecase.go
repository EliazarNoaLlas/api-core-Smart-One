/*
 * File: store_types_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the StoreTypeUseCase interface, which declares methods for interacting with store_type entities.
 * It includes methods for retrieving, creating, updating, and deleting store_type data.
 *
 * Last Modified: 2023-11-03
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type StoreTypeUseCase interface {
	GetStoreTypes(
		ctx context.Context,
		pagination paramsDomain.PaginationParams,
	) (
		[]StoreType,
		*paramsDomain.PaginationResults,
		error,
	)
	CreateStoreType(ctx context.Context, body CreateStoreTypeBody) (*string, error)
	UpdateStoreType(ctx context.Context, body UpdateStoreTypeBody, storeTypeId string) error
	DeleteStoreType(ctx context.Context, storeTypeId string) (bool, error)
}
