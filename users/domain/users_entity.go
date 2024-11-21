/*
 * File: users_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the entities model to users.
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"time"

	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"
)

type UserMultiple struct {
	//Description: user id
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: username of the user
	UserName string `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	//Description: date of created
	CreatedAt *time.Time     `json:"created_at" example:"2023-11-10 08:10:00"`
	UserType  UserTypeByUser `json:"user_type" binding:"required"`
	Role      []Role         `json:"role" binding:"required"`
}

type Role struct {
	//Description: the id of the role
	Id *string `db:"role_id" example:"fcdbfacf-8305-11ee-89fd-0242ac110016"`
	//Description: the id of the role
	Name *string `db:"role_name" example:"Jefe de Area Residual"`
	//Description: the description of the role
	Description *string `db:"role_description" example:"Gerencia del conglomerado"`
	//Description: enable of the role
	Enable *bool `json:"role_enable" example:"true"`
	//Description: the date of created of the role
	CreatedAt *time.Time `db:"role_created_at" example:"2023-11-27 19:47:15"`
	UserRole  UserRole   `json:"user_role" binding:"required"`
}

type UserRole struct {
	//Description: the id of the use role
	Id *string `json:"user_role_id" example:"b36f266d-8492-4f0e-8ecb-fef20e098970"`
}

type User struct {
	//Description: user id
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: username of the user
	UserName string `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	//Description: date of created
	CreatedAt *time.Time     `json:"created_at" example:"2023-11-10 08:10:00"`
	UserType  UserTypeByUser `json:"user_type" binding:"required"`
}

type CreateUserBody struct {
	//Description: the username of the user
	UserName string `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	//Description: the password of the user
	Password string `json:"password" binding:"required" example:"pepitoPass"`
	//Description: the type of the user
	UserTypeId string `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
	//Description: the person id
	PersonId *string `json:"person_id" example:"739bbbc9-7e93-11ee-89fd-0442ac210932"`
	Person   *Person `json:"person"`
}

type UpdateUserBody struct {
	//Description: the username of the user
	UserName string `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	//Description: the type of the user
	UserTypeId string `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0442ac210931"`
	//Description: the person id
	PersonId *string `json:"person_id" example:"739bbbc9-7e93-11ee-89fd-0442ac210932"`
	Person   *Person `json:"person"`
}

type UserTypeByUser struct {
	//Description: the id of the user
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113421"`
	//Description: the description of the user
	Description string `json:"description" binding:"required" example:"Usuario externo"`
	//Description: the code of the user
	Code string `json:"code" binding:"required" example:"USER_EXTERNAL"`
}

type GetUsersParams struct {
	paramsDomain.Params
	//Description: the type of the user
	UserTypeId *string `json:"type_id"`
	//Description: the username of the user
	UserName *string `json:"username"`
	//Description: the role of the user
	RoleId []string `json:"role_id"`
}

type ResetUserPasswordBody struct {
	//Description: the new password of the user
	NewPassword string `json:"new_password" binding:"required" example:"pepitoPass"`
}

type LoginUserBody struct {
	//Description: the username of the user
	UserName string `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	//Description: the password of the user
	Password string `json:"password" binding:"required" example:"pepitoPass"`
}

type ViewMenuUser struct {
	//Description: the id of the view menu user
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110000"`
	//Description: the name of the view menu user
	Name string `json:"name" binding:"required" example:"Requerimientos"`
	//Description: the description of the view menu user
	Description string `json:"description" binding:"required" example:"Vista de requerimientos"`
	//Description: the url of the view menu user
	Url string `json:"url" binding:"required" example:"/logistics/requirements"`
	//Description: the icon in for the view menu user
	Icon string `json:"icon" binding:"required" example:"fa fa-chart"`
	//Description: the date of created the view menu user
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type ModuleMenuUser struct {
	//Description: The id of the menu user
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: The name of the menu user
	Name string `json:"name" binding:"required" example:"Logistic"`
	//Description: The description of the menu user
	Description string `json:"description" binding:"required" example:"Modulo de logística"`
	//Description: The code of the menu user
	Code string `json:"code" binding:"required" example:"logistic"`
	//Description: The icon of the menu user
	Icon string `json:"icon" binding:"required" example:"fa fa-chart"`
	//Description: The position of the menu user
	Position int `json:"position" binding:"required" example:"1"`
	//Description: The date of created the menu user
	CreatedAt *time.Time     `json:"created_at" example:"2023-11-10 08:10:00"`
	Views     []ViewMenuUser `json:"views" binding:"required"`
}

type MenuModule struct {
	ModuleMenuUser
	Modules []MenuModule `json:"modules"`
}

type UserMe struct {
	//Description: user id
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: username of the user
	UserName string `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	//Description: date of created
	CreatedAt *time.Time       `json:"created_at" example:"2023-11-10 08:10:00"`
	Person    *PersonByUser    `json:"person"`
	RoleUser  []RoleUser       `json:"roles" binding:"required"`
	Stores    []StoreByUser    `json:"stores" binding:"required"`
	Merchants []MerchantByUser `json:"merchants" binding:"required"`
}

type UserMeInfo struct {
	//Description: user id
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: username of the user
	UserName string `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	//Description: date of created
	CreatedAt *time.Time    `json:"created_at" example:"2023-11-10 08:10:00"`
	Person    *PersonByUser `json:"person"`
	RoleUser  []RoleUser    `json:"roles" binding:"required"`
}

type RoleUser struct {
	//Description: role user id
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description:user role name
	Name *string `json:"name" example:"Gerencia"`
	//Description: user role description
	Description *string `json:"description" example:"Gerencia general"`
	//Description: user role status
	Enable *bool `json:"enable" example:"1"`
	//Description: date of created
	CreateAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type PersonByUser struct {
	//Description: the id of the people
	Id *string `json:"id" example:"0abbb86f-9836-11ee-a040-0242ac11000e"`
	//Description: the document number of the people
	Document *string `json:"document" example:"77895428"`
	//Description: the name of the people
	Names *string `json:"names" example:"LUCY ANDREA"`
	//Description: the surname of the people
	Surname *string `json:"surname" example:"HANCCO"`
	//Description: the last name of the people
	LastName *string `json:"last_name" example:"HUILLCA"`
	//Description: the phone of the people
	Phone *string `json:"phone" example:"918547496"`
	//Description: the email of the people
	Email *string `json:"email" example:"lucyhancco@gmail.com"`
	//Description: the gender of the people
	Gender *string `json:"gender" example:"MASCULINO"`
	//Description: the status of the people
	Enable *bool `json:"enable" example:"1"`
	//Description: the date of created of the people
	CreatedAt    *time.Time    `json:"created_at" example:"2023-11-10 08:10:00"`
	TypeDocument *TypeDocument `json:"type_document"`
}

type TypeDocument struct {
	//Description: id of document type
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: document type number
	Number *string `json:"number" example:"01"`
	//Description: description of the type of document
	Description *string `json:"description" example:"DOCUMENTO NACIONAL DE IDENTIDAD"`
	//Description: abbreviated description of the type of document
	AbbreviateDescription *string `json:"abbreviate_description" example:"DNI"`
	//Description: abbreviated document type status
	Enable *bool `json:"enable" example:"1"`
	//Description: the creation date of the document type
	CreateAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type Person struct {
	//Description: the type of the document
	TypeDocumentId string `json:"type_document_id" binding:"required" example:"00a58522-93b4-11ee-a040-0242ac11000e"`
	//Description: the document number of the people
	Document string `json:"document" binding:"required" example:"77895428"`
	//Description: the name of the people
	Names string `json:"names" binding:"required" example:"LUCY ANDREA"`
	//Description: the surname of the people
	Surname string `json:"surname" binding:"required" example:"HANCCO"`
	//Description: the last name of the people
	LastName *string `json:"last_name" example:"HUILLCA"`
	//Description: the phone of the people
	Phone string `json:"phone" binding:"required" example:"918547496"`
	//Description: the email of the people
	Email *string `json:"email"  example:"lucyhancco@gmail.com"`
	//Description: the gender of the people
	Gender *string `json:"gender" example:"MASCULINO"`
	//Description: the status of the people
	Enable bool `json:"enable" binding:"required" example:"1"`
}

type UpdatePersonBody struct {
	//Description: the id of the user
	UserId *string `json:"user_id"`
	//Description: the type of the document
	TypeDocumentId string `json:"type_document_id" binding:"required" example:"00a58522-93b4-11ee-a040-0242ac11000e"`
	//Description: the document number of the people
	Document string `json:"document" binding:"required" example:"77895428"`
	//Description: the name of the people
	Names string `json:"names" binding:"required" example:"LUCY ANDREA"`
	//Description: the surname of the people
	Surname string `json:"surname" binding:"required" example:"HANCCO"`
	//Description: the last name of the people
	LastName *string `json:"last_name" example:"HUILLCA"`
	//Description: the phone of the people
	Phone string `json:"phone" binding:"required" example:"918547496"`
	//Description: the email of the people
	Email *string `json:"email" example:"lucyhancco@gmail.com"`
	//Description: the gender of the people
	Gender *string `json:"gender" example:"MASCULINO"`
	//Description: the status of the people
	Enable bool `json:"enable" binding:"required" example:"1"`
}

type UserById struct {
	//Description: user id
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: username of the user
	UserName string `json:"username" binding:"required" example:"pepito.quispe@smartc.pe"`
	//Description: date of created
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type Merchant struct {
	//Description: the id of the merchant
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the merchant
	Name string `json:"name" binding:"required" example:"Almacen Central"`
	//Description: the description of the merchant
	Description string `json:"description" binding:"required" example:"Almacen Central"`
	//Description: the image path of the merchant
	ImagePath string `json:"image_path" binding:"required" example:"/images/almacen-central.jpg"`
}

type StoreByUser struct {
	//Description: the id of the store
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the store
	Name     string   `json:"name" binding:"required" example:"Almacen Central"`
	Merchant Merchant `json:"merchant" binding:"required"`
}

type Store struct {
	//Description: the id of the store
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the store
	Name string `json:"name" binding:"required" example:"Almacen Central"`
}

type MerchantByUser struct {
	//Description: the id of the merchant
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the merchant
	Name string `json:"name" binding:"required" example:"Almacen Central"`
	//Description: the description of the merchant
	Description string `json:"description" binding:"required" example:"Almacen Central"`
	//Description: the image path of the merchant
	ImagePath string  `json:"image_path" binding:"required" example:"/images/almacen-central.jpg"`
	Stores    []Store `json:"stores" binding:"required"`
}

type Permissions struct {
	//Description: user id
	Id string `json:"id" binding:"required" example:"0c4001f3-2dd8-4d9f-820d-db7d7d8c85c0"`
	//Description: The code of the module
	Code string `json:"code" binding:"required" example:"logistics.requirements"`
}

type Module struct {
	//Description: module  id
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: module  name
	Name string `json:"name" binding:"required" example:"Logistic"`
	//Description: module  description
	Description string `json:"description" binding:"required" example:"Modulo de logística"`
	//Description: module  code
	Code string `json:"code" binding:"required" example:"logistic"`
	//Description: module  icon
	Icon string `json:"icon" binding:"required" example:"fa fa-chart"`
	//Description: module  position
	Position int `json:"position" binding:"required" example:"1"`
	//Description: module  created_at
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}
