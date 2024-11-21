/*
 * File: user_types_usecase_test.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of user type.
 *
 * Last Modified: 2023-11-23
 */

package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	mockValidation "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain/mocks"

	userTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain"
	mockUserTypes "gitlab.smartcitiesperu.com/smartone/api-core/user-types/domain/mocks"
)

func TestUseCaseUserTypes_GetUserTypes(t *testing.T) {
	t.Run("When get user types successfully", func(t *testing.T) {
		userTypesRepository := &mockUserTypes.UserTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		userTypesRepository.
			On("GetUserTypes", mock.Anything, mock.Anything).
			Return([]userTypesDomain.UserType{}, nil)
		userTypesRepository.On("GetTotalUserTypes", mock.Anything, mock.Anything).
			Return(&total, nil)
		userTypesUCase := NewUserTypesUseCase(
			userTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		pagination := paramsDomain.NewPaginationParams(nil)
		userTypes, _, err := userTypesUCase.GetUserTypes(context.Background(), pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, userTypes, []userTypesDomain.UserType{})
	})
	t.Run("When get user types error", func(t *testing.T) {
		userTypesRepository := &mockUserTypes.UserTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		userTypesRepository.
			On("GetUserTypes", mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		userTypesRepository.On("GetTotalUserTypes", mock.Anything, mock.Anything).
			Return(&total, nil)
		userTypesUCase := NewUserTypesUseCase(
			userTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		pagination := paramsDomain.NewPaginationParams(nil)
		ok, _, err := userTypesUCase.GetUserTypes(context.Background(), pagination)
		assert.Error(t, err)
		assert.EqualValues(t, ok, []userTypesDomain.UserType(nil))
	})
}

func TestUseCaseUserTypes_CreateUserType(t *testing.T) {
	t.Run("When create user type successfully", func(t *testing.T) {
		userTypesRepository := &mockUserTypes.UserTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		userTypeID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		userTypesRepository.
			On("CreateUserType", mock.Anything, mock.Anything, mock.Anything).
			Return(&userTypeID, nil)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		userTypesUCase := NewUserTypesUseCase(
			userTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := userTypesUCase.CreateUserType(
			context.Background(),
			userTypesDomain.CreateUserTypeBody{},
		)
		assert.NoError(t, err)
	})
	t.Run("When create user type error", func(t *testing.T) {
		userTypesRepository := &mockUserTypes.UserTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		userTypesRepository.
			On("CreateUserType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(true, nil)
		userTypesUCase := NewUserTypesUseCase(
			userTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := userTypesUCase.CreateUserType(
			context.Background(),
			userTypesDomain.CreateUserTypeBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, userTypesDomain.ErrUserTypeDescriptionAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateUserType")
	})

	t.Run("When an error occurs while creating a user type", func(t *testing.T) {
		userTypesRepository := &mockUserTypes.UserTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		errCreate := errDomain.NewErr().SetFunction("CreateUserType").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		userTypesRepository.
			On("CreateUserType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		usersUCase := NewUserTypesUseCase(userTypesRepository, validationRepository, authRepository, 60)
		_, err := usersUCase.CreateUserType(
			context.Background(),
			userTypesDomain.CreateUserTypeBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateUserType")
	})
}

func TestUseCaseUserTypes_UpdateUserType(t *testing.T) {
	t.Run("When update user type successfully", func(t *testing.T) {
		userTypesRepository := &mockUserTypes.UserTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		userTypesRepository.
			On("UpdateUserType", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		userTypesUCase := NewUserTypesUseCase(
			userTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := userTypesUCase.UpdateUserType(
			context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
			userTypesDomain.UpdateUserTypeBody{},
		)
		assert.NoError(t, err)
	})
	t.Run("When update user type error", func(t *testing.T) {
		userTypesRepository := &mockUserTypes.UserTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		validationRepository.On("RecordExists", mock.Anything, mock.Anything).Return(
			false, nil)
		userTypesRepository.
			On("UpdateUserType",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(errors.New("random error"))
		userTypesUCase := NewUserTypesUseCase(
			userTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := userTypesUCase.UpdateUserType(
			context.Background(),
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
			userTypesDomain.UpdateUserTypeBody{},
		)
		assert.Error(t, err)
	})
}

func TestUseCaseUserTypes_DeleteUserType(t *testing.T) {
	t.Run("When delete user type by id successfully", func(t *testing.T) {
		userTypesRepository := &mockUserTypes.UserTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		userTypesRepository.
			On("DeleteUserType", mock.Anything, mock.Anything).
			Return(true, nil)
		userTypesUCase := NewUserTypesUseCase(
			userTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		userTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := userTypesUCase.DeleteUserType(context.Background(), userTypeId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})
	t.Run("When delete user type by id error", func(t *testing.T) {
		userTypesRepository := &mockUserTypes.UserTypeRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		userTypesError := errors.New("random error")
		userTypesRepository.
			On("DeleteUserType", mock.Anything, mock.Anything).
			Return(false, userTypesError)
		userTypesUCase := NewUserTypesUseCase(
			userTypesRepository,
			validationRepository,
			authRepository,
			60,
		)
		userTypeId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := userTypesUCase.DeleteUserType(context.Background(), userTypeId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
