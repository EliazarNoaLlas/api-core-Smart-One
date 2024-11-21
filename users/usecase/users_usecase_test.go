/*
 * File: users_usecase_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of users.
 *
 * Last Modified: 2023-11-23
 */

package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	mockValidation "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain/mocks"

	usersDomain "gitlab.smartcitiesperu.com/smartone/api-core/users/domain"
	mockUsers "gitlab.smartcitiesperu.com/smartone/api-core/users/domain/mocks"
)

func TestUseCaseUsers_GetUser(t *testing.T) {
	t.Run("When user are successfully listed", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		user := usersDomain.User{}
		usersRepository.
			On("GetUser", mock.Anything, mock.Anything).
			Return(&user, nil)
		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, err := userUCase.GetUser(context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016")
		assert.NoError(t, err)
		assert.EqualValues(t, res, &user)
	})

	t.Run("When an error occurs while listing user", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		user := usersDomain.User{}
		expectedError := errors.New("random error")
		usersRepository.
			On("GetUser", mock.Anything, mock.Anything).
			Return(&user, expectedError)
		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, err := userUCase.GetUser(context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016")
		assert.EqualError(t, err, "random error")
		assert.Nil(t, res)
	})
}

func TestUseCaseUsers_GetUsers(t *testing.T) {
	t.Run("When users are successfully listed", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		usersRepository.
			On("GetUsers", mock.Anything, mock.Anything, mock.Anything).
			Return([]usersDomain.UserMultiple{}, nil)
		usersRepository.
			On("GetTotalUsers", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		searchParams := usersDomain.GetUsersParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		users, _, err := usersUCase.GetUsers(context.Background(), searchParams, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, users, []usersDomain.UserMultiple{})
	})

	t.Run("When an error occurs while listing users", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		usersRepository.
			On("GetUsers", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		usersRepository.
			On("GetTotalUsers", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		searchParams := usersDomain.GetUsersParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		users, _, err := usersUCase.GetUsers(context.Background(), searchParams, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, users, []usersDomain.UserMultiple(nil))
	})
}

func TestUseCaseUsers_GetMenuByUser(t *testing.T) {
	t.Run("When get menu of user, successfully", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		modulesByUser := make([]usersDomain.ModuleMenuUser, 0)
		modules := make([]usersDomain.Module, 0)
		usersRepository.
			On("GetMenuByUser", mock.Anything, mock.Anything).
			Return(modulesByUser, nil)
		usersRepository.
			On("GetModules", mock.Anything).
			Return(modules, nil)
		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := userUCase.GetMenuByUser(context.Background(), userId)
		assert.NoError(t, err)
		menu := make([]usersDomain.MenuModule, 0)
		assert.EqualValues(t, res, menu)
	})

	t.Run("When an error occurs while get menu of user", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		modulesByUser := make([]usersDomain.ModuleMenuUser, 0)
		modules := make([]usersDomain.Module, 0)
		expectedError := errors.New("random error")
		usersRepository.
			On("GetMenuByUser", mock.Anything, mock.Anything).
			Return(modulesByUser, expectedError)
		usersRepository.
			On("GetModules", mock.Anything).
			Return(modules, nil)
		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := userUCase.GetMenuByUser(context.Background(), userId)
		assert.EqualError(t, err, "random error")
		assert.Nil(t, res)
	})
}

func TestUseCaseUsers_GetMeByUser(t *testing.T) {
	t.Run("When get me by user, successfully", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		personByUser := usersDomain.UserMeInfo{}
		stores := []usersDomain.StoreByUser{
			{
				Id:       "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Name:     "Benito",
				Merchant: usersDomain.Merchant{},
			},
		}
		merchants := []usersDomain.MerchantByUser{
			{
				Id:     "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Name:   "Benito",
				Stores: []usersDomain.Store{},
			},
		}
		userMe := usersDomain.UserMe{
			Id:        personByUser.Id,
			UserName:  personByUser.UserName,
			CreatedAt: personByUser.CreatedAt,
			Person:    personByUser.Person,
			RoleUser:  personByUser.RoleUser,
			Stores:    stores,
			Merchants: merchants,
		}

		usersRepository.
			On("GetMeByUser", mock.Anything, mock.Anything).
			Return(&personByUser, nil)
		usersRepository.
			On("GetStoresByUser", mock.Anything, mock.Anything).
			Return(stores, nil)
		usersRepository.
			On("GetMerchantsByUser", mock.Anything, mock.Anything).
			Return(merchants, nil)

		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := userUCase.GetMeByUser(context.Background(), userId)
		if err != nil {
			return
		}
		assert.NoError(t, err)
		assert.EqualValues(t, res, &userMe)
	})

	t.Run("When an error occurs while get me by user", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		expectedError := errors.New("random error")
		usersRepository.
			On("GetMeByUser", mock.Anything, mock.Anything).
			Return(nil, expectedError)
		usersRepository.
			On("GetStoresByUser", mock.Anything, mock.Anything).
			Return(nil, expectedError)
		usersRepository.
			On("GetMerchantsByUser", mock.Anything, mock.Anything).
			Return(nil, expectedError)

		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := userUCase.GetMeByUser(context.Background(), userId)
		assert.EqualError(t, err, "random error")
		assert.Nil(t, res)
	})
}

func TestUseCaseUsers_CreateUser(t *testing.T) {
	t.Run("When to successfully create a user", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		usersRepository.On("CreateUserMain", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(userID, nil)
		usersRepository.
			On("CreateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&userID, nil)
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		_, err := usersUCase.CreateUser(
			context.Background(),
			usersDomain.CreateUserBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When an error occurs while creating a user and that user already exists", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(true, nil)
		usersRepository.
			On("CreateUser", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		usersRepository.On("CreateUserMain", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		_, err := usersUCase.CreateUser(
			context.Background(),
			usersDomain.CreateUserBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, "ERR_USER_USERNAME_ALREADY_EXIST")
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateUser")
	})

	t.Run("When an error occurs while creating a user", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		errCreate := errDomain.NewErr().SetFunction("CreateUser").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		errCreateUserMain := errDomain.NewErr().SetFunction("CreateUserMain").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		usersRepository.
			On("CreateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		usersRepository.On("CreateUserMain", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreateUserMain)
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		_, err := usersUCase.CreateUser(
			context.Background(),
			usersDomain.CreateUserBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateUser")
	})
}

func TestUseCaseUsers_UpdateUser(t *testing.T) {
	t.Run("When a user is successfully updated", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		usersRepository.On("VerifyIfUserExist", mock.Anything, mock.Anything).
			Return(nil)
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		usersRepository.
			On("UpdateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		err := usersUCase.UpdateUser(
			context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
			usersDomain.UpdateUserBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When an error occurs while updating a user", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).Return(true, nil)
		usersRepository.On("VerifyIfUserExist", mock.Anything, mock.Anything).
			Return(nil)
		usersRepository.
			On("UpdateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		err := usersUCase.UpdateUser(
			context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
			usersDomain.UpdateUserBody{},
		)
		assert.Error(t, err)
	})
}

func TestUseCaseUsers_DeleteUser(t *testing.T) {
	t.Run("When a user is successfully deleted", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		usersRepository.
			On("DeleteUser", mock.Anything, mock.Anything).
			Return(true, nil)
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := usersUCase.DeleteUser(context.Background(), userId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When an error occurs while deleting a user", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		usersError := errors.New("random error")
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		usersRepository.
			On("DeleteUser", mock.Anything, mock.Anything).
			Return(false, usersError)
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := usersUCase.DeleteUser(context.Background(), userId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}

func TestUseCaseUsers_ResetPasswordUser(t *testing.T) {
	t.Run("When a user's password is updated", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		usersRepository.
			On("ResetPasswordUser", mock.Anything, mock.Anything, mock.Anything).
			Return(true, errors.New("some error"))
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, err := usersUCase.ResetPasswordUser(
			context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
			usersDomain.ResetUserPasswordBody{
				NewPassword: "pepito",
			})
		assert.Error(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When an error occurs while updating a user", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).Return(nil)
		usersRepository.
			On("ResetPasswordUser", mock.Anything, mock.Anything, mock.Anything).
			Return(false, errors.New("random error"))
		usersUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, err := usersUCase.ResetPasswordUser(
			context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
			usersDomain.ResetUserPasswordBody{
				NewPassword: "pepito",
			})
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}

func TestUseCaseUsers_LoginUser(t *testing.T) {
	t.Run("When user are successfully listed", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userName := "pepito.quispe@smartc.pe"
		password := "pepitoPass"
		loginUserBody := usersDomain.LoginUserBody{
			UserName: userName,
			Password: password,
		}
		user := usersDomain.User{
			Id:       userId,
			UserName: userName,
		}
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDEwMjM0NzgsImlhdCI6MTcwMTAxOTg3OCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTAxOTg3OCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.K7hBqmlf-LyxJ_cNUBEEPgh_gpvie7uiQI0HqUrn3rY"
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		usersRepository.
			On("GetUserByUserNameAndPassword", mock.Anything, mock.Anything, mock.Anything).
			Return(&user, &xTenantId, nil)
		authRepository.
			On("GenerateToken", userId).
			Return(&token, nil)
		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, _, err := userUCase.LoginUser(context.Background(), loginUserBody)
		assert.NoError(t, err)
		assert.EqualValues(t, res, &token)
	})

	t.Run("When an error occurs while listing user", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userName := "pepito.quispe@smartc.pe"
		password := "pepitoPass"
		loginUserBody := usersDomain.LoginUserBody{
			UserName: userName,
			Password: password,
		}
		user := usersDomain.User{
			Id:       userId,
			UserName: userName,
		}
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDEwMjM0NzgsImlhdCI6MTcwMTAxOTg3OCwiaXNzIjoiaHR0cDovL21hY3NhbHVkLXYyLnN0Zy5lcnAub25zY3AuY29tL2FwaS9jb3JlL3VzdWFyaW9zL2xvZ2luIiwianRpIjoiRWhUeEU0SXU2SXMwUXNCcCIsIm5iZiI6MTcwMTAxOTg3OCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyIsInN1YiI6IjkxZmI4NmJkLWRhNDYtNDE0Yi05N2ExLWZjZGFhOGNkMzVkMSJ9.K7hBqmlf-LyxJ_cNUBEEPgh_gpvie7uiQI0HqUrn3rY"
		expectedError := errors.New("random error")
		usersRepository.
			On("GetUserByUserNameAndPassword", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, nil, expectedError)
		authRepository.
			On("GenerateToken", userId).
			Return(&token, nil)
		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, _, err := userUCase.LoginUser(context.Background(), loginUserBody)
		assert.EqualError(t, err, "random error")
		assert.Nil(t, res, &user)
	})
}

func TestUseCaseUsers_VerifyPermissionsByUser(t *testing.T) {
	t.Run("When verify permission by user return true", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110018"
		codePermission := "CREATE_PRODUCT"
		usersRepository.
			On("VerifyPermissionsByUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)

		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, err := userUCase.VerifyPermissionsByUser(context.Background(), userId, storeId, codePermission)
		assert.NoError(t, err)
		assert.EqualValues(t, true, res)
	})

	t.Run("When verify permission by user return an error", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110018"
		codePermission := "CREATE_PRODUCT"

		expectedError := errors.New("random error")
		usersRepository.
			On("VerifyPermissionsByUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(false, expectedError)

		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, err := userUCase.VerifyPermissionsByUser(context.Background(), userId, storeId, codePermission)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}

func TestUseCaseUsers_GetModulePermissions(t *testing.T) {
	t.Run("When it returns the list of permissions per module of a user successfully", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		codeModule := "logistics.requirements"

		permissions := []usersDomain.Permissions{
			{
				Id:   "0c4001f3-2dd8-4d9f-820d-db7d7d8c85c0",
				Code: "CREATE_REQUIREMENT",
			},
			{
				Id:   "0c4001f3-2dd8-4d9f-820d-db7d7d8c85c0",
				Code: "UPDATE_REQUIREMENT",
			},
		}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		usersRepository.
			On("GetModulePermissions", mock.Anything, mock.Anything, mock.Anything).
			Return(permissions, nil)

		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, err := userUCase.GetModulePermissions(context.Background(), userId, codeModule)

		assert.NoError(t, err)
		assert.EqualValues(t, res, permissions)
	})

	t.Run("When the list of permissions per module of a user returns an error", func(t *testing.T) {
		usersRepository := &mockUsers.UserRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		codeModule := "logistics.requirements"
		expectedError := errors.New("random error")

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		usersRepository.
			On("GetModulePermissions", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)

		userUCase := NewUsersUseCase(usersRepository, validationRepository, authRepository, 60)
		res, err := userUCase.GetModulePermissions(context.Background(), userId, codeModule)

		assert.EqualError(t, err, "random error")
		assert.Nil(t, res)
	})
}
