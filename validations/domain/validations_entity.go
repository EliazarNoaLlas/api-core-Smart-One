/*
 * File: validations_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implements the entity of the validations.
 *
 * Last Modified: 2023-11-10
 */

package domain

type RecordExistsParams struct {
	Table            string
	IdColumnName     string
	IdValue          interface{}
	StatusColumnName *string
	StatusValue      *int
}
