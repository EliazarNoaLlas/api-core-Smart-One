/*
 * File: policies_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the use cases to policies
 *
 * Last Modified: 2023-11-14
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type PolicyUseCase interface {
	GetPolicies(ctx context.Context, searchParams GetPoliciesParams,
		pagination paramsDomain.PaginationParams) ([]Policy, *paramsDomain.PaginationResults, error)
	CreatePolicy(ctx context.Context, body CreatePolicyBody) (*string, error)
	UpdatePolicy(ctx context.Context, body UpdatePolicyBody, policyId string) error
	DeletePolicy(ctx context.Context, policyId string) (bool, error)
}
