/*
 * File: stores_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the repository to stores.
 *
 * Last Modified: 2023-11-14
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type StoreRepository interface {
	GetStores(ctx context.Context, merchantId string, pagination paramsDomain.PaginationParams) ([]Store, error)
	GetTotalStores(ctx context.Context, merchantId string, pagination paramsDomain.PaginationParams) (
		*int, error)
	CreateStore(ctx context.Context, merchantId string, storeId string, body CreateStoreBody) (*string, error)
	UpdateStore(ctx context.Context, merchantId string, storeId string, body CreateStoreBody) error
	DeleteStore(ctx context.Context, merchantId string, storeId string) (bool, error)
}
