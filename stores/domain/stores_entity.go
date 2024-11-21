/*
 * File: stores_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to stores
 *
 * Last Modified: 2023-11-14
 */

package domain

import (
	"time"
)

type Store struct {
	//Description: the id of the store
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the store
	Name string `json:"name" binding:"required" example:"Obra av. 28 julio"`
	//Description: the shortname of the store
	Shortname string `json:"shortname" binding:"required" example:"Obra 28"`
	//Description: the merchant_id of the store
	MerchantId string `json:"merchant_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
	//Description: the created_at of the store
	CreatedAt *time.Time           `json:"created_at" example:"2023-11-10 08:10:00"`
	StoreType StoreTypeByStore `json:"store_type"`
}

type CreateStoreBody struct {
	//Description: the name of the store
	Name string `json:"name" binding:"required" example:"Obra av. 28 julio"`
	//Description: the shortname of the store
	Shortname string `json:"shortname" binding:"required" example:"Obra 28"`
	//Description: the store_type_id of the store
	StoreTypeId string `json:"store_type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-042hs5278420"`
}

type StoreTypeByStore struct {
	//Description: the id of the store type
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113421"`
	//Description: the description of the store type
	Description string `json:"description" binding:"required" example:"Maquinaria"`
	//Description: the abbreviation of the store type
	Abbreviation string `json:"abbreviation" binding:"required" example:"Maq."`
}
