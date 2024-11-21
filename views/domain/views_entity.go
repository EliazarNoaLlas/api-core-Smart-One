/*
 * File: view_entity.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to view
 *
 * Last Modified: 2023-11-24
 */

package domain

import (
	"time"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type View struct {
	//Description: the id of the view
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the view
	Name string `json:"name" binding:"required" example:"Requerimientos"`
	//Description: the description of the view
	Description string `json:"description" binding:"required" example:"Vista para el registro de requerimientos"`
	//Description: the url of the view
	Url string `json:"url" binding:"required" example:"/logistics/requirements"`
	//Description: the icon of the view
	Icon string `json:"icon" binding:"required" example:"fa fa-table"`
	//Description: the created_at of the view
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type CreateViewBody struct {
	//Description: the name of the view
	Name string `json:"name" binding:"required" example:"Requerimientos"`
	//Description: the description of the view
	Description string `json:"description" binding:"required" example:"Vista para el registro de requerimientos"`
	//Description: the url of the view
	Url string `json:"url" binding:"required" example:"/logistics/requirements"`
	//Description: the icon of the view
	Icon string `json:"icon" binding:"required" example:"fa fa-table"`
}

type UpdateViewBody struct {
	//Description: the name of the view
	Name string `json:"name" binding:"required" example:"Requerimientos"`
	//Description: the description of the view
	Description string `json:"description" binding:"required" example:"Vista para el registro de requerimientos"`
	//Description: the url of the view
	Url string `json:"url" binding:"required" example:"/logistics/requirements"`
	//Description: the icon of the view
	Icon string `json:"icon" binding:"required" example:"fa fa-table"`
}

type GetViewsParams struct {
	paramsDomain.Params
	//Description: the name of the params
	Name *string `json:"name"`
}
