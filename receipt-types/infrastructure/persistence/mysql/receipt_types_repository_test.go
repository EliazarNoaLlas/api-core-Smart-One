/*
 * File: receipt_types_repository_test.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the tests for the receiptTypes repository.
 *
 * Last Modified: 2024-03-06
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

	receiptTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
)

func TestRepositoryReceiptTypes_GetReceiptTypes(t *testing.T) {
	t.Run("When get receipt types called then it should return a list of specialties", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		clock := &mockClock.Clock{}
		var res []receiptTypesDomain.ReceiptType
		mockRegister := []receiptTypesDomain.ReceiptType{
			{
				Id:          "18f7f9c2-b00a-42e4-a469-ea4c01c180dd",
				Description: "Recibo por Honorarios",
				SunatCode:   "02",
				Enable:      true,
				CreatedBy:   "91fb86bd-da46-414b-97a1-fcdaa8cd35d1",
				CreatedAt:   &now,
			},
			{
				Id:          "2e8bfdbb-1a58-4b45-9a2c-8ac54a5db723",
				Description: "Recibo por Arrendamiento",
				SunatCode:   "03",
				Enable:      false,
				CreatedBy:   "c3f92a0d-ef58-4e15-a71b-6f0a8b9d147d",
				CreatedAt:   &now,
			},
		}

		rows := sqlmock.NewRows([]string{
			"receipt_type_id",
			"receipt_type_description",
			"receipt_type_sunat_code",
			"receipt_type_enable",
			"receipt_type_created_by",
			"receipt_type_created_at",
		})
		for _, register := range mockRegister {
			rows = rows.AddRow(
				register.Id,
				register.Description,
				register.SunatCode,
				register.Enable,
				register.CreatedBy,
				register.CreatedAt,
			)
		}

		mock.ExpectQuery(QueryGetReceiptTypes).
			WithArgs().
			WillReturnRows(rows)

		r := NewReceiptTypesRepository(clock, 60)
		res, err = r.GetReceiptTypes(ctx)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, len(mockRegister))
	})

	t.Run("When get receipt types is called then it should return an error", func(t *testing.T) {
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
		var res []receiptTypesDomain.ReceiptType

		mock.ExpectQuery(QueryGetReceiptTypes).
			WithArgs().
			WillReturnError(expectedError)

		r := NewReceiptTypesRepository(clock, 60)
		res, err = r.GetReceiptTypes(ctx)
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
		assert.Equal(t, smartErr.Function, "GetReceiptTypes")
	})
}

func TestRepositoryReceiptTypes_CreateReceiptType(t *testing.T) {
	t.Run("When create an receipt types success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		receiptTypeId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		createdAt := now.Format("2006-01-02 15:04:05")
		userId := uuid.New().String()
		createReceiptTypesBody := receiptTypesDomain.CreateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}

		mock.
			ExpectExec(QueryCreateReceiptType).
			WithArgs(
				receiptTypeId,
				createReceiptTypesBody.Description,
				createReceiptTypesBody.SunatCode,
				createReceiptTypesBody.Enable,
				userId,
				createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		r := NewReceiptTypesRepository(clock, 60)
		err = r.CreateReceiptType(ctx, receiptTypeId, userId, createReceiptTypesBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When create an receipt types return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		clock := &mockClock.Clock{}
		now := time.Now().UTC()
		receiptTypeId := uuid.New().String()
		clock.On("Now").Return(now)

		createdAt := now.Format("2006-01-02 15:04:05")
		userId := uuid.New().String()
		createReceiptTypesBody := receiptTypesDomain.CreateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}

		mock.
			ExpectQuery(QueryCreateReceiptType).
			WithArgs(
				receiptTypeId,
				createReceiptTypesBody.Description,
				createReceiptTypesBody.SunatCode,
				createReceiptTypesBody.Enable,
				userId,
				createdAt).
			WillReturnError(expectedError)

		r := NewReceiptTypesRepository(clock, 60)
		err = r.CreateReceiptType(ctx, receiptTypeId, userId, createReceiptTypesBody)
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
		assert.Equal(t, smartErr.Function, "CreateReceiptType")
	})
}

func TestRepositoryReceiptTypes_UpdateReceiptType(t *testing.T) {
	t.Run("When update receipt types success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		receiptTypeId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		updateReceiptTypesBody := receiptTypesDomain.UpdateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}

		mock.
			ExpectExec(QueryUpdateReceiptType).
			WithArgs(
				updateReceiptTypesBody.Description,
				updateReceiptTypesBody.SunatCode,
				updateReceiptTypesBody.Enable,
				receiptTypeId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		r := NewReceiptTypesRepository(clock, 60)
		err = r.UpdateReceiptType(ctx, receiptTypeId, updateReceiptTypesBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}

		assert.Nil(t, err)
	})

	t.Run("When update receipt types return an error", func(t *testing.T) {
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
		receiptTypeId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		updateReceiptTypesBody := receiptTypesDomain.UpdateReceiptTypeBody{
			Description: "Recibo por Arrendamiento",
			SunatCode:   "02",
			Enable:      true,
		}
		mock.ExpectQuery(QueryUpdateReceiptType).
			WithArgs(
				updateReceiptTypesBody.Description,
				updateReceiptTypesBody.SunatCode,
				updateReceiptTypesBody.Enable,
				receiptTypeId).
			WillReturnError(expectedError)

		r := NewReceiptTypesRepository(clock, 60)
		err = r.UpdateReceiptType(ctx, receiptTypeId, updateReceiptTypesBody)
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
		assert.Equal(t, smartErr.Function, "UpdateReceiptType")
	})
}

func TestRepositoryReceiptTypes_DeleteReceiptType(t *testing.T) {
	t.Run("When delete receipt types successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		receiptTypeId := uuid.New().String()
		deletedAt := now.Format("2006-01-02 15:04:05")
		var res bool

		mock.
			ExpectExec(QueryDeleteReceiptType).
			WithArgs(
				deletedAt,
				receiptTypeId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		r := NewReceiptTypesRepository(clock, 60)
		res, err = r.DeleteReceiptType(ctx, receiptTypeId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}

		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete receipt types error", func(t *testing.T) {
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
		receiptTypeId := uuid.New().String()
		ReceiptTypesId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		deletedAt := now.Format("2006-01-02 15:04:05")
		var res bool
		clock.On("Now").Return(now)

		mock.ExpectExec(QueryDeleteReceiptType).
			WithArgs(deletedAt, ReceiptTypesId).
			WillReturnError(expectedError)

		r := NewReceiptTypesRepository(clock, 60)

		res, err = r.DeleteReceiptType(ctx, receiptTypeId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteReceiptType")
	})
}
