/*
 * File: policyPermissions_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to policyPermissions
 *
 * Last Modified: 2023-11-20
 */

package domain

import (
	"time"
)

type CreatePolicyPermissionBody struct {
	//Description: the permission_id of the created policy permission
	PermissionId string `json:"permission_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-042hs5278420"`
	//Description: enable of the created policy permission
	Enable bool `json:"enable" binding:"required" example:"true"`
}

type CreatePolicyPermissionsMultipleBody []CreatePolicyPermissionBody

type CreatePolicyPermissionMultipleBody struct {
	CreatePolicyPermissionBody
	//Description: the id of the permission policy
	Id string `json:"id" binding:"required" example:"22597e1d-6463-4bf9-ba51-0f8a3967321f"`
}

type PolicyPermission struct {
	//Description: the id of the permission policy
	Id string `json:"id" binding:"required" example:"22597e1d-6463-4bf9-ba51-0f8a3967321f"`
	//Description: the status of the permission policy
	Enable int `json:"enable" binding:"required" example:"1"`
	//Description: date of create of the permission policy
	CreatedAt  *time.Time `json:"created_at" example:"2023-11-30 15:30:49"`
	Permission Permission `json:"permission" binding:"required"`
}

type Permission struct {
	//Description: the id of the permission of the permission
	Id string `json:"id" binding:"required" example:"84305ba9-83d2-11ee-89fd-0242ac110016"`
	//Description: tho code of the permission
	Code string `json:"code" binding:"required" example:"REQUIREMENTS_READ"`
	//Description: the name of the permission
	Name string `json:"name" binding:"required" example:"Aprobar limpiezas"`
	//Description: the description of the permission
	Description string `json:"description" binding:"required" example:"Permiso para listar requerimientos"`
	//Description: the date of created of the permission
	CreatedAt *time.Time `json:"created_at" example:"2023-12-07 17:13:57"`
}

type DeletePolicyPermissionBody struct {
	PolicyPermissionId string `json:"policy_permission_id" binding:"required" example:"22597e1d-6463-4bf9-ba51-0f8a3967321f"`
}

type DeleteMultiplePolicyPermissionBody struct {
	//Description: the permission_id of the created policy permission
	PolicyPermissionIds []string `json:"policy_permission_ids" binding:"required" example:"739bbbc9-7e93-11ee-89fd-042hs5278420"`
}
