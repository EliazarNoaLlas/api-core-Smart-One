/*
 * File: receipt_types_entity.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities of the receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package domain

import "time"

type ReceiptType struct {
	//Description: id of the receipt type
	Id string `json:"id" binding:"required" example:"18f7f9c2-b00a-42e4-a469-ea4c01c180dd"`
	//Description: the description of the receipt type
	Description string `json:"description" binding:"required" example:"activo fijo"`
	//Description: the sunat code of the receipt type
	SunatCode string `json:"sunat_code" binding:"required" example:"2"`
	//Description: the status of the receipt type
	Enable bool `json:"enable" binding:"required" example:"true"`
	//Description: the date of creation of the receipt type
	CreatedBy string `json:"created_by" binding:"required" example:"91fb86bd-da46-414b-97a1-fcdaa8cd35d1"`
	//Description: the date of creation of the receipt type
	CreatedAt *time.Time `json:"created_at" example:"2024-01-31 08:10:00"`
}

type CreateReceiptTypeBody struct {
	//Description: the description of the receipt type
	Description string `json:"description" binding:"required" example:"activo fijo"`
	//Description: the sunat code of the receipt type
	SunatCode string `json:"sunat_code" binding:"required" example:"2"`
	//Description: the status of the receipt type
	Enable bool `json:"enable" binding:"required" example:"true"`
}

type UpdateReceiptTypeBody struct {
	//Description: the description of the receipt type
	Description string `json:"description" binding:"required" example:"activo fijo"`
	//Description: the sunat code of the receipt type
	SunatCode string `json:"sunat_code" binding:"required" example:"2"`
	//Description: the status of the receipt type
	Enable bool `json:"enable" binding:"required" example:"true"`
}
