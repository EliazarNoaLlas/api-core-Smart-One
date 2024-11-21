package domain

import (
	"context"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type ViewUseCase interface {
	GetViews(ctx context.Context, moduleId string, searchParams GetViewsParams,
		pagination paramsDomain.PaginationParams) ([]View, *paramsDomain.PaginationResults, error)
	CreateView(ctx context.Context, moduleId string, body CreateViewBody) (*string, error)
	UpdateView(ctx context.Context, moduleId string, viewId string, body UpdateViewBody) error
	DeleteView(ctx context.Context, moduleId string, viewId string) (bool, error)
}
