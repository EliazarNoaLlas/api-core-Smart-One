/*
 * File: modules_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to modules
 *
 * Last Modified: 2023-11-10
 */

package domain

import (
	"time"

	"gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type Module struct {
	//Description: module  id
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: module  name
	Name string `json:"name" binding:"required" example:"Logistic"`
	//Description: module  description
	Description string `json:"description" binding:"required" example:"Modulo de logística"`
	//Description: module  code
	Code string `json:"code" binding:"required" example:"logistic"`
	//Description: module  icon
	Icon string `json:"icon" binding:"required" example:"fa fa-chart"`
	//Description: module  position
	Position int `json:"position" binding:"required" example:"1"`
	//Description: module  created_at
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type CreateModuleBody struct {
	//Description: module  name
	Name string `json:"name" binding:"required" example:"Logistic"`
	//Description: module  description
	Description string `json:"description" binding:"required" example:"Modulo de logística"`
	//Description: module  code
	Code string `json:"code" binding:"required" example:"logistic"`
	//Description: module  icon
	Icon string `json:"icon" binding:"required" example:"fa fa-chart"`
	//Description: module  position
	Position int `json:"position" binding:"required" example:"1"`
}
type UpdateModuleBody struct {
	//Description: module  name
	Name string `json:"name" binding:"required" example:"Logistic"`
	//Description: module  description
	Description string `json:"description" binding:"required" example:"Modulo de logística"`
	//Description: module  code
	Code string `json:"code" binding:"required" example:"logistic"`
	//Description: module  icon
	Icon string `json:"icon" binding:"required" example:"fa fa-chart"`
	//Description: module  position
	Position int `json:"position" binding:"required" example:"1"`
}

type GetModulesParams struct {
	domain.Params
	// Description: module  name
	Name *string `json:"name" example:"Logistic"`
	// Description: module  code
	Code *string `json:"code" example:"logistic"`
}
