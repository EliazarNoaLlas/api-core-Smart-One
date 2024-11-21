/*
 * File: server_handler_test.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the server handler test.
 *
 * Last Modified: 2024-04-09
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

	serverDomain "gitlab.smartcitiesperu.com/smartone/api-core/server/domain"
	mockServer "gitlab.smartcitiesperu.com/smartone/api-core/server/domain/mocks"
	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	authRest "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
)

const (
	fakeToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9tYWNzYWx1ZC5zdGcuZXJwLm9uc2NwLmNvbVwvYXBpXC9jb3JlXC91c3Vhcmlvc1wvZ2VuZXJhcl90b2tlbl91c3VhcmlvIiwiaWF0IjoxNjkyMjE1NjE3LCJleHAiOjE2OTMwNzk2MTcsIm5iZiI6MTY5MjIxNTYxNywianRpIjoiWUxtbk9iOTMwcHVoY3NGRyIsInN1YiI6MzkwLCJwcnYiOiIyM2JkNWM4OTQ5ZjYwMGFkYjM5ZTcwMWM0MDA4NzJkYjdhNTk3NmY3In0.6xC79TgmyMFTH4TMdljBscs6aRt8VjLgL-wvl4jvpC4"
)

func TestHandlerUsers_GetServerDatetime(t *testing.T) {
	t.Run("When get server datetime successfully", func(t *testing.T) {
		serverUCMock := &mockServer.ServerUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		serverDate := serverDomain.ServerDate{}
		serverUCMock.
			On("GetServerDate", mock.Anything).
			Return(&serverDate, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewServerHandler(serverUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/server/datetime")
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while get server datetime", func(t *testing.T) {
		serverUCMock := &mockServer.ServerUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		serverUCMock.
			On("GetServerDate", mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewServerHandler(serverUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/server/datetime")
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
