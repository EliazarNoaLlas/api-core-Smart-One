/*
 * File: modules_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the use cases to module
 *
 * Last Modified: 2023-11-10
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type ModuleUseCase interface {
	GetModules(ctx context.Context, searchParams GetModulesParams, pagination paramsDomain.PaginationParams) ([]Module,
		*paramsDomain.PaginationResults, error)
	CreateModule(ctx context.Context, body CreateModuleBody) (*string, error)
	UpdateModule(ctx context.Context, moduleId string, body UpdateModuleBody) error
	DeleteModule(ctx context.Context, moduleId string) (bool, error)
}
