package repository

import (
	"context"
	"os"

	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"example.com/mod/src/model"
	"go.uber.org/zap"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser", zap.String("journey", "repository/model/create_user_repository"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		logger.Error("Error trying to get user domain json value", err, zap.String("journey", "repository/model/create_user_repository"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to insert user domain", err, zap.String("journey", "repository/model/create_user_repository"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
