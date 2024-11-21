/*
 * File: modules_func_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation to handlers to modules.
 *
 * Last Modified: 2023-11-10
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

	modulesDomain "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"
)

// GetModules is a method to get modules
// @Summary Get modules
// @Description Get modules
// @Tags Modules
// @Accept json
// @Produce json
// @Param code query string false "Code"
// @Param name query string false "Name"
// @Success 200 {object} modulesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules [get]
// @Security BearerAuth
func (h modulesHandler) GetModules(c *gin.Context) {
	ctx := c.Request.Context()
	searchParams := modulesDomain.GetModulesParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)
	pagination := paramsDomain.NewPaginationParams(c.Request)
	modules, paginationRes, err := h.modulesUseCase.GetModules(ctx, searchParams, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := modulesResult{
		Data:       modules,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateModule is a method to create module
// @Summary Create module
// @Description Create module
// @Tags Modules
// @Accept json
// @Produce json
// @Param createModuleBody body modulesDomain.CreateModuleBody true "Create module body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules [post]
// @Security BearerAuth
func (h modulesHandler) CreateModule(c *gin.Context) {
	ctx := c.Request.Context()
	var modulesValidate createModulesValidate
	if err := c.ShouldBindJSON(&modulesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateModule").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateModule").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createModuleBody = modulesDomain.CreateModuleBody{
		Name:        modulesValidate.Name,
		Description: modulesValidate.Description,
		Code:        modulesValidate.Code,
		Icon:        modulesValidate.Icon,
		Position:    modulesValidate.Position,
	}
	id, err := h.modulesUseCase.CreateModule(ctx, createModuleBody)
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

// UpdateModule is a method to update module
// @Summary Update module
// @Description Update module
// @Tags Modules
// @Accept json
// @Produce json
// @Param moduleId path string true "module id"
// @Param updateModuleBody body modulesDomain.UpdateModuleBody true "Update module body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId} [put]
// @Security BearerAuth
func (h modulesHandler) UpdateModule(c *gin.Context) {
	ctx := c.Request.Context()
	moduleId := c.Param("moduleId")

	var modulesValidate createModulesValidate
	if err := c.ShouldBindJSON(&modulesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateModule").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)

		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateModule").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var moduleBody = modulesDomain.UpdateModuleBody{
		Name:        modulesValidate.Name,
		Description: modulesValidate.Description,
		Code:        modulesValidate.Code,
		Icon:        modulesValidate.Icon,
		Position:    modulesValidate.Position,
	}
	err := h.modulesUseCase.UpdateModule(ctx, moduleId, moduleBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteModule is a method to delete module
// @Summary Delete module
// @Description Delete module
// @Tags Modules
// @Accept json
// @Produce json
// @Param moduleId path string true "module id"
// @Success 200 {object} deleteModulesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/modules/{moduleId} [delete]
// @Security BearerAuth
func (h modulesHandler) DeleteModule(c *gin.Context) {
	ctx := c.Request.Context()
	moduleId := c.Param("moduleId")
	result, err := h.modulesUseCase.DeleteModule(ctx, moduleId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteModulesResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
