/*
 * File: merchant_economic_activities_mysql_repository.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository for merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package infrastructure

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	merchantEconomicActivitiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
)

type merchantEconomicActivitiesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewMerchantEconomicActivitiesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) merchantEconomicActivitiesDomain.MerchantEconomicActivityRepository {
	rep := &merchantEconomicActivitiesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep

}
