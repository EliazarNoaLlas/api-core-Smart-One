/*
 * File: policyPermissions_usecase_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of policyPermissions.
 *
 * Last Modified: 2023-11-20
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

	policyPermissionsDomain "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain"
	mockPolicyPermissions "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/domain/mocks"
)

var (
	policyHasPermission = true
)

func TestUseCasePolicyPermissions_GetPolicyPermissionByPolicy(t *testing.T) {
	t.Run("When policy permissions are successfully listed", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		policyPermissionsRepository.
			On("GetPolicyPermissionsByPolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return([]policyPermissionsDomain.PolicyPermission{}, nil)
		policyPermissionsRepository.
			On("GetTotalPolicyPermissionsByPolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(&total, nil)
		pagination := paramsDomain.NewPaginationParams(nil)
		policyPermissionUCase := NewPolicyPermissionsUseCase(policyPermissionsRepository, validationRepository, authRepository, 60)
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110019"
		policyPermission, _, err := policyPermissionUCase.GetPolicyPermissionsByPolicy(context.Background(), policyId, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, policyPermission, []policyPermissionsDomain.PolicyPermission{})
	})

	t.Run("When an error occurs while listing policy permissions", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		total := 10
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(false, nil)
		policyPermissionsRepository.
			On("GetPolicyPermissionsByPolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, errors.New("random error"))
		policyPermissionsRepository.
			On("GetTotalPolicyPermissionsByPolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(&total, nil)
		policyPermissionUCase := NewPolicyPermissionsUseCase(policyPermissionsRepository, validationRepository, authRepository, 60)
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110019"
		pagination := paramsDomain.NewPaginationParams(nil)
		policyPermission, _, err := policyPermissionUCase.GetPolicyPermissionsByPolicy(context.Background(), policyId, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, policyPermission, []policyPermissionsDomain.PolicyPermission(nil))
	})
}

func TestUseCasePolicyPermissions_CreatePolicyPermission(t *testing.T) {
	policyHasPermission := false
	t.Run("When create policyPermission successfully", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		createPolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		policyPermissionsRepository.
			On("CreatePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(&policyPermissionID, nil)
		policyPermissionsRepository.
			On("VerifyPolicyHasPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(policyHasPermission, nil)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		_, err := policyPermissionsUCase.CreatePolicyPermission(
			context.Background(),
			policyId,
			createPolicyPermissionBody,
		)
		assert.NoError(t, err)
	})

	t.Run("When creating a policy permission error when the permission already exists", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		createPolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       false,
		}
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionsRepository.
			On("CreatePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, errors.New("random error"))
		policyPermissionsRepository.
			On("VerifyPolicyHasPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		_, err := policyPermissionsUCase.CreatePolicyPermission(
			context.Background(),
			policyId,
			createPolicyPermissionBody,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, policyPermissionsDomain.ErrPolicyHasPermissionAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreatePolicyPermission")
	})

	t.Run("When create policyPermission error", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		createPolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       false,
		}
		errCreate := errDomain.NewErr().SetFunction("CreatePolicyPermission").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionsRepository.
			On("CreatePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil, errCreate)
		policyPermissionsRepository.
			On("VerifyPolicyHasPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(policyHasPermission, nil)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		_, err := policyPermissionsUCase.CreatePolicyPermission(
			context.Background(),
			policyId,
			createPolicyPermissionBody,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreatePolicyPermission")
	})
}

func TestUseCasePolicyPermissions_CreatePolicyPermissions(t *testing.T) {
	policyHasPermission := false
	t.Run("When create multiple policy permissions successfully", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		body := []policyPermissionsDomain.CreatePolicyPermissionBody{
			{
				PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
				Enable:       true,
			},
		}
		policyPermissionsRepository.
			On("CreatePolicyPermissions",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(nil)
		policyPermissionsRepository.
			On("VerifyPolicyHasPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(policyHasPermission, nil)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		_, err := policyPermissionsUCase.CreatePolicyPermissions(
			context.Background(),
			policyId,
			body,
		)
		assert.NoError(t, err)
	})

	t.Run("When an error occurs while creating multiple permissions", func(t *testing.T) {

		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		body := []policyPermissionsDomain.CreatePolicyPermissionBody{
			{
				PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
				Enable:       true,
			},
		}
		errCreate := errDomain.NewErr().SetFunction("CreatePolicyPermissions").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionsRepository.
			On("CreatePolicyPermissions",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(errCreate)
		policyPermissionsRepository.
			On("VerifyPolicyHasPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(policyHasPermission, nil)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		_, err := policyPermissionsUCase.CreatePolicyPermissions(
			context.Background(),
			policyId,
			body,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreatePolicyPermissions")
	})
}

func TestUseCasePolicyPermissions_UpdatePolicyPermission(t *testing.T) {
	t.Run("When update policyPermission successfully", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		updatePolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		policyPermissionsRepository.
			On("UpdatePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		err := policyPermissionsUCase.UpdatePolicyPermission(
			context.Background(),
			policyId,
			policyPermissionId,
			updatePolicyPermissionBody,
		)
		assert.NoError(t, err)
	})

	t.Run("When update policyPermission error", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		updatePolicyPermissionBody := policyPermissionsDomain.CreatePolicyPermissionBody{
			PermissionId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
			Enable:       true,
		}
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(false, nil)
		policyPermissionsRepository.
			On("UpdatePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(errors.New("random error"))
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		err := policyPermissionsUCase.UpdatePolicyPermission(
			context.Background(),
			policyId,
			policyPermissionId,
			updatePolicyPermissionBody,
		)
		assert.Error(t, err)
	})
}

func TestUseCasePolicyPermissions_DeletePolicyPermission(t *testing.T) {
	t.Run("When delete policyPermission by id successfully", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		policyPermissionsRepository.
			On("DeletePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := policyPermissionsUCase.DeletePolicyPermission(context.Background(), policyId, policyPermissionId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete policyPermission by id error", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		policyPermissionsError := errors.New("random error")
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(false, nil)
		policyPermissionsRepository.
			On("DeletePolicyPermission",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(false, policyPermissionsError)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := policyPermissionsUCase.DeletePolicyPermission(context.Background(), policyId, policyPermissionId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}

func TestUseCasePolicyPermissions_DeletePolicyPermissions(t *testing.T) {
	t.Run("When delete multiple policy permissions successfully", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		policyPermissionsRepository.
			On("DeletePolicyPermissions",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionIds := []string{"739bbbc9-7e93-11ee-89fd-0242ac110016"}
		err := policyPermissionsUCase.DeletePolicyPermissions(context.Background(), policyId, policyPermissionIds)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
	})

	t.Run("When an error occurs while deleting multiple policy permissions", func(t *testing.T) {
		policyPermissionsRepository := &mockPolicyPermissions.PolicyPermissionRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		policyPermissionsError := errors.New("random error")
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(false, nil)
		policyPermissionsRepository.
			On("DeletePolicyPermissions",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(policyPermissionsError)
		policyPermissionsUCase := NewPolicyPermissionsUseCase(
			policyPermissionsRepository,
			validationRepository,
			authRepository,
			60)
		policyId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyPermissionIds := []string{"739bbbc9-7e93-11ee-89fd-0242ac110016"}
		err := policyPermissionsUCase.DeletePolicyPermissions(context.Background(), policyId, policyPermissionIds)
		assert.Error(t, err)
	})
}
