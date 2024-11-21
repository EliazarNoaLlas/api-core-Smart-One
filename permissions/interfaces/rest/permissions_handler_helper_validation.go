/*
 * File: permissions_handler_helper_validation.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines validation structures for permissions related data.
 *
 * Last Modified: 2023-11-15
 */

package rest

type createPermissionsValidate struct {
	Code        string `json:"code" binding:"required" example:"REQUIREMENTS_READ"`
	Name        string `json:"name" binding:"required" example:"Listar requerimientos"`
	Description string `json:"description" binding:"required" example:"Permiso para listar requerimientos"`
	ModuleId    string `json:"module_id" binding:"required" example:"cddbfacf-8305-11ee-89fd-024255555502"`
}

type updatePermissionsValidate struct {
	Code        string `json:"code" binding:"required" example:"REQUIREMENTS_READ"`
	Name        string `json:"name" binding:"required" example:"Listar requerimientos"`
	Description string `json:"description" binding:"required" example:"Permiso para listar requerimientos"`
}
