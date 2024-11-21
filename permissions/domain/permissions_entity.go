/*
 * File: permissions_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to permissions
 *
 * Last Modified: 2023-11-15
 */

package domain

import (
	"time"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type Permission struct {
	//Description: the id of the permission
	Id string `json:"id" binding:"required" example:"fcdbfacf-8305-11ee-89fd-024255555501"`
	//Description: the code of the permission
	Code string `json:"code" binding:"required" example:"REQUIREMENTS_READ"`
	//Description: the name of the permission
	Name string `json:"name" binding:"required" example:"Listar requerimientos"`
	//Description: the description of the permission
	Description string `json:"description" binding:"required" example:"Permiso para listar requerimientos"`
	//Description: the created_at of the permission
	CreatedAt *time.Time         `json:"created_at" example:"2023-11-10 08:10:00"`
	Module    ModuleByPermission `json:"module" binding:"required"`
}

type CreatePermissionBody struct {
	//Description: the id of the permission
	Id string `json:"id" binding:"required" example:"fcdbfacf-8305-11ee-89fd-024255555501"`
	//Description: the code of the permission
	Code string `json:"code" binding:"required" example:"REQUIREMENTS_READ"`
	//Description: the name of the permission
	Name string `json:"name" binding:"required" example:"Listar requerimientos"`
	//Description: the description of the permission
	Description string `json:"description" binding:"required" example:"Permiso para listar requerimientos"`
	//Description: the module_id of the permission
	ModuleId string `json:"module_id" binding:"required" example:"cddbfacf-8305-11ee-89fd-024255555502"`
}

type UpdatePermissionBody struct {
	//Description: the id of the permission
	Id string `json:"id" binding:"required" example:"fcdbfacf-8305-11ee-89fd-024255555501"`
	//Description: the code of the permission
	Code string `json:"code" binding:"required" example:"REQUIREMENTS_READ"`
	//Description: the name of the permission
	Name string `json:"name" binding:"required" example:"Listar requerimientos"`
	//Description: the description of the permission
	Description string `json:"description" binding:"required" example:"Permiso para listar requerimientos"`
}

type ModuleByPermission struct {
	//Description: the id of the module
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the module
	Name string `json:"name" binding:"required" example:"Logistic"`
	//Description: the description of the module
	Description string `json:"description" binding:"required" example:"Modulo de logística"`
	//Description: the code of the module
	Code string `json:"code" binding:"required" example:"logistic"`
}

type GetPermissionsParams struct {
	paramsDomain.Params
	//Description: the code of the permission
	Code *string `json:"code" example:"REQUIREMENTS_READ"`
	//Description: the name of the permission
	Name *string `json:"name" example:"Listar requerimientos"`
}
