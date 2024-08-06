package model

import (
	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) FindUserById(string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init FindUserById model", zap.String("journey", "FindUserById"))
	return nil, nil
}
