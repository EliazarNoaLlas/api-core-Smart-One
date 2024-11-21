/*
 * File: users_func_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to users.
 *
 * Last Modified: 2023-11-23
 */

package usecase

import (
	"context"
	"strings"
	"sync"

	"github.com/google/uuid"

	logErrorCoreDomain "gitlab.smartcitiesperu.com/smartone/api-shared/error-core/domain"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
	validationsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/validations/domain"

	usersDomain "gitlab.smartcitiesperu.com/smartone/api-core/users/domain"
)

func (u usersUseCase) GetUser(
	ctx context.Context,
	userId string,
) (
	user *usersDomain.User,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err = u.usersRepository.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u usersUseCase) GetUsers(
	ctx context.Context,
	searchParams usersDomain.GetUsersParams,
	pagination paramsDomain.PaginationParams,
) (
	users []usersDomain.UserMultiple,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetUsers, errGetTotalUsers error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		users, errGetUsers = u.usersRepository.GetUsers(ctx, searchParams, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalUsers = u.usersRepository.GetTotalUsers(ctx, searchParams, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetUsers != nil {
		return nil, nil, errGetUsers
	}
	if errGetTotalUsers != nil {
		return nil, nil, errGetTotalUsers
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return users, &paginationRes, nil
}

func (u usersUseCase) GetMenuByUser(
	ctx context.Context,
	userId string,
) (
	menu []usersDomain.MenuModule,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	modules, err := u.usersRepository.GetModules(ctx)
	if err != nil {
		return nil, err
	}

	modulesByUser, err := u.usersRepository.GetMenuByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	menu = make([]usersDomain.MenuModule, 0)
	menuMap := map[string]int{}
	for _, moduleByUser := range modulesByUser {
		if strings.Count(moduleByUser.Code, ".") == 0 {
			menu = append(menu, usersDomain.MenuModule{
				ModuleMenuUser: moduleByUser,
				Modules:        make([]usersDomain.MenuModule, 0),
			})
			menuMap[moduleByUser.Code] = len(menu) - 1
			continue
		}
		moduleBaseParts := strings.Split(moduleByUser.Code, ".")
		if len(moduleBaseParts) == 0 {
			continue
		}
		moduleCodeBase := moduleBaseParts[0]
		iMenu, existModuleInMenu := menuMap[moduleCodeBase]
		if existModuleInMenu {
			menu[iMenu].Modules = append(menu[iMenu].Modules, usersDomain.MenuModule{
				ModuleMenuUser: moduleByUser,
				Modules:        make([]usersDomain.MenuModule, 0),
			})
			continue
		}
		var moduleBase *usersDomain.Module = nil
		for _, module := range modules {
			if module.Code == moduleCodeBase {
				moduleBase = &module
				break
			}
		}
		if moduleBase == nil {
			continue
		}
		newModuleBase := usersDomain.MenuModule{
			ModuleMenuUser: usersDomain.ModuleMenuUser{
				Id:          moduleBase.Id,
				Name:        moduleBase.Name,
				Description: moduleBase.Description,
				Code:        moduleBase.Code,
				Icon:        moduleBase.Icon,
				Position:    moduleBase.Position,
				CreatedAt:   moduleBase.CreatedAt,
				Views:       make([]usersDomain.ViewMenuUser, 0),
			},
			Modules: make([]usersDomain.MenuModule, 0),
		}
		newModuleBase.Modules = append(newModuleBase.Modules, usersDomain.MenuModule{
			ModuleMenuUser: moduleByUser,
			Modules:        make([]usersDomain.MenuModule, 0),
		})
		menu = append(menu, newModuleBase)
		menuMap[moduleCodeBase] = len(menu) - 1
	}
	return menu, nil
}

func (u usersUseCase) GetMeByUser(
	ctx context.Context,
	userId string,
) (
	user *usersDomain.UserMe,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errUserMe, errGetStoresByUser, errGetMerchantsByUser error
	var userMe *usersDomain.UserMeInfo
	var storesByUser []usersDomain.StoreByUser
	var merchantsByUser []usersDomain.MerchantByUser
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		userMe, errUserMe = u.usersRepository.GetMeByUser(ctx, userId)
		wg.Done()
	}()
	go func() {
		storesByUser, errGetStoresByUser = u.usersRepository.GetStoresByUser(ctx, userId)
		wg.Done()
	}()
	go func() {
		merchantsByUser, errGetMerchantsByUser = u.usersRepository.GetMerchantsByUser(ctx, userId)
		wg.Done()
	}()
	wg.Wait()

	if errUserMe != nil {
		return nil, errUserMe
	}
	if errGetStoresByUser != nil {
		return nil, errGetStoresByUser
	}
	if errGetMerchantsByUser != nil {
		return nil, errGetMerchantsByUser
	}

	user = &usersDomain.UserMe{
		Id:        userMe.Id,
		UserName:  userMe.UserName,
		CreatedAt: userMe.CreatedAt,
		Person:    userMe.Person,
		RoleUser:  userMe.RoleUser,
		Stores:    storesByUser,
		Merchants: merchantsByUser,
	}
	return user, nil
}

func (u usersUseCase) CreateUser(
	ctx context.Context,
	body usersDomain.CreateUserBody,
) (
	id *string, err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_users",
		IdColumnName:     "username",
		IdValue:          body.UserName,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, usersDomain.ErrUserUsernameAlreadyExist
	}

	userId := uuid.New().String()
	// case basic
	if body.PersonId == nil && body.Person == nil {
		id, err = u.usersRepository.CreateUser(ctx, nil, userId, body)
	} else if body.PersonId != nil {
		body.Person = nil
		id, err = u.usersRepository.CreateUserMain(ctx, userId, *body.PersonId, body)
	} else {
		err = u.usersRepository.ValidateUniquePersonByDocument(ctx, body.Person.TypeDocumentId, body.Person.Document)
		if err != nil {
			return nil, err
		}
		personId := uuid.New().String()
		id, err = u.usersRepository.CreateUserMain(ctx, userId, personId, body)
	}
	return
}

func (u usersUseCase) UpdateUser(
	ctx context.Context,
	userId string,
	body usersDomain.UpdateUserBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_users",
		IdColumnName:     "id",
		IdValue:          userId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(usersDomain.ErrUserNotFound).SetFunction("UpdateUser")
	}

	if body.PersonId == nil && body.Person == nil {
		err = u.usersRepository.VerifyIfUserExist(ctx, userId)
		if err != nil {
			return err
		}
		err = u.usersRepository.UpdateUser(ctx, userId, body)

	} else if body.PersonId != nil && body.Person != nil {
		err = u.usersRepository.VerifyIfPersonExist(ctx, *body.PersonId)
		if err != nil {
			return err
		}
		err = u.usersRepository.UpdateUserMain(ctx, userId, *body.PersonId, body)
	} else if body.PersonId != nil {
		body.Person = nil
		err = u.usersRepository.UpdateUserMain(ctx, userId, *body.PersonId, body)
	} else {
		personId := uuid.New().String()
		err = u.usersRepository.ValidateUniquePersonByDocument(ctx, body.Person.TypeDocumentId, body.Person.Document)
		if err != nil {
			return err
		}
		err = u.usersRepository.UpdateUserMain(ctx, userId, personId, body)
	}

	return
}

func (u usersUseCase) DeleteUser(
	ctx context.Context,
	userId string,
) (
	update bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	var deleted string
	deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_users",
		IdColumnName:     "id",
		IdValue:          userId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, usersDomain.ErrUserIdHasBeenDeleted
	}

	res, err := u.usersRepository.DeleteUser(ctx, userId)
	return res, err
}

func (u usersUseCase) ResetPasswordUser(
	ctx context.Context,
	userId string,
	body usersDomain.ResetUserPasswordBody,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	hashPassword, err := HashPasswordUser(ctx, body.NewPassword)
	if err != nil {
		return false, err
	}

	updated, err = u.usersRepository.ResetPasswordUser(ctx, userId, hashPassword)
	return
}

func (u usersUseCase) LoginUser(
	ctx context.Context,
	body usersDomain.LoginUserBody,
) (
	tokenString *string,
	xTenantId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	hashPassword, err := HashPasswordUser(ctx, body.Password)
	if err != nil {
		return nil, nil, err
	}

	user, xTenantId, err := u.usersRepository.GetUserByUserNameAndPassword(ctx, body.UserName, hashPassword)
	if err != nil {
		return nil, xTenantId, err
	}

	tokenString, err = u.authRepository.GenerateToken(user.Id)
	return tokenString, xTenantId, nil
}

func HashPasswordUser(
	ctx context.Context,
	password string,
) (
	passwordHash string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	// REVIEW generate hash password in base of password
	return password, nil
}

func (u usersUseCase) VerifyPermissionsByUser(
	ctx context.Context, userId string, storeId string, codePermission string,
) (
	res bool, err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if storeId == "" {
		return res, usersDomain.ErrStoreIdEmpty
	}

	res, err = u.usersRepository.VerifyPermissionsByUser(ctx, userId, storeId, codePermission)
	if err != nil {
		return res, err
	}
	return
}

func (u usersUseCase) GetModulePermissions(
	ctx context.Context, userId string, codeModule string,
) (
	permissions []usersDomain.Permissions, err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_modules",
		IdColumnName:     "code",
		IdValue:          codeModule,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return permissions, err
	}
	if !exist {
		return permissions, u.err.Clone().CopyCodeDescription(usersDomain.ErrInvalidCodeModule).SetFunction("GetModulePermissions")
	}

	permissions, err = u.usersRepository.GetModulePermissions(ctx, userId, codeModule)
	if err != nil {
		return permissions, err
	}
	return permissions, nil
}
