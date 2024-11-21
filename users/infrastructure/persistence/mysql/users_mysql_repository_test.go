/*
 * File: users_mysql_repository_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to user repository.
 *
 * Last Modified: 2023-11-23
 */

package mysql

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	mockClock "gitlab.smartcitiesperu.com/smartone/api-shared/clock/mocks"
	db2 "gitlab.smartcitiesperu.com/smartone/api-shared/db"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	usersDomain "gitlab.smartcitiesperu.com/smartone/api-core/users/domain"
)

func StringToPtr(value string) *string {
	return &value
}

func BoolToPtr(value bool) *bool {
	return &value
}

func TestRepositoryUsers_GetUser(t *testing.T) {
	t.Run("When get from user is called it should return a user ", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		userType := usersDomain.UserTypeByUser{
			Id:          "739bbbc9-7e93-11ee-89fd-0242ac110018",
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
		}

		mockUser := usersDomain.User{
			Id:        "739bbbc9-7e93-11ee-89fd-0242ac110016",
			UserName:  "pepito.quispe@smart.pe",
			CreatedAt: &now,
			UserType:  userType,
		}

		rows := sqlmock.NewRows([]string{"user_id", "user_name", "user_created_at",
			"user_type_id", "user_type_description", "user_type_code"}).
			AddRow(
				mockUser.Id,
				mockUser.UserName,
				mockUser.CreatedAt,
				mockUser.UserType.Id,
				mockUser.UserType.Description,
				mockUser.UserType.Code,
			)
		mock.ExpectQuery(QueryGetUser).WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		res, err := r.GetUser(ctx, "739bbbc9-7e93-11ee-89fd-0242ac110018")
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, res.Id, mockUser.Id)
		assert.Equal(t, res.UserName, mockUser.UserName)
	})

	t.Run("When the user was not found, error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		rows := sqlmock.NewRows([]string{"user_id", "user_name", "user_created_at",
			"user_type_id", "user_type_description", "user_type_code"})
		mock.ExpectQuery(QueryGetUser).WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		res, err := r.GetUser(ctx, "739bbbc9-7e93-11ee-89fd-0242ac110018")
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, usersDomain.ErrUserNotFoundCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetUser")
	})

	t.Run("When get of user is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New(errDomain.ErrUnknownCode)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110018"
		mock.ExpectQuery(QueryGetUser).
			WithArgs(userId).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)

		_, err = r.GetUser(ctx, userId)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetUser")
	})
}

func TestRepositoryUsers_GetUsers(t *testing.T) {
	t.Run("When get users is called then it should return a list of users", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		sizePage := 100
		offset := 0
		now := time.Now().UTC()
		searchParams := usersDomain.GetUsersParams{
			UserTypeId: StringToPtr("739bbbc9-7e93-11ee-89fd-0242ac110018"),
			UserName:   StringToPtr("pepito.quispe@smart.pe"),
			RoleId:     nil,
		}
		rolesIds := strings.Join(searchParams.RoleId, ",")
		userType := usersDomain.UserTypeByUser{
			Id:          "739bbbc9-7e93-11ee-89fd-0242ac110018",
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
		}

		userRoles := usersDomain.UserRole{
			Id: StringToPtr("b36f266d-8492-4f0e-8ecb-fef20e098970"),
		}

		role := []usersDomain.Role{
			{
				Id:          StringToPtr("fcdbfacf-8305-11ee-89fd-0242ac110016"),
				Name:        StringToPtr("Jefe de Area Residual"),
				Description: StringToPtr("Gerencia del conglomerado"),
				Enable:      BoolToPtr(true),
				CreatedAt:   &now,
				UserRole:    userRoles,
			},
		}

		mockUser := usersDomain.UserMultiple{
			Id:        "739bbbc9-7e93-11ee-89fd-0242ac110016",
			UserName:  "pepito.quispe@smart.pe",
			CreatedAt: &now,
			UserType:  userType,
			Role:      role,
		}

		rows := sqlmock.NewRows([]string{
			"user_id",
			"user_name",
			"user_created_at",
			"user_type_id",
			"user_type_description",
			"user_type_code",
			"user_role_id",
			"role_id",
			"role_name",
			"role_description",
			"role_enable",
			"role_created_at",
		}).
			AddRow(
				mockUser.Id,
				mockUser.UserName,
				mockUser.CreatedAt,
				mockUser.UserType.Id,
				mockUser.UserType.Description,
				mockUser.UserType.Code,
				mockUser.Role[0].UserRole.Id,
				mockUser.Role[0].Id,
				mockUser.Role[0].Name,
				mockUser.Role[0].Description,
				mockUser.Role[0].Enable,
				mockUser.Role[0].CreatedAt,
			)
		mock.ExpectQuery(QueryGetUsers).WithArgs(
			"739bbbc9-7e93-11ee-89fd-0242ac110018",
			"739bbbc9-7e93-11ee-89fd-0242ac110018",
			"pepito.quispe@smart.pe",
			"pepito.quispe@smart.pe",
			rolesIds,
			rolesIds,
			sizePage,
			offset).WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		var res []usersDomain.UserMultiple

		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetUsers(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 1)
	})

	t.Run("When get users is called then it should return an error", func(t *testing.T) {
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
		searchParams := usersDomain.GetUsersParams{
			UserTypeId: StringToPtr("739bbbc9-7e93-11ee-89fd-0242ac110018"),
			UserName:   StringToPtr("pepito.quispe@smart.pe"),
			RoleId:     nil,
		}
		rolesIds := strings.Join(searchParams.RoleId, ",")

		mock.ExpectQuery(QueryGetUsers).WithArgs(
			"739bbbc9-7e93-11ee-89fd-0242ac110018",
			"739bbbc9-7e93-11ee-89fd-0242ac110018",
			"pepito.quispe@smart.pe",
			"pepito.quispe@smart.pe",
			rolesIds,
			rolesIds,
			sizePage,
			offset,
		).WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		var res []usersDomain.UserMultiple
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetUsers(ctx, searchParams, pagination)
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
		assert.Equal(t, smartErr.Function, "GetUsers")
	})
}

func TestRepositoryUsers_GetTotalUsers(t *testing.T) {
	t.Run("When get total of users is called then it should return a total", func(t *testing.T) {
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
			ExpectQuery(QueryGetTotalUsers).
			WithArgs().
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)

		searchParams := usersDomain.GetUsersParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalUsers(ctx, searchParams, pagination)
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

		mock.ExpectQuery(QueryGetTotalUsers).
			WithArgs("739bbbc9-7e93-11ee-89fd-0242ac113421", "739bbbc9-7e93-11ee-89fd-0242ac113421").
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		searchParams := usersDomain.GetUsersParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalUsers(ctx, searchParams, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalUsers")
	})
}

func TestRepositoryUsers_GetMenu(t *testing.T) {
	t.Run("When get menu of user is called successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		mockModules := []usersDomain.ModuleMenuUser{
			{
				Id:          "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Name:        "Logistic",
				Description: "Modulo de logística",
				Code:        "logistic",
				Icon:        "fa fa-chart",
				Position:    1,
				CreatedAt:   &now,
				Views: []usersDomain.ViewMenuUser{
					{
						Id:          "739bbbc9-7e93-11ee-89fd-0242ac110000",
						Name:        "Requerimientos",
						Description: "Vista de requerimientos",
						Url:         "/logistics/requirements",
						Icon:        "fa fa-chart",
						CreatedAt:   &now,
					},
					{
						Id:          "739bbbc9-7e93-11ee-89fd-0242ac110001",
						Name:        "Almacenes",
						Description: "Vista de almacenes",
						Url:         "/warehouse/warehouses",
						Icon:        "fa fa-chart",
						CreatedAt:   &now,
					},
				},
			},
		}

		rows := sqlmock.NewRows([]string{
			"module_id",
			"module_name",
			"module_description",
			"module_code",
			"module_icon",
			"module_position",
			"module_created_at",
			"view_id",
			"view_name",
			"view_description",
			"view_url",
			"view_icon",
			"view_created_at",
		}).
			AddRow(
				mockModules[0].Id,
				mockModules[0].Name,
				mockModules[0].Description,
				mockModules[0].Code,
				mockModules[0].Icon,
				mockModules[0].Position,
				mockModules[0].CreatedAt,
				mockModules[0].Views[0].Id,
				mockModules[0].Views[0].Name,
				mockModules[0].Views[0].Description,
				mockModules[0].Views[0].Url,
				mockModules[0].Views[0].Icon,
				mockModules[0].Views[0].CreatedAt,
			).
			AddRow(
				mockModules[0].Id,
				mockModules[0].Name,
				mockModules[0].Description,
				mockModules[0].Code,
				mockModules[0].Icon,
				mockModules[0].Position,
				mockModules[0].CreatedAt,
				mockModules[0].Views[1].Id,
				mockModules[0].Views[1].Name,
				mockModules[0].Views[1].Description,
				mockModules[0].Views[1].Url,
				mockModules[0].Views[1].Icon,
				mockModules[0].Views[1].CreatedAt,
			)
		userId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		mock.ExpectQuery(QueryGetMenu).
			WithArgs(userId).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		res, err := r.GetMenuByUser(ctx, userId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, res[0].Id, mockModules[0].Id)
	})

	t.Run("When get menu of user is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110018"
		mock.ExpectQuery(QueryGetMenu).
			WithArgs(userId).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)

		_, err = r.GetMenuByUser(ctx, userId)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetMenuByUser")
	})
}

func TestRepositoryUsers_GetMeByUser(t *testing.T) {
	t.Run("When get info of user is called successfully", func(t *testing.T) {
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
		userId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		typeDocument := usersDomain.TypeDocument{
			Id:                    StringToPtr("0abbb86f-9836-11ee-a040-0242ac11000e"),
			Number:                StringToPtr("01"),
			Description:           StringToPtr("DOCUMENTO NACIONAL DE IDENTIDAD"),
			AbbreviateDescription: StringToPtr("DNI"),
			Enable:                BoolToPtr(true),
			CreateAt:              &now,
		}
		person := usersDomain.PersonByUser{
			Id:           StringToPtr("0abbb86f-9836-11ee-a040-0242ac11000e"),
			Document:     StringToPtr("77895428"),
			Names:        StringToPtr("LUCY ANDREA"),
			Surname:      StringToPtr("HANCCO"),
			LastName:     StringToPtr("HUILLCA"),
			Phone:        StringToPtr("918547496"),
			Email:        StringToPtr("lucyhancco@gmail.com"),
			Gender:       StringToPtr("MASCULINO"),
			Enable:       BoolToPtr(true),
			CreatedAt:    &now,
			TypeDocument: &typeDocument,
		}
		roleUser := []usersDomain.RoleUser{
			{
				Id:          "0abbb86f-9836-11ee-a040-0242ac11000e",
				Name:        StringToPtr("Gerencia"),
				Description: StringToPtr("Gerencia General"),
				Enable:      BoolToPtr(true),
				CreateAt:    &now,
			},
		}
		mockUserMe := []usersDomain.UserMe{
			{
				Id:        "739bbbc9-7e93-11ee-89fd-0242ac110016",
				UserName:  "pepito.quispe@smart.pe",
				CreatedAt: &now,
				Person:    &person,
				RoleUser:  roleUser,
			},
		}

		rows := sqlmock.NewRows([]string{
			"user_id",
			"user_name",
			"user_created_at",
			"person_id",
			"person_document",
			"person_names",
			"person_surname",
			"person_last_name",
			"person_phone",
			"person_email",
			"person_gender",
			"person_enable",
			"person_created_at",
			"document_type_id",
			"document_type_number",
			"document_type_description",
			"document_type_abbreviated_description",
			"document_type_enable",
			"document_type_created_at",
			"role_id",
			"role_name",
			"role_description",
			"role_enable",
			"role_created_at",
		}).
			AddRow(
				mockUserMe[0].Id,
				mockUserMe[0].UserName,
				mockUserMe[0].CreatedAt,
				mockUserMe[0].Person.Id,
				mockUserMe[0].Person.Document,
				mockUserMe[0].Person.Names,
				mockUserMe[0].Person.Surname,
				mockUserMe[0].Person.LastName,
				mockUserMe[0].Person.Phone,
				mockUserMe[0].Person.Email,
				mockUserMe[0].Person.Gender,
				mockUserMe[0].Person.Enable,
				mockUserMe[0].Person.CreatedAt,
				mockUserMe[0].Person.TypeDocument.Id,
				mockUserMe[0].Person.TypeDocument.Number,
				mockUserMe[0].Person.TypeDocument.Description,
				mockUserMe[0].Person.TypeDocument.AbbreviateDescription,
				mockUserMe[0].Person.TypeDocument.Enable,
				mockUserMe[0].Person.TypeDocument.CreateAt,
				mockUserMe[0].RoleUser[0].Id,
				mockUserMe[0].RoleUser[0].Name,
				mockUserMe[0].RoleUser[0].Description,
				mockUserMe[0].RoleUser[0].Enable,
				mockUserMe[0].RoleUser[0].CreateAt,
			)

		mock.ExpectQuery(QueryGetMeUser).
			WithArgs(userId).
			WillReturnRows(rows)
		r := NewUsersRepository(clock, 60)
		res, err := r.GetMeByUser(ctx, userId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NotNil(t, res)
	})

	t.Run("When get user is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110018"
		clock := &mockClock.Clock{}

		mock.ExpectQuery(QueryGetMeUser).
			WithArgs(userId).
			WillReturnError(expectedError)
		r := NewUsersRepository(clock, 60)

		_, err = r.GetMeByUser(ctx, userId)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetMeByUser")
	})
}

func TestUsersMySQLRepo_GetStoresByUser(t *testing.T) {
	t.Run("When get stores by user and return successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		merchant := usersDomain.Merchant{
			Id:          "0abbb86f-9836-11ee-a040-0242ac11000e",
			Name:        "Smart Cities Peru",
			Description: "Proveedor de software",
		}
		mockStoreByUser := []usersDomain.StoreByUser{
			{
				Id:       "739bbbc9-7e93-11ee-89fd-0242ac110016",
				Name:     "Obra 28 de Julio",
				Merchant: merchant,
			},
		}

		rows := sqlmock.NewRows([]string{
			"store_id",
			"store_name",
			"merchant_id",
			"merchant_name",
			"merchant_description",
			"merchant_image_path",
		}).
			AddRow(
				mockStoreByUser[0].Id,
				mockStoreByUser[0].Name,
				mockStoreByUser[0].Merchant.Id,
				mockStoreByUser[0].Merchant.Name,
				mockStoreByUser[0].Merchant.Description,
				mockStoreByUser[0].Merchant.ImagePath,
			)
		userId := "91fb86bd-da46-414b-97a1-fcdaa8cd35d1"
		mock.ExpectQuery(QueryGetStoresByUser).
			WithArgs(userId).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		res, err := r.GetStoresByUser(ctx, userId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return

		}
		assert.NotNil(t, res)
	})

	t.Run("When get stores by user then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		userId := "91fb86bd-da46-414b-97a1-fcdaa8cd35d1"
		mock.ExpectQuery(QueryGetStoresByUser).
			WithArgs(userId).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)

		_, err = r.GetStoresByUser(ctx, userId)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetStoresByUser")
	})
}

func TestRepositoryUsers_CreateUser(t *testing.T) {
	t.Run("When to successfully create a user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		mock.ExpectBegin()
		var tx *sql.Tx
		personId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		gender := "MASCULINO"
		email := "lucyhancco@gmail.com"
		lastName := "HUILLCA"
		person := usersDomain.Person{
			Document: "77895428",
			Names:    "LUCY ANDREA HANCCO HUILLCA",
			Surname:  "HANCCO",
			LastName: &lastName,
			Phone:    "918547496",
			Email:    &email,
			Gender:   &gender,
			Enable:   true,
		}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		createUserBody := usersDomain.CreateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			Password:   "pepitoPass",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
			PersonId:   &personId,
			Person:     &person,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectExec(QueryCreateUser).
			WithArgs(
				userId,
				createUserBody.UserName,
				createUserBody.Password,
				createUserBody.UserTypeId,
				createdAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		r := NewUsersRepository(clock, 60)
		tx, err = db.Begin()
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		_, err = r.CreateUser(ctx, tx, userId, createUserBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, nil, err)
	})

	t.Run("When an error occurs while creating a user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		userId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		personId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		gender := "MASCULINO"
		email := "lucyhancco@gmail.com"
		lastName := "HUILLCA"
		person := usersDomain.Person{
			Document: "77895428",
			Names:    "LUCY ANDREA HANCCO HUILLCA",
			Surname:  "HANCCO",
			LastName: &lastName,
			Phone:    "918547496",
			Email:    &email,
			Gender:   &gender,
			Enable:   true,
		}
		createUserBody := usersDomain.CreateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			Password:   "pepitoPass",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
			PersonId:   &personId,
			Person:     &person,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectQuery(QueryCreateUser).
			WithArgs(
				userId,
				createUserBody.UserName,
				createUserBody.Password,
				createUserBody.UserTypeId,
				createdAt,
			).
			WillReturnError(expectedError)
		r := NewUsersRepository(clock, 60)
		_, err = r.CreateUser(ctx, nil, userId, createUserBody)
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
		assert.Equal(t, smartErr.Function, "CreateUser")
	})
}

func TestRepositoryUsers_CreateUserMain(t *testing.T) {
	t.Run("When to successfully create a user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		personId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		gender := "MASCULINO"
		email := "lucyhancco@gmail.com"
		lastName := "HUILLCA"
		personBody := usersDomain.Person{
			TypeDocumentId: "00a58522-93b4-11ee-a040-0242ac11000e",
			Document:       "77895428",
			Names:          "LUCY ANDREA HANCCO HUILLCA",
			Surname:        "HANCCO",
			LastName:       &lastName,
			Phone:          "918547496",
			Email:          &email,
			Gender:         &gender,
			Enable:         true,
		}

		userId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		createUserBody := usersDomain.CreateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			Password:   "pepitoPass",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
			PersonId:   &personId,
			Person:     &personBody,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectBegin()

		mock.ExpectExec(QueryCreateUser).
			WithArgs(
				userId,
				createUserBody.UserName,
				createUserBody.Password,
				createUserBody.UserTypeId,
				createdAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec(QueryCreatePerson).WithArgs(
			personId,
			userId,
			personBody.TypeDocumentId,
			personBody.Document,
			personBody.Names,
			personBody.Surname,
			personBody.LastName,
			personBody.Phone,
			personBody.Email,
			personBody.Gender,
			personBody.Enable,
			createdAt).WillReturnResult(sqlmock.NewResult(1, 1))

		total := 1
		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)
		mock.ExpectQuery(QueryValidateUniqueUserExistence).
			WithArgs(userId).
			WillReturnRows(rows)

		// expect a transaction commit
		mock.ExpectCommit()

		r := NewUsersRepository(clock, 60)
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		_, err = r.CreateUserMain(ctx, userId, personId, createUserBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, nil, err)
	})

	t.Run("When an error occurs while creating a user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		userId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		personId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		gender := "MASCULINO"
		email := "lucyhancco@gmail.com"
		lastName := "HUILLCA"
		personBody := usersDomain.Person{
			TypeDocumentId: "00a58522-93b4-11ee-a040-0242ac11000e",
			Document:       "77895428",
			Names:          "LUCY ANDREA HANCCO HUILLCA",
			Surname:        "HANCCO",
			LastName:       &lastName,
			Phone:          "918547496",
			Email:          &email,
			Gender:         &gender,
			Enable:         false,
		}
		createUserBody := usersDomain.CreateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			Password:   "pepitoPass",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
			PersonId:   &personId,
			Person:     &personBody,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectBegin()

		mock.ExpectExec(QueryCreateUser).
			WithArgs(
				userId,
				createUserBody.UserName,
				createUserBody.Password,
				createUserBody.UserTypeId,
				createdAt,
			).
			WillReturnError(expectedError)

		mock.ExpectExec(QueryCreatePerson).WithArgs(
			personId,
			userId,
			personBody.TypeDocumentId,
			personBody.Document,
			personBody.Names,
			personBody.Surname,
			personBody.LastName,
			personBody.Phone,
			personBody.Email,
			personBody.Gender,
			personBody.Enable,
			createdAt).WillReturnError(expectedError)

		mock.ExpectQuery(QueryValidateUniqueUserExistence).
			WithArgs(userId).
			WillReturnError(expectedError)

		r := NewUsersRepository(clock, 60)
		_, err = r.CreateUserMain(ctx, userId, personId, createUserBody)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)

	})
}

func TestRepositoryUsers_UpdateUserMain(t *testing.T) {
	t.Run("When the user update return success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		personId := "249bbbc9-7e93-11ee-89fd-02428c110016"
		gender := "MASCULINO"
		email := "lucyhancco@gmail.com"
		lastName := "HUILLCA"
		updateBody := usersDomain.Person{
			TypeDocumentId: "249bbbc9-7e93-11ee-87fd-02428c110016",
			Document:       "77895428",
			Names:          "LUCY ANDREA HANCCO HUILLCA",
			Surname:        "HANCCO",
			LastName:       &lastName,
			Phone:          "918547496",
			Email:          &email,
			Gender:         &gender,
			Enable:         false,
		}
		now := time.Now().UTC()
		updateUserBody := usersDomain.UpdateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
			PersonId:   &personId,
			Person:     &updateBody,
		}

		userBody := usersDomain.UserById{
			Id:        "739bbbc9-7e93-11ee-89fd-0242ac110016",
			UserName:  "pepito.quispe@smartc.pe",
			CreatedAt: &now,
		}

		rowUser := sqlmock.NewRows([]string{
			"id",
			"username",
			"created_at",
		}).AddRow(
			userBody.Id,
			userBody.UserName,
			userBody.CreatedAt,
		)

		clock := &mockClock.Clock{}

		mock.ExpectBegin()
		mock.ExpectQuery(QueryGetUserById).WithArgs(
			userId).WillReturnRows(rowUser)

		mock.ExpectExec(QueryUpdateUser).
			WithArgs(
				updateUserBody.UserName,
				updateUserBody.UserTypeId,
				userId,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		mock.
			ExpectExec(QueryUpdatePerson).
			WithArgs(
				userId,
				updateBody.TypeDocumentId,
				updateBody.Document,
				updateBody.Names,
				updateBody.Surname,
				updateBody.LastName,
				updateBody.Phone,
				updateBody.Email,
				updateBody.Gender,
				updateBody.Enable,
				personId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		total := 1
		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)
		mock.
			ExpectQuery(QueryValidateUniqueUserExistence).
			WithArgs(userId).
			WillReturnRows(rows)

		mock.ExpectCommit()

		r := NewUsersRepository(clock, 60)
		err = r.UpdateUserMain(ctx, userId, personId, updateUserBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When an error occurs while updating a user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		personId := "249bbbc9-7e93-11ee-89fd-02428c110016"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		updateUserBody := usersDomain.UpdateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
		}

		gender := "MASCULINO"
		email := "lucyhancco@gmail.com"
		lastName := "HUILLCA"
		updateBody := usersDomain.Person{
			TypeDocumentId: "249bbbc9-7e93-11ee-87fd-02428c110016",
			Document:       "77895428",
			Names:          "LUCY ANDREA HANCCO HUILLCA",
			Surname:        "HANCCO",
			LastName:       &lastName,
			Phone:          "918547496",
			Email:          &email,
			Gender:         &gender,
			Enable:         false,
		}

		clock := &mockClock.Clock{}
		expectedError := errors.New("random error")

		mock.ExpectBegin()

		mock.ExpectQuery(QueryGetUserById).WithArgs(
			userId).WillReturnError(expectedError)

		mock.ExpectExec(QueryUpdateUser).
			WithArgs(
				updateUserBody.UserName,
				updateUserBody.UserTypeId,
				userId,
			).WillReturnError(expectedError)

		mock.
			ExpectExec(QueryUpdatePerson).
			WithArgs(
				userId,
				updateBody.TypeDocumentId,
				updateBody.Document,
				updateBody.Names,
				updateBody.Surname,
				updateBody.LastName,
				updateBody.Phone,
				updateBody.Email,
				updateBody.Gender,
				updateBody.Enable,
				personId).
			WillReturnError(expectedError)

		total := 1
		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)
		mock.
			ExpectQuery(QueryValidateUniqueUserExistence).
			WithArgs(userId).
			WillReturnRows(rows)

		mock.ExpectRollback()

		r := NewUsersRepository(clock, 60)
		err = r.UpdateUserMain(ctx, userId, personId, updateUserBody)
		if err == nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)

	})
}

func TestRepositoryUsers_UpdateUser(t *testing.T) {
	t.Run("When the user is prescribed successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		updateUserBody := usersDomain.UpdateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
		}
		clock := &mockClock.Clock{}

		mock.ExpectExec(QueryUpdateUser).
			WithArgs(
				updateUserBody.UserName,
				updateUserBody.UserTypeId,
				userId,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewUsersRepository(clock, 60)
		err = r.UpdateUser(ctx, userId, updateUserBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When an error occurs while updating a user", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		updateUserBody := usersDomain.UpdateUserBody{
			UserName:   "pepito.quispe@smartc.pe",
			UserTypeId: "739bbbc9-7e93-11ee-89fd-0442ac210931",
		}
		clock := &mockClock.Clock{}
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdateUser).
			WithArgs(
				updateUserBody.UserName,
				updateUserBody.UserTypeId,
				userId,
			).
			WillReturnError(expectedError)
		r := NewUsersRepository(clock, 60)
		err = r.UpdateUser(ctx, userId, usersDomain.UpdateUserBody{})
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
		assert.Equal(t, smartErr.Function, "UpdateUser")
	})
}

func TestRepositoryUsers_DeleteUser(t *testing.T) {
	t.Run("When a user is successfully deleted", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		deletedAt := now.Format("2006-01-02 15:04:06")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectExec(QueryDeleteUser).
			WithArgs(deletedAt, userId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewUsersRepository(clock, 60)
		var res bool
		res, err = r.DeleteUser(ctx, userId)

		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("when an error occurs while reentering a password", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:06")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectExec(QueryDeleteUser).
			WithArgs(deletedAt, userId).
			WillReturnError(errors.New("anything"))
		r := NewUsersRepository(clock, 60)
		var res bool
		res, err = r.DeleteUser(ctx, userId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteUser")
	})
}

func TestRepositoryUsers_ResetPasswordUser(t *testing.T) {
	t.Run("When a user's password is reset", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		passwordHash := "$2a$12$9uZ0CpVgFFDFv4MqDX2m6Of0Hll6l5pRnu14Xx5prcBccZ3j0jU72"
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		mock.ExpectExec(QueryResetPasswordUser).
			WithArgs(userId, passwordHash).
			WillReturnResult(sqlmock.NewResult(1, 1))
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		var res bool
		res, err = r.ResetPasswordUser(ctx, passwordHash, userId)
		if res == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When an error occurs while resetting a user's password", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		passwordHash := "$2a$12$9uZ0CpVgFFDFv4MqDX2m6Of0Hll6l5pRnu14Xx5prcBccZ3j0jU72"
		mock.ExpectExec(QueryResetPasswordUser).
			WithArgs(userId, passwordHash).
			WillReturnError(errors.New("anything"))
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		var res bool
		res, err = r.ResetPasswordUser(ctx, passwordHash, userId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "ResetPasswordUser")
	})
}

func TestRepositoryUsers_GetUserByUserNameAndPassword(t *testing.T) {
	t.Run("When we get user by username and password", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		userType := usersDomain.UserTypeByUser{
			Id:          "739bbbc9-7e93-11ee-89fd-0242ac110018",
			Description: "Usuario externo",
			Code:        "USER_EXTERNAL",
		}

		mockUser := usersDomain.User{
			Id:        "739bbbc9-7e93-11ee-89fd-0242ac110016",
			UserName:  "pepito.quispe@smart.pe",
			CreatedAt: &now,
			UserType:  userType,
		}

		rows := sqlmock.NewRows([]string{"user_id", "user_name", "user_created_at",
			"user_type_id", "user_type_description", "user_type_code"}).
			AddRow(
				mockUser.Id,
				mockUser.UserName,
				mockUser.CreatedAt,
				mockUser.UserType.Id,
				mockUser.UserType.Description,
				mockUser.UserType.Code,
			)
		userName := "pepito.quispe@smart.pe"
		passwordHash := "$2a$12$9uZ0CpVgFFDFv4MqDX2m6Of0Hll6l5pRnu14Xx5prcBccZ3j0jU72"
		mock.ExpectQuery(QueryGetUserByPassword).
			WithArgs(userName, passwordHash).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}

		r := NewUsersRepository(clock, 60)
		res, _, err := r.GetUserByUserNameAndPassword(ctx, userName, passwordHash)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, res.Id, mockUser.Id)
		assert.Equal(t, res.UserName, mockUser.UserName)
	})

	t.Run("When calling a user by their username and password it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		userName := "pepito.quispe@smart.pe"
		passwordHash := "$2a$12$9uZ0CpVgFFDFv4MqDX2m6Of0Hll6l5pRnu14Xx5prcBccZ3j0jU72"
		mock.ExpectQuery(QueryGetUser).
			WithArgs(userName, passwordHash).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)

		_, _, err = r.GetUserByUserNameAndPassword(ctx, userName, passwordHash)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetUserByUserNameAndPassword")
	})
}

func TestRepositoryUsers_ValidateUniquePersonByDocument(t *testing.T) {
	t.Run("When verify total of document of the user  is called then it should return a total", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		total := 1
		gender := "MASCULINO"
		email := "lucyhancco@gmail.com"
		lastName := "HUILLCA"
		personBody := usersDomain.Person{
			TypeDocumentId: "00a58522-93b4-11ee-a040-0242ac11000e",
			Document:       "77895428",
			Names:          "LUCY ANDREA HANCCO HUILLCA",
			Surname:        "HANCCO",
			LastName:       &lastName,
			Phone:          "918547496",
			Email:          &email,
			Gender:         &gender,
			Enable:         true,
		}

		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)
		mock.
			ExpectQuery(QueryValidateUniquePersonByDocument).
			WithArgs(personBody.TypeDocumentId, personBody.Document).
			WillReturnRows(rows)

		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)

		err = r.ValidateUniquePersonByDocument(ctx, personBody.TypeDocumentId, personBody.Document)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, nil, err)
	})

	t.Run("When verify total of document of the user is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		gender := "MASCULINO"
		email := "lucyhancco@gmail.com"
		lastName := "HUILLCA"
		personBody := usersDomain.Person{
			TypeDocumentId: "00a58522-93b4-11ee-a040-0242ac11000e",
			Document:       "77895428",
			Names:          "LUCY ANDREA HANCCO HUILLCA",
			Surname:        "HANCCO",
			LastName:       &lastName,
			Phone:          "918547496",
			Email:          &email,
			Gender:         &gender,
			Enable:         true,
		}
		mock.ExpectQuery(QueryValidateUniquePersonByDocument).
			WithArgs(personBody.TypeDocumentId, personBody.Document).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		err = r.ValidateUniquePersonByDocument(ctx, personBody.TypeDocumentId, personBody.Document)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "ValidateUniquePersonByDocument")
	})
}

func TestUsersMySQLRepo_VerifyPermissionsByUser(t *testing.T) {
	t.Run("When verify permission by user return an amount then is true", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		total := 2
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110018"
		codePermission := "CREATE_PRODUCT"
		var totalExpected bool
		rows := sqlmock.NewRows([]string{"total"}).
			AddRow(total)
		mock.
			ExpectQuery(QueryVerifyPermissionsByUser).
			WithArgs(userId,
				storeId,
				storeId,
				storeId,
				codePermission).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)

		totalExpected, err = r.VerifyPermissionsByUser(ctx, userId, storeId, codePermission)
		if totalExpected == false {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, totalExpected)
	})

	t.Run("When verify permission by user return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		storeId := "739bbbc9-7e93-11ee-89fd-0242ac110018"
		codePermission := "CREATE_PRODUCT"
		var totalExpected bool

		mock.
			ExpectQuery(QueryVerifyPermissionsByUser).
			WithArgs(userId,
				storeId,
				storeId,
				storeId,
				codePermission).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		totalExpected, err = r.VerifyPermissionsByUser(ctx, userId, storeId, codePermission)
		assert.Error(t, err)
		assert.Equal(t, false, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "VerifyPermissionsByUser")
	})
}

func TestUsersMySQLRepo_GetModulePermissions(t *testing.T) {
	t.Run("When it returns the list of permissions per module of a user successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		userId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		codePermission := "logistics.requirementss"
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		var res []usersDomain.Permissions
		permissions := []usersDomain.Permissions{
			{
				Id:   "0c4001f3-2dd8-4d9f-820d-db7d7d8c85c0",
				Code: "CREATE_REQUIREMENT",
			},
			{
				Id:   "0c4001f3-2dd8-4d9f-820d-db7d7d8c85c0",
				Code: "UPDATE_REQUIREMENT",
			},
		}
		rows := sqlmock.NewRows([]string{"id", "code"})
		for _, permission := range permissions {
			rows.AddRow(
				permission.Id,
				permission.Code,
			)
		}

		mock.
			ExpectQuery(QueryGetModulePermissions).
			WithArgs(codePermission, userId).
			WillReturnRows(rows)

		res, err = r.GetModulePermissions(ctx, userId, codePermission)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 2)
		assert.Equal(t, "CREATE_REQUIREMENT", permissions[0].Code)
		assert.Equal(t, "UPDATE_REQUIREMENT", permissions[1].Code)
	})

	t.Run("When the list of permissions per module of a user returns an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		userId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		codePermission := "logistics.requirementss"
		clock := &mockClock.Clock{}
		r := NewUsersRepository(clock, 60)
		var res []usersDomain.Permissions

		mock.
			ExpectQuery(QueryGetModulePermissions).
			WithArgs(codePermission, userId).
			WillReturnError(expectedError)

		res, err = r.GetModulePermissions(ctx, userId, codePermission)
		if res != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetModulePermissions")
	})
}
