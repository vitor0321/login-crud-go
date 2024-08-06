package controller

import (
	"net/http"
	"strconv"

	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/validation"
	"example.com/mod/src/controller/model/request"
	"example.com/mod/src/model"
	"example.com/mod/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller", zap.String("journey", "controller/create_user"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "controller/create_user"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "controller/create_user"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created success: email= "+domain.GetEmail()+", name= "+domain.GetName()+", age= "+strconv.Itoa(int(domain.GetAge())), zap.String("journey", "controller/create_user"))
	c.JSON(
		http.StatusOK,
		view.ConvertDomainToResponse(domain),
	)
}
