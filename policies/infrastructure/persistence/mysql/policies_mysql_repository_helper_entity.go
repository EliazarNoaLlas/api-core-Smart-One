/*
 * File: policies_mysql_repository_helper_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to repository for policies
 *
 * Last Modified: 2023-11-14
 */

package mysql

import (
	"time"
)

type ModuleByPolicy struct {
	Id          *string `db:"module_id" `
	Name        *string `db:"module_name"`
	Description *string `db:"module_description"`
	Code        *string `db:"module_code"`
}

type MerchantByPolicy struct {
	Id          *string `db:"merchant_id" `
	Name        *string `db:"merchant_name"`
	Description *string `db:"merchant_description"`
	Document    *string `db:"merchant_document"`
}

type StoreByPolicy struct {
	Id        *string `db:"store_id" `
	Name      *string `db:"store_name"`
	Shortname *string `db:"store_shortname"`
}

type PolicyPermissionByPolicy struct {
	Id     *string `db:"policy_permission_id"`
	Enable *bool   `db:"policy_permission_enable"`
}

type PermissionByPolicy struct {
	Id               *string `db:"permission_id"`
	Code             *string `db:"permission_code"`
	Name             *string `db:"permission_name"`
	Description      *string `db:"permission_description"`
	PolicyPermission *PolicyPermissionByPolicy
}

type Policy struct {
	Id          string     `db:"policy_id" `
	Name        string     `db:"policy_name"`
	Description string     `db:"policy_description"`
	Level       string     `db:"policy_level"`
	Enable      bool       `db:"policy_enable"`
	CreatedAt   *time.Time `db:"policy_created_at"`
	Module      *ModuleByPolicy
	Merchant    *MerchantByPolicy
	Store       *StoreByPolicy
	Permissions []PermissionByPolicy
}
