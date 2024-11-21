/*
 * File: permissions_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file initializes the repository layer to manage permissions related data.
 *
 * Last Modified: 2023-11-15
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	permissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type permissionMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewPermissionsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) permissionsDomain.PermissionRepository {
	rep := &permissionMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
