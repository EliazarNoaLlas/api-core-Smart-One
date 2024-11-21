/*
 * File: merchants_handler_helper_validation.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines validation structures for merchants related data.
 *
 * Last Modified: 2023-11-10
 */

package rest

type createMerchantsValidate struct {
	Name        string `json:"name" binding:"required" example:"Odin Corp"`
	Description string `json:"description" binding:"required" example:"Proveedor de servicios de mantenimiento"`
	Phone       string `json:"phone" binding:"required" example:"+1234567890"`
	Document    string `json:"document" binding:"required" example:"123456789"`
	Address     string `json:"address" binding:"required" example:"123 Main Street"`
	Industry    string `json:"industry" binding:"required" example:"Mantenimiento"`
	ImagePath   string `json:"image_path" binding:"required" example:"https://example.com/images/odin_logo.png"`
}
