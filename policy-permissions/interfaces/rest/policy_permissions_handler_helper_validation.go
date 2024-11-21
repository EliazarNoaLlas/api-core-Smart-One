/*
 * File: policyPermissions_handler_helper_validation.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Validation entity helper to policyPermissions.
 *
 * Last Modified: 2023-11-20
 */

package rest

type createPolicyPermissionsValidate struct {
	PermissionId string `json:"permission_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
	Enable       bool   `json:"enable" example:"true"`
}

type deleteMultiplePolicyPermissionsValidate struct {
	PolicyPermissionIds []string `json:"policy_permission_ids" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
}
