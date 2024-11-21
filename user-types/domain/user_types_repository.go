/*
 * File: user_types_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the repository to UserTypeRepository
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type UserTypeRepository interface {
	GetUserTypes(ctx context.Context, pagination paramsDomain.PaginationParams) ([]UserType, error)
	GetTotalUserTypes(ctx context.Context, pagination paramsDomain.PaginationParams) (*int, error)
	CreateUserType(ctx context.Context, userTypeId string, body CreateUserTypeBody) (*string, error)
	UpdateUserType(ctx context.Context, userTypeId string, body UpdateUserTypeBody) error
	DeleteUserType(ctx context.Context, userTypeId string) (bool, error)
}
