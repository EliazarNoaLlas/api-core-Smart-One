/*
 * File: policies_handler_helper_validation.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Validation entity helper to policies.
 *
 * Last Modified: 2023-11-14
 */

package rest

type createPoliciesValidate struct {
	Name        string  `json:"name" binding:"required" example:"LOGISTICA_REQUERIMIENTOS_CONGLOMERADO"`
	Description string  `json:"description" binding:"required" example:"Politica para accesos a logistica requerimientos en todo el conglomerado"`
	ModuleId    string  `json:"module_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110018"`
	MerchantId  *string `json:"merchant_id" example:"739bbbc9-7e93-11ee-89fd-0242ac110019"`
	StoreId     *string `json:"store_id" example:"739bbbc9-7e93-11ee-89fd-0242ac110020"`
	Level       string  `json:"level" binding:"required" example:"system"`
	Enable      *bool   `json:"enable" example:"true"`
}
