/*
 * File: roles_usecase_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the roles use case.
 *
 * Last Modified: 2023-11-14
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

	rolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/roles/domain"
	mockRoles "gitlab.smartcitiesperu.com/smartone/api-core/roles/domain/mocks"
)

func TestUseCaseGetRoles_GetRoles(t *testing.T) {
	t.Run("When attempting to retrieve roles, the operation is successful.", func(t *testing.T) {
		rolesRepository := &mockRoles.RoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		rolesRepository.
			On("GetRoles", mock.Anything, mock.Anything).
			Return([]rolesDomain.Role{}, nil)
		rolesRepository.
			On("GetTotalRoles", mock.Anything, mock.Anything).
			Return(&total, nil)
		roleUCase := NewRolesUseCase(rolesRepository, validationRepository, authRepository, 60)
		pagination := paramsDomain.NewPaginationParams(nil)
		role, _, err := roleUCase.GetRoles(context.Background(), pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, role, []rolesDomain.Role{})
	})

	t.Run("Encountered an error during role retrieval.", func(t *testing.T) {
		rolesRepository := &mockRoles.RoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		rolesRepository.
			On("GetRoles", mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		rolesRepository.
			On("GetTotalRoles", mock.Anything, mock.Anything).
			Return(&total, nil)
		rolesUCase := NewRolesUseCase(rolesRepository, validationRepository, authRepository, 60)
		pagination := paramsDomain.NewPaginationParams(nil)
		ok, _, err := rolesUCase.GetRoles(context.Background(), pagination)
		assert.Error(t, err)
		assert.EqualValues(t, ok, []rolesDomain.Role(nil))
	})
}

func TestUseCaseRoles_CreateRole(t *testing.T) {
	t.Run("When attempting to create a role, the operation is successful.", func(t *testing.T) {
		rolesRepository := &mockRoles.RoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		roleID := "73900000-7e93-11ee-89fd-0242a500000"
		rolesRepository.
			On("CreateRole", mock.Anything, mock.Anything, mock.Anything).
			Return(&roleID, nil)
		validationRepository.On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		rolesUCase := NewRolesUseCase(rolesRepository, validationRepository, authRepository, 60)
		_, err := rolesUCase.CreateRole(
			context.Background(),
			rolesDomain.CreateRoleBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When attempting to create a role, the operation is successful.", func(t *testing.T) {
		rolesRepository := &mockRoles.RoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		roleID := "73900000-7e93-11ee-89fd-0242a500000"
		rolesRepository.
			On("CreateRole", mock.Anything, mock.Anything, mock.Anything).
			Return(&roleID, errors.New("random error"))
		validationRepository.On("ValidateExistence", mock.Anything, mock.Anything).
			Return(true, nil)
		rolesUCase := NewRolesUseCase(rolesRepository, validationRepository, authRepository, 60)
		_, err := rolesUCase.CreateRole(
			context.Background(),
			rolesDomain.CreateRoleBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, rolesDomain.ErrRoleRoleNameAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateRole")
	})

	t.Run("Encountered an error during role creation.", func(t *testing.T) {
		rolesRepository := &mockRoles.RoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		errCreate := errDomain.NewErr().SetFunction("CreateRole").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		rolesRepository.
			On("CreateRole", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		rolesUCase := NewRolesUseCase(rolesRepository, validationRepository, authRepository, 60)
		_, err := rolesUCase.CreateRole(
			context.Background(),
			rolesDomain.CreateRoleBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateRole")
	})
}

func TestUseCaseRoles_UpdateRole(t *testing.T) {
	t.Run("When attempting to update a role, the operation is successful.", func(t *testing.T) {
		rolesRepository := &mockRoles.RoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		rolesRepository.
			On("UpdateRole", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		rolesUCase := NewRolesUseCase(rolesRepository, validationRepository, authRepository, 60)
		err := rolesUCase.UpdateRole(
			context.Background(),
			"fcdbfacf-8305-11ee-89fd-0242555555",
			rolesDomain.CreateRoleBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("Encountered an error during role update.", func(t *testing.T) {
		rolesRepository := &mockRoles.RoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		rolesRepository.
			On("UpdateRole", mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		rolesUCase := NewRolesUseCase(rolesRepository, validationRepository, authRepository, 60)
		err := rolesUCase.UpdateRole(
			context.Background(),
			"fcdbfacf-8305-11ee-89fd-0242555555",
			rolesDomain.CreateRoleBody{},
		)
		assert.Error(t, err)
	})
}

func TestUseCaseRoles_DeleteRole(t *testing.T) {
	t.Run("When attempting to delete a role by id, the operation is successful.", func(t *testing.T) {
		rolesRepository := &mockRoles.RoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		rolesRepository.
			On("DeleteRole", mock.Anything, mock.Anything).
			Return(true, nil)
		rolesUCase := NewRolesUseCase(rolesRepository, validationRepository, authRepository, 60)
		res, err := rolesUCase.DeleteRole(context.Background(),
			"fcdbfacf-8305-11ee-89fd-0242555555")
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("Encountered an error during role deletion by id.", func(t *testing.T) {
		rolesRepository := &mockRoles.RoleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		rolesError := errors.New("random error")
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		rolesRepository.
			On("DeleteRole", mock.Anything, mock.Anything).
			Return(false, rolesError)
		rolesUCase := NewRolesUseCase(rolesRepository, validationRepository, authRepository, 60)
		res, err := rolesUCase.DeleteRole(context.Background(),
			"fcdbfacf-8305-11ee-89fd-0242555555")
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
