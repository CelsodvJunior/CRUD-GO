package main

import (
	"log"

	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/database/mongdb"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/logger"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/controller"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/controller/routes"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/model/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Abaut to start User app")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongdb.InitConnection()

	//Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
