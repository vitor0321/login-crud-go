package service

import (
	"example.com/mod/src/configuration/rest_err"
	"example.com/mod/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	FindUserById(
		userId string,
	) (*model.UserDomainInterface, *rest_err.RestErr)

	DeleteUser(
		userId string,
	) *rest_err.RestErr
}
