/*
 * File: user_types_handler_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to handler for user types.
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

	userTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain"
	mockUserTypes "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerUserTypes_GetUserType(t *testing.T) {
	t.Run("When get user types successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userTypesUCMock := &mockUserTypes.UserTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		userTypesUCMock.
			On("GetUserTypes", mock.Anything, mock.Anything).
			Return([]userTypesDomain.UserType{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUserTypesHandler(userTypesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/user_types", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When get user types error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userTypesUCMock := &mockUserTypes.UserTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		expectedError := errors.New("random error")
		userTypesUCMock.
			On("GetUserTypes", mock.Anything, mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUserTypesHandler(userTypesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/user_types", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUserTypes_CreateUserType(t *testing.T) {
	t.Run("When create user type successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userTypesUseCaseMock := &mockUserTypes.UserTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = userTypesDomain.CreateUserTypeBody{
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
			Enable:      true,
		}
		userTypeID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userTypesUseCaseMock.
			On(
				"CreateUserType",
				mock.Anything,
				mock.Anything,
			).
			Return(&userTypeID, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserTypesHandler(userTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/user_types", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create user type error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userTypesUseCaseMock := &mockUserTypes.UserTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = userTypesDomain.CreateUserTypeBody{
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
			Enable:      true,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		userTypesUseCaseMock.
			On("CreateUserType", mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserTypesHandler(userTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/user_types", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUserTypes_UpdateUserType(t *testing.T) {
	t.Run("When update user type successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userTypesUseCaseMock := &mockUserTypes.UserTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = userTypesDomain.CreateUserTypeBody{
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
			Enable:      true,
		}
		userTypesUseCaseMock.
			On("UpdateUserType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserTypesHandler(userTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("PUT",
			"/api/v1/core/user_types/739bbbc9-7e93-11ee-89fd-0242ac110016", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update user type error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userTypesUseCaseMock := &mockUserTypes.UserTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = userTypesDomain.CreateUserTypeBody{
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
			Enable:      true,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		userTypesUseCaseMock.
			On("UpdateUserType", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserTypesHandler(userTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("PUT",
			"/api/v1/core/user_types/739bbbc9-7e93-11ee-89fd-0242ac110016", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUserTypes_DeleteUserType(t *testing.T) {
	t.Run("When delete user type by id successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userTypesUseCaseMock := &mockUserTypes.UserTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		userTypesUseCaseMock.
			On("DeleteUserType", mock.Anything, mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserTypesHandler(userTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("DELETE",
			"/api/v1/core/user_types/739bbbc9-7e93-11ee-89fd-0242ac110016", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete user type error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userTypesUseCaseMock := &mockUserTypes.UserTypeUseCase{}
		expectedError := errors.New("any error")

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		userTypesUseCaseMock.
			On("DeleteUserType", mock.Anything, mock.Anything).
			Return(false, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUserTypesHandler(userTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("DELETE",
			"/api/v1/core/user_types/739bbbc9-7e93-11ee-89fd-0242ac110016", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
