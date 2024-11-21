/*
 * File: modules_func_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation to handlers to policyPermissions.
 *
 * Last Modified: 2023-11-20
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

	policyPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
)

// GetPolicyPermissionsByPolicy is a method to get policy permissions
// @Summary get policy permissions
// @Description get policy permissions
// @Tags PolicyPermissions
// @Accept json
// @Produce json
// @Param policyId path string true "policy id"
// @Success 200 {object} policyPermissionsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies/{policyId}/permissions [get]
// @Security BearerAuth
func (h policyPermissionsHandler) GetPolicyPermissionsByPolicy(c *gin.Context) {
	ctx := c.Request.Context()
	policyId := c.Param("policyId")
	pagination := paramsDomain.NewPaginationParams(c.Request)

	banks, paginationRes, err := h.policyPermissionsUseCase.GetPolicyPermissionsByPolicy(ctx, policyId, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := policyPermissionsResult{
		Data:       banks,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreatePolicyPermission is a method to create a policy permission
// @Summary Create a Policy permission
// @Description Create a Policy permission
// @Tags PolicyPermissions
// @Accept json
// @Produce json
// @Param policyId path string true "policy id"
// @Param createPolicyPermissionBody body policyPermissionsDomain.CreatePolicyPermissionBody true "Create  body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies/{policyId}/permissions [post]
// @Security BearerAuth
func (h policyPermissionsHandler) CreatePolicyPermission(c *gin.Context) {
	ctx := c.Request.Context()
	policyId := c.Param("policyId")

	var policyPermissionsValidate createPolicyPermissionsValidate
	if err := c.ShouldBindJSON(&policyPermissionsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreatePolicyPermission").SetRaw(errors.New(
				"casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)

		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreatePolicyPermission").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createPolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
		PermissionId: policyPermissionsValidate.PermissionId,
		Enable:       policyPermissionsValidate.Enable,
	}
	id, err := h.policyPermissionsUseCase.CreatePolicyPermission(ctx, policyId, createPolicyPermissionBody)
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

// CreateMultiplePolicyPermissions is a method to create multiple policy permissions
// @Summary Create multiple policy permissions
// @Description Create multiple policy permissions
// @Tags PolicyPermissions
// @Accept json
// @Produce json
// @Param policyId path string true "policy id"
// @Param createPolicyPermissionsMultipleBody body policyPermissionsDomain.CreatePolicyPermissionsMultipleBody true "Create  body"
// @Success 201 {object} httpResponse.IdsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies/{policyId}/permissions/batch [post]
// @Security BearerAuth
func (h policyPermissionsHandler) CreateMultiplePolicyPermissions(c *gin.Context) {
	ctx := c.Request.Context()
	policyId := c.Param("policyId")

	var policyPermissionsValidate []createPolicyPermissionsValidate
	if err := c.ShouldBindJSON(&policyPermissionsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().
				SetFunction("CreateMultiplePolicyPermissions").
				SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)

		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateMultiplePolicyPermissions").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createPolicyPermissionsBody := make([]policyPermissionsDomain.CreatePolicyPermissionBody, 0)
	for _, policyPermission := range policyPermissionsValidate {
		createPolicyPermissionsBody = append(createPolicyPermissionsBody, policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: policyPermission.PermissionId,
			Enable:       policyPermission.Enable,
		})
	}

	ids, err := h.policyPermissionsUseCase.CreatePolicyPermissions(ctx, policyId, createPolicyPermissionsBody)
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

// UpdatePolicyPermission is a method to update a policy permission
// @Summary Update a Policy permission
// @Description Update a Policy permission
// @Tags PolicyPermissions
// @Accept json
// @Produce json
// @Param policyId path string true "policy id"
// @Param policyPermissionId path string true "policy permission id"
// @Param policyPermissionBody body policyPermissionsDomain.CreatePolicyPermissionBody true "Update policy permission"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies/{policyId}/permissions/{policyPermissionId} [put]
// @Security BearerAuth
func (h policyPermissionsHandler) UpdatePolicyPermission(c *gin.Context) {
	ctx := c.Request.Context()
	policyId := c.Param("policyId")
	policyPermissionId := c.Param("policyPermissionId")

	var policyPermissionsValidate createPolicyPermissionsValidate
	if err := c.ShouldBindJSON(&policyPermissionsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdatePolicyPermission").SetRaw(errors.New(
				"casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdatePolicyPermission").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	policyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
		PermissionId: policyPermissionsValidate.PermissionId,
		Enable:       policyPermissionsValidate.Enable,
	}
	err := h.policyPermissionsUseCase.UpdatePolicyPermission(ctx, policyId, policyPermissionId, policyPermissionBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeletePolicyPermission is a method to delete a policy permission
// @Summary Delete a Policy permission
// @Description Delete a Policy permission
// @Tags PolicyPermissions
// @Accept json
// @Produce json
// @Param policyId path string true "policy id"
// @Param policyPermissionId path string true "policy permission id"
// @Success 200 {object} deletePolicyPermissionsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies/{policyId}/permissions/{policyPermissionId} [delete]
// @Security BearerAuth
func (h policyPermissionsHandler) DeletePolicyPermission(c *gin.Context) {
	ctx := c.Request.Context()
	policyId := c.Param("policyId")
	policyPermissionId := c.Param("policyPermissionId")
	result, err := h.policyPermissionsUseCase.DeletePolicyPermission(ctx, policyId, policyPermissionId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deletePolicyPermissionsResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteMultiplePolicyPermissions is a method to delete multiple policy permissions
// @Summary Delete multiple policy permissions
// @Description Delete multiple policy permissions
// @Tags PolicyPermissions
// @Accept json
// @Produce json
// @Param policyId path string true "policy id"
// @Param deletePolicyPermissionsMultipleBody body policyPermissionsDomain.DeleteMultiplePolicyPermissionBody true "Delete body"
// @Success 200 {object} deletePolicyPermissionsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies/{policyId}/permissions/batch [delete]
// @Security BearerAuth
func (h policyPermissionsHandler) DeleteMultiplePolicyPermissions(c *gin.Context) {
	ctx := c.Request.Context()
	policyId := c.Param("policyId")

	var policyPermissionsValidate deleteMultiplePolicyPermissionsValidate
	if err := c.ShouldBindJSON(&policyPermissionsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().
				SetFunction("DeleteMultiplePolicyPermissions").
				SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)

		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("DeleteMultiplePolicyPermissions").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	err := h.policyPermissionsUseCase.DeletePolicyPermissions(ctx, policyId, policyPermissionsValidate.PolicyPermissionIds)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deletePolicyPermissionsResult{
		Data:   true,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
