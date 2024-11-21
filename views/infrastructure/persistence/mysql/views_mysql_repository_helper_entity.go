/*
 * File: views_mysql_repository_helper_entity.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to repository for view
 *
 * Last Modified: 2023-11-24
 */

package mysql

import (
	"time"
)

type View struct {
	Id          string     `db:"view_id"`
	Name        string     `db:"view_name"`
	Description string     `db:"view_description"`
	Url         string     `db:"view_url"`
	Icon        string     `db:"view_icon"`
	CreatedAt   *time.Time `db:"view_created_at"`
}
