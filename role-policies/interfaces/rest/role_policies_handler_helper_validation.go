/*
 * File: role_policies_handler_helper_validation.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Validation entity helper to rolePolicies.
 *
 * Last Modified: 2023-11-23
 */

package rest

type createRolePolicyValidate struct {
	Enable   bool   `json:"enable" example:"true"`
	PolicyId string `json:"policy_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
}

type createMultipleRolePolicyValidate struct {
	createRolePolicyValidate
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210932"`
}

type createMultipleRolePoliciesValidate struct {
	RolePolicies []createMultipleRolePolicyValidate
}
