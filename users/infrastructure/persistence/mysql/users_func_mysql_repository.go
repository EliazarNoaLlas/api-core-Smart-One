/*
 * File: users_func_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for users
 *
 * Last Modified: 2023-11-23
 */

package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"strings"

	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	"gitlab.smartcitiesperu.com/smartone/api-shared/db"
	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	usersDomain "gitlab.smartcitiesperu.com/smartone/api-core/users/domain"
)

//go:embed sql/get_user.sql
var QueryGetUser string

//go:embed sql/get_users.sql
var QueryGetUsers string

//go:embed sql/get_user_by_id.sql
var QueryGetUserById string

//go:embed sql/get_stores_by_user.sql
var QueryGetStoresByUser string

//go:embed sql/get_total_users.sql
var QueryGetTotalUsers string

//go:embed sql/get_user_me.sql
var QueryGetMeUser string

//go:embed sql/get_menu.sql
var QueryGetMenu string

//go:embed sql/get_user_by_password.sql
var QueryGetUserByPassword string

//go:embed sql/update_user.sql
var QueryUpdateUser string

//go:embed sql/reset_password_user.sql
var QueryResetPasswordUser string

//go:embed sql/delete_user.sql
var QueryDeleteUser string

//go:embed sql/create_user.sql
var QueryCreateUser string

//go:embed sql/create_user_to_people.sql
var QueryCreatePerson string

//go:embed sql/verify_exist_person.sql
var QueryVerifyPersonIdExist string

//go:embed sql/update_user_id_of_person.sql
var QueryUpdatePersonToUser string

//go:embed sql/validate_unique_person_by_document.sql
var QueryValidateUniquePersonByDocument string

//go:embed sql/validate_unique_user.sql
var QueryValidateUniqueUserExistence string

//go:embed sql/update_person.sql
var QueryUpdatePerson string

//go:embed sql/verify_if_the_user_exist.sql
var QueryVerifyIfTheUserExist string

//go:embed sql/verify_permissions_by_user.sql
var QueryVerifyPermissionsByUser string

//go:embed sql/get_module_permissions.sql
var QueryGetModulePermissions string

//go:embed sql/get_modules.sql
var QueryGetModules string

func (r usersMySQLRepo) GetUser(
	ctx context.Context,
	userId string,
) (
	user *usersDomain.User,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUser").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetUser,
			userId,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUser").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	usersTmp := make([]User, 0)
	err = carta.Map(results, &usersTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUser").SetRaw(err)
	}
	var users = make([]usersDomain.User, 0)
	automapper.Map(usersTmp, &users)
	if len(users) == 0 {
		return nil, usersDomain.ErrUserNotFound
	}
	return &users[0], nil
}

func (r usersMySQLRepo) GetUsers(
	ctx context.Context,
	searchParams usersDomain.GetUsersParams,
	pagination paramsDomain.PaginationParams,
) (
	usersRows []usersDomain.UserMultiple,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	rolesIds := strings.Join(searchParams.RoleId, ",")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUsers").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetUsers,
			searchParams.UserTypeId,
			searchParams.UserTypeId,
			searchParams.UserName,
			searchParams.UserName,
			rolesIds,
			rolesIds,
			sizePage,
			offset,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUsers").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	usersTmp := make([]UserMultiple, 0)
	err = carta.Map(results, &usersTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUsers").SetRaw(err)
	}
	var roles = make([]usersDomain.UserMultiple, 0)
	automapper.Map(usersTmp, &roles)
	for iUser, user := range roles {
		usersRole := make([]usersDomain.Role, 0)
		for _, userRole := range user.Role {
			if userRole.Id != nil {
				usersRole = append(usersRole, userRole)
			}

		}
		roles[iUser].Role = usersRole
	}

	return roles, nil
}

func (r usersMySQLRepo) GetTotalUsers(
	ctx context.Context,
	searchParams usersDomain.GetUsersParams,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalUsers").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalUsers,
			searchParams.UserTypeId,
			searchParams.UserTypeId,
			searchParams.UserName,
			searchParams.UserName,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalUsers").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r usersMySQLRepo) GetMenuByUser(
	ctx context.Context,
	userId string,
) (
	user []usersDomain.ModuleMenuUser,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMenuByUser").SetRaw(err)
	}
	results, err := client.
		QueryContext(ctx, QueryGetMenu, userId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMenuByUser").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	modulesTmp := make([]ModuleMenuUser, 0)
	err = carta.Map(results, &modulesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMenuByUser").SetRaw(err)
	}
	var modules = make([]usersDomain.ModuleMenuUser, 0)
	automapper.Map(modulesTmp, &modules)
	return modules, nil
}

func (r usersMySQLRepo) GetMeByUser(
	ctx context.Context,
	userId string,
) (
	info *usersDomain.UserMeInfo,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMeByUser").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetMeUser,
			userId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMeByUser").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	userInfoTmp := make([]UserMe, 0)
	err = carta.Map(results, &userInfoTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMeByUser").SetRaw(err)
	}
	var userInfo = make([]usersDomain.UserMeInfo, 0)
	automapper.Map(userInfoTmp, &userInfo)
	if len(userInfo) == 0 {
		return nil, usersDomain.ErrUserNotFound
	}
	return &userInfo[0], nil
}

func (r usersMySQLRepo) CreateUser(
	ctx context.Context,
	tx *sql.Tx,
	userId string,
	body usersDomain.CreateUserBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")

	if tx != nil {
		_, err = tx.ExecContext(ctx,
			QueryCreateUser,
			userId,
			body.UserName,
			body.Password,
			body.UserTypeId,
			now)
	} else {
		var client *sql.DB
		client, _, err = db.ClientDB(ctx)
		if err != nil {
			return nil, r.err.Clone().SetFunction("CreateUser").SetRaw(err)
		}
		_, err = client.ExecContext(ctx,
			QueryCreateUser,
			userId,
			body.UserName,
			body.Password,
			body.UserTypeId,
			now)
	}

	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateUser").SetRaw(err)
	}
	lastId = &userId
	return
}

func (r usersMySQLRepo) CreateUserMain(
	ctx context.Context,
	userId string,
	personId string,
	user usersDomain.CreateUserBody,
) (lastId *string, err error) {
	var tx *sql.Tx
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateUserMain").SetRaw(err)
	}
	tx, err = client.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	_, err = r.CreateUser(ctx, tx, userId, user)
	if err != nil {
		return nil, err
	}
	if user.Person != nil {
		lastId, err = r.CreatePerson(ctx, tx, userId, personId, user.Person)
		if err != nil {
			return nil, err
		}
	} else {
		err = r.UpdatePersonToUser(ctx, tx, userId, personId)
		if err != nil {
			return nil, err
		}
	}

	err = r.ValidateUniqueUserExistence(ctx, tx, userId)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	lastId = &userId
	return
}

func (r usersMySQLRepo) CreatePerson(
	ctx context.Context,
	tx *sql.Tx,
	userId string,
	personId string,
	body *usersDomain.Person) (lastId *string, err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	if tx != nil {
		_, err = tx.ExecContext(ctx,
			QueryCreatePerson,
			personId,
			userId,
			body.TypeDocumentId,
			body.Document,
			body.Names,
			body.Surname,
			body.LastName,
			body.Phone,
			body.Email,
			body.Gender,
			body.Enable,
			now)
	} else {
		var client *sql.DB
		client, _, err = db.ClientDB(ctx)
		if err != nil {
			return nil, r.err.Clone().SetFunction("CreatePerson").SetRaw(err)
		}
		_, err = client.ExecContext(ctx,
			QueryCreatePerson,
			personId,
			userId,
			body.TypeDocumentId,
			body.Document,
			body.Names,
			body.Surname,
			body.LastName,
			body.Phone,
			body.Email,
			body.Gender,
			body.Enable,
			now)
	}

	if err != nil {
		return nil, r.err.Clone().SetFunction("CreatePerson").SetRaw(err)
	}
	lastId = &userId
	return
}

func (r usersMySQLRepo) UpdatePerson(
	ctx context.Context,
	tx *sql.Tx,
	personId string,
	userId string,
	body *usersDomain.Person,
) (err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	_, err = tx.ExecContext(
		ctx,
		QueryUpdatePerson,
		userId,
		body.TypeDocumentId,
		body.Document,
		body.Names,
		body.Surname,
		body.LastName,
		body.Phone,
		body.Email,
		body.Gender,
		body.Enable,
		personId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdatePerson").SetRaw(err)
	}
	return
}

func (r usersMySQLRepo) UpdateUserMain(
	ctx context.Context,
	userId string,
	personId string,
	body usersDomain.UpdateUserBody,
) (err error) {
	var tx *sql.Tx
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateUserMain").SetRaw(err)
	}
	tx, err = client.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	_, err = r.GetUserById(ctx, tx, userId)
	if err != nil {
		return usersDomain.ErrUserNotExist
	}
	err = r.UpdateUser(ctx, userId, body)
	if err != nil {
		return err
	}

	if err = r.updateOrInsertPerson(ctx, tx, personId, userId, body); err != nil {
		return err
	}

	err = r.ValidateUniqueUserExistence(ctx, tx, userId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return
}

func (r usersMySQLRepo) updateOrInsertPerson(
	ctx context.Context,
	tx *sql.Tx,
	personId string,
	userId string,
	body usersDomain.UpdateUserBody,
) (err error) {

	if body.PersonId != nil && body.Person != nil {
		err = r.UpdatePerson(ctx, tx, personId, userId, body.Person)
		if err != nil {
			return err
		}
	}
	if body.PersonId == nil && body.Person != nil {
		_, err = r.CreatePerson(ctx, tx, userId, personId, body.Person)
		if err != nil {
			return err
		}
	} else if body.PersonId != nil && body.Person == nil {
		err = r.UpdatePersonToUser(ctx, tx, userId, personId)
		if err != nil {
			return err
		}
	}

	return nil
}
func (r usersMySQLRepo) GetUserById(
	ctx context.Context,
	tx *sql.Tx,
	userId string,
) (
	user *usersDomain.UserById,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	results, err := tx.QueryContext(
		ctx,
		QueryGetUserById,
		userId,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUserById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	userTmp := make([]UserById, 0)
	err = carta.Map(results, &userTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetUserById").SetRaw(err)
	}
	userRows := make([]usersDomain.UserById, 0)
	automapper.Map(userTmp, &userRows)
	if len(userRows) == 0 {
		return nil, r.err.CopyCodeDescription(usersDomain.ErrUserNotFound)
	}
	return &userRows[0], nil
}

func (r usersMySQLRepo) UpdateUser(
	ctx context.Context,
	userId string,
	body usersDomain.UpdateUserBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateUser").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdateUser,
		body.UserName,
		body.UserTypeId,
		userId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateUser").SetRaw(err)
	}
	return
}

func (r usersMySQLRepo) DeleteUser(
	ctx context.Context,
	userId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:06")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteUser").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteUser,
		now,
		userId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteUser").SetRaw(err)
	}
	return true, nil
}

func (r usersMySQLRepo) ResetPasswordUser(
	ctx context.Context,
	userId string,
	passwordHash string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("ResetPasswordUser").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryResetPasswordUser,
		passwordHash,
		userId,
	)
	if err != nil {
		return false, r.err.Clone().SetFunction("ResetPasswordUser").SetRaw(err)
	}
	return true, nil
}

func (r usersMySQLRepo) GetUserByUserNameAndPassword(
	ctx context.Context,
	userName string,
	passwordHash string,
) (
	user *usersDomain.User,
	xTenantId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, xTenantId, err := db.ClientDB(ctx)
	if err != nil {
		return nil, xTenantId, err
	}
	results, err := client.QueryContext(
		ctx,
		QueryGetUserByPassword,
		userName,
		passwordHash,
	)
	if err != nil {
		return nil, xTenantId, r.err.Clone().SetFunction("GetUserByUserNameAndPassword").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	usersTmp := make([]User, 0)
	err = carta.Map(results, &usersTmp)
	if err != nil {
		return nil, xTenantId, r.err.Clone().SetFunction("GetUserByUserNameAndPassword").SetRaw(err)
	}
	var users = make([]usersDomain.User, 0)
	automapper.Map(usersTmp, &users)
	if len(users) == 0 {
		return nil, xTenantId, r.err.Clone().SetFunction("GetUserByUserNameAndPassword")
	}
	return &users[0], xTenantId, nil
}

func (r usersMySQLRepo) VerifyIfPersonExist(
	ctx context.Context,
	personId string,
) (err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("VerifyIfPersonExist").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryVerifyPersonIdExist,
		personId,
	).Scan(&totalTmp)
	if err != nil {
		return r.err.Clone().SetFunction("VerifyIfPersonExist").SetRaw(err)
	}
	if totalTmp > 1 {
		return usersDomain.ErrPersonIdNotExist
	}
	return nil
}

func (r usersMySQLRepo) VerifyIfUserExist(
	ctx context.Context,
	userId string,
) (err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("VerifyIfUserExist").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryVerifyIfTheUserExist,
		userId,
	).Scan(&totalTmp)
	if err != nil {
		return r.err.Clone().SetFunction("VerifyIfUserExist").SetRaw(err)
	}
	if totalTmp != 1 {
		return usersDomain.ErrUserIdAlreadyExist
	}
	return nil
}

func (r usersMySQLRepo) UpdatePersonToUser(
	ctx context.Context,
	tx *sql.Tx,
	userId string,
	personId string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	_, err = tx.ExecContext(
		ctx,
		QueryUpdatePersonToUser,
		userId,
		personId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdatePersonToUser").SetRaw(err)
	}
	return
}

func (r usersMySQLRepo) ValidateUniquePersonByDocument(
	ctx context.Context, typeDocumentId string, document string,
) (err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("ValidateUniquePersonByDocument").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryValidateUniquePersonByDocument,
		typeDocumentId,
		document,
	).Scan(&totalTmp)
	if err != nil {
		return r.err.Clone().SetFunction("ValidateUniquePersonByDocument").SetRaw(err)
	}
	if totalTmp > 1 {
		return usersDomain.ErrDocumentOfPersonAlreadyExist
	}
	return nil
}

func (r usersMySQLRepo) ValidateUniqueUserExistence(
	ctx context.Context,
	tx *sql.Tx,
	userId string,
) (err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	err = tx.QueryRowContext(
		ctx,
		QueryValidateUniqueUserExistence,
		userId,
	).Scan(&totalTmp)
	if err != nil {
		return r.err.Clone().SetFunction("ValidateUniqueUserExistence").SetRaw(err)
	}
	if totalTmp != 1 {
		return usersDomain.ErrUserIdAppearsMoreThanOnce
	}
	return nil
}

func (r usersMySQLRepo) VerifyPermissionsByUser(
	ctx context.Context, userId string, storeId string, codePermission string,
) (
	res bool, err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("VerifyPermissionsByUser").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryVerifyPermissionsByUser,
		userId,
		storeId,
		storeId,
		storeId,
		codePermission).
		Scan(&totalTmp)
	if err != nil {
		return res, r.err.Clone().SetFunction("VerifyPermissionsByUser").SetRaw(err)
	}
	if totalTmp > 0 {
		res = true
	}
	return res, nil
}

func (r usersMySQLRepo) GetStoresByUser(
	ctx context.Context,
	userId string,
) (
	storesByUser []usersDomain.StoreByUser,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStoresByUser").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetStoresByUser,
			userId,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStoresByUser").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	storesByUserTemp := make([]StoreByUser, 0)
	err = carta.Map(results, &storesByUserTemp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStoresByUser").SetRaw(err)
	}

	automapper.Map(storesByUserTemp, &storesByUser)
	return storesByUser, nil
}

func (r usersMySQLRepo) GetMerchantsByUser(
	ctx context.Context,
	userId string,
) (
	merchantByUsers []usersDomain.MerchantByUser,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMerchantsByUser").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetStoresByUser,
			userId,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMerchantsByUser").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	merchantsByUserTemp := make([]MerchantByUser, 0)
	err = carta.Map(results, &merchantsByUserTemp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMerchantsByUser").SetRaw(err)
	}

	automapper.Map(merchantsByUserTemp, &merchantByUsers)
	return merchantByUsers, nil
}

func (r usersMySQLRepo) GetModulePermissions(
	ctx context.Context, userId string, codeModule string,
) (
	permissions []usersDomain.Permissions, err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetModulePermissions").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetModulePermissions,
			codeModule,
			userId)
	if err != nil {
		return permissions, r.err.Clone().SetFunction("GetModulePermissions").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	permissionsTmp := make([]Permissions, 0)
	err = carta.Map(results, &permissionsTmp)
	if err != nil {
		return permissions, r.err.Clone().SetFunction("GetModulePermissions").SetRaw(err)
	}

	permissionsAux := make([]usersDomain.Permissions, 0)
	automapper.Map(permissionsTmp, &permissionsAux)
	return permissionsAux, nil
}

func (r usersMySQLRepo) GetModules(
	ctx context.Context,
) (
	modulesRows []usersDomain.Module,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetModules").SetRaw(err)
	}
	results, err := client.QueryContext(
		ctx,
		QueryGetModules,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetModules").SetRaw(err)
	}

	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &err)
		}
	}(results)

	modulesTmp := make([]Module, 0)
	err = carta.Map(results, &modulesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetModules").SetRaw(err)
	}
	automapper.Map(modulesTmp, &modulesRows)
	return modulesRows, nil
}
