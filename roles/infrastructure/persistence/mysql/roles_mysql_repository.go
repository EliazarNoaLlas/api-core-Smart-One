/*
 * File: roles_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file initializes the repository layer to manage roles related data.
 *
 * Last Modified: 2023-11-14
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	rolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/roles/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type roleMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewRolesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) rolesDomain.RoleRepository {
	rep := &roleMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
