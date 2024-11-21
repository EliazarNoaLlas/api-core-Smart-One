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

package domain

import "context"

type ServerUseCase interface {
	GetServerDate(ctx context.Context) (*ServerDate, error)
}
