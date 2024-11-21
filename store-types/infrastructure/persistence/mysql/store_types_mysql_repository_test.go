/*
 * File: store_types_mysql_repository_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the store type repository.
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

	storeTypeDomain "gitlab.smartcitiesperu.com/smartone/api-core/store-types/domain"
)

func TestRepositoryStoreTypes_GetStoreTypes(t *testing.T) {
	t.Run("When get store types called then it should return a list of specialties", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		sizePage := 100
		mockRegister := []storeTypeDomain.StoreType{
			{
				Id:           "73900000-7e93-11ee-89fd-0242a500000",
				Description:  "Maquinaria",
				Abbreviation: "Maq",
			}, {
				Id:           "fcdbfacf-8305-11ee-89fd-0242555555",
				Description:  "Maquinaria",
				Abbreviation: "Maq",
			},
		}
		rows := sqlmock.NewRows([]string{"id", "description", "abbreviation"}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].Description,
				mockRegister[0].Abbreviation,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].Description,
				mockRegister[1].Abbreviation,
			)
		now := time.Now().UTC()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectQuery(QueryGetStoreTypes).
			WillReturnRows(rows)
		r := NewStoreTypesRepository(clock, 60)

		var res []storeTypeDomain.StoreType
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetStoreTypes(ctx, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get store types is called then it should return an error", func(t *testing.T) {
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
		now := time.Now().UTC()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectQuery(QueryGetStoreTypes).WillReturnError(expectedError)
		r := NewStoreTypesRepository(clock, 60)
		var res []storeTypeDomain.StoreType
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetStoreTypes(ctx, pagination)
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
		assert.Equal(t, smartErr.Function, "GetStoreTypes")
	})
}

func TestStoreTypeMySQLRepo_GetTotalStoreTypes(t *testing.T) {
	t.Run("When get total store types called then it should return a total", func(t *testing.T) {
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
		mock.ExpectQuery(QueryGetTotalStoreTypes).
			WithArgs().
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewStoreTypesRepository(clock, 60)

		var res *int
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetTotalStoreTypes(ctx, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, total, *res)
	})

	t.Run("When get total store types is called then it should return an error", func(t *testing.T) {
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
		now := time.Now().UTC()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectQuery(QueryGetTotalStoreTypes).WillReturnError(expectedError)
		r := NewStoreTypesRepository(clock, 60)
		var res *int
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetTotalStoreTypes(ctx, pagination)
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
		assert.Equal(t, smartErr.Function, "GetTotalStoreTypes")
	})
}

func TestRepositoryStoreTypes_CreateStoreType(t *testing.T) {
	t.Run("When create an store type success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		storeTypeId := uuid.New().String()
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createStoreType := storeTypeDomain.CreateStoreTypeBody{
			Description:  "Maquinaria",
			Abbreviation: "Maq",
		}

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryCreateStoreType).
			WithArgs(storeTypeId, createStoreType.Description, createStoreType.Abbreviation, createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewStoreTypesRepository(clock, 60)

		_, err = r.CreateStoreType(ctx, storeTypeId, createStoreType)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When create an store type return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		storeTypeId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createStoreType := storeTypeDomain.CreateStoreTypeBody{
			Description:  "Maquinaria",
			Abbreviation: "Maq",
		}
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreateStoreType).
			WithArgs(storeTypeId, createStoreType.Description, createStoreType.Abbreviation, createdAt).
			WillReturnError(expectedError)
		r := NewStoreTypesRepository(clock, 60)
		_, err = r.CreateStoreType(ctx, storeTypeId, createStoreType)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)
	})
}

func TestRepositoryStoreTypes_UpdateStoreType(t *testing.T) {
	t.Run("When update store type success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		storeTypeId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		now := time.Now().UTC()
		updateStoreType := storeTypeDomain.UpdateStoreTypeBody{
			Description:  "Maquinaria",
			Abbreviation: "Maq",
		}
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryUpdateStoreType).
			WithArgs(updateStoreType.Description, updateStoreType.Abbreviation, storeTypeId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewStoreTypesRepository(clock, 60)
		err = r.UpdateStoreType(ctx, storeTypeId, updateStoreType)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When update store type return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		storeTypeId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		now := time.Now().UTC()
		updateStoreType := storeTypeDomain.UpdateStoreTypeBody{
			Description:  "Maquinaria",
			Abbreviation: "Maq",
		}
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdateStoreType).
			WithArgs(updateStoreType.Description, updateStoreType.Abbreviation, storeTypeId).
			WillReturnError(expectedError)
		r := NewStoreTypesRepository(clock, 60)
		err = r.UpdateStoreType(ctx, storeTypeId, updateStoreType)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)
	})
}

func TestRepositoryStoreTypes_DeleteStoreType(t *testing.T) {
	t.Run("When delete store type successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		storeTypeId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteStoreType).
			WithArgs(deletedAt, storeTypeId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewStoreTypesRepository(clock, 60)
		var res bool
		res, err = r.DeleteStoreType(ctx, storeTypeId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete store type error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		storeTypeId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		now := time.Now().UTC()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteStoreType).
			WithArgs(now, storeTypeId).
			WillReturnError(errors.New("anything"))
		r := NewStoreTypesRepository(clock, 60)
		var res bool
		res, err = r.DeleteStoreType(ctx, "73900000-7e93-11ee-89fd-0242a500000")
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
