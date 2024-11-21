/*
 * File: economic_activities_mysql_repository_test.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the test of the economic activities repository.
 *
 * Last Modified: 2023-12-04
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

	economicActivitiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/domain"
)

func TestRepositoryEconomicActivities_GetEconomicActivities(t *testing.T) {
	t.Run("When get economic activities is called then it should return a list of economic activities", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		var description = "CULTIVO DE ARROZ"
		mockEconomicActivities := []economicActivitiesDomain.EconomicActivity{
			{
				Id:          "703e039e-92be-11ee-a040-0242ac11000e",
				CuuiId:      "0111",
				Description: &description,
				Status:      1,
				CreatedAt:   &now,
			},
		}

		rows := sqlmock.NewRows([]string{"economic_activity_id", "economic_activity_description",
			"economic_activity_cuuiid", "economic_activity_status", "economic_activity_created_at"}).AddRow(
			mockEconomicActivities[0].Id,
			mockEconomicActivities[0].CuuiId,
			mockEconomicActivities[0].Description,
			mockEconomicActivities[0].Status,
			mockEconomicActivities[0].CreatedAt,
		)

		cuuiId := "0111"
		sizePage := 100
		offset := 0
		mock.ExpectQuery(QueryGetEconomicActivities).WithArgs(
			cuuiId,
			cuuiId,
			description,
			description, sizePage, offset).WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewEconomicActivitiesRepository(clock, 60)
		var res []economicActivitiesDomain.EconomicActivity
		searchParams := economicActivitiesDomain.GetEconomicActivitiesParams{
			CuuiId:      "0111",
			Description: &description,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetEconomicActivities(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 1)
	})

	t.Run("When get economic activities is called then it should return error", func(t *testing.T) {
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
		cuuiId := "0111"
		var description = "CULTIVO DE ARROZ"
		mock.ExpectQuery(QueryGetEconomicActivities).WithArgs(
			cuuiId,
			cuuiId,
			description,
			description, sizePage, offset).WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewEconomicActivitiesRepository(clock, 60)
		var res []economicActivitiesDomain.EconomicActivity
		searchParams := economicActivitiesDomain.GetEconomicActivitiesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetEconomicActivities(ctx, searchParams, pagination)
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
		assert.Equal(t, smartErr.Function, "GetEconomicActivities")
	})
}

func TestRepositoryEconomicActivities_GetTotalGetEconomicActivities(t *testing.T) {
	t.Run("When get total of economic activities is called then it should return a total", func(t *testing.T) {
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
			ExpectQuery(QueryGetTotalEconomicActivities).
			WithArgs().
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewEconomicActivitiesRepository(clock, 60)

		searchParams := economicActivitiesDomain.GetEconomicActivitiesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalGetEconomicActivities(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total of economic activities is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryGetTotalEconomicActivities).
			WithArgs().
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewEconomicActivitiesRepository(clock, 60)

		searchParams := economicActivitiesDomain.GetEconomicActivitiesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalGetEconomicActivities(ctx, searchParams, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalGetEconomicActivities")
	})
}
