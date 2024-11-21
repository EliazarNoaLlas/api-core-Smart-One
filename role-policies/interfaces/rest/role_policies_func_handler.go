/*
 * File: modules_func_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation to handlers to rolePolicies.
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

	rolePoliciesDomain "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain"
)

// GetPolicies is a method to get policies by role
// @Summary get policies by role
// @Description get policies by role
// @Tags Role Policy
// @Accept json
// @Produce json
// @Param roleId path string true "role id"
// @Param page query int false "page"
// @Param size_page query int false "size page"
// @Success 200 {object} rolePoliciesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles/{roleId}/policies [get]
// @Security BearerAuth
func (h rolePoliciesHandler) GetPolicies(c *gin.Context) {
	ctx := c.Request.Context()

	roleId := c.Param("roleId")
	searchParams := rolePoliciesDomain.GetRolePoliciesParams{
		RoleId: roleId,
	}
	pagination := paramsDomain.NewPaginationParams(c.Request)

	policies, paginationRes, err := h.rolePoliciesUseCase.GetPolicies(ctx, searchParams, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := rolePoliciesResult{
		Data:       policies,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateRolePolicy is a method to create role policy
// @Summary Create role policy
// @Description Create role policy
// @Tags Role Policy
// @Accept json
// @Produce json
// @Param roleId path string true "role id"
// @Param createRolePolicyBody body rolePoliciesDomain.CreateRolePolicyBody true "Create role policy body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles/{roleId}/policies [post]
// @Security BearerAuth
func (h rolePoliciesHandler) CreateRolePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	roleId := c.Param("roleId")

	var rolePoliciesValidate createRolePolicyValidate
	if err := c.ShouldBindJSON(&rolePoliciesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateRolePolicy").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateRolePolicy").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createRolePolicyBody = rolePoliciesDomain.CreateRolePolicyBody{
		PolicyId: rolePoliciesValidate.PolicyId,
		Enable:   rolePoliciesValidate.Enable,
	}
	id, err := h.rolePoliciesUseCase.CreateRolePolicy(ctx, roleId, createRolePolicyBody)
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

// CreateMultipleRolePolicies is a method to create multiple role policies
// @Summary Create multiple role policies
// @Description Create multiple role policies
// @Tags Role Policy
// @Accept json
// @Produce json
// @Param roleId path string true "role id"
// @Param createRolePolicyBody body createMultipleRolePoliciesValidate true "Create multiple role policy body"
// @Success 201 {object} httpResponse.IdsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles/{roleId}/policies/batch [post]
// @Security BearerAuth
func (h rolePoliciesHandler) CreateMultipleRolePolicies(c *gin.Context) {
	ctx := c.Request.Context()
	roleId := c.Param("roleId")

	var rolePoliciesValidate createMultipleRolePoliciesValidate
	if err := c.ShouldBindJSON(&rolePoliciesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().
				SetFunction("CreateMultipleRolePolicies").
				SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateMultipleRolePolicies").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createRolePoliciesBody := make([]rolePoliciesDomain.CreateRolePolicyBody, 0)
	for _, rolePolicy := range rolePoliciesValidate.RolePolicies {
		createRolePoliciesBody = append(createRolePoliciesBody, rolePoliciesDomain.CreateRolePolicyBody{
			PolicyId: rolePolicy.PolicyId,
			Enable:   rolePolicy.Enable,
		})
	}
	ids, err := h.rolePoliciesUseCase.CreateRolePolicies(ctx, roleId, createRolePoliciesBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdsResult{
		Data:   ids,
		Status: http.StatusCreated,
	}
	restCore.Json(c, http.StatusCreated, res)
}

// UpdateRolePolicy is a method to update role policy
// @Summary Update role policy
// @Description Update role policy
// @Tags Role Policy
// @Accept json
// @Produce json
// @Param roleId path string true "role id"
// @Param rolePolicyId path string true "role policy id"
// @Param updateRolePolicyBody body rolePoliciesDomain.UpdateRolePolicyBody true "Update role policy body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles/{roleId}/policies/{rolePolicyId} [put]
// @Security BearerAuth
func (h rolePoliciesHandler) UpdateRolePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	roleId := c.Param("roleId")
	rolePolicyId := c.Param("rolePolicyId")

	var rolePoliciesValidate createRolePolicyValidate
	if err := c.ShouldBindJSON(&rolePoliciesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateRolePolicy").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateRolePolicy").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var rolePolicyBody = rolePoliciesDomain.UpdateRolePolicyBody{
		PolicyId: rolePoliciesValidate.PolicyId,
		Enable:   rolePoliciesValidate.Enable,
	}
	err := h.rolePoliciesUseCase.UpdateRolePolicy(ctx, roleId, rolePolicyId, rolePolicyBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteRolePolicy is a method to delete role policy
// @Summary Delete role policy
// @Description Delete role policy
// @Tags Role Policy
// @Accept json
// @Produce json
// @Param roleId path string true "role id"
// @Param rolePolicyId path string true "role policy id"
// @Success 200 {object} deleteRolePoliciesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles/{roleId}/policies/{rolePolicyId} [delete]
// @Security BearerAuth
func (h rolePoliciesHandler) DeleteRolePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	roleId := c.Param("roleId")
	rolePolicyId := c.Param("rolePolicyId")
	result, err := h.rolePoliciesUseCase.DeleteRolePolicy(ctx, roleId, rolePolicyId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteRolePoliciesResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteMultipleRolePolicies is a method to delete multiple role policies
// @Summary Delete multiple role policies
// @Description Delete multiple role policies
// @Tags Role Policy
// @Accept json
// @Produce json
// @Param roleId path string true "role id"
// @Success 200 {object} deleteMultipleRolePoliciesValidate "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/roles/{roleId}/policies/batch [delete]
// @Security BearerAuth
func (h rolePoliciesHandler) DeleteMultipleRolePolicies(c *gin.Context) {
	ctx := c.Request.Context()
	roleId := c.Param("roleId")

	var policyPermissionsValidate deleteMultipleRolePoliciesValidate
	if err := c.ShouldBindJSON(&policyPermissionsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().
				SetFunction("DeleteMultipleRolePolicies").
				SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)

		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("DeleteMultipleRolePolicies").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	err := h.rolePoliciesUseCase.DeleteRolePolicies(ctx, roleId, policyPermissionsValidate.RolePolicyIds)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteRolePoliciesResult{
		Data:   true,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
