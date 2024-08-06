package main

import (
	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/controller"
	"example.com/mod/src/controller/routes"
	"example.com/mod/src/model/service"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application", zap.String("journey", "main"))
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err, zap.String("journey", "main"))
	}
	//Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	route := gin.Default()

	routes.InitRoutes(route, userController)

	if err := route.Run(":8080"); err != nil {
		logger.Error("Error running the server", err, zap.String("journey", "main"))
	}
}
