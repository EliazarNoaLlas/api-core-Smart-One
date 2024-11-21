/*
 * File: views_mysql_repository_test.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to views repository.
 *
 * Last Modified: 2023-11-24
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

	viewDomain "gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
)

func TestRepositoryViews_GetViews(t *testing.T) {
	t.Run("When get views by module is called then it should return a list of views", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now()
		views := []viewDomain.View{
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Name:        "Requerimientos",
				Description: "Vista para el registro de requerimientos",
				Url:         "/logistics/requirements",
				Icon:        "fa fa-table",
				CreatedAt:   &now,
			},
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110012",
				Name:        "Requerimientos",
				Description: "Vista para el registro de requerimientos",
				Url:         "/logistics/requirements",
				Icon:        "fa fa-table",
				CreatedAt:   &now,
			},
		}
		rows := sqlmock.NewRows([]string{
			"view_id",
			"view_name",
			"view_description",
			"view_url",
			"view_icon",
			"view_created_at",
		}).
			AddRow(
				views[0].Id,
				views[0].Name,
				views[0].Description,
				views[0].Url,
				views[0].Icon,
				views[0].CreatedAt,
			).
			AddRow(
				views[1].Id,
				views[1].Name,
				views[1].Description,
				views[1].Url,
				views[1].Icon,
				views[1].CreatedAt,
			)

		moduleId := "fcdbfacf-8305-11ee-89fd-024255555501"
		sizePage := 100
		offset := 0
		var name *string = nil
		mock.
			ExpectQuery(QueryGetViews).
			WithArgs(moduleId, name, name, name, sizePage, offset).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewViewRepository(clock, 60)
		var res []viewDomain.View
		searchParams := viewDomain.GetViewsParams{
			Name: name,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetViews(ctx, moduleId, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get views by module is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		moduleId := "fcdbfacf-8305-11ee-89fd-024255555501"
		sizePage := 100
		offset := 0
		var name *string = nil
		mock.ExpectQuery(QueryGetViews).
			WithArgs(moduleId, name, name, name, sizePage, offset).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewViewRepository(clock, 60)
		var res []viewDomain.View
		searchParams := viewDomain.GetViewsParams{
			Name: name,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetViews(ctx, moduleId, searchParams, pagination)
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
		assert.Equal(t, smartErr.Function, "GetViews")
	})
}

func TestRepositoryViews_GetTotalViews(t *testing.T) {
	t.Run("When get total of views is called then it should return a total", func(t *testing.T) {
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

		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac113421"
		var name *string = nil
		mock.
			ExpectQuery(QueryGetTotalViews).
			WithArgs(moduleId, name, name, name).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewViewRepository(clock, 60)

		searchParams := viewDomain.GetViewsParams{
			Name: name,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalViews(ctx, moduleId, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total of views is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac113421"
		var name *string = nil
		mock.ExpectQuery(QueryGetTotalViews).
			WithArgs(moduleId, name, name, name).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewViewRepository(clock, 60)

		searchParams := viewDomain.GetViewsParams{
			Name: name,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalViews(ctx, moduleId, searchParams, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalViews")
	})
}

func TestRepositoryViews_CreateView(t *testing.T) {
	t.Run("When add a view success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		moduleId := uuid.New().String()
		viewId := uuid.New().String()
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createViewBody := viewDomain.CreateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryCreateView).
			WithArgs(
				viewId,
				createViewBody.Name,
				createViewBody.Description,
				createViewBody.Url,
				createViewBody.Icon,
				moduleId,
				createdAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewViewRepository(clock, 60)

		_, err = r.CreateView(
			ctx,
			moduleId,
			viewId,
			createViewBody,
		)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When add a view return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		moduleId := uuid.New().String()
		viewId := uuid.New().String()
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createViewBody := viewDomain.CreateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectExec(QueryCreateView).
			WithArgs(
				viewId,
				createViewBody.Name,
				createViewBody.Description,
				createViewBody.Url,
				createViewBody.Icon,
				moduleId,
				createdAt,
			).
			WillReturnError(expectedError)
		r := NewViewRepository(clock, 60)
		_, err = r.CreateView(ctx, moduleId, viewId, createViewBody)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)
	})
}

func TestRepositoryViews_UpdateView(t *testing.T) {
	t.Run("When update a view successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		moduleId := uuid.New().String()
		viewId := uuid.New().String()
		updateViewBody := viewDomain.UpdateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		clock := &mockClock.Clock{}
		mock.ExpectExec(QueryUpdateView).
			WithArgs(
				updateViewBody.Name,
				updateViewBody.Description,
				updateViewBody.Url,
				updateViewBody.Icon,
				moduleId,
				viewId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewViewRepository(clock, 60)

		err = r.UpdateView(ctx, moduleId, viewId, updateViewBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When update a view return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		moduleId := uuid.New().String()
		viewId := uuid.New().String()
		updateViewBody := viewDomain.UpdateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		expectedError := errors.New("random error")
		clock := &mockClock.Clock{}
		mock.ExpectExec(QueryUpdateView).
			WithArgs(
				updateViewBody.Name,
				updateViewBody.Description,
				updateViewBody.Url,
				updateViewBody.Icon,
				moduleId,
				viewId,
			).
			WillReturnError(expectedError)
		r := NewViewRepository(clock, 60)
		err = r.UpdateView(ctx, moduleId, viewId, updateViewBody)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)
	})
}

func TestRepositoryViews_DeleteView(t *testing.T) {
	t.Run("When delete a view successfully", func(t *testing.T) {
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
		moduleId := uuid.New().String()
		viewId := uuid.New().String()

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteView).
			WithArgs(deletedAt, viewId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewViewRepository(clock, 60)
		var res bool
		res, err = r.DeleteView(ctx, moduleId, viewId)
		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete a view error", func(t *testing.T) {
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
		moduleId := uuid.New().String()
		viewId := "739bbbc9-7e93-11ee-89fd-0442ac210931"

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteView).
			WithArgs(deletedAt, viewId).
			WillReturnError(errors.New("anything"))
		r := NewViewRepository(clock, 60)
		var res bool
		res, err = r.DeleteView(ctx, moduleId, viewId)

		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
