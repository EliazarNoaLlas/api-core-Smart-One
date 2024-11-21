/*
 * File: document_types_handler_helper_validation.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-07
 */

package rest

type CreateDocumentTypeValidate struct {
	Number                 string `json:"number" binding:"required" example:"01"`
	Description            string `json:"description" binding:"required" example:"DOCUMENTO NACIONAL DE IDENTIDAD"`
	AbbreviatedDescription string `json:"abbreviated_description" binding:"required" example:"DNI"`
	Enable                 int    `json:"enable" binding:"required" example:"1"`
}

type UpdateDocumentTypeValidate struct {
	Number                 string `json:"number" binding:"required" example:"01"`
	Description            string `json:"description" binding:"required" example:"DOCUMENTO NACIONAL DE IDENTIDAD"`
	AbbreviatedDescription string `json:"abbreviated_description" binding:"required" example:"DNI"`
	Enable                 int    `json:"enable" binding:"required" example:"1"`
}
