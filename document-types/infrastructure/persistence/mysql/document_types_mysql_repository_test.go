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

	"gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain"
)

func TestRepositoryDocumentTypes_GetDocumentTypes(t *testing.T) {
	t.Run("When get document types is called then it should return a list of document type", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		mockDocumentTypes := []domain.DocumentType{
			{
				Id:                     "00a58296-93b4-11ee-a040-0242ac11000e",
				Number:                 "01",
				Description:            "DOCUMENTO NACIONAL DE IDENTIDAD",
				AbbreviatedDescription: "DNI",
				Enable:                 1,
				CreatedAt:              &now,
			},
		}

		rows := sqlmock.NewRows([]string{"id", "number", "description",
			"abbreviated_description", "enable", "created_at"}).
			AddRow(
				mockDocumentTypes[0].Id,
				mockDocumentTypes[0].Number,
				mockDocumentTypes[0].Description,
				mockDocumentTypes[0].AbbreviatedDescription,
				mockDocumentTypes[0].Enable,
				mockDocumentTypes[0].CreatedAt,
			)

		sizePage := 100
		offset := 0
		description := "DOCUMENTO"
		mock.
			ExpectQuery(QueryGetDocumentTypes).
			WithArgs(
				description,
				description,
				description,
				description,
				sizePage,
				offset).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewDocumentTypesRepository(clock, 60)
		var res []domain.DocumentType
		searchParams := domain.GetDocumentTypeParams{
			SearchDescription: &description,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		pagination.Page = 1
		pagination.SizePage = sizePage
		res, err = r.GetDocumentTypes(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Len(t, res, 1)
	})

	t.Run("When get document types is called then it should return an error", func(t *testing.T) {
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
		mock.ExpectQuery(QueryGetDocumentTypes).
			WithArgs("739bbbc9-7e93-11ee-89fd-0242ac113421", "739bbbc9-7e93-11ee-89fd-0242ac113421",
				sizePage, offset).
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewDocumentTypesRepository(clock, 60)
		var res []domain.DocumentType
		searchParams := domain.GetDocumentTypeParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		res, err = r.GetDocumentTypes(ctx, searchParams, pagination)
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
		assert.Equal(t, smartErr.Function, "GetDocumentTypes")
	})
}

func TestRepositoryDocumentTypes_GetTotalDocumentTypes(t *testing.T) {
	t.Run("When get total of document types is called then it should return a total", func(t *testing.T) {
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

		description := "DOCUMENTO"
		mock.
			ExpectQuery(QueryGetTotalDocumentTypes).
			WithArgs(
				description,
				description,
				description,
				description,
			).
			WillReturnRows(rows)
		clock := &mockClock.Clock{}
		r := NewDocumentTypesRepository(clock, 60)

		searchParams := domain.GetDocumentTypeParams{
			SearchDescription: &description,
		}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalDocumentTypes(ctx, searchParams, pagination)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Equal(t, *totalExpected, total)
	})

	t.Run("When get total of document types is called then it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")

		mock.ExpectQuery(QueryGetTotalDocumentTypes).
			WithArgs().
			WillReturnError(expectedError)
		clock := &mockClock.Clock{}
		r := NewDocumentTypesRepository(clock, 60)

		searchParams := domain.GetDocumentTypeParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		var totalExpected *int
		totalExpected, err = r.GetTotalDocumentTypes(ctx, searchParams, pagination)
		assert.Error(t, err)
		var intPointer *int
		assert.Equal(t, intPointer, totalExpected)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "GetTotalDocumentTypes")
	})
}

func TestRepositoryDocumentTypes_CreateDocumentType(t *testing.T) {
	t.Run("When to successfully create a documen type ", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		createUserBody := domain.CreateDocumentTypeBody{
			Number:                 "01",
			Description:            "DOCUMENTO NACIONAL DE IDENTIDAD",
			AbbreviatedDescription: "DNI",
			Enable:                 1,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectExec(QueryCreateDocumentType).
			WithArgs(
				documentTypeId,
				createUserBody.Number,
				createUserBody.Description,
				createUserBody.AbbreviatedDescription,
				createdAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		r := NewDocumentTypesRepository(clock, 60)
		_, err = r.CreateDocumentType(ctx, documentTypeId, createUserBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When an error occurs while creating a documen type ", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		expectedError := errors.New("random error")
		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		createUserBody := domain.CreateDocumentTypeBody{
			Number:                 "01",
			Description:            "DOCUMENTO NACIONAL DE IDENTIDAD",
			AbbreviatedDescription: "DNI",
			Enable:                 1,
		}
		now := time.Now().UTC()
		createdAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)
		mock.ExpectQuery(QueryCreateDocumentType).
			WithArgs(
				documentTypeId,
				createUserBody.Number,
				createUserBody.Description,
				createUserBody.AbbreviatedDescription,
				createdAt,
			).
			WillReturnError(expectedError)
		r := NewDocumentTypesRepository(clock, 60)
		_, err = r.CreateDocumentType(ctx, documentTypeId, createUserBody)
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
		assert.Equal(t, smartErr.Function, "CreateDocumentType")
	})
}

func TestRepositoryDocumentTypes_UpdateDocumentType(t *testing.T) {
	t.Run("When the documen type  is prescribed successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		createUserBody := domain.UpdateDocumentTypeBody{
			Number:                 "01",
			Description:            "DOCUMENTO NACIONAL DE IDENTIDAD",
			AbbreviatedDescription: "DNI",
			Enable:                 1,
		}
		clock := &mockClock.Clock{}

		mock.ExpectExec(QueryUpdateDocumentType).
			WithArgs(
				createUserBody.Number,
				createUserBody.Description,
				createUserBody.AbbreviatedDescription,
				createUserBody.Enable,
				documentTypeId,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewDocumentTypesRepository(clock, 60)
		err = r.UpdateDocumentType(ctx, documentTypeId, createUserBody)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.Nil(t, err)
	})

	t.Run("When an error occurs while updating a documen type ", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		documentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac117201"
		createUserBody := domain.UpdateDocumentTypeBody{
			Number:                 "01",
			Description:            "DOCUMENTO NACIONAL DE IDENTIDAD",
			AbbreviatedDescription: "DNI",
			Enable:                 1,
		}
		clock := &mockClock.Clock{}
		expectedError := errors.New("random error")
		mock.ExpectQuery(QueryUpdateDocumentType).
			WithArgs(
				createUserBody.Number,
				createUserBody.Description,
				createUserBody.AbbreviatedDescription,
				createUserBody.Enable,
				documentTypeId,
			).
			WillReturnError(expectedError)
		r := NewDocumentTypesRepository(clock, 60)
		err = r.UpdateDocumentType(ctx, documentTypeId, createUserBody)
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
		assert.Equal(t, smartErr.Function, "UpdateDocumentType")
	})
}

func TestRepositoryDocumentTypes_DeleteDocumentType(t *testing.T) {
	t.Run("When a documen type  is successfully deleted", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		xTenantId := "739bbbc9-7e93-11ee-89fd-0242ac110022"
		ctx := context.WithValue(context.Background(), "xTenantId", xTenantId)
		db2.AddClientSchemaDB(xTenantId, db)

		now := time.Now().UTC()
		documenTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		deletedAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectExec(QueryDeleteDocumentType).
			WithArgs(deletedAt, documenTypeId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := NewDocumentTypesRepository(clock, 60)
		var res bool
		res, err = r.DeleteDocumentType(ctx, documenTypeId)

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

		DocumentTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		now := time.Now().UTC()
		deletedAt := now.Format("2006-01-02 15:04:05")
		clock := &mockClock.Clock{}
		clock.On("Now").Return(now)

		mock.ExpectExec(QueryDeleteDocumentType).
			WithArgs(deletedAt, DocumentTypeId).
			WillReturnError(errors.New("anything"))
		r := NewDocumentTypesRepository(clock, 60)
		var res bool
		res, err = r.DeleteDocumentType(ctx, DocumentTypeId)

		assert.Error(t, err)
		assert.Equal(t, false, res)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.Infra)
		assert.Equal(t, smartErr.Function, "DeleteDocumentType")
	})
}
