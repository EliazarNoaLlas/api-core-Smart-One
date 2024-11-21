/*
 * File: receipt_types_handler_helper_entity.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entities helper to handler for receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package rest

import (
	receiptTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
)

type receiptTypesResult struct {
	Data   []receiptTypesDomain.ReceiptType `json:"data" binding:"required"`
	Status int                              `json:"status" binding:"required"`
}

type deleteReceiptTypesResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
