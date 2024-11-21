/*
 * File: view_permissions_handler_test.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the tests for the viewPermissions handler.
 *
 * Last Modified: 2024-02-26
 */

package rest

import (
	"bytes"
	"encoding/json"
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

	ViewPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
	mockViewPermissions "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerViewPermissions_GetViewPermissions(t *testing.T) {
	t.Run("When to successfully get view permissions", func(t *testing.T) {
		viewPermissionsUCMock := &mockViewPermissions.ViewPermissionsUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110016"

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		viewPermissionsUCMock.
			On("GetViewPermissions", mock.Anything, mock.Anything).
			Return([]ViewPermissionsDomain.ViewPermission{}, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewPermissionsHandler(viewPermissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/views/%s/permissions", viewId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When get view permissions error", func(t *testing.T) {
		viewPermissionsUCMock := &mockViewPermissions.ViewPermissionsUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110017"

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		expectedError := errors.New("random error")
		viewPermissionsUCMock.
			On("GetViewPermissions", mock.Anything, mock.Anything).
			Return(nil, expectedError)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewPermissionsHandler(viewPermissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/views/%s/permissions", viewId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerViewPermissions_CreateViewPermission(t *testing.T) {
	t.Run("When create view permission successfully", func(t *testing.T) {
		viewPermissionsUCMock := &mockViewPermissions.ViewPermissionsUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = ViewPermissionsDomain.CreateViewPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-0242ac110015",
		}
		ViewPermissionsID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110017"

		viewPermissionsUCMock.
			On("CreateViewPermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&ViewPermissionsID, nil)

		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewPermissionsHandler(viewPermissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/views/%s/permissions", viewId)
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create view permission error", func(t *testing.T) {
		viewPermissionsUCMock := &mockViewPermissions.ViewPermissionsUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110017"

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = ViewPermissionsDomain.CreateViewPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-0242ac110015",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		viewPermissionsUCMock.
			On("CreateViewPermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewPermissionsHandler(viewPermissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/views/%s/permissions", viewId)
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerViewPermissions_UpdateViewPermission(t *testing.T) {
	t.Run("When update view permission successfully", func(t *testing.T) {
		viewPermissionsUCMock := &mockViewPermissions.ViewPermissionsUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = ViewPermissionsDomain.CreateViewPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-0242ac110015",
		}
		jsonValue, _ := json.Marshal(body)
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		viewPermissionId := "739bbbc9-7e93-11ee-89fd-0242ac110060"

		viewPermissionsUCMock.
			On("UpdateViewPermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewPermissionsHandler(viewPermissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/views/%s/permissions/%s", viewId, viewPermissionId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update view permission error", func(t *testing.T) {
		viewPermissionsUCMock := &mockViewPermissions.ViewPermissionsUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = ViewPermissionsDomain.CreateViewPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-0242ac110015",
		}
		jsonValue, _ := json.Marshal(body)
		viewPermissionsId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110070"
		expectedError := errors.New("random error")

		viewPermissionsUCMock.
			On("UpdateViewPermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewPermissionsHandler(viewPermissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/views/%s/permissions/%s", viewId, viewPermissionsId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerViewPermissions_DeleteViewPermission(t *testing.T) {
	t.Run("When delete view permission by id successfully", func(t *testing.T) {
		viewPermissionsUCMock := &mockViewPermissions.ViewPermissionsUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		viewPermissionsId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110004"
		viewPermissionsUCMock.
			On("DeleteViewPermission", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewPermissionsHandler(viewPermissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/views/%s/permissions/%s", viewId, viewPermissionsId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete view permission error", func(t *testing.T) {
		viewPermissionsUCMock := &mockViewPermissions.ViewPermissionsUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		viewPermissionsId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		commentsError := errors.New("random error")

		viewPermissionsUCMock.
			On("DeleteViewPermission", mock.Anything, mock.Anything, mock.Anything).
			Return(false, commentsError)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewPermissionsHandler(viewPermissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/views/%s/permissions/%s", viewId, viewPermissionsId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
