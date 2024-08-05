package controller

import (
	"log"

	"example.com/mod/src/configuration/validation"
	"example.com/mod/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Create user controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error binding user request: %+v", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	log.Printf("User request: %+v", userRequest)
}
