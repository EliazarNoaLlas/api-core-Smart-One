/*
 * File: receipt_types_mysql_repository.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the repository of the receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package mysql

import (
	"time"

	smartClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	receiptTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
)

type ReceiptTypesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewReceiptTypesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) receiptTypesDomain.ReceiptTypesRepository {
	rep := &ReceiptTypesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
