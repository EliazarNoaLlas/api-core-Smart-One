/*
 * File: user_roles_mysql_repository_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to userRole repository.
 *
 * Last Modified: 2023-11-23
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

	userRolesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/domain"
)

func TestRepositoryUserRoles_GetUserRolesByUser(t *testing.T) {
	t.Run("When get roles is called then it should return a list of user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now()
		roles := userRolesDomain.Role{
			Id:          "476a3664-d0d0-4476-8f12-fb11ae57122a",
			Name:        "Gerencia",
			Description: "Gerencia del conglomerado2221",
			Enable:      false,
			CreatedAt:   &now,
		}
		mockUserRolesType := []userRolesDomain.UserRole{
			{
				Id:        "476a3664-d0d0-4476-8f12-fb11ae57122a",
				Enable:    false,
				CreatedAt: &now,
				Roles:     roles,
			},
			{
				Id:        "22597e1d-6463-4bf9-ba51-0f8a3967321f",
				Enable:    false,
				CreatedAt: &now,
				Roles:     roles,
			},
		}

		rows := sqlmock.NewRows([]string{
			"user_role_id",
			"user_role_enable",
			"user_role_created_at",
			"role_id",
			"role_name",
			"role_description",
			"role_enable",
			"role_created_at",
		}).
			AddRow(
				mockUserRolesType[0].Id,
				mockUserRolesType[0].Enable,
				mockUserRolesType[0].CreatedAt,
				mockUserRolesType[0].Roles.Id,
				mockUserRolesType[0].Roles.Name,
				mockUserRolesType[0].Roles.Description,
				mockUserRolesType[0].Roles.Enable,
				mockUserRolesType[0].Roles.CreatedAt,
			).
			AddRow(
				mockUserRolesType[1].Id,
				mockUserRolesType[1].Enable,
				mockUserRolesType[1].CreatedAt,
				mockUserRolesType[0].Roles.Id,
				mockUserRolesType[1].Roles.Name,
				mockUserRolesType[1].Roles.Description,
				mockUserRolesType[1].Roles.Enable,
				mockUserRolesType[1].Roles.CreatedAt,
			)

		sizePage := 100
		offset := 0
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110020"
		mock.
			ExpectQuery(QueryUserRolesbyUser).
			WithArgs(userId, sizePage, offset).WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUserRolesRepository(clock, 60)
		var res []userRolesDomain.UserRole
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetUserRolesByUser(ctx, userId, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get roles by user is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		sizePage := 100
		offset := 0
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110020"
		mock.ExpectQuery(QueryUserRolesbyUser).
			WithArgs(userId, sizePage, offset).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUserRolesRepository(clock, 60)
		var res []userRolesDomain.UserRole
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetUserRolesByUser(ctx, userId, pagination)
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
		assert.Equal(t, smartErr.Function, "GetUserRolesByUser")
	})
}

func TestRepositoryUserRoles_GetTotalUserRolesByUser(t *testing.T) {
	t.Run("When get total of roles by user is called then it should return a total", func(t *testing.T) {
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
			ExpectQuery(QueryGetTotalRolesbyUser).
			WithArgs().
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUserRolesRepository(clock, 60)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110020"
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalUserRolesByUser(ctx, userId, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total of roles by user is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		mock.ExpectQuery(QueryGetTotalRolesbyUser).
			WithArgs("739bbbc9-7e93-11ee-89fd-0242ac113421", "739bbbc9-7e93-11ee-89fd-0242ac113421").
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUserRolesRepository(clock, 60)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110020"
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalUserRolesByUser(ctx, userId, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalUserRolesByUser")
	})
}

func TestRepositoryUserRoles_CreateUserRole(t *testing.T) {
	var enable = true
	t.Run("When add a role to user success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userRoleId := uuid.New().String()
		userId := "739bbbc9-7e93-11ee-89fd-0442ac219255"
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createUserRoleBody := userRolesDomain.CreateUserRoleBody{
			RoleId: roleId,
			Enable: enable,
		}

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryCreateUserRole).
			WithArgs(userRoleId, userId, roleId, enable, createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewUserRolesRepository(clock, 60)

		_, err = r.CreateUserRole(ctx, userRoleId, userId, createUserRoleBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When add a policy to role return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userRoleId := uuid.New().String()
		userId := "739bbbc9-7e93-11ee-89fd-0442ac219255"
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createUserRoleBody := userRolesDomain.CreateUserRoleBody{
			RoleId: roleId,
			Enable: enable,
		}

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreateUserRole).
			WithArgs(userRoleId, userId, roleId, enable, createdAt).
			WillReturnError(expectedError)
		r := NewUserRolesRepository(clock, 60)
		_, err = r.CreateUserRole(ctx, userRoleId, roleId, createUserRoleBody)
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
		assert.Equal(t, smartErr.Function, "CreateUserRole")
	})
}

func TestRepositoryUserRoles_UpdateUserRole(t *testing.T) {
	var enable = true
	t.Run("When update a policy of role successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0442ac219255"
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		updateUserRoleBody := userRolesDomain.CreateUserRoleBody{
			RoleId: roleId,
			Enable: enable,
		}
		clock := &mockClock.Clock{}
		mock.ExpectExec(QueryUpdateUserRole).
			WithArgs(userId, roleId, enable, userRoleId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewUserRolesRepository(clock, 60)

		err = r.UpdateUserRole(ctx, userId, userRoleId, updateUserRoleBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When update a policy of role return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userId := "739bbbc9-7e93-11ee-89fd-0442ac219255"
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		updateUserRoleBody := userRolesDomain.CreateUserRoleBody{
			RoleId: roleId,
			Enable: enable,
		}
		expectedError := errors.New("random error")
		clock := &mockClock.Clock{}
		mock.ExpectQuery(QueryUpdateUserRole).
			WithArgs(userId, roleId, enable, userRoleId).
			WillReturnError(expectedError)
		r := NewUserRolesRepository(clock, 60)
		err = r.UpdateUserRole(ctx, userId, userRoleId, updateUserRoleBody)
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
		assert.Equal(t, smartErr.Function, "UpdateUserRole")
	})
}

func TestRepositoryUserRoles_DeleteUserRole(t *testing.T) {
	t.Run("When delete a policy of role successfully", func(t *testing.T) {
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
		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteUserRole).
			WithArgs(deletedAt, userRoleId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewUserRolesRepository(clock, 60)
		var res bool
		res, err = r.DeleteUserRole(ctx, userId, userRoleId)
		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete a policy of role error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		userRoleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteUserRole).
			WithArgs(deletedAt, userRoleId).
			WillReturnError(errors.New("anything"))
		r := NewUserRolesRepository(clock, 60)
		var res bool
		res, err = r.DeleteUserRole(ctx, userId, userRoleId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteUserRole")
	})
}
