/*
 * File: stores_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Repository for stores.
 *
 * Last Modified: 2023-11-14
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	storeDomain "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type storesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewStoresRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) storeDomain.StoreRepository {
	rep := &storesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
