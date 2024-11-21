/*
 * File: server_usecase_test.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the use case test of the server
 *
 * Last Modified: 2024-04-09
 */

package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemUseCase_GetServerGetServerDate(t *testing.T) {
	t.Run(
		"Get get date time by series and number successfully", func(t *testing.T) {
			serverAuxUseCase := NewServerUseCase(60)
			res, err := serverAuxUseCase.GetServerDate(context.Background())
			assert.NoError(t, err)
			assert.NotNil(t, res)
		},
	)
	t.Run(
		"Get get date time by series and number with error", func(t *testing.T) {
			serverAuxUseCase := NewServerUseCase(60)
			res, err := serverAuxUseCase.GetServerDate(context.Background())
			assert.NoError(t, err)
			assert.NotNil(t, res)
		},
	)
}
