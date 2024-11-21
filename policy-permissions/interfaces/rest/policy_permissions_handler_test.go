/*
 * File: policyPermissions_handler_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to handler for policyPermissions.
 *
 * Last Modified: 2023-11-20
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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	authRest "gitlab.smartcitiesperu.com/smartone/api-shared/auth/interfaces/rest"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	policyPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
	mockPolicyPermissions "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerPolicyPermissions_GetPolicyPermissionsByPolicy(t *testing.T) {
	t.Run("When policy permissions of id are successfully listed", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policyPermissionsUseCaseMock := &mockPolicyPermissions.PolicyPermissionUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)

		policyPermission := make([]policyPermissionsDomain.PolicyPermission, 0)

		policyPermissionsUseCaseMock.
			On("GetPolicyPermissionsByPolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(policyPermission, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewPolicyPermissionsHandler(policyPermissionsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s/permissions", policyId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while getting policy permissions by id", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policyPermissionsUseCaseMock := &mockPolicyPermissions.PolicyPermissionUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		policyPermissionsUseCaseMock.
			On("GetPolicyPermissionsByPolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewPolicyPermissionsHandler(policyPermissionsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s/permissions", policyId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerPolicyPermissions_CreatePolicyPermission(t *testing.T) {
	t.Run("When create a permission to policy, successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policyPermissionsUseCaseMock := &mockPolicyPermissions.PolicyPermissionUseCase{}
		policyPermissionId := uuid.New().String()
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		policyPermissionsUseCaseMock.
			On("CreatePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(&policyPermissionId, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPolicyPermissionsHandler(policyPermissionsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s/permissions", policyId)
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create a permission to policy, error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policyPermissionsUseCaseMock := &mockPolicyPermissions.PolicyPermissionUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		policyPermissionsUseCaseMock.
			On("CreatePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPolicyPermissionsHandler(policyPermissionsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s/permissions", policyId)
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerPolicyPermissions_UpdatePolicyPermission(t *testing.T) {
	t.Run("When update a permission from policy successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policyPermissionsUseCaseMock := &mockPolicyPermissions.PolicyPermissionUseCase{}
		policyId := uuid.New().String()
		policyPermissionId := uuid.New().String()
		userId := uuid.New().String()
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		policyPermissionsUseCaseMock.
			On("UpdatePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPolicyPermissionsHandler(policyPermissionsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s/permissions/%s", policyId, policyPermissionId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update a permission from policy error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policyPermissionsUseCaseMock := &mockPolicyPermissions.PolicyPermissionUseCase{}
		policyId := uuid.New().String()
		policyPermissionId := uuid.New().String()

		userId := uuid.New().String()
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		var body = policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		policyPermissionsUseCaseMock.
			On("UpdatePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPolicyPermissionsHandler(policyPermissionsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s/permissions/%s", policyId, policyPermissionId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerPolicyPermissions_DeletePolicyPermission(t *testing.T) {
	t.Run("When delete a permissions from policy by id successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policyPermissionsUseCaseMock := &mockPolicyPermissions.PolicyPermissionUseCase{}
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionId := "739bbbc9-7e93-11ee-89fd-0242ac110016"

		userId := uuid.New().String()
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		policyPermissionsUseCaseMock.
			On("DeletePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPolicyPermissionsHandler(policyPermissionsUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s/permissions/%s", merchantId, policyPermissionId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete a permission from policy error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policyPermissionsUseCaseMock := &mockPolicyPermissions.PolicyPermissionUseCase{}
		commentsError := errors.New("random error")
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		policyPermissionsUseCaseMock.
			On("DeletePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPolicyPermissionsHandler(policyPermissionsUseCaseMock, router, authMiddleware)
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/policies/%s/permissions/%s", merchantId, policyPermissionId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
