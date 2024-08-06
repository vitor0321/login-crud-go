package service

import (
	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"example.com/mod/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) FindUserById(
	userId string,
) (
	*model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init FindUserById model", zap.String("journey", "FindUserById"))
	return nil, nil
}
