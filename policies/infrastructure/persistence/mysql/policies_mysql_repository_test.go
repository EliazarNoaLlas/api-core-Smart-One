/*
 * File: policies_mysql_repository_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to policy repository.
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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	mockClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock/mocks"
	db2 "gitlab.smartcitiesperu.com/smartone/api-shared/db"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	policiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain"
)

func pointerToStr(value string) *string {
	return &value
}

func pointerToBool(value bool) *bool {
	return &value
}

func TestRepositoryPolicies_GetPolicies(t *testing.T) {
	t.Run("When get policies is called then it should return a list of policies", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		module := policiesDomain.ModuleByPolicy{
			Id:          pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110018"),
			Name:        pointerToStr("Logistic"),
			Description: pointerToStr("Modulo de logística"),
			Code:        pointerToStr("logistic"),
		}
		merchant := policiesDomain.MerchantByPolicy{
			Id:          pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110016"),
			Name:        pointerToStr("Odin Corp"),
			Description: pointerToStr("Proveedor de servicios de mantenimiento"),
			Document:    pointerToStr("123456789"),
		}
		store := policiesDomain.StoreByPolicy{
			Id:        pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110016"),
			Name:      pointerToStr("Obra av. 28 julio"),
			Shortname: pointerToStr("Obra 28"),
		}
		policyPermission := policiesDomain.PolicyPermissionByPolicy{
			Id:     pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110010"),
			Enable: pointerToBool(false),
		}
		permission := policiesDomain.PermissionByPolicy{
			Id:               pointerToStr("fcdbfacf-8305-11ee-89fd-024255555501"),
			Code:             pointerToStr("REQUIREMENTS_READ"),
			Name:             pointerToStr("Listar requerimientos"),
			Description:      pointerToStr("Permiso para listar requerimientos"),
			PolicyPermission: &policyPermission,
		}

		mockRegister := []policiesDomain.Policy{
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
				Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
				Level:       "system",
				Enable:      true,
				CreatedAt:   &now,
				Module:      &module,
				Merchant:    &merchant,
				Store:       &store,
				Permissions: []policiesDomain.PermissionByPolicy{permission},
			},
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110019",
				Name:        "LOGISTICA_REQUERIMIENTOS_OBRA_28",
				Description: "Politica para accesos a logistica requerimientos en la obra 28 de julio",
				Level:       "system",
				Enable:      true,
				CreatedAt:   &now,
				Module:      &module,
				Merchant:    &merchant,
				Store:       &store,
				Permissions: []policiesDomain.PermissionByPolicy{permission},
			},
		}
		rows := sqlmock.NewRows([]string{"policy_id", "policy_name", "policy_description", "policy_level",
			"policy_enable", "policy_created_at", "module_id", "module_name", "module_description", "module_code",
			"merchant_id", "merchant_name", "merchant_description", "merchant_document", "store_id", "store_name",
			"store_shortname", "permission_id", "permission_code", "permission_name", "permission_description",
			"policy_permission_id", "policy_permission_enable"}).
			AddRow(
				mockRegister[0].Id,
				mockRegister[0].Name,
				mockRegister[0].Description,
				mockRegister[0].Level,
				mockRegister[0].Enable,
				mockRegister[0].CreatedAt,
				mockRegister[0].Module.Id,
				mockRegister[0].Module.Name,
				mockRegister[0].Module.Description,
				mockRegister[0].Module.Code,
				mockRegister[0].Merchant.Id,
				mockRegister[0].Merchant.Name,
				mockRegister[0].Merchant.Description,
				mockRegister[0].Merchant.Document,
				mockRegister[0].Store.Id,
				mockRegister[0].Store.Name,
				mockRegister[0].Store.Shortname,
				mockRegister[0].Permissions[0].Id,
				mockRegister[0].Permissions[0].Code,
				mockRegister[0].Permissions[0].Name,
				mockRegister[0].Permissions[0].Description,
				mockRegister[0].Permissions[0].PolicyPermission.Id,
				mockRegister[0].Permissions[0].PolicyPermission.Enable,
			).
			AddRow(
				mockRegister[1].Id,
				mockRegister[1].Name,
				mockRegister[1].Description,
				mockRegister[1].Level,
				mockRegister[1].Enable,
				mockRegister[1].CreatedAt,
				mockRegister[1].Module.Id,
				mockRegister[1].Module.Name,
				mockRegister[1].Module.Description,
				mockRegister[1].Module.Code,
				mockRegister[1].Merchant.Id,
				mockRegister[1].Merchant.Name,
				mockRegister[1].Merchant.Description,
				mockRegister[1].Merchant.Document,
				mockRegister[1].Store.Id,
				mockRegister[1].Store.Name,
				mockRegister[1].Store.Shortname,
				mockRegister[1].Permissions[0].Id,
				mockRegister[1].Permissions[0].Code,
				mockRegister[1].Permissions[0].Name,
				mockRegister[1].Permissions[0].Description,
				mockRegister[1].Permissions[0].PolicyPermission.Id,
				mockRegister[1].Permissions[0].PolicyPermission.Enable,
			)

		var moduleId *string = nil
		var merchantId *string = nil
		var storeId *string = nil
		var description *string = nil
		sizePage := 100
		offset := 0
		mock.
			ExpectQuery(QueryGetPolicies).
			WithArgs(moduleId, moduleId, merchantId, merchantId, storeId, storeId, description, description, description,
				sizePage, offset).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewPoliciesRepository(clock, 60)
		var res []policiesDomain.Policy
		searchParams := policiesDomain.GetPoliciesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetPolicies(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
	})

	t.Run("When get policies is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		var moduleId *string = nil
		var merchantId *string = nil
		var storeId *string = nil
		var description *string = nil
		sizePage := 100
		offset := 0
		mock.ExpectQuery(QueryGetPolicies).
			WithArgs(moduleId, moduleId, merchantId, merchantId, storeId, storeId, description, description,
				description, sizePage, offset).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewPoliciesRepository(clock, 60)
		var res []policiesDomain.Policy
		searchParams := policiesDomain.GetPoliciesParams{}
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

func TestRepositoryPolicies_GetTotalPolicies(t *testing.T) {
	t.Run("When get total of policies is called then it should return a total", func(t *testing.T) {
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

		var moduleId *string = nil
		var merchantId *string = nil
		var storeId *string = nil
		var description *string = nil
		mock.
			ExpectQuery(QueryGetTotalPolicies).
			WithArgs(moduleId, moduleId, merchantId, merchantId, storeId, storeId, description, description, description).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewPoliciesRepository(clock, 60)

		searchParams := policiesDomain.GetPoliciesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalPolicies(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total of policies is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		var moduleId *string = nil
		var merchantId *string = nil
		var storeId *string = nil
		var description *string = nil
		mock.ExpectQuery(QueryGetTotalPolicies).
			WithArgs(moduleId, moduleId, merchantId, merchantId, storeId, storeId, description, description, description).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewPoliciesRepository(clock, 60)

		searchParams := policiesDomain.GetPoliciesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalPolicies(ctx, searchParams, pagination)
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

func TestRepositoryPolicies_CreatePolicy(t *testing.T) {
	t.Run("When to successfully create a policy", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		createPolicyBody := policiesDomain.CreatePolicyBody{
			Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
			Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
			ModuleId:    "739bbbc9-7e93-11ee-89fd-0242ac110018",
			MerchantId:  pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110019"),
			StoreId:     pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110020"),
			Level:       "system",
			Enable:      pointerToBool(true),
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		policyId := uuid.New().String()
		mock.ExpectExec(QueryCreatePolicy).
			WithArgs(
				policyId,
				createPolicyBody.Name,
				createPolicyBody.Description,
				createPolicyBody.ModuleId,
				createPolicyBody.MerchantId,
				createPolicyBody.StoreId,
				createPolicyBody.Level,
				createPolicyBody.Enable,
				createdAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		r := NewPoliciesRepository(clock, 60)
		_, err = r.CreatePolicy(ctx, createPolicyBody, policyId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When an error occurs while creating a policy", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		createPolicyBody := policiesDomain.CreatePolicyBody{
			Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
			Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
			ModuleId:    "739bbbc9-7e93-11ee-89fd-0242ac110018",
			MerchantId:  pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110019"),
			StoreId:     pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110020"),
			Level:       "system",
			Enable:      pointerToBool(true),
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		policyId := uuid.New().String()
		mock.ExpectQuery(QueryCreatePolicy).
			WithArgs(
				policyId,
				createPolicyBody.Name,
				createPolicyBody.Description,
				createPolicyBody.ModuleId,
				createPolicyBody.MerchantId,
				createPolicyBody.StoreId,
				createPolicyBody.Level,
				createPolicyBody.Enable,
				createdAt).
			WillReturnError(expectedError)
		r := NewPoliciesRepository(clock, 60)
		_, err = r.CreatePolicy(ctx, createPolicyBody, policyId)
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
		assert.Equal(t, smartErr.Function, "CreatePolicy")
	})
}

func TestRepositoryPolicies_UpdatePolicy(t *testing.T) {
	t.Run("When a policy is successfully updated", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		updatePolicyBody := policiesDomain.UpdatePolicyBody{
			Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
			Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
			ModuleId:    "739bbbc9-7e93-11ee-89fd-0242ac110018",
			MerchantId:  pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110019"),
			StoreId:     pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110020"),
			Level:       "system",
			Enable:      pointerToBool(true),
		}
		clock := &mockClock.Clock{}
		policyId := uuid.New().String()
		mock.ExpectExec(QueryUpdatePolicy).
			WithArgs(
				updatePolicyBody.Name,
				updatePolicyBody.Description,
				updatePolicyBody.ModuleId,
				updatePolicyBody.MerchantId,
				updatePolicyBody.StoreId,
				updatePolicyBody.Level,
				updatePolicyBody.Enable,
				policyId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewPoliciesRepository(clock, 60)
		err = r.UpdatePolicy(ctx, updatePolicyBody, policyId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When an error occurs while updating a policy", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		updatePolicyBody := policiesDomain.UpdatePolicyBody{
			Name:        "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO",
			Description: "Politica para accesos a logistica requerimientos en todo el conglomerado",
			ModuleId:    "739bbbc9-7e93-11ee-89fd-0242ac110018",
			MerchantId:  pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110019"),
			StoreId:     pointerToStr("739bbbc9-7e93-11ee-89fd-0242ac110020"),
			Level:       "system",
			Enable:      pointerToBool(true),
		}
		clock := &mockClock.Clock{}
		policyId := uuid.New().String()
		mock.ExpectQuery(QueryUpdatePolicy).
			WithArgs(
				updatePolicyBody.Name,
				updatePolicyBody.Description,
				updatePolicyBody.ModuleId,
				updatePolicyBody.MerchantId,
				updatePolicyBody.StoreId,
				updatePolicyBody.Level,
				updatePolicyBody.Enable,
				policyId).
			WillReturnError(expectedError)
		r := NewPoliciesRepository(clock, 60)
		err = r.UpdatePolicy(ctx, updatePolicyBody, policyId)
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
		assert.Equal(t, smartErr.Function, "UpdatePolicy")
	})
}

func TestRepositoryPolicies_DeletePolicy(t *testing.T) {
	t.Run("When a policy is successfully deleted", func(t *testing.T) {
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
		policyId := uuid.New().String()
		mock.ExpectExec(QueryDeletePolicy).
			WithArgs(deletedAt, policyId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewPoliciesRepository(clock, 60)
		var res bool
		res, err = r.DeletePolicy(ctx, policyId)
		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When an error occurs while deleting a policy", func(t *testing.T) {
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
		policyId := uuid.New().String()
		mock.ExpectExec(QueryDeletePolicy).
			WithArgs(deletedAt, policyId).
			WillReturnError(errors.New("anything"))
		r := NewPoliciesRepository(clock, 60)
		var res bool
		res, err = r.DeletePolicy(ctx, policyId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeletePolicy")
	})
}
