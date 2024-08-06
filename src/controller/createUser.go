package controller

import (
	"net/http"

	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/validation"
	"example.com/mod/src/controller/model/request"
	"example.com/mod/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	if err := domain.CreateUser(); err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "CreateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created success", zap.String("journey", "CreateUser"))
	c.String(http.StatusOK, "")
}
