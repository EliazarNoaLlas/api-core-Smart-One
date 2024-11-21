/*
 * File: role_policies_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to rolePolicies
 *
 * Last Modified: 2023-11-22
 */

package domain

import (
	"time"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type PolicyByRolePolicy struct {
	//Description: the id of the role policies
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the role policies
	Name string `json:"name" binding:"required" example:"LOGISTICA_REQUERIMIENTOS_CONGLOMERADO"`
	//Description: the description of the role policies
	Description string `json:"description" binding:"required" example:"Politica para accesos a logistica requerimientos en todo el conglomerado"`
	//Description: the level of the role policies
	Level string `json:"level" binding:"required" example:"system"`
	//Description: enable of the role policies
	Enable bool `json:"enable" binding:"required" example:"true"`
	//Description: the created_at of the role policies
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type RolePolicy struct {
	//Description: the id of the role policies
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110017"`
	//Description: enable of the role policies
	Enable bool `json:"enable" binding:"required" example:"true"`
	//Description: the created_at of the role policies
	CreatedAt *time.Time         `json:"created_at" example:"2023-11-10 08:10:00"`
	Policy    PolicyByRolePolicy `json:"policy" binding:"required"`
}

type CreateRolePolicyBody struct {
	//Description: the policy_id of the created role policies
	PolicyId string `json:"policy_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-042hs5278420"`
	//Description: enable of the created role policies
	Enable bool `json:"enable" binding:"required" example:"true"`
}

type CreateMultipleRolePolicyBody struct {
	CreateRolePolicyBody
	//Description: the id of the permission policy
	Id string `json:"id" binding:"required" example:"22597e1d-6463-4bf9-ba51-0f8a3967321f"`
}

type CreateMultipleRolePoliciesBody struct {
	RolePolicies []CreateRolePolicyBody `json:"role_policies" binding:"required" example:"true"`
}

type UpdateRolePolicyBody struct {
	//Description: the policy_id of the update role policies
	PolicyId string `json:"policy_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-042hs5278420"`
	//Description: enable of the update role policies
	Enable bool `json:"enable" binding:"required" example:"true"`
}

type GetRolePoliciesParams struct {
	paramsDomain.Params
	//Description: role_id of the params
	RoleId string `json:"role_id"`
}

type DeleteMultipleRolePolicyBody struct {
	//Description: array of role policies to delete
	RolePolicyIds []string `json:"role_policy_ids" binding:"required" example:"739bbbc9-7e93-11ee-89fd-042hs5278420"`
}
