/*
 * File: merchant_economic_activities_error.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the error manage of the application.
 *
 * Last Modified: 2023-12-05
 */

package domain

import (
	"net/http"

	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
)

const (
	ErrMerchantEconomicActivityNotFoundCode         = "ERR_MERCHANT_ECONOMIC_ACTIVITY_NOT_FOUND"
	ErrMerchantEconomicActivityIdAlreadyExistCode   = "ERR_MERCHANT_ECONOMIC_ACTIVITY_ID_ALREADY_EXIST"
	ErrMerchantEconomicActivityIdHasBeenDeletedCode = "ERR_MERCHANT_ECONOMIC_ACTIVITY_ID_HAS_BEEN_DELETED"
)

var (
	ErrMerchantEconomicActivityNotFound = errDomain.NewErr().
						SetCode(ErrMerchantEconomicActivityNotFoundCode).
						SetDescription("MERCHANT ECONOMIC ACTIVITY NOT FOUND").
						SetLevel(errDomain.LevelError).
						SetHttpStatus(http.StatusNotFound).
						SetLayer(errDomain.Infra).
						SetFunction("UpdateEconomicActivity")

	ErrMerchantEconomicActivityIdAlreadyExist = errDomain.NewErr().
							SetCode(ErrMerchantEconomicActivityIdAlreadyExistCode).
							SetDescription("ID ALREADY EXIST").
							SetLevel(errDomain.LevelError).
							SetHttpStatus(http.StatusConflict).
							SetLayer(errDomain.UseCase).
							SetFunction("CreateEconomicActivities")
	ErrMerchantEconomicActivityIdHasBeenDeleted = errDomain.NewErr().
							SetCode(ErrMerchantEconomicActivityIdHasBeenDeletedCode).
							SetDescription("ID HAS BEEN DELETED").
							SetLevel(errDomain.LevelError).
							SetHttpStatus(http.StatusConflict).
							SetLayer(errDomain.UseCase).
							SetFunction("DeleteModule")
)
