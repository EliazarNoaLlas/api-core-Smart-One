/*
 * File: user_types_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to user types
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"time"
)

type UserType struct {
	//Description: the id of the user type
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the description of the user type
	Description string `json:"description" binding:"required" example:"Usuario externo"`
	//Description: the code of the user type
	Code string `json:"code" binding:"required" example:"USER_EXTERNAL"`
	//Description: the id status the user type
	Enable bool `json:"enable" binding:"required" example:"true"`
	//Description: the date of created the user type
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type CreateUserTypeBody struct {
	//Description: the description of the user type
	Description string `json:"description" binding:"required" example:"Usuario externo"`
	//Description: the code of the user type
	Code string `json:"code" binding:"required" example:"USER_EXTERNAL"`
	//Description: the id status the user type
	Enable bool `json:"enable" example:"true"`
}

type UpdateUserTypeBody struct {
	//Description: the description of the user type
	Description string `json:"description" binding:"required" example:"Usuario externo"`
	//Description: the code of the user type
	Code string `json:"code" binding:"required" example:"USER_EXTERNAL"`
	//Description: the id status the user type
	Enable bool `json:"enable" example:"true"`
}
