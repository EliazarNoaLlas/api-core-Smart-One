/*
 * File: views_func_handler.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation to handlers to views.
 *
 * Last Modified: 2023-11-24
 */

package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	restCore "gitlab.smartcitiesperu.com/smartone/api-shared/api-core/interfaces/rest"
	httpResponse "gitlab.smartcitiesperu.com/smartone/api-shared/custom-http/interfaces/rest"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	viewsDomain "gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
)

// GetViews is a method to get views by module
// @Summary get views by module
// @Description get views by module
// @Tags Views
// @Accept json
// @Produce json
// @Param moduleId path string true "role id"
// @Param page query int false "page"
// @Param size_page query int false "size page"
// @Success 200 {object} viewsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId}/views [get]
// @Security BearerAuth
func (h viewsHandler) GetViews(c *gin.Context) {
	ctx := c.Request.Context()

	moduleId := c.Param("moduleId")
	searchParams := viewsDomain.GetViewsParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)
	pagination := paramsDomain.NewPaginationParams(c.Request)

	views, paginationRes, err := h.viewsUseCase.GetViews(ctx, moduleId, searchParams, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := viewsResult{
		Data:       views,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateView is a method to create view
// @Summary Create view
// @Description Create view
// @Tags Views
// @Accept json
// @Produce json
// @Param createViewBody body viewsDomain.CreateViewBody true "Create view body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId}/views [post]
// @Security BearerAuth
func (h viewsHandler) CreateView(c *gin.Context) {
	ctx := c.Request.Context()

	moduleId := c.Param("moduleId")
	var viewsValidate createViewsValidate
	if err := c.ShouldBindJSON(&viewsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateView").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)

		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateView").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createViewBody = viewsDomain.CreateViewBody{
		Name:        viewsValidate.Name,
		Description: viewsValidate.Description,
		Url:         viewsValidate.Url,
		Icon:        viewsValidate.Icon,
	}

	id, err := h.viewsUseCase.CreateView(ctx, moduleId, createViewBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   *id,
		Status: http.StatusCreated,
	}
	restCore.Json(c, http.StatusCreated, res)
}

// UpdateView is a method to update view
// @Summary Update view
// @Description Update view
// @Tags Views
// @Accept json
// @Produce json
// @Param viewId path string true "view id"
// @Param updateViewBody body viewsDomain.UpdateViewBody true "Update view body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId}/views/{viewId} [put]
// @Security BearerAuth
func (h viewsHandler) UpdateView(c *gin.Context) {
	ctx := c.Request.Context()
	moduleId := c.Param("moduleId")
	viewId := c.Param("viewId")

	var viewsValidate updateViewsValidate
	if err := c.ShouldBindJSON(&viewsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateView").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateView").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var updateViewBody = viewsDomain.UpdateViewBody{
		Name:        viewsValidate.Name,
		Description: viewsValidate.Description,
		Url:         viewsValidate.Url,
		Icon:        viewsValidate.Icon,
	}

	err := h.viewsUseCase.UpdateView(ctx, moduleId, viewId, updateViewBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusCreated,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteView is a method to delete view
// @Summary Delete a view
// @Description Delete view
// @Tags Views
// @Accept json
// @Produce json
// @Param viewId path string true "view id"
// @Success 200 {object} deleteViewsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId}/views/{viewId} [delete]
// @Security BearerAuth
func (h viewsHandler) DeleteView(c *gin.Context) {
	ctx := c.Request.Context()
	moduleId := c.Param("moduleId")
	viewId := c.Param("viewId")

	result, err := h.viewsUseCase.DeleteView(ctx, moduleId, viewId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := deleteViewsResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
