/*
 * File: user_types_handler_helper_validation.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Validation entity helper to createUserTypesValidate.
 *
 * Last Modified: 2023-11-23
 */

package rest

type createUserTypesValidate struct {
	Description string `json:"description" binding:"required" example:"Usuario externo"`
	Code        string `json:"code" binding:"required" example:"USER_EXTERNAL"`
	Enable      bool   `json:"enable" example:"true"`
}
