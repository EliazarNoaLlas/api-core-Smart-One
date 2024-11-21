/*
 * File: store_types_handler_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the store types Handler.
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

	storeTypeDomain "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain"
	mockStoreTypes "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerStoreTypes_GetStoreTypes(t *testing.T) {
	t.Run("When get store types  successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		storeTypesUCMock := &mockStoreTypes.StoreTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		storeTypesUCMock.
			On("GetStoreTypes", mock.Anything, mock.Anything).
			Return([]storeTypeDomain.StoreType{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewStoreTypesHandler(storeTypesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/store_types", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When get store types  error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		storeTypesUCMock := &mockStoreTypes.StoreTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		expectedError := errors.New("random error")
		storeTypesUCMock.
			On("GetStoreTypes", mock.Anything, mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewStoreTypesHandler(storeTypesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/store_types", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerStoreTypes_CreateStoreType(t *testing.T) {
	t.Run("When create store type  successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		storeTypesUseCaseMock := &mockStoreTypes.StoreTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = storeTypeDomain.CreateStoreTypeBody{
			Description:  "Maquinaria",
			Abbreviation: "Mac",
		}
		storeTypeID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		storeTypesUseCaseMock.
			On(
				"CreateStoreType",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(&storeTypeID, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoreTypesHandler(storeTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/store_types", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create store type  error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		storeTypesUseCaseMock := &mockStoreTypes.StoreTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = storeTypeDomain.CreateStoreTypeBody{
			Description:  "Maquinaria",
			Abbreviation: "Mac",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		storeTypesUseCaseMock.On(
			"CreateStoreType",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoreTypesHandler(storeTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/store_types", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerStoreTypes_UpdateStoreType(t *testing.T) {
	t.Run("When update store type  successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		storeTypesUseCaseMock := &mockStoreTypes.StoreTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = storeTypeDomain.CreateStoreTypeBody{
			Description:  "Maquinaria",
			Abbreviation: "Mac",
		}
		storeTypesUseCaseMock.
			On(
				"UpdateStoreType",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoreTypesHandler(storeTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("PUT",
			"/api/v1/core/store_types/73900000-7e93-11ee-89fd-0242a500000", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update store type  error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		storeTypesUseCaseMock := &mockStoreTypes.StoreTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = storeTypeDomain.CreateStoreTypeBody{
			Description:  "Maquinaria",
			Abbreviation: "Mac",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		storeTypesUseCaseMock.On(
			"UpdateStoreType",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoreTypesHandler(storeTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("PUT",
			"/api/v1/core/store_types/73900000-7e93-11ee-89fd-0242a500000", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerStoreTypes_DeleteStoreType(t *testing.T) {
	t.Run("When delete store type  by id successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		storeTypesUseCaseMock := &mockStoreTypes.StoreTypeUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		storeTypesUseCaseMock.
			On("DeleteStoreType", mock.Anything, mock.Anything).
			Return(true, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoreTypesHandler(storeTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("DELETE",
			"/api/v1/core/store_types/73900000-7e93-11ee-89fd-0242a500000", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete store type error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		storeTypesUseCaseMock := &mockStoreTypes.StoreTypeUseCase{}
		commentsError := errors.New("any error")

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		storeTypesUseCaseMock.
			On("DeleteStoreType", mock.Anything, mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoreTypesHandler(storeTypesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("DELETE",
			"/api/v1/core/store_types/73900000-7e93-11ee-89fd-0242a500000", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
