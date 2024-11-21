/*
 * File: permissions_mysql_repository_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the permissions repository.
 *
 * Last Modified: 2023-11-15
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
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	permissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/permissions/domain"
)

func TestRepositoryPermissions_GetPermissions(t *testing.T) {
	t.Run("When get permissions is called, it should return a list of permissions", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		moduleByPermission := permissionsDomain.ModuleByPermission{
			Id:          "739bbbc9-7e93-11ee-89fd-0242ac113421",
			Name:        "Logistic",
			Description: "Modulo de logística",
			Code:        "logistic",
		}
		code := "user"
		params := permissionsDomain.GetPermissionsParams{
			Params: paramsDomain.Params{},
			Code:   &code,
			Name:   nil,
		}
		mockRegister := []permissionsDomain.Permission{
			{
				Id:          "fcdbfacf-8305-11ee-89fd-0242555555",
				Code:        "REQUIREMENTS_READ",
				Name:        "Listar requerimientos",
				Description: "Permiso para listar requerimientos",
				CreatedAt:   &now,
				Module:      moduleByPermission,
			}, {
				Id:          "fcdbfacf-8305-11ee-89fd-0242555000",
				Code:        "REQUIREMENTS_READ",
				Name:        "Listar requerimientos",
				Description: "Permiso para listar requerimientos",
				CreatedAt:   &now,
				Module:      moduleByPermission,
			},
		}
		rows := sqlmock.NewRows([]string{"permission_id", "permission_code", "permission_name",
			"permission_description", "permission_created_at", "module_id",
			"module_name", "module_description", "module_code"}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].Code,
				mockRegister[0].Name,
				mockRegister[0].Description,
				mockRegister[0].CreatedAt,
				mockRegister[0].Module.Id,
				mockRegister[0].Module.Name,
				mockRegister[0].Module.Description,
				mockRegister[0].Module.Code,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].Code,
				mockRegister[1].Name,
				mockRegister[1].Description,
				mockRegister[1].CreatedAt,
				mockRegister[1].Module.Id,
				mockRegister[1].Module.Name,
				mockRegister[1].Module.Description,
				mockRegister[1].Module.Code,
			)
		sizePage := 100
		offset := 0
		moduleId := uuid.New().String()
		clock := &mockClock.Clock{}
		mock.
			ExpectQuery(QueryGetPermissions).
			WithArgs(
				moduleId,
				params.Code,
				params.Code,
				params.Name,
				params.Name,
				sizePage,
				offset).
			WillReturnRows(rows)
		r := NewPermissionsRepository(clock, 60)
		var res []permissionsDomain.Permission
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetPermissions(ctx, moduleId, params, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get permissions is called and returns an error, it should handle the error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		moduleId := uuid.New().String()
		clock := &mockClock.Clock{}
		var res []permissionsDomain.Permission
		pagination := paramsDomain.NewPaginationParams(nil)
		code := "user"
		params := permissionsDomain.GetPermissionsParams{
			Params: paramsDomain.Params{},
			Code:   &code,
			Name:   nil,
		}

		mock.ExpectQuery(QueryGetPermissions).WillReturnError(expectedError)
		r := NewPermissionsRepository(clock, 60)

		res, err = r.GetPermissions(ctx, moduleId, params, pagination)
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
		assert.Equal(t, smartErr.Function, "GetPermissions")
	})
}

func TestRepositoryPermissions_GetTotalPermissions(t *testing.T) {
	t.Run("When get total permissions return success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		total := 10
		moduleId := uuid.New().String()
		code := "user"
		params := permissionsDomain.GetPermissionsParams{
			Params: paramsDomain.Params{},
			Code:   &code,
			Name:   nil,
		}
		var totalExpected *int
		clock := &mockClock.Clock{}
		rows := sqlmock.NewRows([]string{"total"}).AddRow(total)

		mock.
			ExpectQuery(QueryGetTotalPermissions).
			WithArgs(
				moduleId,
				params.Code,
				params.Code,
				params.Name,
				params.Name,
			).
			WillReturnRows(rows)
		r := NewPermissionsRepository(clock, 60)

		totalExpected, err = r.GetTotalPermissions(ctx, moduleId, params)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total permissions return error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		moduleId := uuid.New().String()
		clock := &mockClock.Clock{}
		var totalExpected *int
		code := "user"
		params := permissionsDomain.GetPermissionsParams{
			Params: paramsDomain.Params{},
			Code:   &code,
			Name:   nil,
		}

		mock.ExpectQuery(QueryGetTotalPermissions).
			WithArgs(
				moduleId,
				params.Code,
				params.Code,
				params.Name,
				params.Name,
			).
			WillReturnError(expectedError)

		r := NewPermissionsRepository(clock, 60)

		totalExpected, err = r.GetTotalPermissions(ctx, moduleId, params)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalPermissions")

	})
}

func TestRepositoryPermissions_CreatePermission(t *testing.T) {
	t.Run("When creating a permission is successful, it should return without errors.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		moduleId := uuid.New().String()
		permissionId := uuid.New().String()
		createPermissionBody := permissionsDomain.CreatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
			ModuleId:    moduleId,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryCreatePermission).
			WithArgs(permissionId,
				createPermissionBody.Code,
				createPermissionBody.Name,
				createPermissionBody.Description,
				createPermissionBody.ModuleId,
				createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		r := NewPermissionsRepository(clock, 60)
		_, err = r.CreatePermission(ctx, moduleId, permissionId, createPermissionBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When creating a permission returns an error, it should handle the error.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		moduleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		permissionId := "739bbbc9-7e93-11ee-89fd-0442ac219255"
		createPermissionBody := permissionsDomain.CreatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
			ModuleId:    "cddbfacf-8305-11ee-89fd-024255555502",
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(time.Now())
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreatePermission).
			WithArgs(permissionId,
				createPermissionBody.Code,
				createPermissionBody.Name,
				createPermissionBody.Description,
				createPermissionBody.ModuleId,
				createdAt).
			WillReturnError(expectedError)
		r := NewPermissionsRepository(clock, 60)
		_, err = r.CreatePermission(ctx, moduleId, permissionId, createPermissionBody)
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
		assert.Equal(t, smartErr.Function, "CreatePermission")
	})
}

func TestRepositoryPermissions_UpdatePermission(t *testing.T) {
	t.Run("When updating a permission is successful, it should return without errors", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		moduleId := uuid.New().String()
		permissionId := uuid.New().String()
		clock := &mockClock.Clock{}
		updatePermissionBody := permissionsDomain.UpdatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
		}
		mock.ExpectExec(QueryUpdatePermission).
			WithArgs(
				updatePermissionBody.Code,
				updatePermissionBody.Name,
				updatePermissionBody.Description,
				moduleId,
				permissionId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		r := NewPermissionsRepository(clock, 60)
		err = r.UpdatePermission(ctx, moduleId, permissionId, updatePermissionBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When updating a permission returns an error, it should handle the error.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		clock := &mockClock.Clock{}
		moduleId := uuid.New().String()
		permissionId := uuid.New().String()
		updatePermissionBody := permissionsDomain.UpdatePermissionBody{
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
		}

		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdatePermission).
			WithArgs(updatePermissionBody.Code,
				updatePermissionBody.Name,
				updatePermissionBody.Description,
				moduleId,
				permissionId).
			WillReturnError(expectedError)
		r := NewPermissionsRepository(clock, 60)
		err = r.UpdatePermission(ctx, moduleId, permissionId, updatePermissionBody)
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
		assert.Equal(t, smartErr.Function, "UpdatePermission")
	})
}

func TestRepositoryPermissions_DeletePermission(t *testing.T) {
	t.Run("When deleting a permission is successful, it should return without errors.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		permissionId := uuid.New().String()
		moduleId := uuid.New().String()
		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeletePermission).
			WithArgs(deletedAt, permissionId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewPermissionsRepository(clock, 60)
		var res bool

		res, err = r.DeletePermission(ctx, moduleId, permissionId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When deleting a permission returns an error, it should handle the error.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		permissionId := uuid.New().String()
		moduleId := uuid.New().String()
		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeletePermission).
			WithArgs(deletedAt, permissionId).
			WillReturnError(errors.New("anything"))
		r := NewPermissionsRepository(clock, 60)
		var res bool

		res, err = r.DeletePermission(ctx, moduleId, permissionId)
		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeletePermission")
	})
}
