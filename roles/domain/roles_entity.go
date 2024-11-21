/*
 * File: roles_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the structures for roles data: Role and CreateRoleBody.
 *
 * Last Modified: 2023-11-14
 */

package domain

import (
	"time"
)

type Role struct {
	//Description: the id of the role
	Id string `json:"id" binding:"required" example:"fcdbfacf-8305-11ee-89fd-0242555555"`
	//Description: the name of the role
	Name string `json:"name" binding:"required" example:"Gerencia"`
	//Description: the description of the role
	Description string `json:"description" binding:"required" example:"Gerencia del conglomerado"`
	//Description: enable of the role
	Enable bool `json:"enable" example:"true"`
	//Description: the created_at of the role
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type CreateRoleBody struct {
	//Description: the name of the role
	Name string `json:"name" binding:"required" example:"Gerencia"`
	//Description: the description of the role
	Description string `json:"description" binding:"required" example:"Gerencia del conglomerado"`
	//Description: enable of the role
	Enable bool `json:"enable" example:"true"`
}
