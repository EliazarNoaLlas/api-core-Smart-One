/*
 * File: permissions_handler_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the permissions Handler.
 *
 * Last Modified: 2023-11-15
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
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	permissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
	mockPermissions "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerPermissions_GetPermissions(t *testing.T) {
	t.Run("Successful permission retrieval.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		permissionsUCMock := &mockPermissions.PermissionUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)

		permissionsUCMock.
			On("GetPermissions",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return([]permissionsDomain.Permission{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewPermissionsHandler(permissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/permissions", moduleId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("Error during permissions retrieval.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		permissionsUCMock := &mockPermissions.PermissionUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		expectedError := errors.New("random error")
		permissionsUCMock.
			On("GetPermissions",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPermissionsHandler(permissionsUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/permissions", moduleId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerPermissions_CreatePermission(t *testing.T) {
	t.Run("Successful permission creation.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		permissionsUseCaseMock := &mockPermissions.PermissionUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = permissionsDomain.CreatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
			ModuleId:    "cddbfacf-8305-11ee-89fd-024255555502",
		}
		permissionID := "fcdbfacf-8305-11ee-89fd-0242555555"
		permissionsUseCaseMock.
			On("CreatePermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(&permissionID, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPermissionsHandler(permissionsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/permissions", moduleId)
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("Error during permission creation.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		permissionsUseCaseMock := &mockPermissions.PermissionUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = permissionsDomain.CreatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
			ModuleId:    "cddbfacf-8305-11ee-89fd-024255555502",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		permissionsUseCaseMock.
			On("CreatePermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPermissionsHandler(permissionsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/permissions", moduleId)
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerPermissions_UpdatePermission(t *testing.T) {
	t.Run("Successful permission update.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		permissionsUseCaseMock := &mockPermissions.PermissionUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = permissionsDomain.CreatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
			ModuleId:    "cddbfacf-8305-11ee-89fd-024255555502",
		}
		permissionsUseCaseMock.
			On("UpdatePermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPermissionsHandler(permissionsUseCaseMock, router, authMiddleware)
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionId := "7355bbc9-7e93-11ee-89fd-924000000016"
		url := fmt.Sprintf("/api/v1/core/modules/%s/permissions/%s", moduleId, permissionId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("Error during role update.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		permissionsUseCaseMock := &mockPermissions.PermissionUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = permissionsDomain.CreatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
			ModuleId:    "cddbfacf-8305-11ee-89fd-024255555502",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		permissionsUseCaseMock.
			On("UpdatePermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPermissionsHandler(permissionsUseCaseMock, router, authMiddleware)
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionId := "7355bbc9-7e93-11ee-89fd-924000000016"
		url := fmt.Sprintf("/api/v1/core/modules/%s/permissions/%s", moduleId, permissionId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerPermissions_DeletePermission(t *testing.T) {
	t.Run("Successful permission deletion.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		permissionsUseCaseMock := &mockPermissions.PermissionUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		permissionsUseCaseMock.
			On("DeletePermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPermissionsHandler(permissionsUseCaseMock, router, authMiddleware)
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionId := "7355bbc9-7e93-11ee-89fd-924000000016"
		url := fmt.Sprintf("/api/v1/core/modules/%s/permissions/%s", moduleId, permissionId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("Error during permission deletion.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		permissionsUseCaseMock := &mockPermissions.PermissionUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		commentsError := errors.New("random error")
		permissionsUseCaseMock.
			On("DeletePermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPermissionsHandler(permissionsUseCaseMock, router, authMiddleware)
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		url := fmt.Sprintf("/api/v1/core/modules/%s/permissions/%s", moduleId, permissionId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
