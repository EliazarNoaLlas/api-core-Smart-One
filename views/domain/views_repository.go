/*
 * File: views_repository.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for views.
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type ViewRepository interface {
	GetViews(ctx context.Context, moduleId string, searchParams GetViewsParams,
		pagination paramsDomain.PaginationParams) ([]View, error)
	GetTotalViews(ctx context.Context, moduleId string, searchParams GetViewsParams,
		pagination paramsDomain.PaginationParams) (*int, error)
	CreateView(ctx context.Context, moduleId string, viewId string, body CreateViewBody) (*string, error)
	UpdateView(ctx context.Context, moduleId string, viewId string, body UpdateViewBody) error
	DeleteView(ctx context.Context, moduleId string, viewId string) (bool, error)
}
