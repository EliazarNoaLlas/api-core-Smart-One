/*
 * File: store_types_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the structures for store_types data: StoreTypeModel and CreateStoreTypeBody.
 *
 * Last Modified: 2023-11-10
 */

package domain

type StoreType struct {
	//Description: the id of the store type
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the description of the store type
	Description string `json:"description" binding:"required" example:"Maquinaria"`
	//Description: the abbreviation of the store type
	Abbreviation string `json:"abbreviation" binding:"required" example:"Maq"`
}

type CreateStoreTypeBody struct {
	//Description: the description of the store type
	Description string `json:"description" binding:"required" example:"Maquinaria"`
	//Description: the abbreviation of the store type
	Abbreviation string `json:"abbreviation" binding:"required" example:"Maq"`
}

type UpdateStoreTypeBody struct {
	//Description: the description of the store type
	Description string `json:"description" binding:"required" example:"Maquinaria"`
	//Description: the abbreviation of the store type
	Abbreviation string `json:"abbreviation" binding:"required" example:"Maq"`
}
