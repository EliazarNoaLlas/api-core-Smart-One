/*
 * File: store_types_func_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains handler functions for managing store type-related operations.
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

	storeTypeDomain "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain"
)

// GetStoreTypes is a method to get store types
// @Summary Get store types
// @Description Get store types
// @Tags Store Types
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param size_page query int false "size page"
// @Success 200 {object} storeTypesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/store_types [get]
// @Security BearerAuth
func (h storeTypesHandler) GetStoreTypes(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	storeTypes, paginationRes, err := h.storeTypesUseCase.GetStoreTypes(ctx, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := storeTypesResult{
		Data:       storeTypes,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateStoreType is a method to create store type
// @Summary Create store type
// @Description Create store type
// @Tags Store Types
// @Accept json
// @Produce json
// @Param createStoreTypeBody body storeTypeDomain.CreateStoreTypeBody true "Create store type body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/store_types [post]
// @Security BearerAuth
func (h storeTypesHandler) CreateStoreType(c *gin.Context) {
	ctx := c.Request.Context()
	var storeTypesValidate createStoreTypesValidate
	if err := c.ShouldBindJSON(&storeTypesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateStoreType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreatePolicy").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createStoreTypeBody = storeTypeDomain.CreateStoreTypeBody{
		Description:  storeTypesValidate.Description,
		Abbreviation: storeTypesValidate.Abbreviation,
	}
	id, err := h.storeTypesUseCase.CreateStoreType(ctx, createStoreTypeBody)
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

// UpdateStoreType is a method to update store type
// @Summary Update store type
// @Description Update store type
// @Tags Store Types
// @Accept json
// @Produce json
// @Param storeTypeId path string true "store type id"
// @Param updateStoreTypeBody body storeTypeDomain.UpdateStoreTypeBody true "Update store type body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/store_types/{storeTypeId} [put]
// @Security BearerAuth
func (h storeTypesHandler) UpdateStoreType(c *gin.Context) {
	ctx := c.Request.Context()
	storeTypeId := c.Param("storeTypeId")

	var storeTypesValidate createStoreTypesValidate
	if err := c.ShouldBindJSON(&storeTypesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateStoreType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateUser").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var storeTypesBody = storeTypeDomain.UpdateStoreTypeBody{
		Description:  storeTypesValidate.Description,
		Abbreviation: storeTypesValidate.Abbreviation,
	}
	err := h.storeTypesUseCase.UpdateStoreType(ctx, storeTypesBody, storeTypeId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteStoreType is a method to delete store type
// @Summary Delete store type
// @Description Delete store type
// @Tags Store Types
// @Accept json
// @Produce json
// @Param storeTypeId path string true "store type id"
// @Success 200 {object} deleteStoreTypeResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/store_types/{storeTypeId} [delete]
// @Security BearerAuth
func (h storeTypesHandler) DeleteStoreType(c *gin.Context) {
	ctx := c.Request.Context()
	storeTypeId := c.Param("storeTypeId")
	result, err := h.storeTypesUseCase.DeleteStoreType(ctx, storeTypeId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteStoreTypeResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
