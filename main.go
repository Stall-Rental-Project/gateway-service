package main

import (
	"gateway-service/config"
	"gateway-service/server"
	"log"
)

// @title MHMarket REST API
// @version 1.0
// @description This is a documents for Rental System REST API.

// @contact.name Developer: Hoang Ngo
// @contact.email hoanggg2110@gmail.com.vn

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @schemes https
// @securityDefinitions.apikey Bearer Token
// @in header
// @name Authorization
func main() {
	start()
}

func start() {
	err := config.Config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	goServer := server.NewServer()
	goServer.Initialize()
	goServer.InitializeRoutes()
	goServer.Run(config.Config.Port)
}
