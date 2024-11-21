/*
 * File: user_roles_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Repository for userRoles.
 *
 * Last Modified: 2023-11-23
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	userRoleDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type userRolesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewUserRolesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) userRoleDomain.UserRoleRepository {
	rep := &userRolesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
