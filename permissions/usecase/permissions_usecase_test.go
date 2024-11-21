/*
 * File: permissions_usecase_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the permissions use case.
 *
 * Last Modified: 2023-11-15
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

	permissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
	mockPermissions "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain/mocks"
)

func TestUseCaseGetPermissions_GetPermissions(t *testing.T) {
	t.Run("When attempting to retrieve permissions, the operation is successful.", func(t *testing.T) {
		permissionsRepository := &mockPermissions.PermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		searchParams := permissionsDomain.GetPermissionsParams{}
		total := 10
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionsRepository.
			On("GetPermissions", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return([]permissionsDomain.Permission{}, nil)
		permissionsRepository.
			On("GetTotalPermissions", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		permissionUCase := NewPermissionsUseCase(
			permissionsRepository,
			validationRepository,
			authRepository,
			60)
		pagination := paramsDomain.NewPaginationParams(nil)
		permission, _, err := permissionUCase.GetPermissions(context.Background(), moduleId, searchParams, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, permission, []permissionsDomain.Permission{})
	})

	t.Run("Encountered an error during permission retrieval.", func(t *testing.T) {
		permissionsRepository := &mockPermissions.PermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		searchParams := permissionsDomain.GetPermissionsParams{}
		total := 10
		permissionsRepository.
			On("GetPermissions", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		permissionsRepository.
			On("GetTotalPermissions", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		permissionsUCase := NewPermissionsUseCase(permissionsRepository, validationRepository, authRepository, 60)
		pagination := paramsDomain.NewPaginationParams(nil)
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		ok, _, err := permissionsUCase.GetPermissions(context.Background(), moduleId, searchParams, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, ok, []permissionsDomain.Permission(nil))
	})
}

func TestUseCasePermissions_CreatePermission(t *testing.T) {
	t.Run("When attempting to create a permission, the operation is successful.", func(t *testing.T) {
		permissionsRepository := &mockPermissions.PermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionID := "73900000-7e93-11ee-89fd-0242a500000"
		permissionsRepository.
			On("CreatePermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(&permissionID, nil)
		validationRepository.
			On("ValidateExistence",
				mock.Anything,
				mock.Anything).
			Return(false, nil)
		permissionsUCase := NewPermissionsUseCase(
			permissionsRepository,
			validationRepository,
			authRepository,
			60)
		_, err := permissionsUCase.CreatePermission(
			context.Background(),
			moduleId,
			permissionsDomain.CreatePermissionBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When there is an error creating a permission due to previous existence.", func(t *testing.T) {
		permissionsRepository := &mockPermissions.PermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		createPermissionBody := permissionsDomain.CreatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
			ModuleId:    "cddbfacf-8305-11ee-89fd-024255555502",
		}
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionsRepository.
			On("CreatePermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, errors.New("random error"))
		validationRepository.
			On("ValidateExistence",
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		permissionsUCase := NewPermissionsUseCase(
			permissionsRepository,
			validationRepository,
			authRepository,
			60)
		_, err := permissionsUCase.CreatePermission(
			context.Background(),
			moduleId,
			createPermissionBody,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, permissionsDomain.ErrPermissionNameAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreatePermission")
	})

	t.Run("Encountered an error during permission creation.", func(t *testing.T) {
		permissionsRepository := &mockPermissions.PermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		errCreate := errDomain.NewErr().SetFunction("CreatePermission").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		createPermissionBody := permissionsDomain.CreatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
			ModuleId:    "cddbfacf-8305-11ee-89fd-024255555502",
		}
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionsRepository.
			On("CreatePermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, errCreate)
		validationRepository.
			On("ValidateExistence",
				mock.Anything,
				mock.Anything).
			Return(false, nil)
		permissionsUCase := NewPermissionsUseCase(permissionsRepository, validationRepository, authRepository, 60)
		_, err := permissionsUCase.CreatePermission(
			context.Background(),
			moduleId,
			createPermissionBody,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreatePermission")
	})
}

func TestUseCasePermissions_UpdatePermission(t *testing.T) {
	t.Run("When attempting to update a permission, the operation is successful.", func(t *testing.T) {
		permissionsRepository := &mockPermissions.PermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		updatePermissionBody := permissionsDomain.UpdatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
		}
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionId := "fcdbfacf-8305-11ee-89fd-0242555555"
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		permissionsRepository.
			On("UpdatePermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		permissionsUCase := NewPermissionsUseCase(
			permissionsRepository,
			validationRepository,
			authRepository,
			60)
		err := permissionsUCase.UpdatePermission(
			context.Background(),
			moduleId,
			permissionId,
			updatePermissionBody,
		)
		assert.NoError(t, err)
	})

	t.Run("Encountered an error during permission update.", func(t *testing.T) {
		permissionsRepository := &mockPermissions.PermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionId := "fcdbfacf-8305-11ee-89fd-0242555555"
		updatePermissionBody := permissionsDomain.UpdatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
		}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		permissionsRepository.
			On("UpdatePermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		permissionsUCase := NewPermissionsUseCase(
			permissionsRepository,
			validationRepository,
			authRepository,
			60)
		err := permissionsUCase.UpdatePermission(
			context.Background(),
			moduleId,
			permissionId,
			updatePermissionBody,
		)
		assert.Error(t, err)
	})
}

func TestUseCasePermissions_DeletePermission(t *testing.T) {
	t.Run("When attempting to delete a permission by id, the operation is successful.", func(t *testing.T) {
		permissionsRepository := &mockPermissions.PermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		permissionsRepository.
			On("DeletePermission", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		permissionsUCase := NewPermissionsUseCase(
			permissionsRepository,
			validationRepository,
			authRepository,
			60)
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		res, err := permissionsUCase.DeletePermission(context.Background(),
			moduleId, "fcdbfacf-8305-11ee-89fd-0242555555")
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("Encountered an error during permission deletion by id.", func(t *testing.T) {
		permissionsRepository := &mockPermissions.PermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		permissionsError := errors.New("random error")
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		permissionsRepository.
			On("DeletePermission", mock.Anything, mock.Anything, mock.Anything).
			Return(false, permissionsError)
		permissionsUCase := NewPermissionsUseCase(permissionsRepository, validationRepository, authRepository, 60)
		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		res, err := permissionsUCase.DeletePermission(context.Background(),
			moduleId, "fcdbfacf-8305-11ee-89fd-0242555555")
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
