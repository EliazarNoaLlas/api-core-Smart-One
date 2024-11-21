/*
 * File: modules_handler_helper_validation.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Validation entity helper to modules.
 *
 * Last Modified: 2023-11-10
 */

package rest

type createModulesValidate struct {
	Name        string `json:"name" binding:"required" example:"Logistica"`
	Description string `json:"description" binding:"required" example:"Modulo de logistica"`
	Code        string `json:"code" binding:"required" example:"logistic"`
	Icon        string `json:"icon" binding:"required" example:"fa fa-home"`
	Position    int    `json:"position" binding:"required" example:"1"`
}
