/*
 * File: server_entity.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entity of the server
 *
 * Last Modified: 2024-04-09
 */

package domain

import "time"

type ServerDate struct {
	// Description: Date time
	DateTime time.Time `json:"date_time" binding:"required" example:"2023-10-10T00:00:00Z"`
	// Description: Time zone
	TimeZone string `json:"time_zone" binding:"required" example:"UTC"`
}
