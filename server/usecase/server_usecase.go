/*
 * File: server_usecase.go
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
	"time"

	serverDomain "gitlab.smartcitiesperu.com/smartone/api-core/server/domain"
)

type serverUseCase struct {
	contextTimeout time.Duration
}

func NewServerUseCase(timeout time.Duration) serverDomain.ServerUseCase {
	return &serverUseCase{
		contextTimeout: timeout,
	}
}
