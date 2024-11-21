/*
 * File: role_policies_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Repository for rolePolicies.
 *
 * Last Modified: 2023-11-22
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	rolePolicyDomain "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type rolePoliciesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewRolePoliciesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) rolePolicyDomain.RolePolicyRepository {
	rep := &rolePoliciesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
