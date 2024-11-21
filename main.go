/*
 * File: main.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * To define the routes for the core.
 *
 * Last Modified: 2023-11-10
 */

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gitlab.smartcitiesperu.com/smartone/api-shared/config"
	"gitlab.smartcitiesperu.com/smartone/api-shared/db"

	documentTypesSetup "gitlab.smartcitiesperu.com/smartone/api-core/document-types/setup"
	economicActivitiesSetup "gitlab.smartcitiesperu.com/smartone/api-core/economic-activities/setup"
	merchantEconomicActivitiesSetup "gitlab.smartcitiesperu.com/smartone/api-core/merchant-economic-activities/setup"
	merchantsSetup "gitlab.smartcitiesperu.com/smartone/api-core/merchants/setup"
	modulesSetup "gitlab.smartcitiesperu.com/smartone/api-core/modules/setup"
	permissionsSetup "gitlab.smartcitiesperu.com/smartone/api-core/permissions/setup"
	policiesSetup "gitlab.smartcitiesperu.com/smartone/api-core/policies/setup"
	policyPermissionsSetup "gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/setup"
	receiptTypes "gitlab.smartcitiesperu.com/smartone/api-core/receipt-types/setup"
	rolePoliciesSetup "gitlab.smartcitiesperu.com/smartone/api-core/role-policies/setup"
	rolesSetup "gitlab.smartcitiesperu.com/smartone/api-core/roles/setup"
	serverSetup "gitlab.smartcitiesperu.com/smartone/api-core/server/setup"
	storeTypesSetup "gitlab.smartcitiesperu.com/smartone/api-core/store-types/setup"
	storesSetup "gitlab.smartcitiesperu.com/smartone/api-core/stores/setup"
	userRolesSetup "gitlab.smartcitiesperu.com/smartone/api-core/user-roles/setup"
	userTypesSetup "gitlab.smartcitiesperu.com/smartone/api-core/user-types/setup"
	usersSetup "gitlab.smartcitiesperu.com/smartone/api-core/users/setup"
	viewPermissionsSetup "gitlab.smartcitiesperu.com/smartone/api-core/view-permissions/setup"
	viewsSetup "gitlab.smartcitiesperu.com/smartone/api-core/views/setup"
)

func main() {
	cfg := config.Configuration{
		ServerPort:  os.Getenv("SERVER_PORT"),
		StoragePath: os.Getenv("STORAGE_PATH"),
		DB: config.DB{
			DbDatabase: os.Getenv("DB_DATABASE"),
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUsername: os.Getenv("DB_USERNAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}

	err := db.InitClients(cfg)
	if err != nil {
		return
	}
	defer db.Client.Close()
	router := gin.Default()

	documentTypesSetup.LoadDocumentTypes(router)
	economicActivitiesSetup.LoadEconomicActivities(router)
	merchantEconomicActivitiesSetup.LoadMerchantEconomicActivities(router)
	merchantsSetup.LoadMerchants(router)
	modulesSetup.LoadModules(router)
	permissionsSetup.LoadPermissions(router)
	policiesSetup.LoadPolicies(router)
	policyPermissionsSetup.LoadPolicyPermissions(router)
	rolePoliciesSetup.LoadRolePolicies(router)
	rolesSetup.LoadRoles(router)
	storeTypesSetup.LoadStoreTypes(router)
	storesSetup.LoadStores(router)
	userRolesSetup.LoadUserRoles(router)
	userTypesSetup.LoadUserTypes(router)
	usersSetup.LoadUsers(router)
	viewsSetup.LoadViews(router)
	viewPermissionsSetup.LoadViewPermissions(router)
	receiptTypes.LoadReceiptTypes(router)
	serverSetup.LoadServer(router)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	err = router.Run(serverPort)
	if err != nil {
		return
	}
}
