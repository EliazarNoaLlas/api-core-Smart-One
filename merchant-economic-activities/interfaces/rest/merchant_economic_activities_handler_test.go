/*
 * File: merchant_economic_activities_handler_test.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the route handler of the merchant economic activities.
 *
 * Last Modified: 2023-12-05
 */

package interfaces

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"encoding/json"
	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	authRest "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
	mockMerchantEconomicActivities "gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain/mocks"
)

const (
	fakeToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9tYWNzYWx1ZC5zdGcuZXJwLm9uc2NwLmNvbVwvYXBpXC9jb3JlXC91c3Vhcmlvc1wvZ2VuZXJhcl90b2tlbl91c3VhcmlvIiwiaWF0IjoxNjkyMjE1NjE3LCJleHAiOjE2OTMwNzk2MTcsIm5iZiI6MTY5MjIxNTYxNywianRpIjoiWUxtbk9iOTMwcHVoY3NGRyIsInN1YiI6MzkwLCJwcnYiOiIyM2JkNWM4OTQ5ZjYwMGFkYjM5ZTcwMWM0MDA4NzJkYjdhNTk3NmY3In0.6xC79TgmyMFTH4TMdljBscs6aRt8VjLgL-wvl4jvpC4"
)

func TestMerchantEconomicActivities_GetMerchantEconomicActivities(t *testing.T) {
	t.Run("When merchant economic activities are successfully listed", func(t *testing.T) {
		merchantEconomicActivitiesUCMock := &mockMerchantEconomicActivities.MerchantEconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		merchantEconomicActivitiesUCMock.
			On("GetMerchantEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return([]domain.MerchantEconomicActivity{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewMerchantEconomicActivitiesHandler(merchantEconomicActivitiesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET",
			"/api/v1/core/merchant_economic_activities/703e039e-92be-11ee-a040-0242ac11000e", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while listing a merchant economic activities", func(t *testing.T) {
		merchantEconomicActivitiesUCMock := &mockMerchantEconomicActivities.MerchantEconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		merchantEconomicActivitiesUCMock.
			On("GetMerchantEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewMerchantEconomicActivitiesHandler(merchantEconomicActivitiesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET",
			"/api/v1/core/merchant_economic_activities/703e039e-92be-11ee-a040-0242ac11000e", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestMerchantEconomicActivities_CreateMerchantEconomicActivity(t *testing.T) {
	t.Run("When to successfully create a merchant economic activities", func(t *testing.T) {
		merchantEconomicActivitiesUCMock := &mockMerchantEconomicActivities.MerchantEconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantEconomicActivityId := "2d22f1d9-9380-11ee-a040-0242ac11000e"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		var body = domain.CreateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		merchantEconomicActivitiesUCMock.
			On(
				"CreateEconomicActivity",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(&merchantEconomicActivityId, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantEconomicActivitiesHandler(merchantEconomicActivitiesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST",
			"/api/v1/core/merchant_economic_activities/url/22d4b62a-9380-11ee-a040-0242ac11000e", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create merchant economic activities error", func(t *testing.T) {
		merchantEconomicActivitiesUCMock := &mockMerchantEconomicActivities.MerchantEconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		var body = domain.CreateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		merchantEconomicActivitiesUCMock.
			On("CreateEconomicActivity", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantEconomicActivitiesHandler(merchantEconomicActivitiesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST",
			"/api/v1/core/merchant_economic_activities/url/22d4b62a-9380-11ee-a040-0242ac11000e", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestMerchantEconomicActivities_UpdateMerchantEconomicActivity(t *testing.T) {
	t.Run("When a merchant economic activities is successfully updated", func(t *testing.T) {
		merchantEconomicActivitiesUCMock := &mockMerchantEconomicActivities.MerchantEconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantEconomicActivityId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		merchantEconomicActivityBody := domain.UpdateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		merchantEconomicActivitiesUCMock.
			On("UpdateEconomicActivity", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(merchantEconomicActivityBody)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantEconomicActivitiesHandler(merchantEconomicActivitiesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/merchant_economic_activities/%s", merchantEconomicActivityId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while updating a merchant economic activities", func(t *testing.T) {
		merchantEconomicActivitiesUCMock := &mockMerchantEconomicActivities.MerchantEconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantEconomicActivityId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		merchantEconomicActivityBody := domain.UpdateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		jsonValue, _ := json.Marshal(merchantEconomicActivityBody)
		expectedError := errors.New("random error")
		merchantEconomicActivitiesUCMock.
			On("UpdateEconomicActivity", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantEconomicActivitiesHandler(merchantEconomicActivitiesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/merchant_economic_activities/%s", merchantEconomicActivityId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestMerchantEconomicActivities_DeleteMerchantEconomicActivity(t *testing.T) {
	t.Run("When a merchant economic activities is successfully deleted", func(t *testing.T) {
		merchantEconomicActivitiesUCMock := &mockMerchantEconomicActivities.MerchantEconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantEconomicActivityId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		merchantEconomicActivitiesUCMock.
			On("DeleteEconomicActivity", mock.Anything, mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantEconomicActivitiesHandler(merchantEconomicActivitiesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/merchant_economic_activities/%s", merchantEconomicActivityId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while deleting a merchant economic activities", func(t *testing.T) {
		merchantEconomicActivitiesUCMock := &mockMerchantEconomicActivities.MerchantEconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantEconomicActivityId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		commentsError := errors.New("random error")
		merchantEconomicActivitiesUCMock.
			On("DeleteEconomicActivity", mock.Anything, mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantEconomicActivitiesHandler(merchantEconomicActivitiesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/merchant_economic_activities/%s", merchantEconomicActivityId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
