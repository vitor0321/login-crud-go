package controller

import (
	"strconv"

	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/validation"
	"example.com/mod/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller", zap.String("journey", "CreateUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "CreateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	logger.Info("User created success: "+userRequest.Email+" "+userRequest.Name+" "+strconv.Itoa(int(userRequest.Age)), zap.String("journey", "CreateUser"))
}
