/*
 * File: policies_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to policies
 *
 * Last Modified: 2023-11-14
 */

package domain

import (
	"time"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type ModuleByPolicy struct {
	//Description: the id of the module
	Id *string `json:"id"  example:"739bbbc9-7e93-11ee-89fd-0242ac110018"`
	//Description: the name of the module
	Name *string `json:"name"  example:"Logistic"`
	//Description: the description of the module
	Description *string `json:"description" example:"Modulo de logística"`
	//Description: the code of the module
	Code *string `json:"code"  example:"logistic"`
}

type MerchantByPolicy struct {
	//Description: the id of the merchant
	Id *string `json:"id"  example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the merchant
	Name *string `json:"name"  example:"Odin Corp"`
	//Description: the description of the merchant
	Description *string `json:"description"  example:"Proveedor de servicios de mantenimiento"`
	//Description: the document of the merchant
	Document *string `json:"document"  example:"123456789"`
}

type StoreByPolicy struct {
	//Description: the id of the store
	Id *string `json:"id"  example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the store
	Name *string `json:"name"  example:"Obra av. 28 julio"`
	//Description: the shortname of the store
	Shortname *string `json:"shortname"  example:"Obra 28"`
}

type PolicyPermissionByPolicy struct {
	//Description: the id of the policy permission
	Id *string `json:"id"  example:"739bbbc9-7e93-11ee-89fd-0242ac110010"`
	//Description: enable of the policy permission
	Enable *bool `json:"enable" example:"true"`
}

type PermissionByPolicy struct {
	//Description: the id of the permission
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac110010"`
	//Description: the code of the permission
	Code *string `json:"code" example:"REQUIREMENTS_READ"`
	//Description: the name of the permission
	Name *string `json:"name" example:"Listar requerimientos"`
	//Description: the description of the permission
	Description      *string                   `json:"description" example:"Permiso para listar requerimientos"`
	PolicyPermission *PolicyPermissionByPolicy `json:"policy_permission"`
}

type Policy struct {
	//Description: the id of the policy
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the policy
	Name string `json:"name" binding:"required" example:"LOGISTICA_REQUERIMIENTOS_CONGLOMERADO"`
	//Description: the description of the policy
	Description string `json:"description" binding:"required" example:"Politica para accesos a logistica requerimientos en todo el conglomerado"`
	//Description: the level of the policy
	Level string `json:"level" binding:"required" example:"system"`
	//Description: enable of the policy
	Enable bool `json:"enable" binding:"required" example:"true"`
	//Description: the created_at of the policy
	CreatedAt   *time.Time           `json:"created_at" example:"2023-11-10 08:10:00"`
	Module      *ModuleByPolicy      `json:"module"`
	Merchant    *MerchantByPolicy    `json:"merchant"`
	Store       *StoreByPolicy       `json:"store"`
	Permissions []PermissionByPolicy `json:"permissions"`
}

type CreatePolicyBody struct {
	//Description: the name of the created policy
	Name string `json:"name" binding:"required" example:"LOGISTICA_REQUERIMIENTOS_CONGLOMERADO"`
	//Description: the description of the created policy
	Description string `json:"description" binding:"required" example:"Politica para accesos a logistica requerimientos en todo el conglomerado"`
	//Description: the module_id of the created policy
	ModuleId string `json:"module_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110018"`
	//Description: the merchant_id of the created policy
	MerchantId *string `json:"merchant_id" example:"739bbbc9-7e93-11ee-89fd-0242ac110019"`
	//Description: the store_id of the created policy
	StoreId *string `json:"store_id" example:"739bbbc9-7e93-11ee-89fd-0242ac110020"`
	//Description: the level of the created policy
	Level string `json:"level" binding:"required" example:"system"`
	//Description: enable of the created policy
	Enable *bool `json:"enable" example:"true"`
}

type UpdatePolicyBody struct {
	//Description: the name of the update policy
	Name string `json:"name" binding:"required" example:"LOGISTICA_REQUERIMIENTOS_CONGLOMERADO"`
	//Description: the description of the update policy
	Description string `json:"description" binding:"required" example:"Politica para accesos a logistica requerimientos en todo el conglomerado"`
	//Description: the module_id of the update policy
	ModuleId string `json:"module_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110018"`
	//Description: the merchant_id of the update policy
	MerchantId *string `json:"merchant_id" example:"739bbbc9-7e93-11ee-89fd-0242ac110019"`
	//Description: the store_id of the update policy
	StoreId *string `json:"store_id" example:"739bbbc9-7e93-11ee-89fd-0242ac110020"`
	//Description: the level of the update policy
	Level string `json:"level" binding:"required" example:"system"`
	//Description: enable of the update policy
	Enable *bool `json:"enable" example:"true"`
}

type GetPoliciesParams struct {
	paramsDomain.Params
	//Description: module_id of the params
	ModuleId *string `json:"module_id"`
	//Description: merchant_id of the params
	MerchantId *string `json:"merchant_id"`
	//Description: store_id of the params
	StoreId *string `json:"store_id"`
	//Description: description of the params
	Description *string `json:"description"`
}
