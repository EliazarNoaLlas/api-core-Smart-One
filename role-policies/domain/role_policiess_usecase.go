/*
 * File: role_policies_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the use cases to rolePolicies.
 *
 * Last Modified: 2023-11-22
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type RolePolicyUseCase interface {
	GetPolicies(ctx context.Context, searchParams GetRolePoliciesParams, pagination paramsDomain.PaginationParams) (
		[]RolePolicy, *paramsDomain.PaginationResults, error)
	CreateRolePolicy(ctx context.Context, roleId string, body CreateRolePolicyBody) (*string, error)
	CreateRolePolicies(ctx context.Context, roleId string, body []CreateRolePolicyBody) ([]string, error)
	UpdateRolePolicy(ctx context.Context, roleId string, rolePolicyId string, body UpdateRolePolicyBody) error
	DeleteRolePolicy(ctx context.Context, roleId string, rolePolicyId string) (bool, error)
	DeleteRolePolicies(ctx context.Context, roleId string, rolePolicyIds []string) error
}
