/*
 * File: merchants_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file initializes the repository layer to manage merchant related data.
 *
 * Last Modified: 2023-11-10
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	merchantDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type merchantsMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewMerchantsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) merchantDomain.MerchantRepository {
	rep := &merchantsMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
