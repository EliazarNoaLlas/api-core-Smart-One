/*
 * File: role_policies_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the repository to rolePolicies.
 *
 * Last Modified: 2023-11-22
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type RolePolicyRepository interface {
	GetPolicies(ctx context.Context, searchParams GetRolePoliciesParams, pagination paramsDomain.PaginationParams) (
		[]RolePolicy, error)
	GetTotalPolicies(ctx context.Context, searchParams GetRolePoliciesParams, pagination paramsDomain.PaginationParams) (
		*int, error)
	CreateRolePolicy(ctx context.Context, rolePolicyId string, roleId string, body CreateRolePolicyBody) (*string, error)
	CreateRolePolicies(ctx context.Context, roleId string, body []CreateMultipleRolePolicyBody) error
	VerifyRoleHasPolicy(ctx context.Context, roleId string, policyId string) (bool, error)
	UpdateRolePolicy(ctx context.Context, roleId string, rolePolicyId string, body UpdateRolePolicyBody) error
	DeleteRolePolicy(ctx context.Context, roleId string, rolePolicyId string) (bool, error)
	DeleteRolePolicies(ctx context.Context, roleId string, rolePolicyIds []string) error
}
