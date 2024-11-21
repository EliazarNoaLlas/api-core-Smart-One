/*
 * File: server_func_handler.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the server function handler.
 *
 * Last Modified: 2024-04-09
 */

package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"

	restCore "gitlab.smartcitiesperu.com/smartone/api-shared/api-core/interfaces/rest"
)

// GetServerDatetime is a method to get server datetime
// @Summary get datetime
// @Description get server datetime
// @Tags Server
// @Accept json
// @Produce json
// @Success 200 {object} ServerDateTimeResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/server/datetime [get]
// @Security BearerAuth
func (h serverHandler) GetServerDatetime(c *gin.Context) {
	ctx := c.Request.Context()

	result, err := h.serverUseCase.GetServerDate(ctx)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := ServerDateTimeResult{
		Data:   *result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
