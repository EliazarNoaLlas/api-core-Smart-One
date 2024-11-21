/*
 * File: users_handler_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to handler for users.
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

	usersDomain "gitlab.smartcitiesperu.com/smartone/api-core/users/domain"
	mockUsers "gitlab.smartcitiesperu.com/smartone/api-core/users/domain/mocks"
)

const (
	fakeToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9tYWNzYWx1ZC5zdGcuZXJwLm9uc2NwLmNvbVwvYXBpXC9jb3JlXC91c3Vhcmlvc1wvZ2VuZXJhcl90b2tlbl91c3VhcmlvIiwiaWF0IjoxNjkyMjE1NjE3LCJleHAiOjE2OTMwNzk2MTcsIm5iZiI6MTY5MjIxNTYxNywianRpIjoiWUxtbk9iOTMwcHVoY3NGRyIsInN1YiI6MzkwLCJwcnYiOiIyM2JkNWM4OTQ5ZjYwMGFkYjM5ZTcwMWM0MDA4NzJkYjdhNTk3NmY3In0.6xC79TgmyMFTH4TMdljBscs6aRt8VjLgL-wvl4jvpC4"
)

func TestHandlerUsers_GetUser(t *testing.T) {
	t.Run("When user are successfully listed", func(t *testing.T) {
		usersUCMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		user := usersDomain.User{}
		usersUCMock.
			On("GetUser", mock.Anything, mock.Anything).
			Return(&user, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUsersHandler(usersUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s", userId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while listing user", func(t *testing.T) {
		usersUCMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		usersUCMock.
			On("GetUser", mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUsersHandler(usersUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s", userId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUsers_GetUsers(t *testing.T) {
	t.Run("When users are successfully listed", func(t *testing.T) {
		usersUCMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		usersUCMock.
			On("GetUsers", mock.Anything, mock.Anything, mock.Anything).
			Return([]usersDomain.UserMultiple{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUsersHandler(usersUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/users", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while listing users", func(t *testing.T) {
		usersUCMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		usersUCMock.
			On("GetUsers", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, &paramsDomain.PaginationResults{}, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUsersHandler(usersUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET", "/api/v1/core/users", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})

	t.Run("List of users filtered by module, successfully.", func(t *testing.T) {
		usersUCMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		usersUCMock.
			On("GetUsers", mock.Anything, mock.Anything, mock.Anything).
			Return([]usersDomain.UserMultiple{}, &paramsDomain.PaginationResults{}, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUsersHandler(usersUCMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("GET",
			"/api/v1/core/users?type_id=739bbbc9-7e93-11ee-89fd-0242ac110018", nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})
}

func TestHandlerUsers_GetMenuByUser(t *testing.T) {
	t.Run("When menu of user are successfully listed", func(t *testing.T) {
		usersUCMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		menu := make([]usersDomain.MenuModule, 0)

		usersUCMock.
			On("GetMenuByUser", mock.Anything, mock.Anything).
			Return(menu, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUsersHandler(usersUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s/menu", userId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while getting menu by user", func(t *testing.T) {
		usersUCMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		usersUCMock.
			On("GetMenuByUser", mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUsersHandler(usersUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s/menu", userId)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUsers_GetMeByUser(t *testing.T) {
	t.Run("When user me  are successfully listed", func(t *testing.T) {
		usersUCMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		userMe := usersDomain.UserMe{}

		usersUCMock.
			On("GetMeByUser", mock.Anything, mock.Anything).
			Return(&userMe, nil)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUsersHandler(usersUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/me")
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while getting user", func(t *testing.T) {
		usersUCMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		expectedError := errors.New("random error")
		usersUCMock.
			On("GetMeByUser", mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())

		NewUsersHandler(usersUCMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/me")
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUsers_CreateUser(t *testing.T) {
	t.Run("When to successfully create a user", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		var body = usersDomain.CreateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			Password:   "pepitoPass",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
		}
		usersUseCaseMock.
			On(
				"CreateUser",
				mock.Anything,
				mock.Anything,
			).
			Return(&userId, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST",
			"/api/v1/core/users", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusCreated, context.Writer.Status())
	})

	t.Run("When create user error", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		var body = usersDomain.CreateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			Password:   "pepitoPass",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		usersUseCaseMock.
			On("CreateUser", mock.Anything, mock.Anything).
			Return(nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST",
			"/api/v1/core/users", bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUsers_UpdateUser(t *testing.T) {
	t.Run("When a user is successfully updated", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		body := usersDomain.UpdateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210900",
		}
		usersUseCaseMock.
			On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s", userId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while updating a user", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		var body = usersDomain.UpdateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		usersUseCaseMock.
			On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s", userId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUsers_DeleteUser(t *testing.T) {
	t.Run("When a user is successfully deleted", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		usersUseCaseMock.
			On("DeleteUser", mock.Anything, mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s", userId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.NoError(t, nil)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When an error occurs while deleting a user", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		commentsError := errors.New("random error")
		usersUseCaseMock.
			On("DeleteUser", mock.Anything, mock.Anything).
			Return(false, commentsError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s", userId)
		context.Request, _ = http.NewRequest("DELETE", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUsers_ResetPasswordUser(t *testing.T) {
	t.Run("When a user updated their password successfully", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		body := usersDomain.ResetUserPasswordBody{
			NewPassword: "pepitoPass",
		}

		usersUseCaseMock.
			On("ResetPasswordUser", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s/password", userId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("when an error occurs while changing the password", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)

		body := usersDomain.ResetUserPasswordBody{
			NewPassword: "pepitoPass",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		usersUseCaseMock.
			On("ResetPasswordUser", mock.Anything, mock.Anything, mock.Anything).
			Return(false, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/%s/password", userId)
		context.Request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUsers_LoginUser(t *testing.T) {
	t.Run("when a user logs in successfully", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)

		var body = usersDomain.LoginUserBody{
			UserName: "pepito.quispe@smartc.pe",
			Password: "pepitoPass",
		}
		token := fakeToken
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		usersUseCaseMock.
			On("LoginUser", mock.Anything, mock.Anything).
			Return(&token, &xTenantId, nil)
		jsonValue, _ := json.Marshal(body)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonValue))
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When user logs in and shows an error", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		var body = usersDomain.LoginUserBody{
			UserName: "pepito.quispe@smartc.pe",
			Password: "pepitoPass",
		}
		jsonValue, _ := json.Marshal(body)
		expectedError := errors.New("random error")
		usersUseCaseMock.
			On("LoginUser", mock.Anything, mock.Anything).
			Return(nil, nil, expectedError)
		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		context.Request, _ = http.NewRequest("POST",
			"/api/v1/auth/login", bytes.NewBuffer(jsonValue))
		router.ServeHTTP(context.Writer, context.Request)
		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUsers_VerifyPermissionsByUser(t *testing.T) {
	t.Run("When verify permission by user return true", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		codePermission := "CREATE_PRODUCT"

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		usersUseCaseMock.
			On("VerifyPermissionsByUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/me/permissions/%s?store_id=9fa66a3b-d25b-4304-800d-1200735bcc4f", codePermission)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When verify permission by user return an error", func(t *testing.T) {
		usersUseCaseMock := &mockUsers.UserUseCase{}
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		expectedError := errors.New("random error")
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		codePermission := "CREATE_PRODUCT"

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		usersUseCaseMock.
			On("VerifyPermissionsByUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(false, expectedError)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUseCaseMock, router, authMiddleware)
		url := fmt.Sprintf("/api/v1/core/users/me/permissions/%s?store_id=9fa66a3b-d25b-4304-800d-1200735bcc4f", codePermission)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}

func TestHandlerUsers_GetModulePermissions(t *testing.T) {
	t.Run("When it returns the list of permissions per module of a user successfully", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		usersUCMock := &mockUsers.UserUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		codeModule := "logistics.requirements"

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		usersUCMock.
			On("GetModulePermissions", mock.Anything, mock.Anything, mock.Anything).
			Return([]usersDomain.Permissions{}, nil)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUCMock, router, authMiddleware)

		url := fmt.Sprintf("/api/v1/core/users/me/modules/%s/permissions", codeModule)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("When the list of permissions per module of a user returns an error", func(t *testing.T) {
		authUCase := mockAuth.NewAuthUseCase(t)
		authMiddleware := authRest.NewAuthMiddleware(authUCase)
		usersUCMock := &mockUsers.UserUseCase{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		codeModule := "logistics.requirements"
		expectedError := errors.New("random error")

		authUCase.On("DecodeToken", mock.Anything, mock.Anything).
			Return(&userId, nil)
		usersUCMock.
			On("GetModulePermissions", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)

		gin.SetMode(gin.TestMode)
		context, router := gin.CreateTestContext(httptest.NewRecorder())
		NewUsersHandler(usersUCMock, router, authMiddleware)

		url := fmt.Sprintf("/api/v1/core/users/me/modules/%s/permissions", codeModule)
		context.Request, _ = http.NewRequest("GET", url, nil)
		authorizationHeader := fmt.Sprintf("Bearer %s", fakeToken)
		context.Request.Header.Set("Authorization", authorizationHeader)
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		context.Request.Header.Set("x-Tenant-Id", xTenantId)
		router.ServeHTTP(context.Writer, context.Request)

		assert.Equal(t, http.StatusInternalServerError, context.Writer.Status())
	})
}
