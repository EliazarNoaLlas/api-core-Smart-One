/*
 * File: merchant_economic_activities_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the use case domain for the merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type MerchantEconomicActivityUseCase interface {
	GetMerchantEconomicActivities(ctx context.Context, merchantId string, pagination paramsDomain.PaginationParams) (
		[]MerchantEconomicActivity, *paramsDomain.PaginationResults, error)
	CreateEconomicActivity(ctx context.Context, merchantEconomicActivityId string,
		body CreateMerchantEconomicActivityBody) (*string, error)
	UpdateEconomicActivity(ctx context.Context, merchantEconomicActivityId string,
		body UpdateMerchantEconomicActivityBody) error
	DeleteEconomicActivity(ctx context.Context, merchantEconomicActivityId string) (bool, error)
}
