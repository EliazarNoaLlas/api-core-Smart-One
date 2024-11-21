/*
 * File: receipt_types_helper_entity.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities of the receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package mysql

import "time"

type ReceiptType struct {
	Id          string     `db:"receipt_type_id"`
	Description string     `db:"receipt_type_description"`
	SunatCode   string     `db:"receipt_type_sunat_code"`
	Enable      bool       `db:"receipt_type_enable"`
	CreatedBy   string     `db:"receipt_type_created_by"`
	CreatedAt   *time.Time `db:"receipt_type_created_at"`
}
