/*
 * File: receipt_types_repository.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the repository of the receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package domain

import (
	"context"
)

type ReceiptTypesRepository interface {
	GetReceiptTypes(ctx context.Context) ([]ReceiptType, error)
	CreateReceiptType(ctx context.Context, receiptTypeId string, userId string, body CreateReceiptTypeBody) error
	UpdateReceiptType(ctx context.Context, receiptTypeId string, body UpdateReceiptTypeBody) error
	DeleteReceiptType(ctx context.Context, receiptTypeId string) (bool, error)
}
