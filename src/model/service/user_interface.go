package service

import (
	"example.com/mod/src/configuration/rest_err"
	"example.com/mod/src/model"
	"example.com/mod/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository: userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

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
