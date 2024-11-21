/*
 * File: policies_func_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation to handlers to policies.
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

	policiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain"
)

// GetPolicies is a method to get policies
// @Summary get policies
// @Description get policies
// @Tags Policies
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param size_page query int false "size page"
// @Param module_id query string false "module id"
// @Param merchant_id query string false "merchant id"
// @Param store_id  query string false "store id"
// @Success 200 {object} policiesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies [get]
// @Security BearerAuth
func (h policiesHandler) GetPolicies(c *gin.Context) {
	ctx := c.Request.Context()

	searchParams := policiesDomain.GetPoliciesParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)
	pagination := paramsDomain.NewPaginationParams(c.Request)

	policies, paginationRes, err := h.policiesUseCase.GetPolicies(ctx, searchParams, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := policiesResult{
		Data:       policies,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreatePolicy is a method to create a policy
// @Summary Create a policy
// @Description Create a policy
// @Tags Policies
// @Accept json
// @Produce json
// @Param createPolicyBody body policiesDomain.CreatePolicyBody true "Create policy body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies [post]
// @Security BearerAuth
func (h policiesHandler) CreatePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	var policiesValidate createPoliciesValidate
	if err := c.ShouldBindJSON(&policiesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreatePolicy").SetRaw(errors.New("casting ValidationErrors"))
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

	var createPolicyBody = policiesDomain.CreatePolicyBody{
		Name:        policiesValidate.Name,
		Description: policiesValidate.Description,
		ModuleId:    policiesValidate.ModuleId,
		MerchantId:  policiesValidate.MerchantId,
		StoreId:     policiesValidate.StoreId,
		Level:       policiesValidate.Level,
		Enable:      policiesValidate.Enable,
	}
	id, err := h.policiesUseCase.CreatePolicy(ctx, createPolicyBody)
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

// UpdatePolicy is a method to update a policy
// @Summary Update a policy
// @Description Update a policy
// @Tags Policies
// @Accept json
// @Produce json
// @Param policyId path string true "policy id"
// @Param updatePolicyBody body policiesDomain.UpdatePolicyBody true "Update policy body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies/{policyId} [put]
// @Security BearerAuth
func (h policiesHandler) UpdatePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	policyId := c.Param("policyId")

	var policiesValidate createPoliciesValidate
	if err := c.ShouldBindJSON(&policiesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdatePolicy").SetRaw(errors.New("casting ValidationErrors"))
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

	var policyBody = policiesDomain.UpdatePolicyBody{
		Name:        policiesValidate.Name,
		Description: policiesValidate.Description,
		ModuleId:    policiesValidate.ModuleId,
		MerchantId:  policiesValidate.MerchantId,
		StoreId:     policiesValidate.StoreId,
		Level:       policiesValidate.Level,
		Enable:      policiesValidate.Enable,
	}
	err := h.policiesUseCase.UpdatePolicy(ctx, policyBody, policyId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeletePolicy is a method to delete a policy
// @Summary Delete a policy
// @Description Delete a policy
// @Tags Policies
// @Accept json
// @Produce json
// @Param policyId path string true "policy id"
// @Success 200 {object} deletePoliciesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/policies/{policyId} [delete]
// @Security BearerAuth
func (h policiesHandler) DeletePolicy(c *gin.Context) {
	ctx := c.Request.Context()
	policyId := c.Param("policyId")
	result, err := h.policiesUseCase.DeletePolicy(ctx, policyId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deletePoliciesResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
