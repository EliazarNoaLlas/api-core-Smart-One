/*
 * File: economic_activities_entity.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the entity economic activities.
 *
 * Last Modified: 2023-12-04
 */

package domain

import (
	"time"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type EconomicActivity struct {
	//Description: the id og the economic activities
	Id string `json:"id" binding:"required" example:"70402269-92be-11ee-a040-0242ac11000e"`
	//Description: the cuui id
	CuuiId string `json:"cuui_id" binding:"required" example:"0111"`
	//Description: the description of the economic activities
	Description *string `json:"description" example:"CULTIVO DE ARROZ"`
	//Description:the status of the economic activities
	Status int `json:"status" binding:"required" example:"1"`
	//Description: the created at
	CreatedAt *time.Time `json:"created_at" example:"2023-12-04 16:01:51"`
}

type GetEconomicActivitiesParams struct {
	paramsDomain.Params
	//Description: the cuui id
	CuuiId string `json:"cuui_id" binding:"required" example:"0111"`
	//Description: the description of the economic activities
	Description *string `json:"description" example:"CULTIVO DE ARROZ"`
}
