/*
 * File: modules_func_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation to handlers to stores.
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
	_ "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	storesDomain "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain"
)

// GetStores is a method to get stores
// @Summary Get stores
// @Description Get stores
// @Tags Stores
// @Accept json
// @Produce json
// @Param merchantId path string true "merchant id"
// @Success 200 {object} storesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchants/{merchantId}/stores [get]
// @Security BearerAuth
func (h storesHandler) GetStores(c *gin.Context) {
	ctx := c.Request.Context()

	merchantId := c.Param("merchantId")
	pagination := paramsDomain.NewPaginationParams(c.Request)

	stores, paginationRes, err := h.storesUseCase.GetStores(ctx, merchantId, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := storesResult{
		Data:       stores,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateStore is a method to create store
// @Summary Create store
// @Description Create store
// @Tags Stores
// @Accept json
// @Produce json
// @Param merchantId path string true "merchant id"
// @Param createStoreBody body storesDomain.CreateStoreBody true "Create store body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchants/{merchantId}/stores [post]
// @Security BearerAuth
func (h storesHandler) CreateStore(c *gin.Context) {
	ctx := c.Request.Context()
	merchantId := c.Param("merchantId")

	var storesValidate createStoresValidate
	if err := c.ShouldBindJSON(&storesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateStore").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateStore").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createStoreBody = storesDomain.CreateStoreBody{
		Name:        storesValidate.Name,
		Shortname:   storesValidate.Shortname,
		StoreTypeId: storesValidate.StoreTypeId,
	}
	id, err := h.storesUseCase.CreateStore(ctx, merchantId, createStoreBody)
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

// UpdateStore is a method to update store
// @Summary Update store
// @Description Update store
// @Tags Stores
// @Accept json
// @Produce json
// @Param merchantId path string true "merchant id"
// @Param storeId path string true "store id"
// @Param updateStoreBody body storesDomain.CreateStoreBody true "Update store body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError"Bad Request"
// @Router /api/v1/core/merchants/{merchantId}/stores/{storeId} [put]
// @Security BearerAuth
func (h storesHandler) UpdateStore(c *gin.Context) {
	ctx := c.Request.Context()
	merchantId := c.Param("merchantId")
	storeId := c.Param("storeId")

	var storesValidate createStoresValidate
	if err := c.ShouldBindJSON(&storesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateStore").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateStore").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var storeBody = storesDomain.CreateStoreBody{
		Name:        storesValidate.Name,
		Shortname:   storesValidate.Shortname,
		StoreTypeId: storesValidate.StoreTypeId,
	}
	err := h.storesUseCase.UpdateStore(ctx, merchantId, storeId, storeBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteStore is a method to delete store
// @Summary Delete a store
// @Description Delete store
// @Tags Stores
// @Accept json
// @Produce json
// @Param merchantId path string true "merchant id"
// @Param storeId path string true "store id"
// @Success 200 {object} deleteStoresResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchants/{merchantId}/stores/{storeId} [delete]
// @Security BearerAuth
func (h storesHandler) DeleteStore(c *gin.Context) {
	ctx := c.Request.Context()
	merchantId := c.Param("merchantId")
	storeId := c.Param("storeId")
	result, err := h.storesUseCase.DeleteStore(ctx, merchantId, storeId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteStoresResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
