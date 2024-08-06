package service

import (
	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*userDomainService) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init DeleteUser model", zap.String("journey", "DeleteUser"))
	return nil
}
