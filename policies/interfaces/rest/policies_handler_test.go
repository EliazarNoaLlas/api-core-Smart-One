/*
 * File: policies_handler_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to handler for policies.
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

	policiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain"
	mockPolicies "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain/mocks"
)

const (
	fakeToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9tYWNzYWx1ZC5zdGcuZXJwLm9uc2NwLmNvbVwvYXBpXC9jb3JlXC91c3Vhcmlvc1wvZ2VuZXJhcl90b2tlbl91c3VhcmlvIiwiaWF0IjoxNjkyMjE1NjE3LCJleHAiOjE2OTMwNzk2MTcsIm5iZiI6MTY5MjIxNTYxNywianRpIjoiWUxtbk9iOTMwcHVoY3NGRyIsInN1YiI6MzkwLCJwcnYiOiIyM2JkNWM4OTQ5ZjYwMGFkYjM5ZTcwMWM0MDA4NzJkYjdhNTk3NmY3In0.6xC79TgmyMFTH4TMdljBscs6aRt8VjLgL-wvl4jvpC4"
)

func pointerToStr(value string) *string {
	return &value
}

func pointerToBool(value bool) *bool {
	return &value
}

func TestHandlerPolicies_GetPolicy(t *testing.T) {
	t.Run("When policies are successfully listed", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policiesUCMock := &mockPolicies.PolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		policiesUCMock.
			On("GetPolicies",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return([]policiesDomain.Policy{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewPoliciesHandler(policiesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/policies", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while listing policies", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policiesUCMock := &mockPolicies.PolicyUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		policiesUCMock.
			On("GetPolicies",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewPoliciesHandler(policiesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies?%s", moduleId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})

	t.Run("List of policies filtered by module, successfully.", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policiesUCMock := &mockPolicies.PolicyUseCase{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		policiesUCMock.
			On("GetPolicies",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return([]policiesDomain.Policy{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewPoliciesHandler(policiesUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies?%s", moduleId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})
}

func TestHandlerPolicies_CreatePolicy(t *testing.T) {
	t.Run("When to successfully create a policy", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policiesUseCaseMock := &mockPolicies.PolicyUseCase{}
		policyID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		policiesUseCaseMock.
			On(
				"CreatePolicy",
				mock.Anything,
				mock.Anything,
			).
			Return(&policyID, nil)
		var body = policiesDomain.CreatePolicyBody{
			Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
			Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
			ModuleId:    "739bbbc9-7e93-11ee-89fd-0242ac110018",
			MerchantId:  pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110019"),
			StoreId:     pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110020"),
			Level:       "system",
			Enable:      pointerToBool(true),
		}
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPoliciesHandler(policiesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/policies", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create policy error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policiesUseCaseMock := &mockPolicies.PolicyUseCase{}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)

		body := policiesDomain.CreatePolicyBody{
			Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
			Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
			ModuleId:    "739bbbc9-7e93-11ee-89fd-0242ac110018",
			MerchantId:  pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110019"),
			StoreId:     pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110020"),
			Level:       "system",
			Enable:      pointerToBool(true),
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		policiesUseCaseMock.
			On("CreatePolicy",
				mock.Anything,
				mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPoliciesHandler(policiesUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/core/policies", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerPolicies_UpdatePolicy(t *testing.T) {
	t.Run("When a policy is successfully updated", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policiesUseCaseMock := &mockPolicies.PolicyUseCase{}
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)

		body := policiesDomain.CreatePolicyBody{
			Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
			Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
			ModuleId:    "739bbbc9-7e93-11ee-89fd-0242ac110018",
			MerchantId:  pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110019"),
			StoreId:     pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110020"),
			Level:       "system",
			Enable:      pointerToBool(true),
		}
		policiesUseCaseMock.
			On("UpdatePolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPoliciesHandler(policiesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s", policyId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while updating a policy", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policiesUseCaseMock := &mockPolicies.PolicyUseCase{}
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)

		body := policiesDomain.CreatePolicyBody{
			Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
			Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
			ModuleId:    "739bbbc9-7e93-11ee-89fd-0242ac110018",
			MerchantId:  pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110019"),
			StoreId:     pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110020"),
			Level:       "system",
			Enable:      pointerToBool(true),
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		policiesUseCaseMock.
			On("UpdatePolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPoliciesHandler(policiesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s", policyId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerPolicies_DeletePolicy(t *testing.T) {
	t.Run("When a policy is successfully deleted", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policiesUseCaseMock := &mockPolicies.PolicyUseCase{}
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)
		policiesUseCaseMock.
			On("DeletePolicy",
				mock.Anything,
				mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPoliciesHandler(policiesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s", policyId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while deleting a policy", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		policiesUseCaseMock := &mockPolicies.PolicyUseCase{}
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.
			On("DecodeToken",
				mock.Anything,
				mock.Anything).
			Return(&userId, nil)

		commentsError := errors.New("random error")
		policiesUseCaseMock.
			On("DeletePolicy",
				mock.Anything,
				mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewPoliciesHandler(policiesUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/policies/%s", policyId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
