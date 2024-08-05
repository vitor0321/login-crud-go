package main

import (
	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/controller/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	route := gin.Default()

	routes.InitRoutes(route)

	if err := route.Run(":8080"); err != nil {
		logger.Error("Error running the server", err)
	}
}
