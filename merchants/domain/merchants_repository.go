/*
 * File: merchants_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the MerchantRepository interface for merchants data operations.
 *
 * Last Modified: 2023-11-10
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type MerchantRepository interface {
	GetMerchants(ctx context.Context, pagination paramsDomain.PaginationParams) ([]Merchant, error)
	GetTotalMerchants(ctx context.Context, pagination paramsDomain.PaginationParams) (*int, error)
	CreateMerchant(ctx context.Context, merchantId string, body CreateMerchantBody) (*string, error)
	UpdateMerchant(ctx context.Context, merchantId string, body UpdateMerchantBody) error
	DeleteMerchant(ctx context.Context, merchantId string) (bool, error)
}
