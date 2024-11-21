/*
 * File: modules_mysql_repository_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to repository for modules
 *
 * Last Modified: 2023-11-10
 */

package mysql

import "time"

type module struct {
	Id          string     `db:"id" `
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Code        string     `db:"code"`
	Icon        string     `db:"icon"`
	Position    int        `db:"position"`
	CreatedAt   *time.Time `db:"created_at"`
}
