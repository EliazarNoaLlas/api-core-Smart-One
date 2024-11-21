/*
 * File: view_permissions_repository_test.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the tests for the viewPermissions repository.
 *
 * Last Modified: 2024-02-26
 */

package mysql

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	mockClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock/mocks"
	db2 "gitlab.smartcitiesperu.com/smartone/api-shared/db"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	ViewPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/domain"
)

func TestRepositoryViewPermissions_GetViewPermissions(t *testing.T) {
	t.Run("When get classification asset service called then it should return a list of specialties", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		viewId := uuid.New().String()
		clock := &mockClock.Clock{}
		mockRegister := []ViewPermissionsDomain.ViewPermission{
			{
				Id:        "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
				CreatedBy: "91fb86bd-da46-414b-97a1-fcdaa8cd35d1",
				CreatedAt: &now,
				View: ViewPermissionsDomain.View{
					Id:          "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
					Name:        "Bancos",
					Description: "A. de Bancos",
					CreatedAt:   &now,
				},
				Permission: ViewPermissionsDomain.Permission{
					Id:          "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
					Code:        "READ_USERS",
					Name:        "activo fijo",
					Description: "activo fijo",
					CreatedAt:   &now,
					Module: ViewPermissionsDomain.Module{
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
				View: ViewPermissionsDomain.View{
					Id:          "2e8bfdbb-1a58-4b45-9a2c-8ac54a5db723",
					Name:        "Inventory",
					Description: "Inventory Management",
					CreatedAt:   &now,
				},
				Permission: ViewPermissionsDomain.Permission{
					Id:          "2e8bfdbb-1a58-4b45-9a2c-8ac54a5db723",
					Code:        "UPDATE_STOCK",
					Name:        "Update Stock",
					Description: "Modify stock levels",
					CreatedAt:   &now,
					Module: ViewPermissionsDomain.Module{
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

		rows := sqlmock.NewRows([]string{
			"view_permission_id",
			"view_permission_created_by",
			"view_permission_created_at",
			"view_id",
			"view_name",
			"view_description",
			"view_created_at",
			"permission_id",
			"permission_code",
			"permission_name",
			"permission_description",
			"permission_created_at",
			"module_id",
			"module_name",
			"module_description",
			"module_code",
			"module_icon",
			"module_position",
			"module_created_at",
		}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].CreatedBy,
				mockRegister[0].CreatedAt,
				mockRegister[0].View.Id,
				mockRegister[0].View.Name,
				mockRegister[0].View.Description,
				mockRegister[0].View.CreatedAt,
				mockRegister[0].Permission.Id,
				mockRegister[0].Permission.Code,
				mockRegister[0].Permission.Name,
				mockRegister[0].Permission.Description,
				mockRegister[0].Permission.CreatedAt,
				mockRegister[0].Permission.Module.Id,
				mockRegister[0].Permission.Module.Name,
				mockRegister[0].Permission.Module.Description,
				mockRegister[0].Permission.Module.Code,
				mockRegister[0].Permission.Module.Icon,
				mockRegister[0].Permission.Module.Position,
				mockRegister[0].Permission.Module.CreatedAt,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].CreatedBy,
				mockRegister[1].CreatedAt,
				mockRegister[1].View.Id,
				mockRegister[1].View.Name,
				mockRegister[1].View.Description,
				mockRegister[1].View.CreatedAt,
				mockRegister[1].Permission.Id,
				mockRegister[1].Permission.Code,
				mockRegister[1].Permission.Name,
				mockRegister[1].Permission.Description,
				mockRegister[1].Permission.CreatedAt,
				mockRegister[1].Permission.Module.Id,
				mockRegister[1].Permission.Module.Name,
				mockRegister[1].Permission.Module.Description,
				mockRegister[1].Permission.Module.Code,
				mockRegister[1].Permission.Module.Icon,
				mockRegister[1].Permission.Module.Position,
				mockRegister[1].Permission.Module.CreatedAt,
			)

		mock.ExpectQuery(QueryGetViewPermissions).
			WithArgs(viewId).
			WillReturnRows(rows)
		r := NewViewPermissionsRepository(clock, 60)
		var res []ViewPermissionsDomain.ViewPermission
		res, err = r.GetViewPermissions(ctx, viewId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get classification asset service is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		now := time.Now().UTC()
		viewId := uuid.New().String()

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		var res []ViewPermissionsDomain.ViewPermission

		mock.ExpectQuery(QueryGetViewPermissions).
			WithArgs().
			WillReturnError(expectedError)

		r := NewViewPermissionsRepository(clock, 60)
		res, err = r.GetViewPermissions(ctx, viewId)
		if res != nil {
			t.Errorf("this is the error getting the registers: %v\n", res)
			return
		}
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetViewPermissions")
	})
}

func TestRepositoryViewPermissions_CreateViewPermission(t *testing.T) {
	t.Run("When create an classification asset service success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		ViewPermissionsId := uuid.New().String()
		viewId := uuid.New().String()

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		createdAt := now.Format("2006-01-02 15:04:05")
		userId := uuid.New().String()
		createViewPermissionsBody := ViewPermissionsDomain.CreateViewPermissionBody{
			PermissionId: "18f7f9c2-b00a-42e4-a469-ea4c01c180dc",
		}

		mock.ExpectExec(QueryCreateViewPermission).
			WithArgs(
				ViewPermissionsId,
				viewId,
				createViewPermissionsBody.PermissionId,
				userId,
				createdAt).
			WillReturnResult(
				sqlmock.NewResult(1, 1))

		r := NewViewPermissionsRepository(clock, 60)
		_, err = r.CreateViewPermission(ctx, viewId, userId, ViewPermissionsId, createViewPermissionsBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When create an classification asset service return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		clock := &mockClock.Clock{}
		now := time.Now().UTC()
		viewId := uuid.New().String()
		ViewPermissionsId := uuid.New().String()
		clock.On("Now").Return(now)

		createdAt := now.Format("2006-01-02 15:04:05")
		userId := uuid.New().String()
		createViewPermissionsBody := ViewPermissionsDomain.CreateViewPermissionBody{
			PermissionId: "18f7f9c2-b00a-42e4-a469-ea4c01c180dc",
		}

		mock.ExpectQuery(QueryCreateViewPermission).
			WithArgs(
				ViewPermissionsId,
				viewId,
				createViewPermissionsBody.PermissionId,
				userId,
				createdAt).
			WillReturnError(expectedError)
		r := NewViewPermissionsRepository(clock, 60)
		_, err = r.CreateViewPermission(ctx, viewId, userId, ViewPermissionsId, createViewPermissionsBody)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "CreateViewPermission")
	})
}

func TestRepositoryViewPermissions_UpdateViewPermission(t *testing.T) {
	t.Run("When update classification asset service success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		viewId := uuid.New().String()
		ViewPermissionsId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		updateViewPermissionsBody := ViewPermissionsDomain.UpdateViewPermissionBody{
			PermissionId: "18f7f9c2-b00a-42e4-a469-ea4c01c180dc",
		}
		mock.ExpectExec(QueryUpdateViewPermission).
			WithArgs(
				updateViewPermissionsBody.PermissionId,
				ViewPermissionsId).
			WillReturnResult(
				sqlmock.NewResult(1, 1))
		r := NewViewPermissionsRepository(clock, 60)
		err = r.UpdateViewPermission(
			ctx,
			viewId,
			ViewPermissionsId,
			updateViewPermissionsBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When update classification asset service return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		now := time.Now().UTC()
		viewId := uuid.New().String()
		ViewPermissionsId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		updateViewPermissionsBody := ViewPermissionsDomain.UpdateViewPermissionBody{
			PermissionId: "18f7f9c2-b00a-42e4-a469-ea4c01c180dc",
		}
		mock.ExpectQuery(QueryUpdateViewPermission).
			WithArgs(
				updateViewPermissionsBody.PermissionId,
				ViewPermissionsId).
			WillReturnError(expectedError)

		r := NewViewPermissionsRepository(clock, 60)
		err = r.UpdateViewPermission(
			ctx,
			viewId,
			ViewPermissionsId,
			updateViewPermissionsBody)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "UpdateViewPermission")
	})
}

func TestRepositoryViewPermissions_DeleteViewPermission(t *testing.T) {
	t.Run("When delete classification asset service successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		viewId := uuid.New().String()
		ViewPermissionsId := uuid.New().String()
		deletedAt := now.Format("2006-01-02 15:04:05")
		mock.ExpectExec(QueryDeleteViewPermission).
			WithArgs(deletedAt, ViewPermissionsId, viewId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewViewPermissionsRepository(clock, 60)
		var res bool
		res, err = r.DeleteViewPermission(ctx, viewId, ViewPermissionsId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete classification asset service error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		now := time.Now().UTC()
		viewId := uuid.New().String()
		ViewPermissionsId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		deletedAt := now.Format("2006-01-02 15:04:05")

		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteViewPermission).
			WithArgs(deletedAt, ViewPermissionsId).
			WillReturnError(expectedError)
		r := NewViewPermissionsRepository(clock, 60)
		var res bool
		res, err = r.DeleteViewPermission(ctx, viewId, ViewPermissionsId)
		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteViewPermission")
	})
}
