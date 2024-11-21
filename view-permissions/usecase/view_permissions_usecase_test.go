/*
 * File: view_permissions_usecase_test.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the test for the viewPermissions use case.
 *
 * Last Modified: 2024-02-26
 */

package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	mockValidation "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain/mocks"

	viewPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
	mockViewPermissions "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain/mocks"
)

func TestViewPermissionsUseCase_GetViewPermissions(t *testing.T) {
	t.Run("When get view permission  successfully", func(t *testing.T) {
		ViewPermissionsRepository := &mockViewPermissions.ViewPermissionsRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		now := time.Now().UTC()
		viewId := "73900000-7e93-11ee-89fd-0242a500000"
		viewPermissions := []viewPermissionsDomain.ViewPermission{
			{
				Id:        "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
				CreatedBy: "91fb86bd-da46-414b-97a1-fcdaa8cd35d1",
				CreatedAt: &now,
				View: viewPermissionsDomain.View{
					Id:          "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
					Name:        "Bancos",
					Description: "A. de Bancos",
					CreatedAt:   &now,
				},
				Permission: viewPermissionsDomain.Permission{
					Id:          "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
					Code:        "READ_USERS",
					Name:        "activo fijo",
					Description: "activo fijo",
					CreatedAt:   &now,
					Module: viewPermissionsDomain.Module{
						Id:          "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
						Name:        "listado de activos fijos",
						Description: "listado de activos fijos",
						Icon:        "core",
						Position:    1,
						CreatedAt:   &now,
					},
				},
			},
			{
				Id:        "2e8bfdbb-1a58-4b45-9a2c-8ac54a5db723",
				CreatedBy: "c3f92a0d-ef58-4e15-a71b-6f0a8b9d147d",
				CreatedAt: &now,
				View: viewPermissionsDomain.View{
					Id:          "2e8bfdbb-1a58-4b45-9a2c-8ac54a5db723",
					Name:        "Inventory",
					Description: "Inventory Management",
					CreatedAt:   &now,
				},
				Permission: viewPermissionsDomain.Permission{
					Id:          "2e8bfdbb-1a58-4b45-9a2c-8ac54a5db723",
					Code:        "UPDATE_STOCK",
					Name:        "Update Stock",
					Description: "Modify stock levels",
					CreatedAt:   &now,
					Module: viewPermissionsDomain.Module{
						Id:          "2e8bfdbb-1a58-4b45-9a2c-8ac54a5db723",
						Name:        "Stock Management",
						Description: "Manage inventory stock",
						Icon:        "inventory",
						Position:    1,
						CreatedAt:   &now,
					},
				},
			},
		}

		ViewPermissionsRepository.
			On("GetViewPermissions", mock.Anything, mock.Anything).
			Return(viewPermissions, nil)

		ViewPermissionsUCase := NewViewPermissionsUseCase(
			ViewPermissionsRepository,
			validationRepository,
			authRepository,
			60,
		)

		res, err := ViewPermissionsUCase.GetViewPermissions(context.Background(), viewId)

		assert.NoError(t, err)
		assert.EqualValues(t, res, viewPermissions)
	})

	t.Run("When get view permissions error", func(t *testing.T) {
		ViewPermissionsRepository := &mockViewPermissions.ViewPermissionsRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		viewId := "73900000-7e93-11ee-89fd-0242a500000"
		ViewPermissionsRepository.
			On("GetViewPermissions", mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		ViewPermissionsUCase := NewViewPermissionsUseCase(
			ViewPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		res, err := ViewPermissionsUCase.GetViewPermissions(context.Background(), viewId)
		assert.Error(t, err)
		assert.EqualValues(t, res, []viewPermissionsDomain.ViewPermission(nil))
	})
}

func TestViewPermissionsUseCase_CreateViewPermission(t *testing.T) {
	t.Run("When create view permission successfully", func(t *testing.T) {
		ViewPermissionsRepository := &mockViewPermissions.ViewPermissionsRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		viewId := "73900000-7e93-11ee-89fd-0242a500000"
		ViewPermissionsId := "73900000-7e93-11ee-89fd-0242a500000"
		userId := "91fb86bd-da46-414b-97a1-fcdaa8cd35d1"
		body := viewPermissionsDomain.CreateViewPermissionBody{
			PermissionId: "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
		}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		ViewPermissionsRepository.
			On("CreateViewPermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&ViewPermissionsId, nil)

		ViewPermissionsUCase := NewViewPermissionsUseCase(
			ViewPermissionsRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := ViewPermissionsUCase.CreateViewPermission(
			context.Background(),
			viewId,
			userId,
			body,
		)
		assert.NoError(t, err)
	})

	t.Run("When view permission not fount o", func(t *testing.T) {
		ViewPermissionsRepository := &mockViewPermissions.ViewPermissionsRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		viewId := "73900000-7e93-11ee-89fd-0242a500000"
		userId := "91fb86bd-da46-414b-97a1-fcdaa8cd35d1"
		body := viewPermissionsDomain.CreateViewPermissionBody{
			PermissionId: "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
		}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		ViewPermissionsRepository.
			On("CreateViewPermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		ViewPermissionsUCase := NewViewPermissionsUseCase(
			ViewPermissionsRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := ViewPermissionsUCase.CreateViewPermission(
			context.Background(),
			viewId,
			userId,
			body,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, viewPermissionsDomain.ErrViewNotFoundCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateViewPermission")
	})
}

func TestViewPermissionsUseCase_UpdateViewPermission(t *testing.T) {
	t.Run("When update view permission successfully", func(t *testing.T) {
		ViewPermissionsRepository := &mockViewPermissions.ViewPermissionsRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		viewId := "73900000-7e93-11ee-89fd-0242a500000"
		viewPermissionId := "73900000-7e93-11ee-89fd-0242a500000"
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		ViewPermissionsRepository.
			On("UpdateViewPermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		ViewPermissionsUCase := NewViewPermissionsUseCase(
			ViewPermissionsRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := ViewPermissionsUCase.UpdateViewPermission(
			context.Background(),
			viewId,
			viewPermissionId,
			viewPermissionsDomain.UpdateViewPermissionBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When update view permission error", func(t *testing.T) {
		ViewPermissionsRepository := &mockViewPermissions.ViewPermissionsRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		viewId := "73900000-7e93-11ee-89fd-0242a500000"
		viewPermissionId := "73900000-7e93-11ee-89fd-0242a500000"
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		ViewPermissionsRepository.
			On("UpdateViewPermission", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		ViewPermissionsUCase := NewViewPermissionsUseCase(
			ViewPermissionsRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := ViewPermissionsUCase.UpdateViewPermission(
			context.Background(),
			viewId,
			viewPermissionId,
			viewPermissionsDomain.UpdateViewPermissionBody{},
		)
		assert.Error(t, err)
	})
}

func TestViewPermissionsUseCase_DeleteViewPermission(t *testing.T) {
	t.Run("When delete view permission by id successfully", func(t *testing.T) {
		ViewPermissionsRepository := &mockViewPermissions.ViewPermissionsRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		viewId := "73900000-7e93-11ee-89fd-0242a500000"
		viewPermissionId := "73900000-7e93-11ee-89fd-0242a500000"
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		ViewPermissionsRepository.
			On("DeleteViewPermission", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		ViewPermissionsUCase := NewViewPermissionsUseCase(
			ViewPermissionsRepository,
			validationRepository,
			authRepository,
			60,
		)
		res, err := ViewPermissionsUCase.DeleteViewPermission(
			context.Background(),
			viewId,
			viewPermissionId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete classifications view permissions by id error", func(t *testing.T) {
		ViewPermissionsRepository := &mockViewPermissions.ViewPermissionsRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		viewId := "73900000-7e93-11ee-89fd-0242a500000"
		viewPermissionId := "73900000-7e93-11ee-89fd-0242a500000"
		ViewPermissionsError := errors.New("random error")
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		ViewPermissionsRepository.
			On("DeleteViewPermission", mock.Anything, mock.Anything, mock.Anything).
			Return(false, ViewPermissionsError)
		ViewPermissionsUCase := NewViewPermissionsUseCase(
			ViewPermissionsRepository,
			validationRepository,
			authRepository,
			60,
		)
		res, err := ViewPermissionsUCase.DeleteViewPermission(
			context.Background(),
			viewId,
			viewPermissionId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
