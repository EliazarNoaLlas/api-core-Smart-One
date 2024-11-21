/*
 * File: view_permissions_handler_helper_validation.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities helper to handler for viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package rest

type createViewPermissionBodyValidated struct {
	PermissionId string `json:"permission_id" binding:"required" example:"18f7f9c2-b00a-42e4-a469-ea4c01c180dd"`
}

type updateViewPermissionBodyValidated struct {
	PermissionId string `json:"permission_id" binding:"required" example:"18f7f9c2-b00a-42e4-a469-ea4c01c180dd"`
}
