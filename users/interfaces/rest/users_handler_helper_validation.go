/*
 * File: users_handler_helper_validation.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Validation entity helper to users.
 *
 * Last Modified: 2023-11-23
 */

package rest

type UpdateUserValidate struct {
	UserName   string  `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	UserTypeId string  `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
	PersonId   *string `json:"person_id" example:"739bbbc9-7e93-11ee-89fd-0442ac210932"`
	Person     *Person `json:"person"`
}

type createUsersValidate struct {
	UserName   string  `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	Password   string  `json:"password" binding:"required" example:"pepitoPass"`
	UserTypeId string  `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
	PersonId   *string `json:"person_id" example:"739bbbc9-7e93-11ee-89fd-0442ac210932"`
	Person     *Person `json:"person"`
}

type Person struct {
	TypeDocumentId string  `json:"type_document_id" binding:"required" example:"00a58522-93b4-11ee-a040-0242ac11000e"`
	Document       string  `json:"document" binding:"required" example:"77895428"`
	Names          string  `json:"names" binding:"required" example:"LUCY ANDREA"`
	Surname        string  `json:"surname" binding:"required" example:"HANCCO"`
	LastName       *string `json:"last_name" example:"HUILLCA"`
	Phone          string  `json:"phone" binding:"required" example:"918547496"`
	Email          *string `json:"email"  example:"lucyhancco@gmail.com"`
	Gender         *string `json:"gender" example:"MASCULINO"`
	Enable         bool    `json:"enable" binding:"required" example:"1"`
}

type resetUserPasswordValidate struct {
	NewPassword string `json:"new_password" binding:"required" example:"pepitoPass"`
}

type loginUserValidate struct {
	UserName string `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	Password string `json:"password" binding:"required" example:"pepitoPass"`
}
