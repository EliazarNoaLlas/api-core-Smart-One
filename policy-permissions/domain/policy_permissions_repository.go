/*
 * File: policyPermissions_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the repository to policyPermissions.
 *
 * Last Modified: 2023-11-20
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type PolicyPermissionRepository interface {
	GetPolicyPermissionsByPolicy(ctx context.Context, policyId string, pagination paramsDomain.PaginationParams) (
		[]PolicyPermission, error)
	GetTotalPolicyPermissionsByPolicy(ctx context.Context, policyId string,
		pagination paramsDomain.PaginationParams) (*int, error)
	CreatePolicyPermission(ctx context.Context, policyId string, policyPermissionId string,
		body CreatePolicyPermissionBody) (*string, error)
	CreatePolicyPermissions(ctx context.Context, policyId string, body []CreatePolicyPermissionMultipleBody) error
	VerifyPolicyHasPermission(ctx context.Context, policyId string, permissionId string) (
		bool, error)
	UpdatePolicyPermission(ctx context.Context, policyId string, policyPermissionId string,
		body CreatePolicyPermissionBody) error
	DeletePolicyPermission(ctx context.Context, policyId string, policyPermissionId string) (bool, error)
	DeletePolicyPermissions(ctx context.Context, policyId string, policyPermissionIds []string) error
}
