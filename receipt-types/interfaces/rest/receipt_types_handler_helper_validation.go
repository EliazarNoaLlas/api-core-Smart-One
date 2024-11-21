/*
 * File: receipt_types_handler_helper_validation.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities helper to handler for receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package rest

type createReceiptTypeBodyValidated struct {
	Description string `json:"description" binding:"required" example:"activo fijo"`
	SunatCode   string `json:"sunat_code" binding:"required" example:"2"`
	Enable      bool   `json:"enable" example:"true"`
}

type updateReceiptTypeBodyValidated struct {
	Description string `json:"description" binding:"required" example:"activo fijo"`
	SunatCode   string `json:"sunat_code" binding:"required" example:"2"`
	Enable      bool   `json:"enable" example:"true"`
}
