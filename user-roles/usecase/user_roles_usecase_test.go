/*
 * File: user_roles_usecase_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of userRoles.
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

	userRolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
	mockUserRoles "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain/mocks"
)

func TestUseCaseUserRoles_GetUserRolesByUser(t *testing.T) {
	t.Run("When roles by user are successfully listed", func(t *testing.T) {
		userRolesRepository := &mockUserRoles.UserRoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		userRolesRepository.
			On("GetUserRolesByUser", mock.Anything, mock.Anything, mock.Anything).
			Return([]userRolesDomain.UserRole{}, nil)
		userRolesRepository.
			On("GetTotalUserRolesByUser", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		pagination := paramsDomain.NewPaginationParams(nil)
		userRolesUCase := NewUserRolesUseCase(userRolesRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermission, _, err := userRolesUCase.GetUserRolesByUser(context.Background(), userId, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, policyPermission, []userRolesDomain.UserRole{})
	})

	t.Run("When an error occurs while listing roles by user", func(t *testing.T) {
		userRolesRepository := &mockUserRoles.UserRoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		userRolesRepository.
			On("GetUserRolesByUser", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		userRolesRepository.
			On("GetTotalUserRolesByUser", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		userRolesUCase := NewUserRolesUseCase(userRolesRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		pagination := paramsDomain.NewPaginationParams(nil)
		policyPermission, _, err := userRolesUCase.GetUserRolesByUser(context.Background(), userId, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, policyPermission, []userRolesDomain.UserRole(nil))
	})
}

func TestUseCaseUserRoles_CreateUserRole(t *testing.T) {
	roleHasPolicy := false
	t.Run("When add a role to user successfully", func(t *testing.T) {
		userRolesRepository := &mockUserRoles.UserRoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		roleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		createUserRoleBody := userRolesDomain.CreateUserRoleBody{
			RoleId: roleId,
			Enable: true,
		}
		userRolesRepository.
			On("CreateUserRole", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&userRoleId, nil)
		userRolesRepository.
			On("VerifyUserHasRole", mock.Anything, mock.Anything, mock.Anything).
			Return(roleHasPolicy, nil)
		userRolesUCase := NewUserRolesUseCase(userRolesRepository, validationRepository, authRepository, 60)
		_, err := userRolesUCase.CreateUserRole(context.Background(), userId, createUserRoleBody)
		assert.NoError(t, err)
	})

	t.Run("When adding a role to the user, and the user already exists, it shows us an error",
		func(t *testing.T) {
			userRolesRepository := &mockUserRoles.UserRoleRepository{}
			validationRepository := &mockValidation.ValidationRepository{}
			authRepository := &mockAuth.AuthRepository{}
			userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
			roleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
			createUserRoleBody := userRolesDomain.CreateUserRoleBody{
				RoleId: roleId,
				Enable: true,
			}
			userRolesRepository.
				On("CreateUserRole", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(nil, errors.New("random error"))
			userRolesRepository.
				On("VerifyUserHasRole", mock.Anything, mock.Anything, mock.Anything).
				Return(true, nil)
			userRolesUCase := NewUserRolesUseCase(userRolesRepository, validationRepository, authRepository, 60)
			_, err := userRolesUCase.CreateUserRole(context.Background(), userId, createUserRoleBody)
			assert.Error(t, err)

			var smartErr *errDomain.SmartError
			ok := errors.As(err, &smartErr)
			assert.Equal(t, ok, true)
			assert.Equal(t, smartErr.Code, userRolesDomain.ErrUserHasRoleAlreadyExistCode)
			assert.Equal(t, smartErr.Layer, errDomain.UseCase)
			assert.Equal(t, smartErr.Function, "CreateUserRole")
		})

	t.Run("When add a role to user, error", func(t *testing.T) {
		userRolesRepository := &mockUserRoles.UserRoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		roleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		createUserRoleBody := userRolesDomain.CreateUserRoleBody{
			RoleId: roleId,
			Enable: true,
		}

		errCreate := errDomain.NewErr().SetFunction("CreateUserRole").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		userRolesRepository.
			On("CreateUserRole", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		userRolesRepository.
			On("VerifyUserHasRole", mock.Anything, mock.Anything, mock.Anything).
			Return(roleHasPolicy, nil)
		userRolesUCase := NewUserRolesUseCase(userRolesRepository, validationRepository, authRepository, 60)
		_, err := userRolesUCase.CreateUserRole(context.Background(), userId, createUserRoleBody)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateUserRole")
	})

}

func TestUseCaseUserRoles_UpdateUserRole(t *testing.T) {
	t.Run("When update a role of user successfully", func(t *testing.T) {
		userRolesRepository := &mockUserRoles.UserRoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		userRolesRepository.
			On("UpdateUserRole", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		userRolesUCase := NewUserRolesUseCase(userRolesRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		roleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		createUserRoleBody := userRolesDomain.CreateUserRoleBody{
			RoleId: roleId,
			Enable: true,
		}
		err := userRolesUCase.UpdateUserRole(context.Background(), userId, userRoleId, createUserRoleBody)
		assert.NoError(t, err)
	})

	t.Run("When update a role of user, error", func(t *testing.T) {
		userRolesRepository := &mockUserRoles.UserRoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		userRolesRepository.
			On("UpdateUserRole", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		userRolesUCase := NewUserRolesUseCase(userRolesRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		roleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		createUserRoleBody := userRolesDomain.CreateUserRoleBody{
			RoleId: roleId,
			Enable: true,
		}
		err := userRolesUCase.UpdateUserRole(context.Background(), userId, userRoleId, createUserRoleBody)
		assert.Error(t, err)
	})
}

func TestUseCaseUserRoles_DeleteUserRole(t *testing.T) {
	t.Run("When delete a role of user successfully", func(t *testing.T) {
		userRolesRepository := &mockUserRoles.UserRoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		userRolesRepository.
			On("DeleteUserRole", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		userRolesUCase := NewUserRolesUseCase(userRolesRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := userRolesUCase.DeleteUserRole(context.Background(), userId, userRoleId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete a role of user, error", func(t *testing.T) {
		userRolesRepository := &mockUserRoles.UserRoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		userRolesError := errors.New("random error")
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		userRolesRepository.
			On("DeleteUserRole", mock.Anything, mock.Anything, mock.Anything).
			Return(false, userRolesError)
		userRolesUCase := NewUserRolesUseCase(userRolesRepository, validationRepository, authRepository, 60)
		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := userRolesUCase.DeleteUserRole(context.Background(), userId, userRoleId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
