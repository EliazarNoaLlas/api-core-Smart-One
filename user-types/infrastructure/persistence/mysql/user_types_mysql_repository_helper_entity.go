/*
 * File: user_types_mysql_repository_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to repository for userType
 *
 * Last Modified: 2023-11-23
 */

package mysql

import (
	"time"
)

type userType struct {
	Id          string     `db:"id" `
	Description string     `db:"description"`
	Code        string     `db:"code"`
	Enable      bool       `db:"enable"`
	CreatedAt   *time.Time `db:"created_at"`
}
