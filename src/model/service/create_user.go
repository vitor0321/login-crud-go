package service

import (
	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"example.com/mod/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser", zap.String("journey", "service/model/create_user"))

	userDomain.EncryptPassword()

	logger.Info("User created success", zap.String("journey", "service/model/create_user"))
	return nil
}
