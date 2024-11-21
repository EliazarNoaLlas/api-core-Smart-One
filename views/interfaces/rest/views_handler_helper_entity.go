/*
 * File: views_handler_helper_entity.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to handler for views.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	viewsDomain "gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
)

type createViewsValidate struct {
	Name        string `json:"name" binding:"required" example:"Requerimientos"`
	Description string `json:"description" binding:"required" example:"Vista para el registro de requerimientos"`
	Url         string `json:"url" binding:"required" example:"/logistics/requirements"`
	Icon        string `json:"icon" binding:"required" example:"fa fa-table"`
}

type updateViewsValidate struct {
	Name        string `json:"name" binding:"required" example:"Requerimientos"`
	Description string `json:"description" binding:"required" example:"Vista para el registro de requerimientos"`
	Url         string `json:"url" binding:"required" example:"/logistics/requirements"`
	Icon        string `json:"icon" binding:"required" example:"fa fa-table"`
}

type viewsResult struct {
	Data       []viewsDomain.View                 `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}
