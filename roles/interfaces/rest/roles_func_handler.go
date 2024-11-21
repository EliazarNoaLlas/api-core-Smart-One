/*
 * File: roles_func_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains handler functions for managing roles-related operations.
 *
 * Last Modified: 2023-11-14
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

	rolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/roles/domain"
)

// GetRoles is a method to get roles
// @Summary Get roles
// @Description Get roles
// @Tags Roles
// @Accept json
// @Produce json
// @Success 200 {object} rolesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles [get]
// @Security BearerAuth
func (h rolesHandler) GetRoles(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	roles, paginationRes, err := h.rolesUseCase.GetRoles(ctx, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := rolesResult{
		Data:       roles,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateRole is a method to create role
// @Summary Create role
// @Description Create role
// @Tags Roles
// @Accept json
// @Produce json
// @Param createRoleBody body rolesDomain.CreateRoleBody true "Create role body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles [post]
// @Security BearerAuth
func (h rolesHandler) CreateRole(c *gin.Context) {
	ctx := c.Request.Context()
	var rolesValidate createRoleValidate
	if err := c.ShouldBindJSON(&rolesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateRole").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}

		err = h.err.Clone().SetFunction("CreateRole").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createRoleBody = rolesDomain.CreateRoleBody{
		Description: rolesValidate.Description,
		Name:        rolesValidate.Name,
		Enable:      rolesValidate.Enable,
	}
	id, err := h.rolesUseCase.CreateRole(ctx, createRoleBody)
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

// UpdateRole is a method to update role
// @Summary Update role
// @Description Update role
// @Tags Roles
// @Accept json
// @Produce json
// @Param roleId path string true "role id"
// @Param rolesBody body rolesDomain.CreateRoleBody true "Update role body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles/{roleId} [put]
// @Security BearerAuth
func (h rolesHandler) UpdateRole(c *gin.Context) {
	ctx := c.Request.Context()
	roleId := c.Param("roleId")

	var rolesValidate createRoleValidate
	if err := c.ShouldBindJSON(&rolesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateRole").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateRole").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var rolesBody = rolesDomain.CreateRoleBody{
		Description: rolesValidate.Description,
		Name:        rolesValidate.Name,
		Enable:      rolesValidate.Enable,
	}
	err := h.rolesUseCase.UpdateRole(ctx, roleId, rolesBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteRole is a method to delete role
// @Summary Delete role
// @Description Delete role
// @Tags Roles
// @Accept json
// @Produce json
// @Param roleId path string true "role id"
// @Success 200 {object} deleteRoleResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles/{roleId} [delete]
// @Security BearerAuth
func (h rolesHandler) DeleteRole(c *gin.Context) {
	ctx := c.Request.Context()
	roleId := c.Param("roleId")
	result, err := h.rolesUseCase.DeleteRole(ctx, roleId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteRoleResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
