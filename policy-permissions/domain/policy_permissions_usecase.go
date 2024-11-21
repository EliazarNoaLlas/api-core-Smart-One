/*
 * File: policyPermissions_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the use cases to policyPermissions.
 *
 * Last Modified: 2023-11-20
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type PolicyPermissionUseCase interface {
	GetPolicyPermissionsByPolicy(ctx context.Context, policyId string, pagination paramsDomain.PaginationParams) (
		[]PolicyPermission, *paramsDomain.PaginationResults, error)
	CreatePolicyPermission(ctx context.Context, policyId string, body CreatePolicyPermissionBody) (*string, error)
	CreatePolicyPermissions(ctx context.Context, policyId string, body []CreatePolicyPermissionBody) ([]string, error)
	UpdatePolicyPermission(ctx context.Context, policyId string, policyPermissionId string,
		body CreatePolicyPermissionBody) error
	DeletePolicyPermission(ctx context.Context, policyId string, policyPermissionId string) (bool, error)
	DeletePolicyPermissions(ctx context.Context, policyId string, policyPermissionIds []string) error
}
