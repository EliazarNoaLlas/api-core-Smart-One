/*
 * File: policies_mysql_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Repository for policies.
 *
 * Last Modified: 2023-11-14
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"

	policyDomain "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

type policiesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewPoliciesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) policyDomain.PolicyRepository {
	rep := &policiesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
