/*
 * File: modules_usecase_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of modules.
 *
 * Last Modified: 2023-11-10
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

	modulesDomain "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain"
	mockModules "gitlab.smartcitiesperu.com/smartone/api-core/modules/domain/mocks"
)

func TestUseCaseModules_GetModules(t *testing.T) {
	t.Run("When get modules successfully", func(t *testing.T) {
		modulesRepository := &mockModules.ModuleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		searchParams := modulesDomain.GetModulesParams{}
		total := 10

		modulesRepository.
			On("GetModules", mock.Anything, mock.Anything, mock.Anything).
			Return([]modulesDomain.Module{}, nil)
		modulesRepository.
			On("GetTotalModules", mock.Anything, mock.Anything).
			Return(&total, nil)
		modulesUCase := NewModulesUseCase(
			modulesRepository,
			validationRepository,
			authRepository,
			60,
		)
		pagination := paramsDomain.NewPaginationParams(nil)
		modules, _, err := modulesUCase.GetModules(context.Background(), searchParams, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, modules, []modulesDomain.Module{})
	})

	t.Run("When get modules error", func(t *testing.T) {
		modulesRepository := &mockModules.ModuleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		searchParams := modulesDomain.GetModulesParams{}
		total := 10

		modulesRepository.
			On("GetModules", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		modulesRepository.
			On("GetTotalModules", mock.Anything, mock.Anything).
			Return(&total, nil)
		modulesUCase := NewModulesUseCase(
			modulesRepository,
			validationRepository,
			authRepository,
			60,
		)
		pagination := paramsDomain.NewPaginationParams(nil)
		ok, _, err := modulesUCase.GetModules(context.Background(), searchParams, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, ok, []modulesDomain.Module(nil))
	})
}

func TestUseCaseModules_CreateModule(t *testing.T) {
	t.Run("When create module successfully", func(t *testing.T) {
		modulesRepository := &mockModules.ModuleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		modulesRepository.
			On("CreateModule", mock.Anything, mock.Anything, mock.Anything).
			Return(&moduleId, nil)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		modulesUCase := NewModulesUseCase(
			modulesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := modulesUCase.CreateModule(
			context.Background(),
			modulesDomain.CreateModuleBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When create module error", func(t *testing.T) {
		modulesRepository := &mockModules.ModuleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		modulesRepository.
			On("CreateModule", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(true, nil)
		modulesUCase := NewModulesUseCase(
			modulesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := modulesUCase.CreateModule(
			context.Background(),
			modulesDomain.CreateModuleBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, modulesDomain.ErrModuleCodeAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateModule")
	})

	t.Run("When create module error", func(t *testing.T) {
		modulesRepository := &mockModules.ModuleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		errCreate := errDomain.NewErr().SetFunction("CreateModule").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		modulesRepository.
			On("CreateModule", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		modulesUCase := NewModulesUseCase(
			modulesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := modulesUCase.CreateModule(
			context.Background(),
			modulesDomain.CreateModuleBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateModule")
	})
}

func TestUseCaseModules_UpdateModule(t *testing.T) {
	t.Run("When update module successfully", func(t *testing.T) {
		modulesRepository := &mockModules.ModuleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		modulesRepository.
			On("UpdateModule", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		modulesUCase := NewModulesUseCase(
			modulesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := modulesUCase.UpdateModule(
			context.Background(),
			moduleId,
			modulesDomain.UpdateModuleBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When update module error", func(t *testing.T) {
		modulesRepository := &mockModules.ModuleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		modulesRepository.
			On("UpdateModule",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(errors.New("random error"))
		modulesUCase := NewModulesUseCase(
			modulesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := modulesUCase.UpdateModule(
			context.Background(),
			moduleId,
			modulesDomain.UpdateModuleBody{},
		)
		assert.Error(t, err)
	})
}

func TestUseCaseModules_DeleteModule(t *testing.T) {
	t.Run("When delete module by id successfully", func(t *testing.T) {
		modulesRepository := &mockModules.ModuleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		modulesRepository.
			On("DeleteModule", mock.Anything, mock.Anything).
			Return(true, nil)
		modulesUCase := NewModulesUseCase(
			modulesRepository,
			validationRepository,
			authRepository,
			60,
		)
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := modulesUCase.DeleteModule(context.Background(), moduleId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete module by id error", func(t *testing.T) {
		modulesRepository := &mockModules.ModuleRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		modulesError := errors.New("random error")
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		modulesRepository.
			On("DeleteModule", mock.Anything, mock.Anything).
			Return(false, modulesError)
		modulesUCase := NewModulesUseCase(
			modulesRepository,
			validationRepository,
			authRepository, 60)
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := modulesUCase.DeleteModule(context.Background(), moduleId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
