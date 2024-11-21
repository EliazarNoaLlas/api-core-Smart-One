/*
 * File: users_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the UsersUserRepository interface for users data operations.
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"context"
	"database/sql"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type UserRepository interface {
	GetUser(ctx context.Context, userId string) (*User, error)
	GetUsers(ctx context.Context, searchParams GetUsersParams, pagination paramsDomain.PaginationParams) (
		[]UserMultiple, error)
	GetTotalUsers(ctx context.Context, searchParams GetUsersParams, pagination paramsDomain.PaginationParams) (
		*int, error)
	GetMenuByUser(ctx context.Context, userId string) ([]ModuleMenuUser, error)
	GetMeByUser(ctx context.Context, userId string) (*UserMeInfo, error)
	GetStoresByUser(ctx context.Context, userId string) ([]StoreByUser, error)
	GetMerchantsByUser(ctx context.Context, userId string) ([]MerchantByUser, error)
	CreateUser(ctx context.Context, tx *sql.Tx, userId string, body CreateUserBody) (*string, error)
	CreateUserMain(ctx context.Context, userId string, personId string, body CreateUserBody) (*string, error)
	CreatePerson(ctx context.Context, tx *sql.Tx, userId string, PersonId string, body *Person) (*string, error)
	UpdatePerson(ctx context.Context, tx *sql.Tx, personId string, userId string, body *Person) error
	UpdateUserMain(ctx context.Context, userId string, personId string, body UpdateUserBody) error
	GetUserById(ctx context.Context, tx *sql.Tx, userId string) (*UserById, error)
	UpdateUser(ctx context.Context, userId string, body UpdateUserBody) error
	DeleteUser(ctx context.Context, userId string) (bool, error)
	ResetPasswordUser(ctx context.Context, userId string, passwordHash string) (bool, error)
	GetUserByUserNameAndPassword(ctx context.Context, userName string, passwordHash string) (*User, *string, error)
	VerifyIfPersonExist(ctx context.Context, personId string) error
	VerifyIfUserExist(ctx context.Context, userId string) error
	UpdatePersonToUser(ctx context.Context, tx *sql.Tx, peopleId string, userId string) error
	ValidateUniquePersonByDocument(ctx context.Context, typeDocumentId string, document string) error
	ValidateUniqueUserExistence(ctx context.Context, tx *sql.Tx, userId string) error
	VerifyPermissionsByUser(ctx context.Context, userId string, storeId string, codePermission string) (bool, error)
	GetModulePermissions(ctx context.Context, userId string, codeModule string) ([]Permissions, error)
	GetModules(ctx context.Context) ([]Module, error)
}
