/*
 * File: merchants_mysql_repository_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the merchantModel entity.
 *
 * Last Modified: 2023-11-10
 */

package mysql

import "time"

type merchantModel struct {
	Id          string     `db:"id" `
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Phone       string     `db:"phone"`
	Document    string     `db:"document"`
	Address     string     `db:"address"`
	Industry    string     `db:"industry"`
	ImagePath   string     `db:"image_path"`
	CreatedAt   *time.Time `db:"created_at"`
}
