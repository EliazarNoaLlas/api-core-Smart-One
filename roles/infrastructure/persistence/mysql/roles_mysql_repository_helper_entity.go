/*
 * File: roles_mysql_repository_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the roleModel entity.
 *
 * Last Modified: 2023-11-14
 */

package mysql

import (
	"time"
)

type RoleModel struct {
	Id          string     `db:"id" `
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Enable      bool       `db:"enable"`
	CreatedAt   *time.Time `db:"created_at"`
}
