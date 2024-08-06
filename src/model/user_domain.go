package model

import (
	"crypto/md5"
	"encoding/hex"

	"example.com/mod/src/configuration/rest_err"
)

func NewUserDomain(
	email,
	password,
	name string,
	age int16,
) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int16
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUserById(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
