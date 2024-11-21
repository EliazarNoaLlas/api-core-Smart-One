/*
 * File: role_policies_usecase_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of rolePolicies.
 *
 * Last Modified: 2023-11-23
 */

package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	mockAuth "gitlab.smartcitiesperu.com/smartone/api-shared/auth/domain/mocks"
	errDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	mockValidation "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain/mocks"

	rolePoliciesDomain "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain"
	mockRolePolicies "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/domain/mocks"
)

func TestUseCaseRolePolicies_GetPolicies(t *testing.T) {
	t.Run("When policies by role are successfully listed", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		total := 10
		rolePoliciesRepository.
			On("GetPolicies", mock.Anything, mock.Anything, mock.Anything).
			Return([]rolePoliciesDomain.RolePolicy{}, nil)
		rolePoliciesRepository.
			On("GetTotalPolicies", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		policiesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60,
		)
		searchParams := rolePoliciesDomain.GetRolePoliciesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		policies, _, err := policiesUCase.GetPolicies(context.Background(), searchParams, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, policies, []rolePoliciesDomain.RolePolicy{})
	})

	t.Run("When an error occurs while listing policies", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		total := 10
		rolePoliciesRepository.
			On("GetPolicies", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		rolePoliciesRepository.
			On("GetTotalPolicies", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		policiesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60,
		)
		searchParams := rolePoliciesDomain.GetRolePoliciesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		policies, _, err := policiesUCase.GetPolicies(context.Background(), searchParams, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, policies, []rolePoliciesDomain.RolePolicy(nil))
	})
}

func TestUseCaseRolePolicies_CreateRolePolicy(t *testing.T) {
	roleHasPolicy := false
	t.Run("When add a policy to role successfully", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		rolePoliciesRepository.
			On("CreateRolePolicy", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&rolePolicyId, nil)
		rolePoliciesRepository.
			On("VerifyRoleHasPolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(roleHasPolicy, nil)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := rolePoliciesUCase.CreateRolePolicy(
			context.Background(),
			roleId,
			rolePoliciesDomain.CreateRolePolicyBody{
				PolicyId: policyId,
				Enable:   true,
			},
		)
		assert.NoError(t, err)
	})

	t.Run("When add a policy to role, error", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		rolePoliciesRepository.
			On("CreateRolePolicy", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		rolePoliciesRepository.
			On("VerifyRoleHasPolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := rolePoliciesUCase.CreateRolePolicy(
			context.Background(),
			roleId,
			rolePoliciesDomain.CreateRolePolicyBody{
				PolicyId: policyId,
				Enable:   true,
			},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, rolePoliciesDomain.ErrRoleAlreadyHasThePolicyCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateRolePolicy")
	})

	t.Run("When add a policy to role, error", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		errCreate := errDomain.NewErr().SetFunction("CreateRolePolicy").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		rolePoliciesRepository.
			On("CreateRolePolicy", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		rolePoliciesRepository.
			On("VerifyRoleHasPolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(false, nil)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := rolePoliciesUCase.CreateRolePolicy(
			context.Background(),
			roleId,
			rolePoliciesDomain.CreateRolePolicyBody{
				PolicyId: policyId,
				Enable:   true,
			},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateRolePolicy")
	})
}

func TestUseCasePolicyPermissions_CreateRolePolicies(t *testing.T) {
	roleHasPolicy := false
	t.Run("When create multiple role policies successfully", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		body := []rolePoliciesDomain.CreateRolePolicyBody{
			{
				PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
				Enable:   true,
			},
		}
		rolePoliciesRepository.
			On("CreateRolePolicies",
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).
			Return(nil)
		rolePoliciesRepository.
			On("VerifyRoleHasPolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(roleHasPolicy, nil)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60)
		_, err := rolePoliciesUCase.CreateRolePolicies(
			context.Background(),
			roleId,
			body,
		)
		assert.NoError(t, err)
	})

	t.Run("When an error occurs while creating multiple role policies", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		body := []rolePoliciesDomain.CreateRolePolicyBody{
			{
				PolicyId: "739bbbc9-7e93-11ee-89fd-042hs5278420",
				Enable:   true,
			},
		}
		errCreate := errDomain.NewErr().SetFunction("CreateRolePolicies").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePoliciesRepository.
			On("CreateRolePolicies",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(errCreate)
		rolePoliciesRepository.
			On("VerifyRoleHasPolicy",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(roleHasPolicy, nil)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60)
		_, err := rolePoliciesUCase.CreateRolePolicies(
			context.Background(),
			roleId,
			body,
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreateRolePolicies")
	})
}

func TestUseCaseRolePolicies_UpdateRolePolicy(t *testing.T) {
	t.Run("When update a policy of role successfully", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		rolePoliciesRepository.
			On("UpdateRolePolicy", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60,
		)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110017"
		err := rolePoliciesUCase.UpdateRolePolicy(
			context.Background(),
			roleId,
			rolePolicyId,
			rolePoliciesDomain.UpdateRolePolicyBody{
				PolicyId: policyId,
				Enable:   true,
			},
		)
		assert.NoError(t, err)
	})

	t.Run("When update a policy of role, error", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		rolePoliciesRepository.
			On("UpdateRolePolicy", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60,
		)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		err := rolePoliciesUCase.UpdateRolePolicy(
			context.Background(),
			roleId,
			rolePolicyId,
			rolePoliciesDomain.UpdateRolePolicyBody{
				PolicyId: policyId,
				Enable:   true,
			},
		)
		assert.Error(t, err)
	})
}

func TestUseCaseRolePolicies_DeleteRolePolicy(t *testing.T) {
	t.Run("When delete a policy of role successfully", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		rolePoliciesRepository.
			On("DeleteRolePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(true, nil)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60,
		)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := rolePoliciesUCase.DeleteRolePolicy(context.Background(), roleId, rolePolicyId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When delete a policy of role, error", func(t *testing.T) {
		authRepository := &mockAuth.AuthRepository{}
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		rolePoliciesError := errors.New("random error")
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		rolePoliciesRepository.
			On("DeleteRolePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(false, rolePoliciesError)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60,
		)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := rolePoliciesUCase.DeleteRolePolicy(context.Background(), roleId, rolePolicyId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}

func TestUseCasePolicyPermissions_DeleteRolePolicies(t *testing.T) {
	t.Run("When delete multiple role policies successfully", func(t *testing.T) {
		rolePolicyRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(true, nil)
		rolePolicyRepository.
			On("DeleteRolePolicies",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(nil)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePolicyRepository,
			validationRepository,
			authRepository,
			60)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyIds := []string{"739bbbc9-7e93-11ee-89fd-0242ac110016"}
		err := rolePoliciesUCase.DeleteRolePolicies(context.Background(), roleId, rolePolicyIds)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
	})

	t.Run("When an error occurs while deleting multiple role policies", func(t *testing.T) {
		rolePoliciesRepository := &mockRolePolicies.RolePolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		policyPermissionsError := errors.New("random error")
		validationRepository.
			On("RecordExists",
				mock.Anything,
				mock.Anything).
			Return(false, nil)
		rolePoliciesRepository.
			On("DeleteRolePolicies",
				mock.Anything,
				mock.Anything,
				mock.Anything).
			Return(policyPermissionsError)
		rolePoliciesUCase := NewRolePoliciesUseCase(
			rolePoliciesRepository,
			validationRepository,
			authRepository,
			60)
		roleId := "739bbbc9-7e93-11ee-89fd-0442ac210931"
		rolePolicyIds := []string{"739bbbc9-7e93-11ee-89fd-0242ac110016"}
		err := rolePoliciesUCase.DeleteRolePolicies(context.Background(), roleId, rolePolicyIds)
		assert.Error(t, err)
	})
}
