package model

import (
	"strconv"

	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "CreateUser"))
	ud.EncryptPassword()
	logger.Info("User created success: email= "+ud.Email+", name= "+ud.Name+", age= "+strconv.Itoa(int(ud.Age)), zap.String("journey", "CreateUser"))
	return nil
}
