/*
 * File: modules_func_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation to handlers to userRoles.
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
	_ "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	userRolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
)

// GetUserRolesByUser is a method to get roles by user
// @Summary get roles by user
// @Description get roles by user
// @Tags UserRoles
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Success 200 {object} userRolesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/users/{userId}/roles [get]
// @Security BearerAuth
func (h userRolesHandler) GetUserRolesByUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	pagination := paramsDomain.NewPaginationParams(c.Request)

	banks, paginationRes, err := h.userRolesUseCase.GetUserRolesByUser(ctx, userId, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := userRolesResult{
		Data:       banks,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateUserRole is a method to create user role
// @Summary Create user role
// @Description Create user role
// @Tags UserRoles
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Param createUserRoleBody body userRolesDomain.CreateUserRoleBody true "Create user role body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/users/{userId}/roles [post]
// @Security BearerAuth
func (h userRolesHandler) CreateUserRole(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	var userRolesValidate createUserRoleValidate
	if err := c.ShouldBindJSON(&userRolesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateUserRole").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreatePolicy").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createUserRoleBody = userRolesDomain.CreateUserRoleBody{
		RoleId: userRolesValidate.RoleId,
		Enable: userRolesValidate.Enable,
	}
	id, err := h.userRolesUseCase.CreateUserRole(ctx, userId, createUserRoleBody)
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

// UpdateUserRole is a method to update user role
// @Summary Update user role
// @Description Update user role
// @Tags UserRoles
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Param userRoleId path string true "user role id"
// @Param userRoleBody body userRolesDomain.CreateUserRoleBody true "Update user role body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/users/{userId}/roles{userRoleId} [put]
// @Security BearerAuth
func (h userRolesHandler) UpdateUserRole(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")
	userRoleId := c.Param("userRoleId")

	var userRolesValidate createUserRoleValidate
	if err := c.ShouldBindJSON(&userRolesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateUserRole").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateUser").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var userRoleBody = userRolesDomain.CreateUserRoleBody{
		RoleId: userRolesValidate.RoleId,
		Enable: userRolesValidate.Enable,
	}
	err := h.userRolesUseCase.UpdateUserRole(ctx, userId, userRoleId, userRoleBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteUserRole is a method to delete user role
// @Summary Delete a user role
// @Description Delete user role
// @Tags UserRoles
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Param userRoleId path string true "user role id"
// @Success 200 {object} deleteUserRolesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/users/{userId}/roles{userRoleId} [delete]
// @Security BearerAuth
func (h userRolesHandler) DeleteUserRole(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")
	userRoleId := c.Param("userRoleId")
	result, err := h.userRolesUseCase.DeleteUserRole(ctx, userId, userRoleId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteUserRolesResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
