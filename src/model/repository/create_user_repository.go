package repository

import (
	"example.com/mod/src/configuration/rest_err"
	"example.com/mod/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(databaseConnection *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: databaseConnection,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
