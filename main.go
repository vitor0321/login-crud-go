package main

import (
	"fmt"
	"log"

	"example.com/mod/src/controller/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	route := gin.Default()

	routes.InitRoutes(route)

	if err := route.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
