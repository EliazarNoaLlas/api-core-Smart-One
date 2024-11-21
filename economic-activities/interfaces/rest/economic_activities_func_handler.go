/*
 * File: economic_activities_func_handler.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the handler function.
 *
 * Last Modified: 2023-12-04
 */

package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	restCore "gitlab.smartcitiesperu.com/smartone/api-shared/api-core/interfaces/rest"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	economicActivityDomain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
)

// GetEconomicActivities is a method to get economic activities
// @Summary get economic activities
// @Description get economic activities
// @Tags EconomicActivities
// @Accept json
// @Produce json
// @Param cuui_id query string false "the cuui id"
// @Param description query string false "the description of the economic activities"
// @Success 200 {object} economicActivitiesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/economic_activities/ [get]
// @Security BearerAuth
func (h economicActivitiesHandler) GetEconomicActivities(c *gin.Context) {
	ctx := c.Request.Context()

	searchParams := economicActivityDomain.GetEconomicActivitiesParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)
	pagination := paramsDomain.NewPaginationParams(c.Request)

	economicActivities, paginationRes, err := h.economicActivitiesUseCase.GetEconomicActivities(ctx, searchParams, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := economicActivitiesResult{
		Data:       economicActivities,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)

}
