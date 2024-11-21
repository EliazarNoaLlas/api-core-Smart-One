/*
 * File: role_policies_handler_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to handler for rolePolicies.
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

	rolePoliciesDomain "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain"
	mockRolePolicies "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain/mocks"
)

const (
	fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExMDE1NTAsImlhdCI6MTcwMTA5Nzk1MCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTA5Nzk1MCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.Dmx0qHToCxFj73cmU-ouSp9zN78GRwFnC4Cy_LOR1cU"
)

func TestHandlerRolePolicies_CreateRolePolicy(t *testing.T) {
	t.Run("When add a policy to role, successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = rolePoliciesDomain.CreateRolePolicyBody{
			PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:   true,
		}
		rolePolicyID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		rolePoliciesUseCaseMock.
			On("CreateRolePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(&rolePolicyID, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST",
			"/api/v1/core/roles/739bbbc9-7e93-11ee-89fd-0442ac210931/policies", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When add a policy to role, error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = rolePoliciesDomain.CreateRolePolicyBody{
			PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:   true,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		rolePoliciesUseCaseMock.
			On("CreateRolePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST",
			"/api/v1/core/roles/739bbbc9-7e93-11ee-89fd-0442ac210931/policies", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerRolePolicies_CreateRolePolicies(t *testing.T) {
	t.Run("When add multiple policies to role, successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = rolePoliciesDomain.CreateMultipleRolePoliciesBody{
			RolePolicies: []rolePoliciesDomain.CreateRolePolicyBody{
				{
					PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
					Enable:   true,
				},
			},
		}
		rolePolicyIDs := []string{"739bbbc9-7e93-11ee-89fd-0242ac110016"}
		rolePoliciesUseCaseMock.
			On("CreateRolePolicies", mock.Anything, mock.Anything, mock.Anything).
			Return(rolePolicyIDs, nil)

		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		url := fmt.Sprintf("/api/v1/core/roles/%s/policies/batch", roleId)

		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When add multiple policies to role, error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = rolePoliciesDomain.CreateMultipleRolePoliciesBody{
			RolePolicies: []rolePoliciesDomain.CreateRolePolicyBody{
				{
					PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
					Enable:   true,
				},
			},
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		rolePoliciesUseCaseMock.
			On("CreateRolePolicies", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		url := fmt.Sprintf("/api/v1/core/roles/%s/policies/batch", roleId)
		context.Request, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerRolePolicies_UpdateRolePolicy(t *testing.T) {
	t.Run("When update a policy of role successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = rolePoliciesDomain.CreateRolePolicyBody{
			PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:   true,
		}
		rolePoliciesUseCaseMock.
			On("UpdateRolePolicy", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/roles/%s/policies/%s", roleId, rolePolicyId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When update a policy of role error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		var body = rolePoliciesDomain.CreateRolePolicyBody{
			PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:   true,
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		rolePoliciesUseCaseMock.
			On("UpdateRolePolicy", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/roles/%s/policies/%s", roleId, rolePolicyId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerRolePolicies_DeleteRolePolicy(t *testing.T) {
	t.Run("When delete a policy from role by id successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		rolePoliciesUseCaseMock.
			On("DeleteRolePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/roles/%s/policies/%s", roleId, rolePolicyId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete a policy from role error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		commentsError := errors.New("random error")
		rolePoliciesUseCaseMock.
			On("DeleteRolePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		url := fmt.Sprintf("/api/v1/core/roles/%s/policies/%s", roleId, rolePolicyId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerRolePolicies_DeleteRolePolicies(t *testing.T) {
	t.Run("When delete multiple policies by role successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		rolePoliciesUseCaseMock.
			On("DeleteRolePolicies", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		var body = deleteMultipleRolePoliciesValidate{
			RolePolicyIds: []string{"739bbbc9-7e93-11ee-89fd-042hs5278420"},
		}
		jsonValue, _ := json.Marshal(body)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		url := fmt.Sprintf("/api/v1/core/roles/%s/policies/batch", roleId)

		context.Request, _ = http.NewRequest("DELETE", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When delete multiple policies of a role, error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		rolePoliciesUseCaseMock := &mockRolePolicies.RolePolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		commentsError := errors.New("random error")
		rolePoliciesUseCaseMock.
			On("DeleteRolePolicies", mock.Anything, mock.Anything, mock.Anything).
			Return(commentsError)

		var body = deleteMultipleRolePoliciesValidate{
			RolePolicyIds: []string{"739bbbc9-7e93-11ee-89fd-042hs5278420"},
		}
		jsonValue, _ := json.Marshal(body)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewRolePoliciesHandler(rolePoliciesUseCaseMock, router, authMiddleware)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		url := fmt.Sprintf("/api/v1/core/roles/%s/policies/batch", roleId)
		context.Request, _ = http.NewRequest("DELETE", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
