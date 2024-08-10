package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"example.com/mod/src/configuration/logger"
	"go.uber.org/zap"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int16

	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword()
}

func NewUserDomain(
	email,
	password,
	name string,
	age int16,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int16  `json:"age"`
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		logger.Error("Error trying to marshal user domain", err, zap.String("journey", "model/user_domain"))
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int16 {
	return ud.Age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
