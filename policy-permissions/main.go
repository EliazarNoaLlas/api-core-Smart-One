/*
 * File: main.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * microservice to manage permissions of a policy
 *
 * Last Modified: 2023-11-23
 */

package main

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"gitlab.smartcitiesperu.com/smartone/api-shared/config"
	"gitlab.smartcitiesperu.com/smartone/api-shared/db"

	"gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/docs"
	"gitlab.smartcitiesperu.com/smartone/api-core/policy-permissions/setup"
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

	setup.LoadPolicyPermissions(router)
	loadSwagger(router)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	err = router.Run(serverPort)
	if err != nil {
		return
	}
}

func loadSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Swagger policy permissions API"
	docs.SwaggerInfo.Description = "This is a policy permissions microservice."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9009"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.StaticFile("/docs/swagger3.json", "./docs/swagger3.json")
}
