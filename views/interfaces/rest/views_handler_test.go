/*
 * File: views_handler_test.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to handler for views.
 *
 * Last Modified: 2023-11-24
 */

package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	authRest "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	viewsDomain "gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
	mockViews "gitlab.smartcitiesperu.com/smartone/api-core/views/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestViewsHandler_GetViews(t *testing.T) {
	t.Run("When views are successfully listed", func(t *testing.T) {
		viewUseCaseMock := &mockViews.ViewUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		moduleId := uuid.New().String()
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		viewUseCaseMock.
			On("GetViews", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return([]viewsDomain.View{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewViewsHandler(viewUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/views", moduleId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while listing views", func(t *testing.T) {
		viewUseCaseMock := &mockViews.ViewUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		moduleId := uuid.New().String()
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		viewUseCaseMock.
			On("GetViews", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewViewsHandler(viewUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/views", moduleId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestViewsHandler_CreateView(t *testing.T) {
	t.Run("When add a view, successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		viewUseCaseMock := &mockViews.ViewUseCase{}

		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = viewsDomain.CreateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		viewUseCaseMock.
			On("CreateView", mock.Anything, mock.Anything, mock.Anything).
			Return(&viewId, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewsHandler(viewUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/views", moduleId)
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When add a view, error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		viewUseCaseMock := &mockViews.ViewUseCase{}

		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = viewsDomain.CreateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		viewUseCaseMock.
			On("CreateView", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewsHandler(viewUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/views", moduleId)
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestViewsHandler_UpdateView(t *testing.T) {
	t.Run("When update a view successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		viewUseCaseMock := &mockViews.ViewUseCase{}

		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = viewsDomain.UpdateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		viewUseCaseMock.
			On("UpdateView", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewsHandler(viewUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/views/%s", moduleId, viewId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update a view error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		viewUseCaseMock := &mockViews.ViewUseCase{}

		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = viewsDomain.UpdateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		viewUseCaseMock.
			On("UpdateView", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewsHandler(viewUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/views/%s", moduleId, viewId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestViewsHandler_DeleteView(t *testing.T) {
	t.Run("When delete a view by id successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		viewUseCaseMock := &mockViews.ViewUseCase{}

		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		viewUseCaseMock.
			On("DeleteView", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewsHandler(viewUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/views/%s", moduleId, viewId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete a view error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		viewUseCaseMock := &mockViews.ViewUseCase{}

		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		commentsError := errors.New("random error")
		viewUseCaseMock.
			On("DeleteView", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewViewsHandler(viewUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/modules/%s/views/%s", moduleId, viewId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
