/*
 * File: validations_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the functions of the validations.
 *
 * Last Modified: 2023-11-10
 */

package mysql

import (
	"context"
	"fmt"
	"time"

	"gitlab.smartcitiesperu.com/smartone/api-shared/db"

	"gitlab.smartcitiesperu.com/smartone/api-core/validations/domain"
)

type validationsMySQLRepo struct {
	timeout time.Duration
}

func NewValidationsRepository(mongoTimeout int) domain.ValidationRepository {
	rep := &validationsMySQLRepo{
		timeout: time.Duration(mongoTimeout) * time.Second,
	}
	return rep
}

func (r validationsMySQLRepo) RecordExists(
	ctx context.Context,
	params domain.RecordExistsParams,
) error {
	var exists int
	var query string
	var args []interface{}

	query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", params.Table, params.IdColumnName)
	args = append(args, params.IdValue)

	if params.StatusColumnName != nil && params.StatusValue != nil {
		query += fmt.Sprintf(" AND %s = ?", *params.StatusColumnName)
		args = append(args, *params.StatusValue)
	}

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return err
	}
	err = client.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		return err
	}
	if exists == 0 {
		return err
	}

	return nil
}

func (r validationsMySQLRepo) ValidateExistence(
	ctx context.Context,
	params domain.RecordExistsParams,
) (bool, error) {
	var exists int
	var query string
	var args []interface{}

	query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", params.Table, params.IdColumnName)
	args = append(args, params.IdValue)

	if params.StatusColumnName != nil && params.StatusValue != nil {
		query += fmt.Sprintf(" AND %s = ?", *params.StatusColumnName)
		args = append(args, *params.StatusValue)
	}

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, err
	}
	err = client.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists != 0 {
		return true, nil
	}
	return false, nil
}
