package domain

import (
	"golang.org/x/crypto/bcrypt"
	"server-go/internal/config/rest_errors"
)

func NewUserDomain(firstName, lastName, email, password, phone, role string, isActive bool) UserDomainInterface {
	return &userDomain{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Phone:     phone,
		Role:      role,
		IsActive:  isActive,
	}
}

type userDomain struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      string
	Phone     string
	IsActive  bool
}

type UserDomainInterface interface {
	CreateUser() *rest_errors.RestErr
	//UpdateUser(string) *rest_errors.RestErr
	//FindUser(string) (*userDomain, *rest_errors.RestErr)
	//DeleteUser(string) *rest_errors.RestErr
}

func (ud *userDomain) EncryptPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	ud.Password = string(bytes)
}
