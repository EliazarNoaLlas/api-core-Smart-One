/*
 * File: view_permissions_mysql_repository.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the repository of the viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	ViewPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
)

type ViewPermissionsMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewViewPermissionsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) ViewPermissionsDomain.ViewPermissionsRepository {
	rep := &ViewPermissionsMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
