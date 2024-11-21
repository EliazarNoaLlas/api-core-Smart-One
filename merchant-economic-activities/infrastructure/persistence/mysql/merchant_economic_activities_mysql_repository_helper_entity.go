/*
 * File: merchant_economic_activities_mysql_repository_helper_entity.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the helper entity of the merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package infrastructure

import (
	"time"
)

type MerchantActivity struct {
	Id               string                     `db:"merchant_economic_id"`
	Sequence         int                        `db:"merchant_economic_sequence"`
	CreatedAt        *time.Time                 `db:"merchant_economic_created_at"`
	EconomicActivity EconomicActivityByMerchant `db:"economic_activity"`
}

type EconomicActivityByMerchant struct {
	Id          string     `db:"activity_id"`
	CuuiId      string     `db:"activity_cuui_id"`
	Description *string    `db:"activity_description"`
	Status      int        `db:"activity_status"`
	CreatedAt   *time.Time `db:"activity_created_at"`
}
