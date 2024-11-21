/*
 * File: policyPermissions_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Repository for policyPermissions.
 *
 * Last Modified: 2023-11-20
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	policyPermissionDomain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type policyPermissionsMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewPolicyPermissionsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) policyPermissionDomain.PolicyPermissionRepository {
	rep := &policyPermissionsMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
