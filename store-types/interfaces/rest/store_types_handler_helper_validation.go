/*
 * File: store_types_handler_helper_validation.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines validation structures for store type related data.
 *
 * Last Modified: 2023-11-10
 */

package rest

type createStoreTypesValidate struct {
	Description  string `json:"description" binding:"required" example:"Maquinaria"`
	Abbreviation string `json:"abbreviation" binding:"required" example:"Maq"`
}
