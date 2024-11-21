/*
 * File: view_permissions_entity.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities of the viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package domain

import "time"

type ViewPermission struct {
	//Description: id of the view permission
	Id string `json:"id" binding:"required" example:"18f7f9c2-b00a-42e4-a469-ea4c01c180dd"`
	//Description: the description of the view permission
	CreatedBy string `json:"created_by" binding:"required" example:"91fb86bd-da46-414b-97a1-fcdaa8cd35d1"`
	//Description: the date of creation of the view permission
	CreatedAt  *time.Time `json:"created_at" example:"2024-01-31 08:10:00"`
	View       View       `json:"view" binding:"required"`
	Permission Permission `json:"permissions" binding:"required"`
}

type View struct {
	//Description: the id of the view
	Id string `json:"id" binding:"required" example:"18f7f9c2-b00a-42e4-a469-ea4c01c180dd"`
	//Description: the name of the view
	Name string `json:"name" binding:"required" example:"activo fijo"`
	//Description: the description of the view
	Description string `json:"description" binding:"required" example:"activo fijo"`
	//Description: the date of creation of the view
	CreatedAt *time.Time `json:"created_at" example:"2024-01-31 08:10:00"`
}

type Permission struct {
	//Description: the id of the permission
	Id string `json:"id" binding:"required" example:"18f7f9c2-b00a-42e4-a469-ea4c01c180dd"`
	//Description: the code of the permission
	Code string `json:"code" binding:"required" example:"2"`
	//Description: the name of the permission
	Name string `json:"name" binding:"required" example:"activo fijo"`
	//Description: the description of the permission
	Description string `json:"description" binding:"required" example:"activo fijo"`
	//Description: the date of creation of the permission
	CreatedAt *time.Time `json:"created_at" example:"2024-01-31 08:10:00"`
	Module    Module     `json:"module" binding:"required"`
}

type Module struct {
	//Description: the id of the module
	Id string `json:"id" binding:"required" example:"18f7f9c2-b00a-42e4-a469-ea4c01c180dd"`
	//Description: the name of the module
	Name string `json:"name" binding:"required" example:"activo fijo"`
	//Description: the description of the module
	Description string `json:"description" binding:"required" example:"activo fijo"`
	//Description: the code of the module
	Code string `json:"code" binding:"required" example:"2"`
	//Description: the icon of the module
	Icon string `json:"icon" binding:"required" example:"activo fijo"`
	//Description: the position of the module
	Position int `json:"position" binding:"required" example:"2"`
	//Description: the date of creation of the module
	CreatedAt *time.Time `json:"created_at" example:"2024-01-31 08:10:00"`
}

type CreateViewPermissionBody struct {
	//Description: the id of the permission
	PermissionId string `json:"permission_id" binding:"required" example:"18f7f9c2-b00a-42e4-a469-ea4c01c180dd"`
}

type UpdateViewPermissionBody struct {
	//Description: the id of the permission
	PermissionId string `json:"permission_id" binding:"required" example:"18f7f9c2-b00a-42e4-a469-ea4c01c180dd"`
}
