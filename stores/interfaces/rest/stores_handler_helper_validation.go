/*
 * File: stores_handler_helper_validation.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Validation entity helper to stores.
 *
 * Last Modified: 2023-11-14
 */

package rest

type createStoresValidate struct {
	Name        string `json:"name" binding:"required" example:"Obra av. 28 julio"`
	Shortname   string `json:"shortname" binding:"required" example:"Obra 28"`
	StoreTypeId string `json:"store_type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
}
