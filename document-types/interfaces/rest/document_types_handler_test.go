/*
 * File: document_types_handler_test.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-07
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

	"gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain"
	mocksDocumentTypes "gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain/mocks"
)

const (
	fakeToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9tYWNzYWx1ZC5zdGcuZXJwLm9uc2NwLmNvbVwvYXBpXC9jb3JlXC91c3Vhcmlvc1wvZ2VuZXJhcl90b2tlbl91c3VhcmlvIiwiaWF0IjoxNjkyMjE1NjE3LCJleHAiOjE2OTMwNzk2MTcsIm5iZiI6MTY5MjIxNTYxNywianRpIjoiWUxtbk9iOTMwcHVoY3NGRyIsInN1YiI6MzkwLCJwcnYiOiIyM2JkNWM4OTQ5ZjYwMGFkYjM5ZTcwMWM0MDA4NzJkYjdhNTk3NmY3In0.6xC79TgmyMFTH4TMdljBscs6aRt8VjLgL-wvl4jvpC4"
)

func TestHandlerDocumentTypes_GetDocumentTypes(t *testing.T) {
	t.Run("When document types are successfully listed", func(t *testing.T) {
		DocumentTypesUCMock := &mocksDocumentTypes.DocumentTypeUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		DocumentTypesUCMock.
			On("GetDocumentTypes", mock.Anything, mock.Anything, mock.Anything).
			Return([]domain.DocumentType{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewDocumentTypesHandler(DocumentTypesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/document_types/", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while listing document types", func(t *testing.T) {
		DocumentTypesUCMock := &mocksDocumentTypes.DocumentTypeUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		DocumentTypesUCMock.
			On("GetDocumentTypes", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewDocumentTypesHandler(DocumentTypesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/document_types/", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})

	t.Run("List of document types should be return successfully.", func(t *testing.T) {
		DocumentTypesUCMock := &mocksDocumentTypes.DocumentTypeUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		DocumentTypesUCMock.
			On("GetDocumentTypes", mock.Anything, mock.Anything, mock.Anything).
			Return([]domain.DocumentType{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewDocumentTypesHandler(DocumentTypesUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET",
			"/api/v1/core/document_types/?description=DOCUMENTO NACIONAL DE IDENTIDAD", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})
}

func TestHandlerDocumentTypes_CreateDocumentType(t *testing.T) {
	t.Run("When to successfully create a document type", func(t *testing.T) {
		usersUseCaseMock := &mocksDocumentTypes.DocumentTypeUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		createUserBody := domain.CreateDocumentTypeBody{
			Number:                 "01",
			Description:            "DOCUMENTO NACIONAL DE IDENTIDAD",
			AbbreviatedDescription: "DNI",
			Enable:                 1,
		}
		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		usersUseCaseMock.
			On(
				"CreateDocumentType",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(&documentTypeId, nil)
		jsonValue, _ := json.Marshal(createUserBody)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewDocumentTypesHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/document_types/create_document_types/%s", documentTypeId)
		context.Request, _ = http.NewRequest("POST",
			url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create document type error", func(t *testing.T) {
		usersUseCaseMock := &mocksDocumentTypes.DocumentTypeUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		createUserBody := domain.CreateDocumentTypeBody{
			Number:                 "01",
			Description:            "DOCUMENTO NACIONAL DE IDENTIDAD",
			AbbreviatedDescription: "DNI",
			Enable:                 1,
		}
		jsonValue, _ := json.Marshal(createUserBody)
		expectedError := errors.New("random error")
		usersUseCaseMock.
			On("CreateDocumentType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewDocumentTypesHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/document_types/create_document_types/%s", documentTypeId)
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

func TestHandlerDocumentTypes_UpdateDocumentType(t *testing.T) {
	t.Run("When a document type is successfully updated", func(t *testing.T) {
		usersUseCaseMock := &mocksDocumentTypes.DocumentTypeUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		updateUserBody := domain.CreateDocumentTypeBody{
			Number:                 "01",
			Description:            "DOCUMENTO NACIONAL DE IDENTIDAD",
			AbbreviatedDescription: "DNI",
			Enable:                 1,
		}
		usersUseCaseMock.
			On("UpdateDocumentType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(updateUserBody)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewDocumentTypesHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/document_types/update_document_types/%s", userId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while updating a document type", func(t *testing.T) {
		usersUseCaseMock := &mocksDocumentTypes.DocumentTypeUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		updateUserBody := domain.CreateDocumentTypeBody{
			Number:                 "01",
			Description:            "DOCUMENTO NACIONAL DE IDENTIDAD",
			AbbreviatedDescription: "DNI",
			Enable:                 1,
		}
		jsonValue, _ := json.Marshal(updateUserBody)
		expectedError := errors.New("random error")
		usersUseCaseMock.
			On("UpdateDocumentType", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewDocumentTypesHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/document_types/update_document_types/%s", userId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerDocumentTypes_DeleteDocumentType(t *testing.T) {
	t.Run("When a document type is successfully deleted", func(t *testing.T) {
		usersUseCaseMock := &mocksDocumentTypes.DocumentTypeUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		usersUseCaseMock.
			On("DeleteDocumentType", mock.Anything, mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewDocumentTypesHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/document_types/delete_document_types/%s", userId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while deleting a document type", func(t *testing.T) {
		usersUseCaseMock := &mocksDocumentTypes.DocumentTypeUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		commentsError := errors.New("random error")
		usersUseCaseMock.
			On("DeleteDocumentType", mock.Anything, mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewDocumentTypesHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/document_types/delete_document_types/%s", userId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
