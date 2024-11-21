/*
 * File: user_roles_handler_helper_validation.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Validation entity helper to userRoles.
 *
 * Last Modified: 2023-11-23
 */

package rest

type createUserRoleValidate struct {
	RoleId string `json:"role_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
	Enable bool   `json:"enable" example:"true"`
}
