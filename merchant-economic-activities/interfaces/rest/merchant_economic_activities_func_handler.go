/*
 * File: merchant_economic_activities_func_hendler.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the handler function of the merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package interfaces

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	restCore "gitlab.smartcitiesperu.com/smartone/api-shared/api-core/interfaces/rest"
	httpResponse "gitlab.smartcitiesperu.com/smartone/api-shared/custom-http/interfaces/rest"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
)

// GetMerchantEconomicActivities is a method to get merchant economic activities
// @Summary get merchant economic activities
// @Description get merchant economic activities
// @Tags MerchantEconomicActivities
// @Accept json
// @Produce json
// @Param merchantId path string true "the merchant id"
// @Success 200 {object} merchantEconomicActivitiesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchant_economic_activities/{merchantId} [get]
// @Security BearerAuth
func (h merchantEconomicActivitiesHandler) GetMerchantEconomicActivities(c *gin.Context) {
	ctx := c.Request.Context()

	pagination := paramsDomain.NewPaginationParams(c.Request)
	merchantId := c.Param("merchantId")

	merchantEconomicActivities, paginationRes, err := h.merchantEconomicActivitiesUseCase.
		GetMerchantEconomicActivities(ctx, merchantId, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := merchantEconomicActivitiesResult{
		Data:       merchantEconomicActivities,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateEconomicActivity is a method to create a merchant economic activities
// @Summary Create a merchant economic activities
// @Description Create a merchant economic activities
// @Tags MerchantEconomicActivities
// @Accept json
// @Produce json
// @Param merchantEconomicActivityId path string true "the merchant economic activity id"
// @Param createMerchantEconomicActivityBody body domain.CreateMerchantEconomicActivityBody true "Create user body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchant_economic_activities/url/{merchantEconomicActivityId} [post]
// @Security BearerAuth
func (h merchantEconomicActivitiesHandler) CreateEconomicActivity(c *gin.Context) {
	ctx := c.Request.Context()
	merchantEconomicActivityId := c.Param("merchantEconomicActivityId")
	var merchantEconomicActivityValidate CreateMerchantEconomicActivityValidate
	if err := c.ShouldBindJSON(&merchantEconomicActivityValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateEconomicActivity").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateEconomicActivity").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	CreateMerchantEconomicActivityBody := domain.CreateMerchantEconomicActivityBody{
		MerchantId:         merchantEconomicActivityValidate.MerchantId,
		EconomicActivityId: merchantEconomicActivityValidate.EconomicActivityId,
		Sequence:           merchantEconomicActivityValidate.Sequence,
	}
	id, err := h.merchantEconomicActivitiesUseCase.CreateEconomicActivity(ctx, merchantEconomicActivityId, CreateMerchantEconomicActivityBody)
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

// UpdateEconomicActivity is a method to update a user
// @Summary Update a merchant economic activities
// @Description Update a merchant economic activities
// @Tags MerchantEconomicActivities
// @Accept json
// @Produce json
// @Param merchantEconomicActivityId path string true "the merchant economic activity id"
// @Param updateMerchantEconomicActivityBody body domain.UpdateMerchantEconomicActivityBody true "Update user body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchant_economic_activities/{merchantEconomicActivityId} [put]
// @Security BearerAuth
func (h merchantEconomicActivitiesHandler) UpdateEconomicActivity(c *gin.Context) {
	ctx := c.Request.Context()
	merchantEconomicActivityId := c.Param("merchantEconomicActivityId")

	var merchantEconomicActivityValidate UpdateMerchantEconomicActivityValidate
	if err := c.ShouldBindJSON(&merchantEconomicActivityValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateEconomicActivity").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateEconomicActivity").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	merchantEconomicActivityBody := domain.UpdateMerchantEconomicActivityBody{
		MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
		EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
		Sequence:           1,
	}
	err := h.merchantEconomicActivitiesUseCase.UpdateEconomicActivity(ctx, merchantEconomicActivityId,
		merchantEconomicActivityBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteEconomicActivity is a method to delete a user
// @Summary Delete a merchant economic activities
// @Description Delete a merchant economic activities
// @Tags MerchantEconomicActivities
// @Accept json
// @Produce json
// @Param merchantEconomicActivityId path string true "merchant economic activity id"
// @Success 200 {object} deleteMerchantEconomicActivityResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/merchant_economic_activities/{merchantEconomicActivityId} [delete]
// @Security BearerAuth
func (h merchantEconomicActivitiesHandler) DeleteEconomicActivity(c *gin.Context) {
	ctx := c.Request.Context()
	merchantEconomicActivityId := c.Param("merchantEconomicActivityId")
	result, err := h.merchantEconomicActivitiesUseCase.DeleteEconomicActivity(ctx, merchantEconomicActivityId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteMerchantEconomicActivityResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
