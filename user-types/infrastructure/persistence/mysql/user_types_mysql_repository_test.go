/*
 * File: user_types_mysql_repository_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to user types repository.
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

	userTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain"
)

func TestRepositoryUserTypes_GetUserTypes(t *testing.T) {
	t.Run("When get user types is called then it should return a list of user types", func(t *testing.T) {
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
		mockRegister := []userTypesDomain.UserType{
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Description: "Usuario externo",
				Code:        "USER_EXTERNAL",
				Enable:      true,
				CreatedAt:   &now,
			},
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110000",
				Description: "Usuario externo",
				Code:        "USER_EXTERNAL",
				Enable:      false,
				CreatedAt:   &now,
			},
		}
		rows := sqlmock.NewRows([]string{"id", "description", "code", "enable", "created_at"}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].Description,
				mockRegister[0].Code,
				mockRegister[0].Enable,
				mockRegister[0].CreatedAt,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].Description,
				mockRegister[1].Code,
				mockRegister[1].Enable,
				mockRegister[1].CreatedAt,
			)

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectQuery(QueryGetUserTypes).WithArgs(sizePage, offset).WillReturnRows(rows)
		r := NewUserTypesRepository(clock, 60)
		var res []userTypesDomain.UserType
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetUserTypes(ctx, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get user types is called then it should return an error", func(t *testing.T) {
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
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectQuery(QueryGetUserTypes).WillReturnError(expectedError)
		r := NewUserTypesRepository(clock, 60)
		var res []userTypesDomain.UserType
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetUserTypes(ctx, pagination)
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
		assert.Equal(t, smartErr.Function, "GetUserTypes")
	})
}

func TestRepositoryUserTypes_GetTotalUserTypes(t *testing.T) {
	t.Run("When get total of users types is called then it should return a total", func(t *testing.T) {
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
			ExpectQuery(QueryGetTotalUserTypes).
			WithArgs().
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUserTypesRepository(clock, 60)

		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalUserTypes(ctx, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total of users types is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		mock.ExpectQuery(QueryGetTotalUserTypes).
			WithArgs("739bbbc9-7e93-11ee-89fd-0242ac113421", "739bbbc9-7e93-11ee-89fd-0242ac113421").
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUserTypesRepository(clock, 60)

		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalUserTypes(ctx, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalUserTypes")
	})
}

func TestRepositoryUserTypes_CreateUserType(t *testing.T) {
	t.Run("When create user type success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		createUserTypeBody := userTypesDomain.CreateUserTypeBody{
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
			Enable:      true,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		userTypeId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryCreateUserType).
			WithArgs(
				userTypeId,
				createUserTypeBody.Description,
				createUserTypeBody.Code,
				createUserTypeBody.Enable,
				createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewUserTypesRepository(clock, 60)
		_, err = r.CreateUserType(ctx, userTypeId, createUserTypeBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When create user type return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		createUserTypeBody := userTypesDomain.CreateUserTypeBody{
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
			Enable:      true,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		userTypeId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreateUserType).
			WithArgs(
				userTypeId,
				createUserTypeBody.Description,
				createUserTypeBody.Code,
				createUserTypeBody.Enable,
				createdAt).
			WillReturnError(expectedError)
		r := NewUserTypesRepository(clock, 60)
		_, err = r.CreateUserType(ctx, userTypeId, createUserTypeBody)
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
		assert.Equal(t, smartErr.Function, "CreateUserType")
	})
}

func TestRepositoryUserTypes_UpdateUserType(t *testing.T) {
	t.Run("When update user type success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		updateUserTypeBody := userTypesDomain.UpdateUserTypeBody{
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
			Enable:      true,
		}
		now := time.Now().UTC()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		userTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		mock.ExpectExec(QueryUpdateUserType).
			WithArgs(
				updateUserTypeBody.Description,
				updateUserTypeBody.Code,
				updateUserTypeBody.Enable,
				userTypeId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewUserTypesRepository(clock, 60)

		err = r.UpdateUserType(ctx, userTypeId, updateUserTypeBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When update user type return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		updateUserTypeBody := userTypesDomain.UpdateUserTypeBody{
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
			Enable:      true,
		}
		userTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		now := time.Now().UTC()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectQuery(QueryUpdateUserType).
			WithArgs(
				updateUserTypeBody.Description,
				updateUserTypeBody.Code,
				updateUserTypeBody.Enable,
				userTypeId).
			WillReturnError(expectedError)
		r := NewUserTypesRepository(clock, 60)
		err = r.UpdateUserType(ctx, userTypeId, updateUserTypeBody)
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
		assert.Equal(t, smartErr.Function, "UpdateUserType")
	})
}

func TestRepositoryUserTypes_DeleteUserType(t *testing.T) {
	t.Run("When delete user type successfully", func(t *testing.T) {
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
		userTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteUserType).
			WithArgs(deletedAt, userTypeId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewUserTypesRepository(clock, 60)
		var res bool
		res, err = r.DeleteUserType(ctx, userTypeId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete user type error", func(t *testing.T) {
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
		userTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteUserType).
			WithArgs(deletedAt, userTypeId).
			WillReturnError(errors.New("anything"))
		r := NewUserTypesRepository(clock, 60)
		var res bool
		res, err = r.DeleteUserType(ctx, userTypeId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteUserType")
	})
}
