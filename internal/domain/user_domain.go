package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetRole() string
	GetPhone() string
	GetFirstName() string
	GetLastName() string
	GetIsActive() bool

	EncryptPassword()
}

func NewUserDomain(firstName, lastName, email, password, phone, role string, isActive bool) UserDomainInterface {
	return &userDomain{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  password,
		phone:     phone,
		role:      role,
		isActive:  isActive,
	}
}

type userDomain struct {
	firstName string
	lastName  string
	email     string
	password  string
	role      string
	phone     string
	isActive  bool
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetRole() string {
	return ud.role
}
func (ud *userDomain) GetPhone() string {
	return ud.phone
}
func (ud *userDomain) GetFirstName() string {
	return ud.firstName
}
func (ud *userDomain) GetLastName() string {
	return ud.lastName
}
func (ud *userDomain) GetIsActive() bool {
	return ud.isActive
}

func (ud *userDomain) EncryptPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	ud.password = string(bytes)
}
