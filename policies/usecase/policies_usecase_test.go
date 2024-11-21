/*
 * File: policies_usecase_test.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Unit tests to use case of policies.
 *
 * Last Modified: 2023-11-14
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

	policiesDomain "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain"
	mockPolicies "gitlab.smartcitiesperu.com/smartone/api-core/policies/domain/mocks"
)

func TestUseCasePolicies_GetPolicies(t *testing.T) {
	t.Run("When policies are successfully listed", func(t *testing.T) {
		policiesRepository := &mockPolicies.PolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		total := 10
		policiesRepository.
			On("GetPolicies", mock.Anything, mock.Anything, mock.Anything).
			Return([]policiesDomain.Policy{}, nil)
		policiesRepository.
			On("GetTotalPolicies", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		policiesUCase := NewPoliciesUseCase(
			policiesRepository,
			validationRepository,
			authRepository,
			60,
		)
		searchParams := policiesDomain.GetPoliciesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		policies, _, err := policiesUCase.GetPolicies(context.Background(), searchParams, pagination)
		assert.NoError(t, err)
		assert.EqualValues(t, policies, []policiesDomain.Policy{})
	})

	t.Run("When an error occurs while listing policies", func(t *testing.T) {
		policiesRepository := &mockPolicies.PolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		total := 10
		policiesRepository.
			On("GetPolicies", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		policiesRepository.
			On("GetTotalPolicies", mock.Anything, mock.Anything, mock.Anything).
			Return(&total, nil)
		policiesUCase := NewPoliciesUseCase(
			policiesRepository,
			validationRepository,
			authRepository,
			60,
		)
		searchParams := policiesDomain.GetPoliciesParams{}
		pagination := paramsDomain.NewPaginationParams(nil)
		policies, _, err := policiesUCase.GetPolicies(context.Background(), searchParams, pagination)
		assert.Error(t, err)
		assert.EqualValues(t, policies, []policiesDomain.Policy(nil))
	})
}

func TestUseCasePolicies_CreatePolicy(t *testing.T) {
	t.Run("When to successfully create a policy", func(t *testing.T) {
		policiesRepository := &mockPolicies.PolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		policyID := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		policiesRepository.
			On("CreatePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(&policyID, nil)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		policiesUCase := NewPoliciesUseCase(
			policiesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := policiesUCase.CreatePolicy(
			context.Background(),
			policiesDomain.CreatePolicyBody{},
		)
		assert.NoError(t, err)
	})

	t.Run("When an error occurs while creating a policy and the policy already exists", func(t *testing.T) {
		policiesRepository := &mockPolicies.PolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}
		policiesRepository.
			On("CreatePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("random error"))
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(true, nil)
		policiesUCase := NewPoliciesUseCase(
			policiesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := policiesUCase.CreatePolicy(
			context.Background(),
			policiesDomain.CreatePolicyBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, policiesDomain.ErrPolicyNameAlreadyExistCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreatePolicy")
	})

	t.Run("When an error occurs while creating a policy", func(t *testing.T) {
		policiesRepository := &mockPolicies.PolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		errCreate := errDomain.NewErr().SetFunction("CreatePolicy").
			SetLayer(errDomain.UseCase).
			SetRaw(errors.New("random error"))
		policiesRepository.
			On("CreatePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errCreate)
		validationRepository.
			On("ValidateExistence", mock.Anything, mock.Anything).
			Return(false, nil)
		policiesUCase := NewPoliciesUseCase(
			policiesRepository,
			validationRepository,
			authRepository,
			60,
		)
		_, err := policiesUCase.CreatePolicy(
			context.Background(),
			policiesDomain.CreatePolicyBody{},
		)
		assert.Error(t, err)

		var smartErr *errDomain.SmartError
		ok := errors.As(err, &smartErr)
		assert.Equal(t, ok, true)
		assert.Equal(t, smartErr.Code, errDomain.ErrUnknownCode)
		assert.Equal(t, smartErr.Layer, errDomain.UseCase)
		assert.Equal(t, smartErr.Function, "CreatePolicy")
	})
}

func TestUseCasePolicies_UpdatePolicy(t *testing.T) {
	t.Run("When a policy is successfully updated", func(t *testing.T) {
		policiesRepository := &mockPolicies.PolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		policiesRepository.
			On("UpdatePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		policiesUCase := NewPoliciesUseCase(
			policiesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := policiesUCase.UpdatePolicy(
			context.Background(),
			policiesDomain.UpdatePolicyBody{},
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
		)
		assert.NoError(t, err)
	})

	t.Run("When an error occurs while updating a policy", func(t *testing.T) {
		policiesRepository := &mockPolicies.PolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		policiesRepository.
			On("UpdatePolicy", mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("random error"))
		policiesUCase := NewPoliciesUseCase(
			policiesRepository,
			validationRepository,
			authRepository,
			60,
		)
		err := policiesUCase.UpdatePolicy(
			context.Background(),
			policiesDomain.UpdatePolicyBody{},
			"739bbbc9-7e93-11ee-89fd-0242ac110016",
		)
		assert.Error(t, err)
	})
}

func TestUseCasePolicies_DeletePolicy(t *testing.T) {
	t.Run("When a policy is successfully deleted", func(t *testing.T) {
		policiesRepository := &mockPolicies.PolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		policiesRepository.
			On("DeletePolicy", mock.Anything, mock.Anything).
			Return(true, nil)
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(true, nil)
		policiesUCase := NewPoliciesUseCase(
			policiesRepository,
			validationRepository,
			authRepository,
			60,
		)
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := policiesUCase.DeletePolicy(context.Background(), policyId)
		if err != nil {
			t.Errorf("this is the error getting the registers: %v\n", err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("When an error occurs while deleting a policy", func(t *testing.T) {
		policiesRepository := &mockPolicies.PolicyRepository{}
		validationRepository := &mockValidation.ValidationRepository{}
		authRepository := &mockAuth.AuthRepository{}

		policiesError := errors.New("random error")
		policiesRepository.
			On("DeletePolicy", mock.Anything, mock.Anything).
			Return(false, policiesError)
		validationRepository.
			On("RecordExists", mock.Anything, mock.Anything).
			Return(false, nil)
		policiesUCase := NewPoliciesUseCase(
			policiesRepository,
			validationRepository,
			authRepository,
			60,
		)
		policyId := "739bbbc9-7e93-11ee-89fd-0242ac110016"
		res, err := policiesUCase.DeletePolicy(context.Background(), policyId)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
