/*
 * File: view_permissions_helper_entity.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities of the viewPermissions.
 *
 * Last Modified: 2024-02-26
 */

package mysql

import "time"

type ViewPermission struct {
	Id         string     `db:"view_permission_id"`
	CreatedBy  string     `db:"view_permission_created_by"`
	CreatedAt  *time.Time `db:"view_permission_created_at"`
	View       View
	Permission Permission
}

type View struct {
	Id          string     `db:"view_id"`
	Name        string     `db:"view_name"`
	Description string     `db:"view_description"`
	CreatedAt   *time.Time `db:"view_created_at"`
}

type Permission struct {
	Id          string     `db:"permission_id"`
	Code        string     `db:"permission_code"`
	Name        string     `db:"permission_name"`
	Description string     `db:"permission_description"`
	CreatedAt   *time.Time `db:"permission_created_at"`
	Module      Module
}

type Module struct {
	Id          string     `db:"module_id"`
	Name        string     `db:"module_name"`
	Description string     `db:"module_description"`
	Code        string     `db:"module_code"`
	Icon        string     `db:"module_icon"`
	Position    int        `db:"module_position"`
	CreatedAt   *time.Time `db:"module_created_at"`
}
