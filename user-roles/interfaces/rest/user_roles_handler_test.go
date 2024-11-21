/*
 * File: user_roles_handler_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to handler for userRoles.
 *
 * Last Modified: 2023-11-23
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

	userRolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
	mockUserRoles "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerUserRoles_GetUserRolesByUser(t *testing.T) {
	t.Run("When roles by user of id are successfully listed", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userRolesUseCaseMock := &mockUserRoles.UserRoleUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		userRoles := make([]userRolesDomain.UserRole, 0)

		userRolesUseCaseMock.
			On("GetUserRolesByUser", mock.Anything, mock.Anything, mock.Anything).
			Return(userRoles, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUserRolesHandler(userRolesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s/roles", userId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while getting policy roles by user", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userRolesUseCaseMock := &mockUserRoles.UserRoleUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		userRolesUseCaseMock.
			On("GetUserRolesByUser", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUserRolesHandler(userRolesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s/roles", userId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUserRoles_CreateUserRole(t *testing.T) {
	t.Run("When add a role to user, successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userRolesUseCaseMock := &mockUserRoles.UserRoleUseCase{}

		userRoleID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		var body = userRolesDomain.CreateUserRoleBody{
			RoleId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable: true,
		}

		authUCase.
			On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		userRolesUseCaseMock.
			On("CreateUserRole", mock.Anything, mock.Anything, mock.Anything).
			Return(&userRoleID, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserRolesHandler(userRolesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST",
			"/api/v1/core/users/739bbbc9-7e93-11ee-89fd-0442ac210931/roles", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When add a role to user, error", func(t *testing.T) {
		userRolesUseCaseMock := &mockUserRoles.UserRoleUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = userRolesDomain.CreateUserRoleBody{
			RoleId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable: true,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		userRolesUseCaseMock.
			On("CreateUserRole", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserRolesHandler(userRolesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST",
			"/api/v1/core/users/739bbbc9-7e93-11ee-89fd-0442ac210931/roles", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUserRoles_UpdateUserRole(t *testing.T) {
	t.Run("When update a role of user successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userRolesUseCaseMock := &mockUserRoles.UserRoleUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = userRolesDomain.CreateUserRoleBody{
			RoleId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable: true,
		}
		userRolesUseCaseMock.
			On("UpdateUserRole", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserRolesHandler(userRolesUseCaseMock, router, authMiddleware)
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/users/%s/roles/%s", userId, userRoleId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update a role of user error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userRolesUseCaseMock := &mockUserRoles.UserRoleUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = userRolesDomain.CreateUserRoleBody{
			RoleId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable: true,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		userRolesUseCaseMock.
			On("UpdateUserRole", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserRolesHandler(userRolesUseCaseMock, router, authMiddleware)
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/users/%s/roles/%s", userId, userRoleId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUserRoles_DeleteUserRole(t *testing.T) {
	t.Run("When delete a role from user successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userRolesUseCaseMock := &mockUserRoles.UserRoleUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		userRolesUseCaseMock.
			On("DeleteUserRole", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserRolesHandler(userRolesUseCaseMock, router, authMiddleware)
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/users/%s/roles/%s", userId, userRoleId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete a role from user error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userRolesUseCaseMock := &mockUserRoles.UserRoleUseCase{}
		commentsError := errors.New("random error")
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		userRolesUseCaseMock.
			On("DeleteUserRole", mock.Anything, mock.Anything, mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserRolesHandler(userRolesUseCaseMock, router, authMiddleware)
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/users/%s/roles/%s", userId, userRoleId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
