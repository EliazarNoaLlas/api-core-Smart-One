/*
 * File: role_policies_mysql_repository_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to repository for rolePolicies
 *
 * Last Modified: 2023-11-22
 */

package mysql

import (
	"time"
)

type PolicyByRolePolicy struct {
	Id          string     `db:"policy_id"`
	Name        string     `db:"policy_name"`
	Description string     `db:"policy_description"`
	Level       string     `db:"policy_level"`
	Enable      bool       `db:"policy_enable"`
	CreatedAt   *time.Time `db:"policy_created_at"`
}

type RolePolicy struct {
	Id        string     `db:"role_policy_id"`
	Enable    bool       `db:"role_policy_enable"`
	CreatedAt *time.Time `db:"role_policy_created_at"`
	Policy    PolicyByRolePolicy
}
