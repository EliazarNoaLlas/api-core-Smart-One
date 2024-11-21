/*
 * File: receipt_types_func_mysql_repository.go
 * Author: euridice
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the func mysql repository of the receiptTypes.
 *
 * Last Modified: 2024-03-06
 */

package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	"gitlab.smartcitiesperu.com/smartone/api-shared/db"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"

	receiptTypesDomain "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/domain"
)

//go:embed sql/get_receipt_types.sql
var QueryGetReceiptTypes string

//go:embed sql/update_receipt_type.sql
var QueryUpdateReceiptType string

//go:embed sql/create_receipt_type.sql
var QueryCreateReceiptType string

//go:embed sql/delete_receipt_type.sql
var QueryDeleteReceiptType string

func (r ReceiptTypesMySQLRepo) GetReceiptTypes(
	ctx context.Context,
) (
	ReceiptTypeRows []receiptTypesDomain.ReceiptType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetReceiptTypes").SetRaw(err)
	}
	results, err := client.QueryContext(
		ctx,
		QueryGetReceiptTypes)

	if err != nil {
		return nil, r.err.Clone().SetFunction("GetReceiptTypes").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	ReceiptTypeTmp := make([]ReceiptType, 0)
	err = carta.Map(results, &ReceiptTypeTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetReceiptTypes").SetRaw(err)
	}
	automapper.Map(ReceiptTypeTmp, &ReceiptTypeRows)
	return ReceiptTypeRows, nil
}

func (r ReceiptTypesMySQLRepo) CreateReceiptType(
	ctx context.Context,
	receiptTypeId string,
	userId string,
	body receiptTypesDomain.CreateReceiptTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateReceiptType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateReceiptType,
		receiptTypeId,
		body.Description,
		body.SunatCode,
		body.Enable,
		userId,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateReceiptType").SetRaw(err)
	}
	return
}

func (r ReceiptTypesMySQLRepo) UpdateReceiptType(
	ctx context.Context,
	receiptTypeId string,
	body receiptTypesDomain.UpdateReceiptTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateReceiptType").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateReceiptType,
		body.Description,
		body.SunatCode,
		body.Enable,
		receiptTypeId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateReceiptType").SetRaw(err)
	}
	return
}

func (r ReceiptTypesMySQLRepo) DeleteReceiptType(
	ctx context.Context,
	ReceiptTypeId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteReceiptType").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteReceiptType,
		now,
		ReceiptTypeId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteReceiptType").SetRaw(err)
	}
	return true, nil
}
