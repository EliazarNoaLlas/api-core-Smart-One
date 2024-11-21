/*
 * File: receipt_types_handler_test.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the tests for the receiptTypes handler.
 *
 * Last Modified: 2024-03-06
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

	receiptTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
	mockReceiptTypes "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerReceiptTypes_GetReceiptTypes(t *testing.T) {
	t.Run("When to successfully get receipt types", func(t *testing.T) {
		receiptTypesUCMock := &mockReceiptTypes.ReceiptTypesUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		receiptTypesUCMock.
			On("GetReceiptTypes", mock.Anything, mock.Anything).
			Return([]receiptTypesDomain.ReceiptType{}, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewReceiptTypesHandler(receiptTypesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/receipt_types")
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When get receipt types error", func(t *testing.T) {
		receiptTypesUCMock := &mockReceiptTypes.ReceiptTypesUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		expectedError := errors.New("random error")
		receiptTypesUCMock.
			On("GetReceiptTypes", mock.Anything, mock.Anything).
			Return(nil, expectedError)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewReceiptTypesHandler(receiptTypesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/receipt_types")
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerReceiptTypes_CreateReceiptType(t *testing.T) {
	t.Run("When create receipt type successfully", func(t *testing.T) {
		receiptTypesUCMock := &mockReceiptTypes.ReceiptTypesUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = receiptTypesDomain.CreateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}
		receiptTypesId := "739bbbc9-7e93-11ee-89fd-0242ac110016"

		receiptTypesUCMock.
			On("CreateReceiptType", mock.Anything, mock.Anything, mock.Anything).
			Return(&receiptTypesId, nil)

		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewReceiptTypesHandler(receiptTypesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/receipt_types")
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create receipt type error", func(t *testing.T) {
		receiptTypesUCMock := &mockReceiptTypes.ReceiptTypesUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = receiptTypesDomain.CreateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		receiptTypesUCMock.
			On("CreateReceiptType", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewReceiptTypesHandler(receiptTypesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/receipt_types")
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerReceiptTypes_UpdateReceiptType(t *testing.T) {
	t.Run("When update receipt type successfully", func(t *testing.T) {
		receiptTypesUCMock := &mockReceiptTypes.ReceiptTypesUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = receiptTypesDomain.UpdateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}
		jsonValue, _ := json.Marshal(body)
		receiptTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110060"

		receiptTypesUCMock.
			On("UpdateReceiptType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewReceiptTypesHandler(receiptTypesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/receipt_types/%s", receiptTypeId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update receipt type error", func(t *testing.T) {
		receiptTypesUCMock := &mockReceiptTypes.ReceiptTypesUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = receiptTypesDomain.UpdateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}
		jsonValue, _ := json.Marshal(body)
		receiptTypesId := "739bbbc9-7e93-11ee-89fd-0242ac110000"
		expectedError := errors.New("random error")

		receiptTypesUCMock.
			On("UpdateReceiptType", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewReceiptTypesHandler(receiptTypesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/receipt_types/%s", receiptTypesId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerReceiptTypes_DeleteReceiptType(t *testing.T) {
	t.Run("When delete receipt type by id successfully", func(t *testing.T) {
		receiptTypesUCMock := &mockReceiptTypes.ReceiptTypesUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		receiptTypesId := "739bbbc9-7e93-11ee-89fd-0242ac110000"

		receiptTypesUCMock.
			On("DeleteReceiptType", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewReceiptTypesHandler(receiptTypesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/receipt_types/%s", receiptTypesId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete receipt type error", func(t *testing.T) {
		receiptTypesUCMock := &mockReceiptTypes.ReceiptTypesUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		commentsError := errors.New("random error")
		receiptTypesId := "739bbbc9-7e93-11ee-89fd-0242ac110000"

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		receiptTypesUCMock.
			On("DeleteReceiptType", mock.Anything, mock.Anything, mock.Anything).
			Return(false, commentsError)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewReceiptTypesHandler(receiptTypesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/receipt_types/%s", receiptTypesId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
