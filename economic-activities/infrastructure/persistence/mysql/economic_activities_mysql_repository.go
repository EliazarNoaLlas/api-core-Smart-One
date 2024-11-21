/*
 * File: economic_activities_mysql_repository.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-04
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	economicActivityDomain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type economicActivitiesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewEconomicActivitiesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) economicActivityDomain.EconomicActivityRepository {
	rep := &economicActivitiesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
