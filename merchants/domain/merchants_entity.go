/*
 * File: merchants_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the MerchantModel and CreateMerchantBody structs for merchants data.
 *
 * Last Modified: 2023-11-10
 */

package domain

import (
	"time"
)

type Merchant struct {
	//Description: the id of the merchant
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the merchant
	Name string `json:"name" binding:"required" example:"Odin Corp"`
	//Description: the description of the merchant
	Description string `json:"description" binding:"required" example:"Proveedor de servicios de mantenimiento"`
	//Description: the phone of the merchant
	Phone string `json:"phone" binding:"required" example:"+1234567890"`
	//Description: the document of the merchant
	Document string `json:"document" binding:"required" example:"123456789"`
	//Description: the address of the merchant
	Address string `json:"address" binding:"required" example:"123 Main Street"`
	//Description: the industry of the merchant
	Industry string `json:"industry" binding:"required" example:"Mantenimiento"`
	//Description: the image_path of the merchant
	ImagePath string `json:"image_path" binding:"required" example:"https://example.com/images/odin_logo.png"`
	//Description: the created_at of the merchant
	CreatedAt *time.Time `json:"created_at" binding:"required" example:"2023-11-10 08:10:00"`
}

type CreateMerchantBody struct {
	//Description: the name of the merchant
	Name string `json:"name" binding:"required" example:"Odin Corp"`
	//Description: the description of the merchant
	Description string `json:"description" binding:"required" example:"Proveedor de servicios de mantenimiento"`
	//Description: the phone of the merchant
	Phone string `json:"phone" binding:"required" example:"+1234567890"`
	//Description: the document of the merchant
	Document string `json:"document" binding:"required" example:"123456789"`
	//Description: the address of the merchant
	Address string `json:"address" binding:"required" example:"123 Main Street"`
	//Description: the industry of the merchant
	Industry string `json:"industry" binding:"required" example:"Mantenimiento"`
	//Description: the image_path of the merchant
	ImagePath string `json:"image_path" binding:"required" example:"https://example.com/images/odin_logo.png"`
}

type UpdateMerchantBody struct {
	//Description: the name of the merchant
	Name string `json:"name" binding:"required" example:"Odin Corp"`
	//Description: the description of the merchant
	Description string `json:"description" binding:"required" example:"Proveedor de servicios de mantenimiento"`
	//Description: the phone of the merchant
	Phone string `json:"phone" binding:"required" example:"+1234567890"`
	//Description: the document of the merchant
	Document string `json:"document" binding:"required" example:"123456789"`
	//Description: the address of the merchant
	Address string `json:"address" binding:"required" example:"123 Main Street"`
	//Description: the industry of the merchant
	Industry string `json:"industry" binding:"required" example:"Mantenimiento"`
	//Description: the image_path of the merchant
	ImagePath string `json:"image_path" binding:"required" example:"https://example.com/images/odin_logo.png"`
}
