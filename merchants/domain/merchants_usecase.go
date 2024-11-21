/*
 * File: `merchants_usecase.go`
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the MerchantUseCase interface, which declares methods for interacting with merchants entities.
 * It includes methods for retrieving, creating, updating, and deleting merchants data.
 *
 * Last Modified: 2023-11-10
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type MerchantUseCase interface {
	GetMerchants(ctx context.Context, pagination paramsDomain.PaginationParams) ([]Merchant,
		*paramsDomain.PaginationResults, error)
	CreateMerchant(ctx context.Context, body CreateMerchantBody) (*string, error)
	UpdateMerchant(ctx context.Context, merchantId string, body UpdateMerchantBody) error
	DeleteMerchant(ctx context.Context, merchantId string) (bool, error)
}
