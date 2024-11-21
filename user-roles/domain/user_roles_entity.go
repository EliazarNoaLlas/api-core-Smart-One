/*
 * File: user_roles_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to userRoles
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"time"
)

type CreateUserRoleBody struct {
	//Description: the role_id of the user role
	RoleId string `json:"role_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-042hs5278420"`
	//Description: enable of the user role
	Enable bool `json:"enable" binding:"required" example:"true"`
}

type UserRole struct {
	//Description:the id of the user role
	Id string `json:"id" binding:"required" binding:"required" example:"476a3664-d0d0-4476-8f12-fb11ae57122a"`
	//Description: the status of the user role
	Enable bool `json:"enable" binding:"required" example:"0"`
	//Description: the date of create the user role
	CreatedAt *time.Time `json:"created_at" example:"2023-11-24 16:39:25"`
	Roles     Role       `json:"roles"`
}

type Role struct {
	//Description: the id of the role
	Id string `json:"id" binding:"required" example:"476a3664-d0d0-4476-8f12-fb11ae57122a"`
	//Description: the name of the role
	Name string `json:"name" binding:"required" example:"Gerencia"`
	//Description: the description of the role
	Description string `json:"description" binding:"required" example:"Gerencia del conglomerado2221"`
	//Description: the status of the role
	Enable bool `json:"enable" binding:"required" example:"1"`
	//Description: the date of created of the role
	CreatedAt *time.Time `json:"created_at" binding:"required" example:"0000-00-00 00:00:00"`
}
