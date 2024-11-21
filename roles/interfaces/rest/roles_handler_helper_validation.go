/*
 * File: roles_handler_helper_validation.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines validation structures for roles related data.
 *
 * Last Modified: 2023-11-14
 */

package rest

type createRoleValidate struct {
	Name        string `json:"name" binding:"required" example:"Gerencia"`
	Description string `json:"description" binding:"required" example:"Gerencia del conglomerado"`
	Enable      bool   `json:"enable" example:"true"`
}
