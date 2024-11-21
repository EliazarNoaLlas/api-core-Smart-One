/*
 * File: stores_mysql_repository_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to repository for stores
 *
 * Last Modified: 2023-11-14
 */

package mysql

import "time"

type StoreTypeByStore struct {
	Id           string `db:"type_id" `
	Description  string `db:"type_description"`
	Abbreviation string `db:"type_abbreviation"`
}

type Store struct {
	Id         string     `db:"store_id" `
	Name       string     `db:"store_name"`
	Shortname  string     `db:"store_shortname"`
	MerchantId string     `db:"store_merchant_id"`
	CreatedAt  *time.Time `db:"store_created_at"`
	StoreType  StoreTypeByStore
}
