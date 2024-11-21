/*
 * File: merchant_economic_activities_repository.go_repository.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository domain of the merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type MerchantEconomicActivityRepository interface {
	GetMerchantEconomicActivities(ctx context.Context, merchantId string, pagination paramsDomain.PaginationParams) (
		[]MerchantEconomicActivity, error)
	GetTotalEconomicActivities(ctx context.Context, merchantId string, pagination paramsDomain.PaginationParams) (
		*int, error)
	CreateEconomicActivity(ctx context.Context, merchantEconomicActivityId string,
		body CreateMerchantEconomicActivityBody) (*string, error)
	UpdateEconomicActivity(ctx context.Context, merchantEconomicActivityId string,
		body UpdateMerchantEconomicActivityBody) error
	DeleteEconomicActivity(ctx context.Context, merchantEconomicActivityId string) (bool, error)
}
