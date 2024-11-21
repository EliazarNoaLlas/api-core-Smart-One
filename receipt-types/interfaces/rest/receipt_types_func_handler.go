/*
 * File: receipt_types_func_handler.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities of the receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	restCore "gitlab.smartcitiesperu.com/smartone/api-shared/api-core/interfaces/rest"
	httpResponse "gitlab.smartcitiesperu.com/smartone/api-shared/custom-http/interfaces/rest"

	receiptTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
)

// GetReceiptTypes is a method to get classifications receipt types
// @Summary Get classifications receipt types
// @Description Get receipt types
// @Tags Receipt Types
// @Accept json
// @Produce json
// @Success 200 {object} receiptTypesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/receipt_types [get]
// @Security BearerAuth
func (h ReceiptTypesHandler) GetReceiptTypes(c *gin.Context) {
	ctx := c.Request.Context()
	receiptTypes, err := h.ReceiptTypesUseCase.GetReceiptTypes(ctx)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := receiptTypesResult{
		Data:   receiptTypes,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateReceiptType is a method to create receipt type
// @Summary Create a receipt type
// @Description Create receipt type
// @Tags Receipt Types
// @Accept json
// @Produce json
// @Param ReceiptTypeBody body receiptTypesDomain.CreateReceiptTypeBody true "Create receipt type body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/receipt_types [post]
// @Security BearerAuth
func (h ReceiptTypesHandler) CreateReceiptType(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	var receiptTypesValidate createReceiptTypeBodyValidated
	if err := c.ShouldBindJSON(&receiptTypesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateReceiptType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateReceiptType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createReceiptTypesBody = receiptTypesDomain.CreateReceiptTypeBody{
		Description: receiptTypesValidate.Description,
		SunatCode:   receiptTypesValidate.SunatCode,
		Enable:      receiptTypesValidate.Enable,
	}

	id, err := h.ReceiptTypesUseCase.CreateReceiptType(ctx, userId, createReceiptTypesBody)
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

// UpdateReceiptType is a method to update receipt type
// @Summary Update a receipt type
// @Description Update receipt type
// @Tags Receipt Types
// @Accept json
// @Produce json
// @Param receiptTypeId path string true "Receipt Types id"
// @Param ReceiptTypeBody body receiptTypesDomain.UpdateReceiptTypeBody true "Update receipt type body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/receipt_types/{receiptTypeId} [put]
// @Security BearerAuth
func (h ReceiptTypesHandler) UpdateReceiptType(c *gin.Context) {
	ctx := c.Request.Context()
	receiptTypeId := c.Param("receiptTypeId")

	var receiptTypesValidated updateReceiptTypeBodyValidated
	if err := c.ShouldBindJSON(&receiptTypesValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateReceiptType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateReceiptType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var updateReceiptTypesBody = receiptTypesDomain.UpdateReceiptTypeBody{
		Description: receiptTypesValidated.Description,
		SunatCode:   receiptTypesValidated.SunatCode,
		Enable:      receiptTypesValidated.Enable,
	}
	err := h.ReceiptTypesUseCase.UpdateReceiptType(ctx, receiptTypeId, updateReceiptTypesBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteReceiptType is a method to delete receipt type
// @Summary Delete a receipt type
// @Description Delete receipt type
// @Tags Receipt Types
// @Accept json
// @Produce json
// @Param receiptTypeId path string true "Receipt Types id"
// @Success 200 {object} deleteReceiptTypesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/receipt_types/{receiptTypeId} [delete]
// @Security BearerAuth
func (h ReceiptTypesHandler) DeleteReceiptType(c *gin.Context) {
	ctx := c.Request.Context()
	receiptTypeId := c.Param("receiptTypeId")
	result, err := h.ReceiptTypesUseCase.DeleteReceiptType(ctx, receiptTypeId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteReceiptTypesResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
