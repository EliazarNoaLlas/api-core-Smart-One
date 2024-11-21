/*
 * File: view_usecase_test.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of rolePolicies.
 *
 * Last Modified: 2023-11-24
 */

package usecases

import (
	"context"
	"errors"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	mockValidation "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain/mocks"

	viewDomain "gitlab.smartcitiesperu.com/smartone/api-core/views/domain"
	mockViews "gitlab.smartcitiesperu.com/smartone/api-core/views/domain/mocks"
)

func TestViewsUseCase_GetViews(t *testing.T) {
	t.Run("When views are successfully listed", func(t *testing.T) {
		viewsRepository := &mockViews.ViewRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		total := 10
		viewsRepository.
			On("GetViews", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return([]viewDomain.View{}, nil)
		viewsRepository.
			On("GetTotalViews", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		viewsUCase := NewViewUseCase(
			viewsRepository,
			validationRepository,
			authRepository,
			60,
		)
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		searchParams := viewDomain.GetViewsParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		views, _, err := viewsUCase.GetViews(context.Background(), moduleId, searchParams, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, views, []viewDomain.View{})
	})

	t.Run("When an error occurs while listing views", func(t *testing.T) {
		viewsRepository := &mockViews.ViewRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		total := 10
		viewsRepository.
			On("GetViews", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		viewsRepository.
			On("GetTotalViews", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		viewsUCase := NewViewUseCase(
			viewsRepository,
			validationRepository,
			authRepository,
			60,
		)
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		searchParams := viewDomain.GetViewsParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		views, _, err := viewsUCase.GetViews(context.Background(), moduleId, searchParams, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, views, []viewDomain.View(nil))
	})
}

func TestViewsUseCase_CreateView(t *testing.T) {
	t.Run("When add a view successfully", func(t *testing.T) {
		viewRepository := &mockViews.ViewRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		createViewBody := viewDomain.CreateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		viewRepository.
			On("CreateView", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&viewId, nil)

		viewUCase := NewViewUseCase(
			viewRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := viewUCase.CreateView(context.Background(), moduleId, createViewBody)
		assert.NoError(t, err)
	})

	t.Run("When add a view, error", func(t *testing.T) {
		viewRepository := &mockViews.ViewRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		createViewBody := viewDomain.CreateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		errCreate := errDomain.NewErr().SetFunction("CreateView").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		viewRepository.
			On("CreateView", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		viewUCase := NewViewUseCase(
			viewRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := viewUCase.CreateView(
			context.Background(),
			moduleId,
			createViewBody,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateView")

	})
}

func TestViewsUseCase_UpdateView(t *testing.T) {
	t.Run("When update a view successfully", func(t *testing.T) {
		viewRepository := &mockViews.ViewRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		viewId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		updateViewBody := viewDomain.UpdateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		viewRepository.
			On("UpdateView", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		viewUCase := NewViewUseCase(
			viewRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := viewUCase.UpdateView(
			context.Background(),
			moduleId,
			viewId,
			updateViewBody,
		)
		assert.NoError(t, err)
	})

	t.Run("When update a view, error", func(t *testing.T) {
		viewRepository := &mockViews.ViewRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		viewRepository.
			On("UpdateView", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		viewUCase := NewViewUseCase(
			viewRepository,
			validationRepository,
			authRepository,
			60,
		)
		viewId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		updateViewBody := viewDomain.UpdateViewBody{
			Name:        "logistics",
			Description: "View about logistics",
			Url:         "/logistics/requirements",
			Icon:        "fa fa-table",
		}
		err := viewUCase.UpdateView(
			context.Background(),
			moduleId,
			viewId,
			updateViewBody,
		)
		assert.Error(t, err)
	})
}

func TestViewsUseCase_DeleteView(t *testing.T) {
	t.Run("When delete a view successfully", func(t *testing.T) {
		viewRepository := &mockViews.ViewRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		viewRepository.
			On("DeleteView", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		viewUCase := NewViewUseCase(
			viewRepository,
			validationRepository,
			authRepository,
			60,
		)
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		viewId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		res, err := viewUCase.DeleteView(context.Background(), moduleId, viewId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete a view, error", func(t *testing.T) {
		viewRepository := &mockViews.ViewRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		rolePoliciesError := errors.New("random error")
		validationRepository.On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		viewRepository.
			On("DeleteView", mock.Anything, mock.Anything, mock.Anything).
			Return(false, rolePoliciesError)
		viewUCase := NewViewUseCase(
			viewRepository,
			validationRepository,
			authRepository,
			60,
		)
		moduleId := "739bbbc9-7e93-11ee-89fd-0242ac110097"
		viewId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		res, err := viewUCase.DeleteView(context.Background(), moduleId, viewId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
