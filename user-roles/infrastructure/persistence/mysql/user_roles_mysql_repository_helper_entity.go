/*
 * File: user_roles_mysql_repository_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to repository for userRoles
 *
 * Last Modified: 2023-11-23
 */

package mysql

import (
	"time"
)

type UserRole struct {
	Id        string     `db:"user_role_id"`
	Enable    bool       `db:"user_role_enable"`
	CreatedAt *time.Time `db:"user_role_created_at"`
	Roles     Role       `db:"roles"`
}

type Role struct {
	Id          string     `db:"role_id"`
	Name        string     `db:"role_name"`
	Description string     `db:"role_description"`
	Enable      bool       `db:"role_enable"`
	CreatedAt   *time.Time `db:"role_created_at"`
}
