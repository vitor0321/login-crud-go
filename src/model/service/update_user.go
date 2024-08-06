package service

import (
	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"example.com/mod/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateUser model", zap.String("journey", "UpdateUser"))
	return nil
}
