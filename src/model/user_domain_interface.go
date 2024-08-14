package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int16
	GetID() string

	SetID(string)

	EncryptPassword()
}

func NewUserDomain(
	email,
	password,
	name string,
	age int16,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}
