/*
 * File: modules_mysql_repository_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to module repository.
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
	"github.com/stretchr/testify/assert"

	mockClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock/mocks"
	db2 "gitlab.smartcitiesperu.com/smartone/api-shared/db"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	modulesDomain "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"
)

func TestRepositoryModules_GetModules(t *testing.T) {
	t.Run("When get modules is called then it should return a list of modules", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		code := "logistic"
		searchParams := modulesDomain.GetModulesParams{
			Params: paramsDomain.Params{},
			Name:   nil,
			Code:   &code,
		}
		mockRegister := []modulesDomain.Module{
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Name:        "Logistica",
				Description: "Modulo de logistica",
				Code:        "logistic",
				Icon:        "fa fa-chart",
				Position:    1,
				CreatedAt:   &now,
			},
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110017",
				Name:        "Tesoreria",
				Description: "Modulo de tesoreria",
				Code:        "treasury",
				Icon:        "fa fa-home",
				Position:    2,
				CreatedAt:   &now,
			},
		}
		rows := sqlmock.NewRows([]string{"id", "name", "description", "code", "icon", "position",
			"created_at"}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].Name,
				mockRegister[0].Description,
				mockRegister[0].Code,
				mockRegister[0].Icon,
				mockRegister[0].Position,
				mockRegister[0].CreatedAt,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].Name,
				mockRegister[1].Description,
				mockRegister[1].Code,
				mockRegister[1].Icon,
				mockRegister[1].Position,
				mockRegister[1].CreatedAt,
			)
		sizePage := 100
		offset := 0
		clock := &mockClock.Clock{}
		mock.ExpectQuery(QueryGetModules).
			WithArgs(
				searchParams.Code,
				searchParams.Code,
				searchParams.Name,
				searchParams.Name,
				sizePage,
				offset).
			WillReturnRows(rows)
		r := NewModulesRepository(clock, 60)
		var res []modulesDomain.Module
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetModules(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get modules is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		code := "logistic"
		searchParams := modulesDomain.GetModulesParams{
			Params: paramsDomain.Params{},
			Name:   nil,
			Code:   &code,
		}
		sizePage := 100
		offset := 0
		clock := &mockClock.Clock{}
		mock.
			ExpectQuery(QueryGetModules).
			WithArgs(
				searchParams.Code,
				searchParams.Code,
				searchParams.Name,
				searchParams.Name,
				sizePage,
				offset).
			WillReturnError(expectedError)
		r := NewModulesRepository(clock, 60)
		var res []modulesDomain.Module
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetModules(ctx, searchParams, pagination)
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
		assert.Equal(t, smartErr.Function, "GetModules")
	})
}

func TestRepositoryModules_GetTotalModules(t *testing.T) {
	t.Run("When get total modules return success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		total := 10
		var totalExpected *int
		code := "logistic"
		searchParams := modulesDomain.GetModulesParams{
			Params: paramsDomain.Params{},
			Name:   nil,
			Code:   &code,
		}
		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)

		mock.
			ExpectQuery(QueryGetTotalModules).
			WithArgs(
				searchParams.Code,
				searchParams.Code,
				searchParams.Name,
				searchParams.Name,
			).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewModulesRepository(clock, 60)

		totalExpected, err = r.GetTotalModules(ctx, searchParams)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total modules return error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		var totalExpected *int
		code := "logistic"
		searchParams := modulesDomain.GetModulesParams{
			Params: paramsDomain.Params{},
			Name:   nil,
			Code:   &code,
		}
		mock.ExpectQuery(QueryGetTotalModules).
			WithArgs(
				searchParams.Code,
				searchParams.Code,
				searchParams.Name,
				searchParams.Name,
			).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewModulesRepository(clock, 60)

		totalExpected, err = r.GetTotalModules(ctx, searchParams)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalModules")

	})
}

func TestRepositoryModules_CreateModule(t *testing.T) {
	t.Run("When create module success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		createModuleBody := modulesDomain.CreateModuleBody{
			Name:        "Logistic",
			Description: "Modulo de logística",
			Code:        "logistic",
			Icon:        "fa fa-chart",
			Position:    1,
		}
		now := time.Now().UTC()
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		createdAt := now.Format("2006-01-02 15:04:05")
		mock.ExpectExec(QueryCreateModule).
			WithArgs(moduleId,
				createModuleBody.Name,
				createModuleBody.Description,
				createModuleBody.Code,
				createModuleBody.Icon,
				createModuleBody.Position,
				createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewModulesRepository(clock, 60)
		_, err = r.CreateModule(ctx, moduleId, createModuleBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When create modules return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		createModuleBody := modulesDomain.CreateModuleBody{
			Name:        "Logistic",
			Description: "Modulo de logística",
			Code:        "logistic",
			Icon:        "fa fa-chart",
			Position:    1,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreateModule).
			WithArgs(moduleId,
				createModuleBody.Name,
				createModuleBody.Description,
				createModuleBody.Code,
				createModuleBody.Icon,
				createModuleBody.Position,
				createdAt).
			WillReturnError(expectedError)
		r := NewModulesRepository(clock, 60)
		_, err = r.CreateModule(ctx, moduleId, createModuleBody)
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
		assert.Equal(t, smartErr.Function, "CreateModule")
	})
}

func TestRepositoryModules_UpdateModule(t *testing.T) {
	t.Run("When update modules success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		updateModuleBody := modulesDomain.UpdateModuleBody{
			Name:        "Logistic",
			Description: "Modulo de logística",
			Code:        "logistic",
			Icon:        "fa fa-chart",
			Position:    1,
		}
		clock := &mockClock.Clock{}
		mock.ExpectExec(QueryUpdateModule).
			WithArgs(
				updateModuleBody.Name,
				updateModuleBody.Description,
				updateModuleBody.Code,
				updateModuleBody.Icon,
				updateModuleBody.Position,
				moduleId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewModulesRepository(clock, 60)

		err = r.UpdateModule(ctx, moduleId, updateModuleBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When update module return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		updateModuleBody := modulesDomain.UpdateModuleBody{
			Name:        "Logistic",
			Description: "Modulo de logística",
			Code:        "logistic",
			Icon:        "fa fa-chart",
			Position:    1,
		}
		clock := &mockClock.Clock{}
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdateModule).
			WithArgs(
				updateModuleBody.Name,
				updateModuleBody.Description,
				updateModuleBody.Code,
				updateModuleBody.Icon,
				updateModuleBody.Position,
				moduleId).
			WillReturnError(expectedError)
		r := NewModulesRepository(clock, 60)
		err = r.UpdateModule(ctx, moduleId, updateModuleBody)
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
		assert.Equal(t, smartErr.Function, "UpdateModule")
	})
}

func TestRepositoryModules_DeleteModule(t *testing.T) {
	t.Run("When delete module successfully", func(t *testing.T) {
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
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteModule).
			WithArgs(deletedAt, moduleId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewModulesRepository(clock, 60)
		var res bool
		res, err = r.DeleteModule(ctx, moduleId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete module error", func(t *testing.T) {
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
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		mock.ExpectExec(QueryDeleteModule).
			WithArgs(deletedAt, moduleId).
			WillReturnError(errors.New("anything"))
		r := NewModulesRepository(clock, 60)
		var res bool
		res, err = r.DeleteModule(ctx, moduleId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteModule")
	})
}
