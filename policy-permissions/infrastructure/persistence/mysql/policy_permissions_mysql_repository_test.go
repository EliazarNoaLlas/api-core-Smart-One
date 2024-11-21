/*
 * File: policyPermissions_mysql_repository_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to policyPermission repository.
 *
 * Last Modified: 2023-11-20
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

	policyPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
)

func TestRepositoryPolicyPermissions_GetPolicyPermissionByPolicy(t *testing.T) {
	t.Run("When get policy permissions is called then it should return a list of policy permission", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now()
		permission := policyPermissionsDomain.Permission{
			Id:          "84305ba9-83d2-11ee-89fd-0242ac110016",
			Code:        "REQUIREMENTS_READ",
			Name:        "Listar requerimientos",
			Description: "Permiso para listar requerimientos",
			CreatedAt:   &now,
		}
		mockPolicyPermissionType := []policyPermissionsDomain.PolicyPermission{
			{
				Id:         "22597e1d-6463-4bf9-ba51-0f8a3967321f",
				Enable:     1,
				CreatedAt:  &now,
				Permission: permission,
			},
			{
				Id:         "22597e1d-6463-4bf9-ba51-0f8a3967321g",
				Enable:     1,
				CreatedAt:  &now,
				Permission: permission,
			},
		}

		rows := sqlmock.NewRows([]string{
			"policy_permission_id",
			"policy_permission_enable",
			"policy_permission_created_at",
			"permissions_id",
			"permissions_code",
			"permissions_name",
			"permissions_description",
			"permissions_created_at",
		}).
			AddRow(
				mockPolicyPermissionType[0].Id,
				mockPolicyPermissionType[0].Enable,
				mockPolicyPermissionType[0].CreatedAt,
				mockPolicyPermissionType[0].Permission.Id,
				mockPolicyPermissionType[0].Permission.Code,
				mockPolicyPermissionType[0].Permission.Name,
				mockPolicyPermissionType[0].Permission.Description,
				mockPolicyPermissionType[0].Permission.CreatedAt,
			).
			AddRow(
				mockPolicyPermissionType[1].Id,
				mockPolicyPermissionType[1].Enable,
				mockPolicyPermissionType[1].CreatedAt,
				mockPolicyPermissionType[1].Permission.Id,
				mockPolicyPermissionType[1].Permission.Code,
				mockPolicyPermissionType[1].Permission.Name,
				mockPolicyPermissionType[1].Permission.Description,
				mockPolicyPermissionType[1].Permission.CreatedAt,
			)

		sizePage := 100
		offset := 0
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110019"
		mock.
			ExpectQuery(QueryGetPermissionsByPolicy).
			WithArgs(policyId, sizePage, offset).WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewPolicyPermissionsRepository(clock, 60)
		var res []policyPermissionsDomain.PolicyPermission
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetPolicyPermissionsByPolicy(ctx, policyId, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get policy permissions is called then it should return an error", func(t *testing.T) {
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
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110019"
		mock.ExpectQuery(QueryGetPermissionsByPolicy).
			WithArgs(policyId,
				sizePage, offset).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewPolicyPermissionsRepository(clock, 60)
		var res []policyPermissionsDomain.PolicyPermission
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetPolicyPermissionsByPolicy(ctx, policyId, pagination)
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
		assert.Equal(t, smartErr.Function, "GetPolicyPermissionsByPolicy")
	})
}

func TestRepositoryPolicyPermissions_GetTotalPermissionByPolicy(t *testing.T) {
	t.Run("When get total of policy permissions is called then it should return a total", func(t *testing.T) {
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
			ExpectQuery(QueryTotalPermissionByPolicy).
			WithArgs().
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewPolicyPermissionsRepository(clock, 60)

		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110019"
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalPolicyPermissionsByPolicy(ctx, policyId, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total of policy permissions is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		mock.ExpectQuery(QueryTotalPermissionByPolicy).
			WithArgs("739bbbc9-7e93-11ee-89fd-0242ac113421", "739bbbc9-7e93-11ee-89fd-0242ac113421").
			WillReturnError(expectedError)
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110019"

		clock := &mockClock.Clock{}
		r := NewPolicyPermissionsRepository(clock, 60)

		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalPolicyPermissionsByPolicy(ctx, policyId, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalPermissionByPolicy")
	})
}

func TestRepositoryPolicyPermissions_CreatePolicyPermission(t *testing.T) {
	t.Run("When add a permission to policy success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		policyId := uuid.New().String()
		policyPermissionId := uuid.New().String()
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createPolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryCreatePolicyPermission).
			WithArgs(
				policyPermissionId,
				policyId,
				createPolicyPermissionBody.PermissionId,
				createPolicyPermissionBody.Enable,
				createdAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewPolicyPermissionsRepository(clock, 60)

		_, err = r.CreatePolicyPermission(
			ctx,
			policyId,
			policyPermissionId,
			createPolicyPermissionBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When add a permission to policy return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		policyPermissionId := uuid.New().String()
		policyId := uuid.New().String()
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createPolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreatePolicyPermission).
			WithArgs(
				policyPermissionId,
				policyId,
				createPolicyPermissionBody.PermissionId,
				createPolicyPermissionBody.Enable,
				createdAt,
			).WillReturnError(expectedError)
		r := NewPolicyPermissionsRepository(clock, 60)
		_, err = r.CreatePolicyPermission(
			ctx,
			policyId,
			policyPermissionId,
			createPolicyPermissionBody)
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
		assert.Equal(t, smartErr.Function, "CreatePolicyPermission")
	})
}

func TestPolicyPermissionsMySQLRepo_CreatePolicyPermissions(t *testing.T) {
	t.Run("When add multiple permissions to policy success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		policyId := uuid.New().String()
		createPolicyPermissionsBody := []policyPermissionsDomain.CreatePolicyPermissionMultipleBody{
			{
				Id: uuid.New().String(),
				CreatePolicyPermissionBody: policyPermissionsDomain.CreatePolicyPermissionBody{
					PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
					Enable:       true,
				},
			},
			{
				Id: uuid.New().String(),
				CreatePolicyPermissionBody: policyPermissionsDomain.CreatePolicyPermissionBody{
					PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278421",
					Enable:       true,
				},
			},
		}
		clock := &mockClock.Clock{}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock.On("Now").Return(now)

		// expect transaction begin
		mock.ExpectBegin()

		for _, policyPermission := range createPolicyPermissionsBody {
			mock.ExpectExec(QueryCreatePolicyPermission).
				WithArgs(
					policyPermission.Id,
					policyId,
					policyPermission.PermissionId,
					policyPermission.Enable,
					createdAt,
				).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
		r := NewPolicyPermissionsRepository(clock, 60)

		// expect a transaction commit
		mock.ExpectCommit()

		err = r.CreatePolicyPermissions(
			ctx,
			policyId,
			createPolicyPermissionsBody,
		)
		if err != nil {
			t.Errorf("this is the error creating the registers: %v\n", err)
			return
		}

		assert.Nil(t, err)

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("When add multiple permissions to policy return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		policyId := uuid.New().String()
		createPolicyPermissionsBody := []policyPermissionsDomain.CreatePolicyPermissionMultipleBody{
			{
				Id: uuid.New().String(),
				CreatePolicyPermissionBody: policyPermissionsDomain.CreatePolicyPermissionBody{
					PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
					Enable:       true,
				},
			},
			{
				Id: uuid.New().String(),
				CreatePolicyPermissionBody: policyPermissionsDomain.CreatePolicyPermissionBody{
					PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278421",
					Enable:       true,
				},
			},
		}
		clock := &mockClock.Clock{}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock.On("Now").Return(now)
		expectedError := errors.New("any error")

		// expect transaction begin
		mock.ExpectBegin()

		for _, policyPermission := range createPolicyPermissionsBody {
			mock.ExpectExec(QueryCreatePolicyPermission).
				WithArgs(
					policyPermission.Id,
					policyId,
					policyPermission.PermissionId,
					policyPermission.Enable,
					createdAt,
				).
				WillReturnError(expectedError)
		}

		// expect a transaction commit
		mock.ExpectRollback()

		r := NewPolicyPermissionsRepository(clock, 60)
		err = r.CreatePolicyPermissions(
			ctx,
			policyId,
			createPolicyPermissionsBody,
		)
		if err == nil {
			t.Errorf("this is the error creating the registers: %v\n", err)
			return
		}
		assert.Error(t, err)
	})
}

func TestRepositoryPolicyPermissions_UpdatePolicyPermission(t *testing.T) {
	t.Run("When update a permission of policy success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		policyPermissionId := uuid.New().String()
		policyId := uuid.New().String()
		updatePolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		clock := &mockClock.Clock{}
		mock.ExpectExec(QueryUpdatePolicyPermission).
			WithArgs(
				policyId,
				updatePolicyPermissionBody.PermissionId,
				updatePolicyPermissionBody.Enable,
				policyPermissionId,
			).WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewPolicyPermissionsRepository(clock, 60)
		err = r.UpdatePolicyPermission(ctx, policyId, policyPermissionId, updatePolicyPermissionBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When update a permission of policy return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		policyPermissionId := uuid.New().String()
		policyId := uuid.New().String()
		updatePolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		clock := &mockClock.Clock{}
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdatePolicyPermission).
			WithArgs(
				policyId,
				updatePolicyPermissionBody.PermissionId,
				updatePolicyPermissionBody.Enable,
				policyPermissionId,
			).WillReturnError(expectedError)
		r := NewPolicyPermissionsRepository(clock, 60)
		err = r.UpdatePolicyPermission(ctx, policyId, policyPermissionId, updatePolicyPermissionBody)
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
		assert.Equal(t, smartErr.Function, "UpdatePolicyPermission")
	})
}

func TestRepositoryPolicyPermissions_DeletePolicyPermission(t *testing.T) {
	t.Run("When delete a permission of policy successfully", func(t *testing.T) {
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
		policyPermissionId := uuid.New().String()
		policyId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeletePolicyPermission).
			WithArgs(deletedAt, policyPermissionId, policyId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewPolicyPermissionsRepository(clock, 60)
		var res bool
		res, err = r.DeletePolicyPermission(ctx, policyId, policyPermissionId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete a permission of policy error", func(t *testing.T) {
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
		policyPermissionId := uuid.New().String()
		policyId := uuid.New().String()
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeletePolicyPermission).
			WithArgs(deletedAt, policyPermissionId, policyId).
			WillReturnError(errors.New("anything"))
		r := NewPolicyPermissionsRepository(clock, 60)
		var res bool
		res, err = r.DeletePolicyPermission(ctx, policyId, policyPermissionId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeletePolicyPermission")
	})
}

func TestRepositoryPolicyPermissions_DeletePolicyPermissions(t *testing.T) {
	t.Run("When delete multiple permissions of policy successfully", func(t *testing.T) {
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
		policyId := uuid.New().String()
		policyPermissionIds := []string{
			uuid.New().String(),
			uuid.New().String(),
		}

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		// expect transaction begin
		mock.ExpectBegin()

		for _, policyPermissionId := range policyPermissionIds {
			mock.ExpectExec(QueryDeletePolicyPermission).
				WithArgs(deletedAt, policyPermissionId, policyId).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
		r := NewPolicyPermissionsRepository(clock, 60)

		// expect a transaction commit
		mock.ExpectCommit()

		err = r.DeletePolicyPermissions(ctx, policyId, policyPermissionIds)

		assert.NoError(t, err)
		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("When delete multiple permissions of policy error", func(t *testing.T) {
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
		policyId := uuid.New().String()
		policyPermissionIds := []string{
			uuid.New().String(),
			uuid.New().String(),
		}

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		// expect transaction begin
		mock.ExpectBegin()

		for _, policyPermissionId := range policyPermissionIds {
			mock.ExpectExec(QueryDeletePolicyPermission).
				WithArgs(deletedAt, policyPermissionId).
				WillReturnError(errors.New("anything"))
		}

		r := NewPolicyPermissionsRepository(clock, 60)

		// expect a transaction commit
		mock.ExpectRollback()

		err = r.DeletePolicyPermissions(ctx, policyId, policyPermissionIds)

		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeletePolicyPermissions")
	})
}
