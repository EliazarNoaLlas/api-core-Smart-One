/*
 * File: policyPermissions_mysql_repository_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to repository for policyPermissions
 *
 * Last Modified: 2023-11-20
 */

package mysql

import (
	"time"
)

type PermissionPolicy struct {
	Id         string      `db:"policy_permission_id"`
	Enable     int         `db:"policy_permission_enable"`
	CreatedAt  *time.Time  `db:"policy_permission_created_at"`
	Permission Permissions `db:"permission"`
}

type Permissions struct {
	Id          string     `db:"permissions_id"`
	Code        string     `db:"permissions_code"`
	Name        string     `db:"permissions_name"`
	Description string     `db:"permissions_description"`
	CreatedAt   *time.Time `db:"permissions_created_at"`
}
