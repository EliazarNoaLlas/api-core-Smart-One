/*
 * File: user_types_func_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation to handlers to user types.
 *
 * Last Modified: 2023-11-23
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

	userTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain"
)

// GetUserTypes is a method to get user types
// @Summary Get user types
// @Description Get user types
// @Tags User Types
// @Accept json
// @Produce json
// @Success 200 {object} userTypesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/user_types [get]
// @Security BearerAuth
func (h userTypesHandler) GetUserTypes(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	userTypes, paginationRes, err := h.userTypesUseCase.GetUserTypes(ctx, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := userTypesResult{
		Data:       userTypes,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateUserType is a method to create user type
// @Summary Create user type
// @Description Create user type
// @Tags User Types
// @Accept json
// @Produce json
// @Param createUserTypeBody body userTypesDomain.CreateUserTypeBody true "Create user type body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/user_types [post]
// @Security BearerAuth
func (h userTypesHandler) CreateUserType(c *gin.Context) {
	ctx := c.Request.Context()
	var userTypesValidate createUserTypesValidate
	if err := c.ShouldBindJSON(&userTypesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateUserType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)

		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateUserType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createUserTypeBody = userTypesDomain.CreateUserTypeBody{
		Description: userTypesValidate.Description,
		Code:        userTypesValidate.Code,
		Enable:      userTypesValidate.Enable,
	}
	id, err := h.userTypesUseCase.CreateUserType(ctx, createUserTypeBody)
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

// UpdateUserType is a method to update user type
// @Summary Update user type
// @Description Update user type
// @Tags User Types
// @Accept json
// @Produce json
// @Param userTypeId path string true "user type id"
// @Param updateUserTypeBody body userTypesDomain.UpdateUserTypeBody true "Update user type body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/user_types/{userTypeId} [put]
// @Security BearerAuth
func (h userTypesHandler) UpdateUserType(c *gin.Context) {
	ctx := c.Request.Context()
	userTypeId := c.Param("userTypeId")

	var userTypesValidate createUserTypesValidate
	if err := c.ShouldBindJSON(&userTypesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateUserType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateUserType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var userTypeBody = userTypesDomain.UpdateUserTypeBody{
		Description: userTypesValidate.Description,
		Code:        userTypesValidate.Code,
		Enable:      userTypesValidate.Enable,
	}
	err := h.userTypesUseCase.UpdateUserType(ctx, userTypeId, userTypeBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteUserType is a method to delete user type
// @Summary Delete user type
// @Description Delete user type
// @Tags User Types
// @Accept json
// @Produce json
// @Param userTypeId path string true "store type id"
// @Success 200 {object} deleteUserTypesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/user_types/{userTypeId} [delete]
// @Security BearerAuth
func (h userTypesHandler) DeleteUserType(c *gin.Context) {
	ctx := c.Request.Context()
	userTypeId := c.Param("userTypeId")
	result, err := h.userTypesUseCase.DeleteUserType(ctx, userTypeId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteUserTypesResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
