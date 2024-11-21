/*
 * File: stores_mysql_repository_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to store repository.
 *
 * Last Modified: 2023-11-14
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

	storesDomain "gitlab.smartcitiesperu.com/smartone/api-core/stores/domain"
)

func TestRepositoryStores_GetStores(t *testing.T) {
	t.Run("When get stores is called then it should return a list of stores", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storeTypeByStore := storesDomain.StoreTypeByStore{
			Id:           "739bbbc9-7e93-11ee-89fd-0242ac113421",
			Description:  "Maquinaria",
			Abbreviation: "Maq.",
		}

		mockRegister := []storesDomain.Store{
			{
				Id:         "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Name:       "Obra av. 28 julio",
				Shortname:  "Obra 28",
				MerchantId: merchantId,
				CreatedAt:  &now,
				StoreType:  storeTypeByStore,
			},
			{
				Id:         "739bbbc9-7e93-11ee-89fd-0242ac110017",
				Name:       "Obra av. 29 julio",
				Shortname:  "Obra 29",
				MerchantId: merchantId,
				CreatedAt:  &now,
				StoreType:  storeTypeByStore,
			},
		}
		rows := sqlmock.NewRows([]string{"store_id", "store_name", "store_shortname", "store_merchant_id",
			"store_created_at", "type_id", "type_description", "type_abbreviation"}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].Name,
				mockRegister[0].Shortname,
				mockRegister[0].MerchantId,
				mockRegister[0].CreatedAt,
				mockRegister[0].StoreType.Id,
				mockRegister[0].StoreType.Description,
				mockRegister[0].StoreType.Abbreviation,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].Name,
				mockRegister[1].Shortname,
				mockRegister[1].MerchantId,
				mockRegister[1].CreatedAt,
				mockRegister[1].StoreType.Id,
				mockRegister[1].StoreType.Description,
				mockRegister[1].StoreType.Abbreviation,
			)
		sizePage := 100
		offset := 0
		mock.ExpectQuery(QueryGetStores).
			WithArgs(
				merchantId,
				sizePage,
				offset).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewStoresRepository(clock, 60)
		var res []storesDomain.Store
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetStores(ctx, merchantId, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})
	t.Run("When get stores is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		sizePage := 100
		offset := 0
		mock.ExpectQuery(QueryGetStores).
			WithArgs(
				merchantId,
				sizePage,
				offset).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewStoresRepository(clock, 60)
		var res []storesDomain.Store
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetStores(ctx, merchantId, pagination)
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
		assert.Equal(t, smartErr.Function, "GetStores")
	})
}

func TestRepositoryStores_GetTotalStores(t *testing.T) {
	t.Run("When get total of stores is called then it should return a total", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		total := 10
		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)
		mock.
			ExpectQuery(QueryGetTotalStores).
			WithArgs(merchantId).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewStoresRepository(clock, 60)

		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalStores(ctx, merchantId, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})
	t.Run("When get total of users is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		mock.ExpectQuery(QueryGetTotalStores).
			WithArgs(merchantId).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewStoresRepository(clock, 60)
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalStores(ctx, merchantId, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalStores")
	})
}

func TestRepositoryStores_CreateStore(t *testing.T) {
	t.Run("When create a store success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		storeId := "739bbbc9-7e93-11ee-89fd-042hs5278000"
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		createStoreBody := storesDomain.CreateStoreBody{
			Name:        "Obra av. 28 julio",
			Shortname:   "Obra 28",
			StoreTypeId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectExec(QueryCreateStore).
			WithArgs(
				storeId,
				createStoreBody.Name,
				createStoreBody.Shortname,
				merchantId,
				createStoreBody.StoreTypeId,
				createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewStoresRepository(clock, 60)

		_, err = r.CreateStore(
			ctx,
			merchantId,
			storeId,
			createStoreBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})
	t.Run("When create a store return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		storeId := "739bbbc9-7e93-11ee-89fd-042hs5278000"
		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		createStoreBody := storesDomain.CreateStoreBody{
			Name:        "Obra av. 28 julio",
			Shortname:   "Obra 28",
			StoreTypeId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreateStore).
			WithArgs(
				storeId,
				createStoreBody.Name,
				createStoreBody.Shortname,
				merchantId,
				createStoreBody.StoreTypeId,
				createdAt).
			WillReturnError(expectedError)
		r := NewStoresRepository(clock, 60)
		_, err = r.CreateStore(
			ctx,
			merchantId,
			storeId,
			createStoreBody)
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
		assert.Equal(t, smartErr.Function, "CreateStore")
	})
}

func TestRepositoryStores_UpdateStore(t *testing.T) {
	t.Run("When update a store success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		updateStoreBody := storesDomain.CreateStoreBody{
			Name:        "Obra av. 28 julio",
			Shortname:   "Obra 28",
			StoreTypeId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
		}
		mock.ExpectExec(QueryUpdateStore).
			WithArgs(
				updateStoreBody.Name,
				updateStoreBody.Shortname,
				merchantId,
				updateStoreBody.StoreTypeId,
				storeId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		clock := &mockClock.Clock{}
		r := NewStoresRepository(clock, 60)

		err = r.UpdateStore(
			ctx,
			merchantId,
			storeId,
			updateStoreBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})
	t.Run("When update a store return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchantId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		updateStoreBody := storesDomain.CreateStoreBody{
			Name:        "Obra av. 28 julio",
			Shortname:   "Obra 28",
			StoreTypeId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
		}
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdateStore).
			WithArgs(
				updateStoreBody.Name,
				updateStoreBody.Shortname,
				merchantId,
				updateStoreBody.StoreTypeId,
				storeId).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewStoresRepository(clock, 60)
		err = r.UpdateStore(ctx, merchantId, storeId, updateStoreBody)
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
		assert.Equal(t, smartErr.Function, "UpdateStore")
	})
}

func TestRepositoryStores_DeleteStore(t *testing.T) {
	t.Run("When delete a store successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteStore).
			WithArgs(deletedAt, storeId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewStoresRepository(clock, 60)
		var res bool
		res, err = r.DeleteStore(ctx, merchantId, storeId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})
	t.Run("When delete a store error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		merchantId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		deletedAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteStore).
			WithArgs(deletedAt, storeId).
			WillReturnError(errors.New("anything"))
		r := NewStoresRepository(clock, 60)

		var res bool
		res, err = r.DeleteStore(ctx, merchantId, storeId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteStore")
	})
}
