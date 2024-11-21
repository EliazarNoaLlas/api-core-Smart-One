/*
 * File: validations_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the repository of the validations.
 *
 * Last Modified: 2023-11-10
 */

package domain

import "context"

type ValidationRepository interface {
	RecordExists(ctx context.Context, params RecordExistsParams) error
	ValidateExistence(ctx context.Context, params RecordExistsParams) (bool, error)
}
