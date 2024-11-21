/*
 * File: users_mysql_repository_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to repository for users
 *
 * Last Modified: 2023-11-23
 */

package mysql

import (
	"time"
)

type UserTypeByUser struct {
	Id          string `db:"user_type_id" `
	Description string `db:"user_type_description"`
	Code        string `db:"user_type_code"`
}

type User struct {
	Id        string     `db:"user_id" `
	UserName  string     `db:"user_name"`
	CreatedAt *time.Time `db:"user_created_at"`
	UserType  UserTypeByUser
}

type UserMultiple struct {
	Id        string     `db:"user_id" `
	UserName  string     `db:"user_name"`
	CreatedAt *time.Time `db:"user_created_at"`
	UserType  UserTypeByUser
	Role      []Role
}

type Role struct {
	Id          *string    `db:"role_id" `
	Name        *string    `db:"role_name"`
	Description *string    `db:"role_description"`
	Enable      *bool      `db:"role_enable"`
	CreatedAt   *time.Time `db:"role_created_at"`
	UserRole    UserRole
}

type UserRole struct {
	Id *string `db:"user_role_id"`
}

type ViewMenuUser struct {
	Id          string     `db:"view_id"`
	Name        string     `db:"view_name"`
	Description string     `db:"view_description"`
	Url         string     `db:"view_url"`
	Icon        string     `db:"view_icon"`
	CreatedAt   *time.Time `db:"view_created_at"`
}

type ModuleMenuUser struct {
	Id          string     `db:"module_id"`
	Name        string     `db:"module_name"`
	Description string     `db:"module_description"`
	Code        string     `db:"module_code"`
	Icon        string     `db:"module_icon"`
	Position    int        `db:"module_position"`
	CreatedAt   *time.Time `db:"module_created_at"`
	Views       []ViewMenuUser
}

type UserById struct {
	Id        string     `db:"id"`
	UserName  string     `db:"username"`
	CreatedAt *time.Time `db:"created_at"`
}

type UserMe struct {
	Id        string     `db:"user_id" `
	UserName  string     `db:"user_name"`
	CreatedAt *time.Time `db:"user_created_at"`
	Person    *PersonByUser
	RoleUser  []RoleUser
}

type RoleUser struct {
	Id          string     `db:"role_id"`
	Name        *string    `db:"role_name"`
	Description *string    `db:"role_description"`
	Enable      *bool      `db:"role_enable"`
	CreateAt    *time.Time `db:"role_created_at"`
}

type PersonByUser struct {
	Id           *string    `db:"person_id"`
	Document     *string    `db:"person_document"`
	Names        *string    `db:"person_names"`
	Surname      *string    `db:"person_surname"`
	LastName     *string    `db:"person_last_name"`
	Phone        *string    `db:"person_phone"`
	Email        *string    `db:"person_email"`
	Gender       *string    `db:"person_gender"`
	Enable       *bool      `db:"person_enable"`
	CreatedAt    *time.Time `db:"person_created_at"`
	TypeDocument *TypeDocument
}

type TypeDocument struct {
	Id                    *string    `db:"document_type_id"`
	Number                *string    `db:"document_type_number"`
	Description           *string    `db:"document_type_description" `
	AbbreviateDescription *string    `db:"document_type_abbreviated_description" `
	Enable                *bool      `db:"document_type_enable"`
	CreateAt              *time.Time `db:"document_type_created_at"`
}

type StoreByUser struct {
	Id       string `db:"store_id"`
	Name     string `db:"store_name"`
	Merchant Merchant
}

type Merchant struct {
	Id          string `db:"merchant_id"`
	Name        string `db:"merchant_name"`
	Description string `db:"merchant_description"`
	ImagePath   string `db:"merchant_image_path"`
}

type MerchantByUser struct {
	Id          string `db:"merchant_id"`
	Name        string `db:"merchant_name"`
	Description string `db:"merchant_description"`
	ImagePath   string `db:"merchant_image_path"`
	Stores      []Store
}

type Store struct {
	Id   string `db:"store_id"`
	Name string `db:"store_name"`
}

type Permissions struct {
	Id   string `db:"id"`
	Code string `db:"code"`
}

type Module struct {
	Id          string     `db:"id" `
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Code        string     `db:"code"`
	Icon        string     `db:"icon"`
	Position    int        `db:"position"`
	CreatedAt   *time.Time `db:"created_at"`
}
