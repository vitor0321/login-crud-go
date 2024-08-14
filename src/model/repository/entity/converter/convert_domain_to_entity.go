package converter

import (
	"example.com/mod/src/model"
	"example.com/mod/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertDomainToEntity(
	ud model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		ID:       primitive.NewObjectID(),
		Email:    ud.GetEmail(),
		Password: ud.GetPassword(),
		Name:     ud.GetName(),
		Age:      ud.GetAge(),
	}
}
