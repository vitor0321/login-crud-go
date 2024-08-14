package service

import (
	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"example.com/mod/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser", zap.String("journey", "service/model/create_user"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "service/model/create_user"))
		return nil, err
	}

	logger.Info("User created success", zap.String("journey", "service/model/create_user"))
	return userDomainRepository, nil
}
