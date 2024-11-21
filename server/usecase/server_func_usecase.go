/*
 * File: server_func_usecase.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the use case of the server
 *
 * Last Modified: 2024-04-09
 */

package usecase

import (
	"context"
	"time"

	serverDomain "gitlab.smartcitiesperu.com/smartone/api-core/server/domain"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

func (u serverUseCase) GetServerDate(
	ctx context.Context,
) (
	configurationDateTime *serverDomain.ServerDate, err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	dateTime := serverDomain.ServerDate{
		DateTime: time.Now().UTC(),
		TimeZone: "UTC",
	}
	configurationDateTime = &dateTime

	return configurationDateTime, err
}
