/*
 * File: views_mysql_repository.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Repository for views.
 *
 * Last Modified: 2023-11-24
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	viewDomain "gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type viewMySqlRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewViewRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) viewDomain.ViewRepository {
	rep := &viewMySqlRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
