/*
 * File: modules_handler_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to handler for modules.
 *
 * Last Modified: 2023-11-10
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

	modulesDomain "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"
	mockModules "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerModules_GetModules(t *testing.T) {
	t.Run("When get modules successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		modulesUCMock := &mockModules.ModuleUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		modulesUCMock.
			On("GetModules",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return([]modulesDomain.Module{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewModulesHandler(modulesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/modules", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When get modules error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		modulesUCMock := &mockModules.ModuleUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		expectedError := errors.New("random error")
		modulesUCMock.
			On("GetModules",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewModulesHandler(modulesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/modules", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})

	t.Run("List of modules successfully.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		modulesUCMock := &mockModules.ModuleUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)

		modulesUCMock.
			On("GetModules",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return([]modulesDomain.Module{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewModulesHandler(modulesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/modules", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})
}

func TestHandlerModules_CreateModule(t *testing.T) {
	t.Run("When create module successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		modulesUseCaseMock := &mockModules.ModuleUseCase{}
		moduleID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = modulesDomain.CreateModuleBody{
			Name:        "Logistica",
			Description: "Modulo de logistica",
			Code:        "logistic",
			Icon:        "fa fa-home",
			Position:    1,
		}
		modulesUseCaseMock.
			On(
				"CreateModule",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(&moduleID, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewModulesHandler(modulesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/modules", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create module error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		modulesUseCaseMock := &mockModules.ModuleUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = modulesDomain.CreateModuleBody{
			Name:        "Logistica",
			Description: "Modulo de logistica",
			Code:        "logistic",
			Icon:        "fa fa-home",
			Position:    1,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		modulesUseCaseMock.
			On("CreateModule",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewModulesHandler(modulesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/modules", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerModules_UpdateModule(t *testing.T) {
	t.Run("When update module successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		modulesUseCaseMock := &mockModules.ModuleUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = modulesDomain.UpdateModuleBody{
			Name:        "Logistica",
			Description: "Modulo de logistica",
			Code:        "logistic",
			Icon:        "fa fa-home",
			Position:    1,
		}
		modulesUseCaseMock.
			On("UpdateModule",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewModulesHandler(modulesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s", moduleId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update module error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		modulesUseCaseMock := &mockModules.ModuleUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken",
			mock.Anything,
			mock.Anything).
			Return(&userId, nil)
		var body = modulesDomain.CreateModuleBody{
			Name:        "Logistica",
			Description: "Modulo de logistica",
			Code:        "logistic",
			Icon:        "fa fa-home",
			Position:    1,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		modulesUseCaseMock.
			On("UpdateModule",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewModulesHandler(modulesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s", moduleId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))

		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerModules_DeleteModule(t *testing.T) {
	t.Run("When delete modules by id successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		modulesUseCaseMock := &mockModules.ModuleUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		modulesUseCaseMock.
			On("DeleteModule",
				mock.Anything,
				mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewModulesHandler(modulesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s", moduleId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete module error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		modulesUseCaseMock := &mockModules.ModuleUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		commentsError := errors.New("random error")
		modulesUseCaseMock.
			On("DeleteModule",
				mock.Anything,
				mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewModulesHandler(modulesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s", moduleId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
