/*
 * File: stores_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the use cases to stores.
 *
 * Last Modified: 2023-11-14
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type StoreUseCase interface {
	GetStores(ctx context.Context, merchantId string, pagination paramsDomain.PaginationParams) (
		[]Store, *paramsDomain.PaginationResults, error)
	CreateStore(ctx context.Context, merchantId string, body CreateStoreBody) (*string, error)
	UpdateStore(ctx context.Context, merchantId string, storeId string, body CreateStoreBody) error
	DeleteStore(ctx context.Context, merchantId string, storeId string) (bool, error)
}
