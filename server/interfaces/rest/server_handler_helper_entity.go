/*
 * File: server_handler_helper_entity.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2024-04-09
 */

package rest

import (
	serverDomain "gitlab.smartcitiesperu.com/smartone/api-core/server/domain"
)

type ServerDateTimeResult struct {
	Data   serverDomain.ServerDate `json:"data" binding:"required"`
	Status int                     `json:"status" binding:"required"`
}
