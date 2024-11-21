/*
 * File: roles_mysql_repository_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the roles repository.
 *
 * Last Modified: 2023-11-10
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

	rolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/roles/domain"
)

func TestRepositoryRoles_GetRoles(t *testing.T) {
	t.Run("When get roles is called, it should return a list of roles", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		sizePage := 100
		offset := 0
		now := time.Now().UTC()
		mockRegister := []rolesDomain.Role{
			{
				Id:          "fcdbfacf-8305-11ee-89fd-0242555555",
				Name:        "Gerencia",
				Description: "Gerencia del conglomerado",
				Enable:      true,
				CreatedAt:   &now,
			}, {
				Id:          "fcdbfacf-8305-11ee-89fd-0242555500",
				Name:        "Gerencia",
				Description: "Gerencia del conglomerado",
				Enable:      false,
				CreatedAt:   &now,
			},
		}
		rows := sqlmock.NewRows([]string{"id", "name", "description", "enable", "created_at"}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].Name,
				mockRegister[0].Description,
				mockRegister[0].Enable,
				mockRegister[0].CreatedAt,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].Name,
				mockRegister[1].Description,
				mockRegister[1].Enable,
				mockRegister[1].CreatedAt,
			)

		clock := &mockClock.Clock{}
		mock.ExpectQuery(QueryGetRoles).WithArgs(sizePage, offset).WillReturnRows(rows)
		r := NewRolesRepository(clock, 60)
		var res []rolesDomain.Role
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetRoles(ctx, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get roles is called and returns an error, it should handle the error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		sizePage := 100
		offset := 0
		expectedError := errors.New("random error")
		clock := &mockClock.Clock{}
		mock.ExpectQuery(QueryGetRoles).WithArgs(sizePage, offset).WillReturnError(expectedError)
		r := NewRolesRepository(clock, 60)
		var res []rolesDomain.Role
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetRoles(ctx, pagination)
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
		assert.Equal(t, smartErr.Function, "GetRoles")
	})
}

func TestRepositoryRoles_GetTotalRoles(t *testing.T) {
	t.Run("When get total roles is called then it should return a total success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		total := 10
		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)

		mock.
			ExpectQuery(QueryGetTotalRoles).
			WithArgs().
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewRolesRepository(clock, 60)

		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalRoles(ctx, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total roles is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		mock.ExpectQuery(QueryGetTotalRoles).
			WithArgs("739bbbc9-7e93-11ee-89fd-0242ac113421", "739bbbc9-7e93-11ee-89fd-0242ac113421").
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewRolesRepository(clock, 60)

		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalRoles(ctx, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalRoles")
	})
}

func TestRepositoryRoles_CreateRole(t *testing.T) {
	t.Run("When creating a role is successful, it should return without errors.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		roleId := uuid.New().String()
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createRoleBody := rolesDomain.CreateRoleBody{
			Name:        "Gerencia",
			Description: "Gerencia del conglomerado",
			Enable:      true,
		}
		mock.ExpectExec(QueryCreateRole).
			WithArgs(
				roleId,
				createRoleBody.Name,
				createRoleBody.Description,
				createRoleBody.Enable,
				createdAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		r := NewRolesRepository(clock, 60)
		_, err = r.CreateRole(ctx, roleId, createRoleBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When creating a role returns an error, it should handle the error.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		roleId := uuid.New().String()
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createRoleBody := rolesDomain.CreateRoleBody{
			Name:        "Gerencia",
			Description: "Gerencia del conglomerado",
			Enable:      true,
		}
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreateRole).
			WithArgs(
				roleId,
				createRoleBody.Name,
				createRoleBody.Description,
				createRoleBody.Enable,
				createdAt,
			).WillReturnError(expectedError)
		r := NewRolesRepository(clock, 60)
		_, err = r.CreateRole(ctx, roleId, createRoleBody)
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
		assert.Equal(t, smartErr.Function, "CreateRole")
	})
}

func TestRepositoryRoles_UpdateRole(t *testing.T) {
	t.Run("When updating a role is successful, it should return without errors", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		roleId := uuid.New().String()
		clock := &mockClock.Clock{}
		updateRoleBody := rolesDomain.CreateRoleBody{
			Name:        "Gerencia",
			Description: "Gerencia del conglomerado",
			Enable:      true,
		}
		mock.ExpectExec(QueryUpdateRole).
			WithArgs(
				updateRoleBody.Name,
				updateRoleBody.Description,
				updateRoleBody.Enable,
				roleId).
			WillReturnResult(
				sqlmock.NewResult(
					1,
					1,
				),
			)
		r := NewRolesRepository(clock, 60)
		err = r.UpdateRole(ctx, roleId, updateRoleBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When updating a role returns an error, it should handle the error.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		roleId := uuid.New().String()
		clock := &mockClock.Clock{}
		updateRoleBody := rolesDomain.CreateRoleBody{
			Name:        "Gerencia",
			Description: "Gerencia del conglomerado",
			Enable:      true,
		}
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdateRole).
			WithArgs(
				updateRoleBody.Name,
				updateRoleBody.Description,
				updateRoleBody.Enable,
				roleId).
			WillReturnError(expectedError)
		r := NewRolesRepository(clock, 60)
		err = r.UpdateRole(
			ctx,
			roleId,
			updateRoleBody)
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
		assert.Equal(t, smartErr.Function, "UpdateRole")
	})
}

func TestRepositoryRoles_DeleteRole(t *testing.T) {
	t.Run("When deleting a role is successful, it should return without errors.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteRole).
			WithArgs(deletedAt, roleId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewRolesRepository(clock, 60)
		var res bool
		res, err = r.DeleteRole(ctx, roleId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When deleting a role returns an error, it should handle the error.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		mock.ExpectExec(QueryDeleteRole).
			WithArgs(deletedAt, roleId).
			WillReturnError(errors.New("anything"))
		r := NewRolesRepository(clock, 60)
		var res bool
		res, err = r.DeleteRole(ctx, roleId)
		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteRole")
	})
}
