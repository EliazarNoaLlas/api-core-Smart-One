/*
 * File: merchant_economic_activities_mysql_repository_test.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository test merchant economic activities
 *
 * Last Modified: 2023-12-05
 */

package infrastructure

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	mockClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock/mocks"
	db2 "gitlab.smartcitiesperu.com/smartone/api-shared/db"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	merchantEconomicActivitiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/domain"
)

func TestRepositoryMerchantEconomicActivities_GetMerchantEconomicActivities(t *testing.T) {
	t.Run("When get merchant activities is called it should return a merchant activities", func(t *testing.T) {
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
		activity := merchantEconomicActivitiesDomain.EconomicActivityByMerchant{
			Id:          "70402269-92be-11ee-a040-0242ac11000e",
			CuuiId:      "0111",
			Description: &description,
			Status:      1,
			CreatedAt:   &now,
		}

		mockMerchantEconomicActivity := []merchantEconomicActivitiesDomain.MerchantEconomicActivity{
			{
				Id:               "22d4b62a-9380-11ee-a040-0242ac11000e",
				Sequence:         1,
				CreatedAt:        &now,
				EconomicActivity: activity,
			},
		}

		rows := sqlmock.NewRows([]string{
			"merchant_economic_id",
			"merchant_economic_sequence",
			"merchant_economic_created_at",
			"activity_id",
			"activity_cuui_id",
			"activity_description",
			"activity_status",
			"activity_created_at",
		}).AddRow(
			mockMerchantEconomicActivity[0].Id,
			mockMerchantEconomicActivity[0].Sequence,
			mockMerchantEconomicActivity[0].CreatedAt,
			mockMerchantEconomicActivity[0].EconomicActivity.Id,
			mockMerchantEconomicActivity[0].EconomicActivity.CuuiId,
			mockMerchantEconomicActivity[0].EconomicActivity.Description,
			mockMerchantEconomicActivity[0].EconomicActivity.Status,
			mockMerchantEconomicActivity[0].EconomicActivity.CreatedAt,
		)

		sizePage := 100
		offset := 0
		merchantId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		mock.ExpectQuery(QueryGetMerchantEconomicActivities).WithArgs(merchantId, sizePage, offset).WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewMerchantEconomicActivitiesRepository(clock, 60)
		var res []merchantEconomicActivitiesDomain.MerchantEconomicActivity
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetMerchantEconomicActivities(ctx, merchantId, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 1)
	})

	t.Run("When get merchant activities is called it should return a merchant activities ", func(t *testing.T) {
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
		mock.ExpectQuery(QueryGetMerchantEconomicActivities).
			WithArgs(sizePage, offset).WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewMerchantEconomicActivitiesRepository(clock, 60)
		var res []merchantEconomicActivitiesDomain.MerchantEconomicActivity
		pagination := paramsDomain.NewPaginationParams(nil)

		merchantId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		res, err = r.GetMerchantEconomicActivities(ctx, merchantId, pagination)
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
		assert.Equal(t, smartErr.Function, "GetMerchantEconomicActivities")
	})
}

func TestRepositoryMerchantEconomicActivities_GetTotalEconomicActivities(t *testing.T) {
	t.Run("When get total of merchant activities is called then it should return a total", func(t *testing.T) {
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
			ExpectQuery(QueryGetTotalMerchantEconomicActivities).WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewMerchantEconomicActivitiesRepository(clock, 60)
		merchantId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalEconomicActivities(ctx, merchantId, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total of merchant activities is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		mock.ExpectQuery(QueryGetTotalMerchantEconomicActivities).WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewMerchantEconomicActivitiesRepository(clock, 60)

		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		merchantId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		totalExpected, err = r.GetTotalEconomicActivities(ctx, merchantId, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalEconomicActivities")
	})
}

func TestRepositoryMerchantEconomicActivities_CreateEconomicActivities(t *testing.T) {
	t.Run("When creating a merchant economic activities is successful, it should return without errors.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantEconomicActivityId := uuid.New().String()
		createMerchantEconomicActivityBody := merchantEconomicActivitiesDomain.CreateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}

		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryCreateMerchantEconomicActivity).
			WithArgs(merchantEconomicActivityId,
				createMerchantEconomicActivityBody.MerchantId,
				createMerchantEconomicActivityBody.EconomicActivityId,
				createMerchantEconomicActivityBody.Sequence,
				createdAt,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		r := NewMerchantEconomicActivitiesRepository(clock, 60)
		_, err = r.CreateEconomicActivity(ctx, merchantEconomicActivityId, createMerchantEconomicActivityBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When creating a merchant economic activities returns an error, it should handle the error.", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantEconomicActivityId := uuid.New().String()
		createMerchantEconomicActivityBody := merchantEconomicActivitiesDomain.CreateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(time.Now())
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreateMerchantEconomicActivity).
			WithArgs(merchantEconomicActivityId,
				createMerchantEconomicActivityBody.MerchantId,
				createMerchantEconomicActivityBody.EconomicActivityId,
				createMerchantEconomicActivityBody.Sequence,
				createdAt).
			WillReturnError(expectedError)
		r := NewMerchantEconomicActivitiesRepository(clock, 60)
		_, err = r.CreateEconomicActivity(ctx, merchantEconomicActivityId, createMerchantEconomicActivityBody)
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
		assert.Equal(t, smartErr.Function, "CreateEconomicActivity")
	})
}

func TestRepositoryMerchantEconomicActivities_UpdateEconomicActivity(t *testing.T) {
	t.Run("When the merchant economic activities is prescribed successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantEconomicActivityId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		updateMerchantEconomicActivityBody := merchantEconomicActivitiesDomain.UpdateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		clock := &mockClock.Clock{}
		mock.ExpectExec(QueryUpdateMerchantEconomicActivity).
			WithArgs(
				updateMerchantEconomicActivityBody.MerchantId,
				updateMerchantEconomicActivityBody.EconomicActivityId,
				updateMerchantEconomicActivityBody.Sequence,
				merchantEconomicActivityId,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewMerchantEconomicActivitiesRepository(clock, 60)
		err = r.UpdateEconomicActivity(ctx, merchantEconomicActivityId, updateMerchantEconomicActivityBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When an error occurs while updating a merchant economic activities", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantEconomicActivityId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		updateMerchantEconomicActivityBody := merchantEconomicActivitiesDomain.UpdateMerchantEconomicActivityBody{
			MerchantId:         "cf6e4017-f918-4ef0-b641-236d89901a5c",
			EconomicActivityId: "70402269-92be-11ee-a040-0242ac11000e",
			Sequence:           1,
		}
		clock := &mockClock.Clock{}
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdateMerchantEconomicActivity).
			WithArgs(
				updateMerchantEconomicActivityBody.MerchantId,
				updateMerchantEconomicActivityBody.EconomicActivityId,
				updateMerchantEconomicActivityBody.Sequence,
				merchantEconomicActivityId,
			).
			WillReturnError(expectedError)
		r := NewMerchantEconomicActivitiesRepository(clock, 60)
		err = r.UpdateEconomicActivity(ctx, merchantEconomicActivityId, updateMerchantEconomicActivityBody)
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
		assert.Equal(t, smartErr.Function, "UpdateEconomicActivity")
	})

}

func TestRepositoryMerchantEconomicActivities_DeleteEconomicActivity(t *testing.T) {
	t.Run("When deleting a merchant economic activities is successful", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantEconomicActivityId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		clock := &mockClock.Clock{}
		now := time.Now().UTC().Format("2006-01-02 15:04:05")
		mock.ExpectExec(QueryDeleteMerchantEconomicActivity).
			WithArgs(now, merchantEconomicActivityId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewMerchantEconomicActivitiesRepository(clock, 60)
		var res bool
		res, err = r.DeleteEconomicActivity(ctx, merchantEconomicActivityId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When deleting a merchant economic activities returns an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		clock := &mockClock.Clock{}
		merchantEconomicActivityId := "22d4b62a-9380-11ee-a040-0242ac11000e"
		now := time.Now().UTC().Format("2006-01-02 15:04:05")
		mock.ExpectExec(QueryDeleteMerchantEconomicActivity).
			WithArgs(now, merchantEconomicActivityId).
			WillReturnError(errors.New("anything"))
		r := NewMerchantEconomicActivitiesRepository(clock, 60)
		var res bool
		res, err = r.DeleteEconomicActivity(ctx, merchantEconomicActivityId)
		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteEconomicActivity")
	})
}
