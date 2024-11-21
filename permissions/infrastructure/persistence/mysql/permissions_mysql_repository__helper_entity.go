/*
 * File: permissions_mysql_repository_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the permission entity.
 *
 * Last Modified: 2023-11-15
 */

package mysql

import (
	"time"
)

type ModuleByPermission struct {
	Id          string `db:"module_id" `
	Name        string `db:"module_name"`
	Description string `db:"module_description"`
	Code        string `db:"module_code"`
}

type Permission struct {
	Id          string     `db:"permission_id" `
	Code        string     `db:"permission_code"`
	Name        string     `db:"permission_name"`
	Description string     `db:"permission_description"`
	CreatedAt   *time.Time `db:"permission_created_at"`
	Module      ModuleByPermission
}
