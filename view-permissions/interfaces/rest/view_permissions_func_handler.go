/*
 * File: view_permissions_func_handler.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities of the viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	restCore "gitlab.smartcitiesperu.com/smartone/api-shared/api-core/interfaces/rest"
	httpResponse "gitlab.smartcitiesperu.com/smartone/api-shared/custom-http/interfaces/rest"

	ViewPermissions "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
)

// GetViewPermissions is a method to get classifications view permissions
// @Summary Get classifications view permissions
// @Description Get view permissions
// @Tags View Permissions
// @Accept json
// @Produce json
// @Param viewId path string true "View id"
// @Success 200 {object} viewPermissionsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/views/{viewId}/permissions [get]
// @Security BearerAuth
func (h ViewPermissionsHandler) GetViewPermissions(c *gin.Context) {
	ctx := c.Request.Context()
	viewId := c.Param("viewId")
	viewPermissions, err := h.ViewPermissionsUseCase.GetViewPermissions(ctx, viewId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := viewPermissionsResult{
		Data:   viewPermissions,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateViewPermission is a method to create view permission
// @Summary Create a view permission
// @Description Create view permission
// @Tags View Permissions
// @Accept json
// @Produce json
// @Param viewId path string true "View id"
// @Param ViewPermissionBody body ViewPermissions.CreateViewPermissionBody true "Create view permission body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/views/{viewId}/permissions [post]
// @Security BearerAuth
func (h ViewPermissionsHandler) CreateViewPermission(c *gin.Context) {
	ctx := c.Request.Context()
	viewId := c.Param("viewId")
	userId := c.GetString("userId")

	var ViewPermissionsValidate createViewPermissionBodyValidated
	if err := c.ShouldBindJSON(&ViewPermissionsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateViewPermission").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateViewPermission").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createViewPermissionsBody = ViewPermissions.CreateViewPermissionBody{
		PermissionId: ViewPermissionsValidate.PermissionId,
	}
	id, err := h.ViewPermissionsUseCase.CreateViewPermission(ctx, viewId, userId, createViewPermissionsBody)
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

// UpdateViewPermission is a method to update view permission
// @Summary Update a view permission
// @Description Update view permission
// @Tags View Permissions
// @Accept json
// @Produce json
// @Param viewId path string true "View id"
// @Param viewPermissionId path string true "View Permissions id"
// @Param ViewPermissionBody body ViewPermissions.UpdateViewPermissionBody true "Update view permission body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/views/{viewId}/permissions/{viewPermissionId} [put]
// @Security BearerAuth
func (h ViewPermissionsHandler) UpdateViewPermission(c *gin.Context) {
	ctx := c.Request.Context()
	viewId := c.Param("viewId")
	viewPermissionId := c.Param("viewPermissionId")

	var ViewPermissionsValidate updateViewPermissionBodyValidated
	if err := c.ShouldBindJSON(&ViewPermissionsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateViewPermission").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateViewPermission").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var updateViewPermissionsBody = ViewPermissions.UpdateViewPermissionBody{
		PermissionId: ViewPermissionsValidate.PermissionId,
	}
	err := h.ViewPermissionsUseCase.UpdateViewPermission(ctx, viewId, viewPermissionId, updateViewPermissionsBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteViewPermission is a method to delete view permission
// @Summary Delete a view permission
// @Description Delete view permission
// @Tags View Permissions
// @Accept json
// @Produce json
// @Param viewId path string true "View id"
// @Param viewPermissionId path string true "View Permissions id"
// @Success 200 {object} deleteViewPermissionsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/views/{viewId}/permissions/{viewPermissionId} [delete]
// @Security BearerAuth
func (h ViewPermissionsHandler) DeleteViewPermission(c *gin.Context) {
	ctx := c.Request.Context()
	viewId := c.Param("viewId")
	viewPermissionId := c.Param("viewPermissionId")
	result, err := h.ViewPermissionsUseCase.DeleteViewPermission(ctx, viewId, viewPermissionId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteViewPermissionsResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
