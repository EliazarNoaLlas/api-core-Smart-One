/*
 * File: main.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entry point for the application.
 *
 * Last Modified: 2024-04-09
 */

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.smartcitiesperu.com/smartone/api-shared/config"
	"gitlab.smartcitiesperu.com/smartone/api-shared/db"

	"gitlab.smartcitiesperu.com/smartone/api-core/server/setup"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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

	setup.LoadServer(router)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	err = router.Run(serverPort)
	if err != nil {
		return
	}
}
