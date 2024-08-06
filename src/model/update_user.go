package model

import (
	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) UpdateUser(string) *rest_err.RestErr {
	logger.Info("Init UpdateUser model", zap.String("journey", "UpdateUser"))
	return nil
}
