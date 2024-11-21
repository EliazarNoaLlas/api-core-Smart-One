/*
 * File: merchant_economic_activities_handler_helper_validation.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-05
 */

package interfaces

type CreateMerchantEconomicActivityValidate struct {
	MerchantId         string `json:"merchant_id" binding:"required" example:"ac868f48-9448-11ee-a040-0242ac11000e"`
	EconomicActivityId string `json:"economic_activity_id" binding:"required" example:"b4fd2b56-9448-11ee-a040-0242ac11000e"`
	Sequence           int    `json:"sequence" binding:"required" example:"1"`
}

type UpdateMerchantEconomicActivityValidate struct {
	MerchantId         string `json:"merchant_id" binding:"required" example:"70402269-92be-11ee-a040-0242ac11000e"`
	EconomicActivityId string `json:"economic_activity_id" binding:"required" example:"70402269-92be-11ee-a040-0242ac11000e"`
	Sequence           int    `json:"sequence" binding:"required" example:"1"`
}

type deleteMerchantEconomicActivityResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
