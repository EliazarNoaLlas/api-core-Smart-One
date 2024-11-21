/*
 * File: merchants_func_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains handler functions for managing merchants related operations.
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
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	merchantsDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
)

// GetMerchants is a method to get merchant
// @Summary Get merchants
// @Description Get merchant
// @Tags Merchants
// @Accept json
// @Produce json
// @Success 200 {object} merchantsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchants [get]
// @Security BearerAuth
func (h merchantsHandler) GetMerchants(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	merchants, paginationRes, err := h.merchantsUseCase.GetMerchants(ctx, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := merchantsResult{
		Data:       merchants,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateMerchant is a method to create merchant
// @Summary Create merchant
// @Description Create merchant
// @Tags Merchants
// @Accept json
// @Produce json
// @Param createMerchantBody body merchantsDomain.CreateMerchantBody true "Create merchant body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchants [post]
// @Security BearerAuth
func (h merchantsHandler) CreateMerchant(c *gin.Context) {
	ctx := c.Request.Context()
	var merchantsValidate createMerchantsValidate
	if err := c.ShouldBindJSON(&merchantsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateMerchant").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateMerchant").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createMerchantBody = merchantsDomain.CreateMerchantBody{
		Name:        merchantsValidate.Name,
		Description: merchantsValidate.Description,
		Phone:       merchantsValidate.Phone,
		Document:    merchantsValidate.Document,
		Address:     merchantsValidate.Address,
		Industry:    merchantsValidate.Industry,
		ImagePath:   merchantsValidate.ImagePath,
	}
	id, err := h.merchantsUseCase.CreateMerchant(ctx, createMerchantBody)
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

// UpdateMerchant is a method to update merchant
// @Summary Update merchant
// @Description Update merchant
// @Tags Merchants
// @Accept json
// @Produce json
// @Param merchantId path string true "merchant id"
// @Param updateMerchantBody body merchantsDomain.UpdateMerchantBody true "Update merchant body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchants/{merchantId} [put]
// @Security BearerAuth
func (h merchantsHandler) UpdateMerchant(c *gin.Context) {
	ctx := c.Request.Context()
	merchantId := c.Param("merchantId")

	var merchantsValidate createMerchantsValidate
	if err := c.ShouldBindJSON(&merchantsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateMerchant").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateMerchant").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var merchantsBody = merchantsDomain.UpdateMerchantBody{
		Name:        merchantsValidate.Name,
		Description: merchantsValidate.Description,
		Phone:       merchantsValidate.Phone,
		Document:    merchantsValidate.Document,
		Address:     merchantsValidate.Address,
		Industry:    merchantsValidate.Industry,
		ImagePath:   merchantsValidate.ImagePath,
	}
	err := h.merchantsUseCase.UpdateMerchant(ctx, merchantId, merchantsBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteMerchant is a method to delete merchant
// @Summary Delete a merchant
// @Description Delete merchant
// @Tags Merchants
// @Accept json
// @Produce json
// @Param merchantId path string true "merchant id"
// @Success 200 {object} deleteMerchantsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchants/{merchantId} [delete]
// @Security BearerAuth
func (h merchantsHandler) DeleteMerchant(c *gin.Context) {
	ctx := c.Request.Context()
	merchantId := c.Param("merchantId")
	result, err := h.merchantsUseCase.DeleteMerchant(ctx, merchantId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteMerchantsResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
