/*
 * File: role_policies_mysql_repository_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to rolePolicy repository.
 *
 * Last Modified: 2023-11-22
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

	rolePoliciesDomain "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain"
)

func TestRepositoryRolePolicies_GetPolicies(t *testing.T) {
	t.Run("When get policies by role is called then it should return a list of policies", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now()
		policy := rolePoliciesDomain.PolicyByRolePolicy{
			Id:          "fcdbfacf-8305-11ee-89fd-024255555501",
			Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
			Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
			Level:       "system",
			Enable:      true,
			CreatedAt:   &now,
		}

		mockRegister := []rolePoliciesDomain.RolePolicy{
			{
				Id:        "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Enable:    true,
				CreatedAt: &now,
				Policy:    policy,
			},
			{
				Id:        "739bbbc9-7e93-11ee-89fd-0242ac110019",
				Enable:    true,
				CreatedAt: &now,
				Policy:    policy,
			},
		}
		rows := sqlmock.NewRows([]string{
			"role_policy_id",
			"role_policy_enable",
			"role_policy_created_at",
			"policy_id",
			"policy_name",
			"policy_description",
			"policy_level",
			"policy_enable",
			"policy_created_at",
		}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].Enable,
				&now,
				mockRegister[0].Policy.Id,
				mockRegister[0].Policy.Name,
				mockRegister[0].Policy.Description,
				mockRegister[0].Policy.Level,
				mockRegister[0].Policy.Enable,
				&now,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].Enable,
				&now,
				mockRegister[1].Policy.Id,
				mockRegister[1].Policy.Name,
				mockRegister[1].Policy.Description,
				mockRegister[1].Policy.Level,
				mockRegister[1].Policy.Enable,
				&now,
			)

		roleId := "fcdbfacf-8305-11ee-89fd-024255555501"
		sizePage := 100
		offset := 0
		mock.
			ExpectQuery(QueryGetRolePolicies).
			WithArgs(roleId, sizePage, offset).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewRolePoliciesRepository(clock, 60)
		var res []rolePoliciesDomain.RolePolicy
		searchParams := rolePoliciesDomain.GetRolePoliciesParams{
			RoleId: roleId,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetPolicies(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get policies by role is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		roleId := "fcdbfacf-8305-11ee-89fd-024255555501"
		sizePage := 100
		offset := 0
		mock.ExpectQuery(QueryGetRolePolicies).
			WithArgs(roleId, sizePage, offset).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewRolePoliciesRepository(clock, 60)
		var res []rolePoliciesDomain.RolePolicy
		searchParams := rolePoliciesDomain.GetRolePoliciesParams{
			RoleId: roleId,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetPolicies(ctx, searchParams, pagination)
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
		assert.Equal(t, smartErr.Function, "GetPolicies")
	})
}

func TestRepositoryRolePolicies_GetTotalPolicies(t *testing.T) {
	t.Run("When get total of policies by role is called then it should return a list of policies", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		total := 2
		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)
		roleId := "fcdbfacf-8305-11ee-89fd-024255555501"
		mock.
			ExpectQuery(QueryGetTotalRolePolicies).
			WithArgs(roleId).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewRolePoliciesRepository(clock, 60)
		searchParams := rolePoliciesDomain.GetRolePoliciesParams{
			RoleId: roleId,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalPolicies(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total of policies by role is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		roleId := "fcdbfacf-8305-11ee-89fd-024255555501"
		mock.ExpectQuery(QueryGetTotalRolePolicies).
			WithArgs(roleId).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewRolePoliciesRepository(clock, 60)
		var res []rolePoliciesDomain.RolePolicy
		searchParams := rolePoliciesDomain.GetRolePoliciesParams{
			RoleId: roleId,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalPolicies(ctx, searchParams, pagination)
		if res != nil {
			t.Errorf("this is the error getting the registers: %v\n", res)
			return
		}
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalPolicies")
	})
}

func TestRepositoryRolePolicies_CreateRolePolicy(t *testing.T) {
	var enable = true
	t.Run("When add a policy to role success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		rolePolicyId := uuid.New().String()
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac219255"
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createRolePolicyBody := rolePoliciesDomain.CreateRolePolicyBody{
			PolicyId: policyId,
			Enable:   enable,
		}

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryCreateRolePolicy).
			WithArgs(
				rolePolicyId,
				policyId,
				roleId,
				enable,
				createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewRolePoliciesRepository(clock, 60)

		_, err = r.CreateRolePolicy(ctx, rolePolicyId, roleId, createRolePolicyBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When add a policy to role return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		rolePolicyId := uuid.New().String()
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac219255"
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		createRolePolicyBody := rolePoliciesDomain.CreateRolePolicyBody{
			PolicyId: policyId,
			Enable:   enable,
		}
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryCreateRolePolicy).
			WithArgs(rolePolicyId, policyId, roleId, enable, createdAt).
			WillReturnError(expectedError)
		r := NewRolePoliciesRepository(clock, 60)
		_, err = r.CreateRolePolicy(ctx, rolePolicyId, roleId, createRolePolicyBody)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)
	})
}

func TestRepositoryRolePolicies_CreateRolePolicies(t *testing.T) {
	var enable = true
	t.Run("When add multiple policies to role success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"

		createMultipleRolePolicyBody := []rolePoliciesDomain.CreateMultipleRolePolicyBody{
			{
				Id: uuid.New().String(),
				CreateRolePolicyBody: rolePoliciesDomain.CreateRolePolicyBody{
					PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
					Enable:   enable,
				},
			},
			{
				Id: uuid.New().String(),
				CreateRolePolicyBody: rolePoliciesDomain.CreateRolePolicyBody{
					PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278421",
					Enable:   enable,
				},
			},
		}
		clock := &mockClock.Clock{}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock.On("Now").Return(now)

		// expect transaction begin
		mock.ExpectBegin()

		for _, rolePolicy := range createMultipleRolePolicyBody {
			mock.ExpectExec(QueryCreateRolePolicy).
				WithArgs(
					rolePolicy.Id,
					rolePolicy.PolicyId,
					roleId,
					rolePolicy.Enable,
					createdAt,
				).WillReturnResult(sqlmock.NewResult(1, 1))
		}
		r := NewRolePoliciesRepository(clock, 60)

		// expect a transaction commit
		mock.ExpectCommit()

		err = r.CreateRolePolicies(ctx, roleId, createMultipleRolePolicyBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("When add multiple policies to role return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		createMultipleRolePolicyBody := []rolePoliciesDomain.CreateMultipleRolePolicyBody{
			{
				Id: uuid.New().String(),
				CreateRolePolicyBody: rolePoliciesDomain.CreateRolePolicyBody{
					PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
					Enable:   enable,
				},
			},
			{
				Id: uuid.New().String(),
				CreateRolePolicyBody: rolePoliciesDomain.CreateRolePolicyBody{
					PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278421",
					Enable:   enable,
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

		for _, rolePolicy := range createMultipleRolePolicyBody {
			mock.ExpectQuery(QueryCreateRolePolicy).
				WithArgs(
					rolePolicy.Id,
					rolePolicy.PolicyId,
					roleId,
					rolePolicy.Enable,
					createdAt,
				).WillReturnError(expectedError)
		}
		r := NewRolePoliciesRepository(clock, 60)

		// expect a transaction commit
		mock.ExpectCommit()

		err = r.CreateRolePolicies(ctx, roleId, createMultipleRolePolicyBody)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)
	})
}

func TestRepositoryRolePolicies_UpdateRolePolicy(t *testing.T) {
	var enable = true
	t.Run("When update a policy of role successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac219255"
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		updateRolePolicyBody := rolePoliciesDomain.UpdateRolePolicyBody{
			PolicyId: policyId,
			Enable:   enable,
		}
		clock := &mockClock.Clock{}
		mock.ExpectExec(QueryUpdateRolePolicy).
			WithArgs(roleId, policyId, enable, rolePolicyId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewRolePoliciesRepository(clock, 60)

		err = r.UpdateRolePolicy(ctx, roleId, rolePolicyId, updateRolePolicyBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When update a policy of role return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac219255"
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		updateRolePolicyBody := rolePoliciesDomain.UpdateRolePolicyBody{
			PolicyId: policyId,
			Enable:   enable,
		}
		expectedError := errors.New("random error")
		clock := &mockClock.Clock{}
		mock.ExpectQuery(QueryUpdateRolePolicy).
			WithArgs(roleId, policyId, enable, rolePolicyId).
			WillReturnError(expectedError)
		r := NewRolePoliciesRepository(clock, 60)
		err = r.UpdateRolePolicy(ctx, roleId, rolePolicyId, updateRolePolicyBody)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)
	})
}

func TestRepositoryRolePolicies_DeleteRolePolicy(t *testing.T) {
	t.Run("When delete a policy of role successfully", func(t *testing.T) {
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
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteRolePolicy).
			WithArgs(deletedAt, rolePolicyId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewRolePoliciesRepository(clock, 60)
		var res bool
		res, err = r.DeleteRolePolicy(ctx, roleId, rolePolicyId)
		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete a policy of role error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectExec(QueryDeleteRolePolicy).
			WithArgs(deletedAt, rolePolicyId).
			WillReturnError(errors.New("anything"))
		r := NewRolePoliciesRepository(clock, 60)
		var res bool
		res, err = r.DeleteRolePolicy(ctx, roleId, rolePolicyId)

		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}

func TestRepositoryRolePolicies_DeleteRolePolicies(t *testing.T) {
	t.Run("When delete multiple policies of role successfully", func(t *testing.T) {
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
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePoliciesIds := []string{
			uuid.New().String(),
			uuid.New().String(),
		}

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		// expect transaction begin
		mock.ExpectBegin()

		for _, rolePolicyId := range rolePoliciesIds {
			mock.ExpectExec(QueryDeleteRolePolicy).
				WithArgs(deletedAt, rolePolicyId).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
		r := NewRolePoliciesRepository(clock, 60)

		// expect a transaction commit
		mock.ExpectCommit()

		err = r.DeleteRolePolicies(ctx, roleId, rolePoliciesIds)
		assert.NoError(t, err)
		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("When delete multiple policies of role error", func(t *testing.T) {
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
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyIds := []string{
			uuid.New().String(),
			uuid.New().String(),
		}

		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		// expect transaction begin
		mock.ExpectBegin()

		for _, rolePolicyId := range rolePolicyIds {
			mock.ExpectExec(QueryDeleteRolePolicy).
				WithArgs(deletedAt, rolePolicyId).
				WillReturnError(errors.New("anything"))
		}
		r := NewRolePoliciesRepository(clock, 60)

		// expect a transaction commit
		mock.ExpectRollback()

		err = r.DeleteRolePolicies(ctx, roleId, rolePolicyIds)

		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteRolePolicies")
	})
}
