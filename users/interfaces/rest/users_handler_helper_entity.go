/*
 * File: users_handler_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to handler for users.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	paginationDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	usersDomain "gitlab.smartcitiesperu.com/smartone/api-core/users/domain"
)

type userResult struct {
	Data   usersDomain.User `json:"data" binding:"required"`
	Status int              `json:"status" binding:"required"`
}

type usersResult struct {
	Data       []usersDomain.User                 `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type multipleUsersResult struct {
	Data       []usersDomain.UserMultiple         `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type menuByUserResult struct {
	Data   []usersDomain.MenuModule `json:"data" binding:"required"`
	Status int                      `json:"status" binding:"required"`
}

type GetMeByUser struct {
	Data   usersDomain.UserMe `json:"data" binding:"required"`
	Status int                `json:"status" binding:"required"`
}

type deleteUsersResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}

type HashPasswordUserResult struct {
	Data   string `json:"data" binding:"required"`
	Status int    `json:"status" binding:"required"`
}

type ResetPasswordUserResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}

type LoginUserResult struct {
	Data   string `json:"data" binding:"required"`
	Status int    `json:"status" binding:"required"`
}

type PermissionsResult struct {
	Data   []usersDomain.Permissions `json:"data" binding:"required"`
	Status int                       `json:"status" binding:"required"`
}
