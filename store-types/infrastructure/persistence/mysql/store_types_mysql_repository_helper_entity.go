/*
 * File: store_types_mysql_repository_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the storeTypeModel entity.
 *
 * Last Modified: 2023-11-10
 */

package mysql

type storeType struct {
	Id           string `db:"id" `
	Description  string `db:"description"`
	Abbreviation string `db:"abbreviation"`
}
