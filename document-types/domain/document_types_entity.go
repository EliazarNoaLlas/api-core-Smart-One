/*
 * File: document_types_entity.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the document_types entity.
 *
 * Last Modified: 2023-12-06
 */

package domain

import (
	"time"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type DocumentType struct {
	//Description: the id of the type of document
	Id string `json:"id" binding:"required" example:"00a58296-93b4-11ee-a040-0242ac11000e"`
	//Description: the number of the type of document
	Number string `json:"number" binding:"required" example:"01"`
	//Description: the description of the type of document
	Description string `json:"description" binding:"required" example:"DOCUMENTO NACIONAL DE IDENTIDAD"`
	//Description: the abbreviation of the type of document
	AbbreviatedDescription string `json:"abbreviated_description" binding:"required" example:"DNI"`
	//Description: the status of the type of document
	Enable int `json:"enable" binding:"required" example:"1"`
	//Description: the date of the type of document
	CreatedAt *time.Time `json:"created_at" example:"2023-12-05 15:49:56"`
}

type GetDocumentTypeParams struct {
	paramsDomain.Params
	//Description: the description of the type of document
	SearchDescription *string `json:"search_description" binding:"required" example:"DOCUMENTO"`
}

type CreateDocumentTypeBody struct {
	//Description: the number of the type of document
	Number string `json:"number" binding:"required" example:"01"`
	//Description: the description of the type of document
	Description string `json:"description" binding:"required" example:"DOCUMENTO NACIONAL DE IDENTIDAD"`
	//Description: the abbreviation of the type of document
	AbbreviatedDescription string `json:"abbreviated_description" binding:"required" example:"DNI"`
	//Description: the status of the type of document
	Enable int `json:"enable" binding:"required" example:"1"`
}

type UpdateDocumentTypeBody struct {
	//Description: the number of the type of document
	Number string `json:"number" binding:"required" example:"01"`
	//Description: the description of the type of document
	Description string `json:"description" binding:"required" example:"DOCUMENTO NACIONAL DE IDENTIDAD"`
	//Description: the abbreviation of the type of document
	AbbreviatedDescription string `json:"abbreviated_description" binding:"required" example:"DNI"`
	//Description: the status of the type of document
	Enable int `json:"enable" binding:"required" example:"1"`
}
