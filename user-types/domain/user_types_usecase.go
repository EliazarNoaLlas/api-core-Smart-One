/*
 * File: user_types_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the use cases to UserTypeUseCase
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type UserTypeUseCase interface {
	GetUserTypes(ctx context.Context, pagination paramsDomain.PaginationParams) ([]UserType,
		*paramsDomain.PaginationResults, error)
	CreateUserType(ctx context.Context, body CreateUserTypeBody) (*string, error)
	UpdateUserType(ctx context.Context, userTypeId string, body UpdateUserTypeBody) error
	DeleteUserType(ctx context.Context, userTypeId string) (bool, error)
}
