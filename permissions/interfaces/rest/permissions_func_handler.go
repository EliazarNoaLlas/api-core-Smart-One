/*
 * File: permissions_func_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains handler functions for managing permissions-related operations.
 *
 * Last Modified: 2023-11-15
 */

package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	restCore "gitlab.smartcitiesperu.com/smartone/api-shared/api-core/interfaces/rest"
	httpResponse "gitlab.smartcitiesperu.com/smartone/api-shared/custom-http/interfaces/rest"
	_ "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	permissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
)

// GetPermissions is method to get the permission
// @Summary Get permissions
// @Description Get permissions
// @Tags Permissions
// @Accept json
// @Produce json
// @Param moduleId path string false "module id"
// @Param code query string false "code"
// @Param name query string false "name"
// @Success 200 {object} permissionsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId}/permissions [get]
// @Security BearerAuth
func (h permissionsHandler) GetPermissions(c *gin.Context) {
	ctx := c.Request.Context()
	searchParams := permissionsDomain.GetPermissionsParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)
	pagination := paramsDomain.NewPaginationParams(c.Request)
	moduleId := c.Param("moduleId")

	permissions, paginationRes, err := h.permissionsUseCase.GetPermissions(ctx, moduleId, searchParams, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := permissionsResult{
		Data:       permissions,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreatePermission is a method to create permission
// @Summary Create a permission
// @Description Create a permission
// @Tags Permissions
// @Accept json
// @Produce json
// @Param moduleId path string false "module id"
// @Param createPermissionBody body permissionsDomain.CreatePermissionBody true "Create permission body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId}/permissions [post]
// @Security BearerAuth
func (h permissionsHandler) CreatePermission(c *gin.Context) {
	ctx := c.Request.Context()
	moduleId := c.Param("moduleId")

	var permissionsValidate createPermissionsValidate
	if err := c.ShouldBindJSON(&permissionsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreatePermission").SetRaw(errors.New(
				"casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreatePermission").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createPermissionsBody := permissionsDomain.CreatePermissionBody{
		Code:        permissionsValidate.Code,
		Description: permissionsValidate.Description,
		Name:        permissionsValidate.Name,
		ModuleId:    permissionsValidate.ModuleId,
	}
	id, err := h.permissionsUseCase.CreatePermission(ctx, moduleId, createPermissionsBody)
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

// UpdatePermission is a method to update a permission
// @Summary Update a permission
// @Description Update a permission
// @Tags Permissions
// @Accept json
// @Produce json
// @Param moduleId path string true "module id"
// @Param permissionId path string true "permission id"
// @Param updatePermissionBody body permissionsDomain.UpdatePermissionBody true "Update permission body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId}/permissions/{permissionId} [put]
// @Security BearerAuth
func (h permissionsHandler) UpdatePermission(c *gin.Context) {
	ctx := c.Request.Context()
	moduleId := c.Param("moduleId")
	permissionId := c.Param("permissionId")

	var permissionsValidate updatePermissionsValidate
	if err := c.ShouldBindJSON(&permissionsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdatePermission").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdatePermission").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	permissionsBody := permissionsDomain.UpdatePermissionBody{
		Code:        permissionsValidate.Code,
		Description: permissionsValidate.Description,
		Name:        permissionsValidate.Name,
	}
	err := h.permissionsUseCase.UpdatePermission(ctx, moduleId, permissionId, permissionsBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeletePermission is a method to delete a permission
// @Summary Delete a permission
// @Description Delete a permission
// @Tags Permissions
// @Accept json
// @Produce json
// @Param moduleId path string true "module id"
// @Param permissionId path string true "permission id"
// @Success 200 {object} deletePermissionResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId}/permissions/{permissionId} [delete]
// @Security BearerAuth
func (h permissionsHandler) DeletePermission(c *gin.Context) {
	ctx := c.Request.Context()
	moduleId := c.Param("moduleId")
	permissionId := c.Param("permissionId")

	result, err := h.permissionsUseCase.DeletePermission(ctx, moduleId, permissionId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deletePermissionResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
