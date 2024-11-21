/*
 * File: merchants_mysql_repository_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains tests for the merchant repository.
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

	merchantsDomain "gitlab.smartcitiesperu.com/smartone/api-core/merchants/domain"
)

func TestRepositoryMerchants_GetMerchant(t *testing.T) {
	t.Run("When get merchants is called then it should return a list", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		mockRegister := []merchantsDomain.Merchant{
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Name:        "Odin Corp",
				Description: "Proveedor de servicios de mantenimiento",
				Phone:       "+1234567890",
				Document:    "123456789",
				Address:     "123 Main Street",
				Industry:    "Mantenimiento",
				ImagePath:   "https://example.com/images/odin_logo.png",
				CreatedAt:   &now,
			}, {
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110000",
				Name:        "Odin Corp",
				Description: "Proveedor de servicios de mantenimiento",
				Phone:       "+1234567890",
				Document:    "123456789",
				Address:     "123 Main Street",
				Industry:    "Mantenimiento",
				ImagePath:   "https://example.com/images/odin_logo.png",
				CreatedAt:   &now,
			},
		}
		rows := sqlmock.NewRows([]string{"id", "name", "description", "phone", "document", "address",
			"industry", "image_path", "created_at"}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].Name,
				mockRegister[0].Description,
				mockRegister[0].Phone,
				mockRegister[0].Document,
				mockRegister[0].Address,
				mockRegister[0].Industry,
				mockRegister[0].ImagePath,
				mockRegister[0].CreatedAt,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].Name,
				mockRegister[1].Description,
				mockRegister[1].Phone,
				mockRegister[1].Document,
				mockRegister[1].Address,
				mockRegister[1].Industry,
				mockRegister[1].ImagePath,
				mockRegister[1].CreatedAt,
			)
		sizePage := 100
		offset := 0
		clock := &mockClock.Clock{}
		mock.ExpectQuery(QueryGetMerchants).
			WithArgs(sizePage, offset).
			WillReturnRows(rows)
		r := NewMerchantsRepository(clock, 60)
		var res []merchantsDomain.Merchant
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetMerchants(ctx, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get merchant is called then it should return an error", func(t *testing.T) {
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
		clock := &mockClock.Clock{}
		mock.ExpectQuery(QueryGetMerchants).
			WithArgs(sizePage, offset).
			WillReturnError(expectedError)
		r := NewMerchantsRepository(clock, 60)
		var res []merchantsDomain.Merchant
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetMerchants(ctx, pagination)
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
		assert.Equal(t, smartErr.Function, "GetMerchants")
	})
}

func TestRepositoryMerchants_GetTotalMerchants(t *testing.T) {
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
		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)

		mock.
			ExpectQuery(QueryGetTotalMerchants).
			WithArgs().
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewMerchantsRepository(clock, 60)
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalMerchants(ctx, pagination)
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

		mock.ExpectQuery(QueryGetTotalMerchants).
			WithArgs("739bbbc9-7e93-11ee-89fd-0242ac113421", "739bbbc9-7e93-11ee-89fd-0242ac113421").
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewMerchantsRepository(clock, 60)

		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalMerchants(ctx, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalMerchants")

	})
}

func TestMerchant_CreateMerchant(t *testing.T) {
	t.Run("When create merchant success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		createMerchantBody := merchantsDomain.CreateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		now := time.Now().UTC()
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		createdAt := now.Format("2006-01-02 15:04:05")

		mock.ExpectExec(QueryCreateMerchant).
			WithArgs(merchantId,
				createMerchantBody.Name,
				createMerchantBody.Description,
				createMerchantBody.Phone,
				createMerchantBody.Document,
				createMerchantBody.Address,
				createMerchantBody.Industry,
				createMerchantBody.ImagePath,
				createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewMerchantsRepository(clock, 60)
		_, err = r.CreateMerchant(ctx, merchantId, createMerchantBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When create merchants return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		createMerchantBody := merchantsDomain.CreateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		now := time.Now().UTC()
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		createdAt := now.Format("2006-01-02 15:04:05")
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreateMerchant).
			WithArgs(merchantId,
				createMerchantBody.Name,
				createMerchantBody.Description,
				createMerchantBody.Phone,
				createMerchantBody.Document,
				createMerchantBody.Address,
				createMerchantBody.Industry,
				createMerchantBody.ImagePath,
				createdAt).
			WillReturnError(expectedError)
		r := NewMerchantsRepository(clock, 60)
		_, err = r.CreateMerchant(ctx, merchantId, createMerchantBody)
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
		assert.Equal(t, smartErr.Function, "CreateMerchant")
	})
}

func TestMerchant_UpdateMerchant(t *testing.T) {
	t.Run("When update merchants success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		updateMerchantBody := merchantsDomain.UpdateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		mock.ExpectExec(QueryUpdateMerchant).
			WithArgs(
				updateMerchantBody.Name,
				updateMerchantBody.Description,
				updateMerchantBody.Phone,
				updateMerchantBody.Document,
				updateMerchantBody.Address,
				updateMerchantBody.Industry,
				updateMerchantBody.ImagePath,
				merchantId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		clock := &mockClock.Clock{}
		r := NewMerchantsRepository(clock, 60)
		err = r.UpdateMerchant(ctx, merchantId, updateMerchantBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When update merchant return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		updateMerchantBody := merchantsDomain.UpdateMerchantBody{
			Name:        "Odin Corp",
			Description: "Proveedor de servicios de mantenimiento",
			Phone:       "+1234567890",
			Document:    "123456789",
			Address:     "123 Main Street",
			Industry:    "Mantenimiento",
			ImagePath:   "https://example.com/images/odin_logo.png",
		}
		clock := &mockClock.Clock{}
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdateMerchant).
			WithArgs(
				updateMerchantBody.Name,
				updateMerchantBody.Description,
				updateMerchantBody.Phone,
				updateMerchantBody.Document,
				updateMerchantBody.Address,
				updateMerchantBody.Industry,
				updateMerchantBody.ImagePath,
				merchantId).
			WillReturnError(expectedError)
		r := NewMerchantsRepository(clock, 60)
		err = r.UpdateMerchant(ctx, merchantId, updateMerchantBody)
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
		assert.Equal(t, smartErr.Function, "UpdateMerchant")
	})
}

func TestMerchant_DeleteMerchant(t *testing.T) {
	t.Run("When delete merchant successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteMerchant).
			WithArgs(deletedAt, merchantId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewMerchantsRepository(clock, 60)
		var res bool
		res, err = r.DeleteMerchant(ctx, merchantId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete merchant error", func(t *testing.T) {
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
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		mock.ExpectExec(QueryDeleteMerchant).
			WithArgs(deletedAt, merchantId).
			WillReturnError(errors.New("anything"))
		r := NewMerchantsRepository(clock, 60)
		var res bool
		res, err = r.DeleteMerchant(ctx, merchantId)

		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
