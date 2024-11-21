/*
 * File: stores_handler_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to handler for stores.
 *
 * Last Modified: 2023-11-14
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

	storesDomain "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain"
	mockStores "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerStores_GetStore(t *testing.T) {
	t.Run("When get stores successfully", func(t *testing.T) {
		storesUCMock := &mockStores.StoreUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac1100100"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&userId, nil)
		storesUCMock.
			On("GetStores", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return([]storesDomain.Store{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewStoresHandler(storesUCMock, router, authMiddleware)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		url := fmt.Sprintf("/api/v1/core/merchants/%s/stores", merchantId)
		context.Request, _ = http.NewRequest("GET",
			url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When get stores error", func(t *testing.T) {
		storesUCMock := &mockStores.StoreUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&userId, nil)
		expectedError := errors.New("random error")
		storesUCMock.
			On("GetStores", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewStoresHandler(storesUCMock, router, authMiddleware)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		url := fmt.Sprintf("/api/v1/core/merchants/%s/stores", merchantId)
		context.Request, _ = http.NewRequest("GET",
			url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerStores_CreateStore(t *testing.T) {
	t.Run("When create store successfully", func(t *testing.T) {
		storesUseCaseMock := &mockStores.StoreUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = storesDomain.CreateStoreBody{
			Name:        "Obra av. 28 julio",
			Shortname:   "Obra 28",
			StoreTypeId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
		}
		storeID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		storesUseCaseMock.
			On("CreateStore", mock.Anything, mock.Anything, mock.Anything).
			Return(&storeID, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoresHandler(storesUseCaseMock, router, authMiddleware)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		url := fmt.Sprintf("/api/v1/core/merchants/%s/stores", merchantId)
		context.Request, _ = http.NewRequest("POST",
			url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create store error", func(t *testing.T) {
		storesUseCaseMock := &mockStores.StoreUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = storesDomain.CreateStoreBody{
			Name:        "Obra av. 28 julio",
			Shortname:   "Obra 28",
			StoreTypeId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		storesUseCaseMock.
			On("CreateStore", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoresHandler(storesUseCaseMock, router, authMiddleware)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		url := fmt.Sprintf("/api/v1/core/merchants/%s/stores", merchantId)
		context.Request, _ = http.NewRequest("POST",
			url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerStores_UpdateStore(t *testing.T) {
	t.Run("When update store successfully", func(t *testing.T) {
		storesUseCaseMock := &mockStores.StoreUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = storesDomain.CreateStoreBody{
			Name:        "Obra av. 28 julio",
			Shortname:   "Obra 28",
			StoreTypeId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
		}
		storesUseCaseMock.
			On("UpdateStore", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoresHandler(storesUseCaseMock, router, authMiddleware)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/merchants/%s/stores/%s", merchantId, storeId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update store error", func(t *testing.T) {
		storesUseCaseMock := &mockStores.StoreUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = storesDomain.CreateStoreBody{
			Name:        "Obra av. 28 julio",
			Shortname:   "Obra 28",
			StoreTypeId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		storesUseCaseMock.
			On("UpdateStore", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoresHandler(storesUseCaseMock, router, authMiddleware)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/merchants/%s/stores/%s", merchantId, storeId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerStores_DeleteStore(t *testing.T) {
	t.Run("When delete stores by id successfully", func(t *testing.T) {
		storesUseCaseMock := &mockStores.StoreUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		storesUseCaseMock.
			On("DeleteStore", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoresHandler(storesUseCaseMock, router, authMiddleware)
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/merchants/%s/stores/%s", storeId, merchantId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete store error", func(t *testing.T) {
		storesUseCaseMock := &mockStores.StoreUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		commentsError := errors.New("any error")
		storesUseCaseMock.
			On("DeleteStore", mock.Anything, mock.Anything, mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewStoresHandler(storesUseCaseMock, router, authMiddleware)
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/merchants/%s/stores/%s", storeId, merchantId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
