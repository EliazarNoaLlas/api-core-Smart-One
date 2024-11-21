/*
 * File: merchants_handler_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the merchants Handler.
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

	merchantsDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
	mockMerchants "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestMerchant_GetMerchant(t *testing.T) {
	t.Run("When trying to retrieve a merchants list successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantsUCMock := &mockMerchants.MerchantUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		merchantsUCMock.
			On("GetMerchants",
				mock.Anything,
				mock.Anything).
			Return([]merchantsDomain.Merchant{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewMerchantsHandler(merchantsUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/merchants", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run(
		"When an error occurs during the attempt to retrieve merchants", func(t *testing.T) {
			authUCase := mockAuth.NewAuthUseCase(t)
			authMiddleware := authRest.NewAuthMiddleware(authUCase)
			merchantsUCMock := &mockMerchants.MerchantUseCase{}

			userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
			authUCase.
				On("DecodeToken",
					mock.Anything,
					mock.Anything).
				Return(&userId, nil)
			expectedError := errors.New("random error")
			merchantsUCMock.
				On("GetMerchants",
					mock.Anything,
					mock.Anything).
				Return(nil, &paramsDomain.PaginationResults{}, expectedError)
			gin.SetMode(gin.TestMode)
			context, router := gin.CreateTestContext(httptest.NewRecorder())

			NewMerchantsHandler(merchantsUCMock, router, authMiddleware)
			context.Request, _ = http.NewRequest("GET", "/api/v1/core/merchants", nil)
			authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
			context.Request.Header.Set("Authorization", authorizationHeader)
			xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
			context.Request.Header.Set("x-Tenant-Id", xTenantId)
			router.ServeHTTP(context.Writer, context.Request)
			assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
		})
}

func TestMerchant_CreateMerchant(t *testing.T) {
	t.Run("When you successfully create a new merchant", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantsUseCaseMock := &mockMerchants.MerchantUseCase{}
		merchantID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		merchantsUseCaseMock.
			On("CreateMerchant",
				mock.Anything,
				mock.Anything,
			).Return(&merchantID, nil)
		body := merchantsDomain.CreateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantsHandler(merchantsUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/merchants", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create merchant error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantsUseCaseMock := &mockMerchants.MerchantUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		expectedError := errors.New("random error")
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		merchantsUseCaseMock.
			On("CreateMerchant",
				mock.Anything,
				mock.Anything).
			Return(nil, expectedError)
		body := merchantsDomain.CreateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantsHandler(merchantsUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/merchants", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestMerchant_UpdateMerchant(t *testing.T) {
	t.Run("When you successfully update a new merchant", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantsUseCaseMock := &mockMerchants.MerchantUseCase{}
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		merchantsUseCaseMock.
			On("UpdateMerchant",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil)
		var body = merchantsDomain.UpdateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantsHandler(merchantsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/merchants/%s", merchantId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update merchant error", func(t *testing.T) {
		merchantsUseCaseMock := &mockMerchants.MerchantUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		expectedError := errors.New("random error")
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		merchantsUseCaseMock.
			On("UpdateMerchant",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(expectedError)
		var body = merchantsDomain.UpdateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantsHandler(merchantsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/merchants/%s", merchantId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestMerchant_DeleteMerchant(t *testing.T) {
	t.Run("When delete merchants by id successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantsUseCaseMock := &mockMerchants.MerchantUseCase{}
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		merchantsUseCaseMock.
			On("DeleteMerchant",
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantsHandler(merchantsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/merchants/%s", merchantId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete merchant error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		merchantsUseCaseMock := &mockMerchants.MerchantUseCase{}
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		commentsError := errors.New("random error")
		merchantsUseCaseMock.
			On("DeleteMerchant",
				mock.Anything,
				mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewMerchantsHandler(merchantsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/merchants/%s", merchantId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
