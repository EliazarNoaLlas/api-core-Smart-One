/*
 * File: economic_activities_handler_test.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the test handler.
 *
 * Last Modified: 2023-12-04
 */

package rest

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	authRest "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	economicActivitiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
	mockEconomicActivities "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain/mocks"
)

const (
	fakeToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9tYWNzYWx1ZC5zdGcuZXJwLm9uc2NwLmNvbVwvYXBpXC9jb3JlXC91c3Vhcmlvc1wvZ2VuZXJhcl90b2tlbl91c3VhcmlvIiwiaWF0IjoxNjkyMjE1NjE3LCJleHAiOjE2OTMwNzk2MTcsIm5iZiI6MTY5MjIxNTYxNywianRpIjoiWUxtbk9iOTMwcHVoY3NGRyIsInN1YiI6MzkwLCJwcnYiOiIyM2JkNWM4OTQ5ZjYwMGFkYjM5ZTcwMWM0MDA4NzJkYjdhNTk3NmY3In0.6xC79TgmyMFTH4TMdljBscs6aRt8VjLgL-wvl4jvpC4"
)

func TestHandlerEconomicActivities_GetEconomicActivities(t *testing.T) {
	t.Run("When economic activities are successfully listed", func(t *testing.T) {
		economicActivitiesUCMock := &mockEconomicActivities.EconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		economicActivitiesUCMock.
			On("GetEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return([]economicActivitiesDomain.EconomicActivity{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewEconomicActivitiesHandler(economicActivitiesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/economic_activities", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while listing economic activities", func(t *testing.T) {
		economicActivitiesUCMock := &mockEconomicActivities.EconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		economicActivitiesUCMock.
			On("GetEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewEconomicActivitiesHandler(economicActivitiesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/economic_activities", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})

	t.Run("List of economic activities filtered by module, successfully.", func(t *testing.T) {
		economicActivitiesUCMock := &mockEconomicActivities.EconomicActivityUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		economicActivitiesUCMock.
			On("GetEconomicActivities", mock.Anything, mock.Anything, mock.Anything).
			Return([]economicActivitiesDomain.EconomicActivity{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewEconomicActivitiesHandler(economicActivitiesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET",
			"/api/v1/core/economic_activities?cuui_id=0111", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})
}
