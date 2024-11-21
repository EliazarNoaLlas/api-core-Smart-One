/*
 * File: merchant_economic_activities_entity.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the merchant economic activities entity.
 *
 * Last Modified: 2023-12-05
 */

package domain

import (
	"time"
)

type MerchantEconomicActivity struct {
	//Description: the id of the merchant economic activities
	Id string `json:"id" binding:"required" example:"22d4b62a-9380-11ee-a040-0242ac11000e"`
	//Description: the position of the merchant economic activities
	Sequence int `json:"sequence" binding:"required" example:"1"`
	//Description: the date of create of the merchant economic activities
	CreatedAt        *time.Time                 `json:"created_at" example:"2023-12-05 16:01:51"`
	EconomicActivity EconomicActivityByMerchant `json:"economic_activity"`
}

type EconomicActivityByMerchant struct {
	//Description: the id of the economic activities
	Id string `json:"id" binding:"required" example:"70402269-92be-11ee-a040-0242ac11000e"`
	//Description: the cuui id of the economic activities
	CuuiId string `json:"cuui_id" binding:"required" example:"0111"`
	//Description: the description of the economic activities
	Description *string `json:"description"  example:"CULTIVO DE ARROZ"`
	//Description: the status of the economic activities
	Status int `json:"status" binding:"required" example:"1"`
	//Description: the date of create of the economic activities
	CreatedAt *time.Time `json:"created_at" example:"2023-12-05 16:01:51"`
}

type CreateMerchantEconomicActivityBody struct {
	//Description: the id of the merchants
	MerchantId string `json:"merchant_id" binding:"required" example:"70402269-92be-11ee-a040-0242ac11000e"`
	//Description: the id of the economic activities
	EconomicActivityId string `json:"economic_activity_id" binding:"required" example:"70402269-92be-11ee-a040-0242ac11000e"`
	//Description: the position of the merchant economic activities
	Sequence int `json:"sequence" binding:"required" example:"1"`
}

type UpdateMerchantEconomicActivityBody struct {
	//Description: the id of the merchants
	MerchantId string `json:"merchant_id" binding:"required" example:"70402269-92be-11ee-a040-0242ac11000e"`
	//Description: the id of the economic activities
	EconomicActivityId string `json:"economic_activity_id" binding:"required" example:"70402269-92be-11ee-a040-0242ac11000e"`
	//Description: the position of the merchant economic activities
	Sequence int `json:"sequence" binding:"required" example:"1"`
}
