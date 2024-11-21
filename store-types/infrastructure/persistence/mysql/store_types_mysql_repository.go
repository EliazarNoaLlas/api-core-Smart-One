/*
 * File: store_types_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file initializes the repository layer to manage store_types related data.
 *
 * Last Modified: 2023-11-10
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	storeTypeDomain "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type storeTypeMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewStoreTypesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) storeTypeDomain.StoreTypeRepository {
	rep := &storeTypeMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
