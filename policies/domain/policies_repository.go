/*
 * File: policies_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the repository to policies
 *
 * Last Modified: 2023-11-14
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type PolicyRepository interface {
	GetPolicies(ctx context.Context, searchParams GetPoliciesParams, pagination paramsDomain.PaginationParams) (
		[]Policy, error)
	GetTotalPolicies(ctx context.Context, searchParams GetPoliciesParams, pagination paramsDomain.PaginationParams) (
		*int, error)
	CreatePolicy(ctx context.Context, body CreatePolicyBody, policyId string) (*string, error)
	UpdatePolicy(ctx context.Context, body UpdatePolicyBody, policyId string) error
	DeletePolicy(ctx context.Context, policyId string) (bool, error)
}
