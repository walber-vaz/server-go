package domain

import (
	"encoding/json"

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
	GetJsonValue() (string, error)
	SetID(id string)

	EncryptPassword()
}

func NewUserDomain(firstName, lastName, email, password, phone, role string, isActive bool) UserDomainInterface {
	return &userDomain{
		ID:        "",
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
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      string
	Phone     string
	IsActive  bool
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

func (ud *userDomain) GetJsonValue() (string, error) {
	data, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}
func (ud *userDomain) GetPassword() string {
	return ud.Password
}
func (ud *userDomain) GetRole() string {
	return ud.Role
}
func (ud *userDomain) GetPhone() string {
	return ud.Phone
}
func (ud *userDomain) GetFirstName() string {
	return ud.FirstName
}
func (ud *userDomain) GetLastName() string {
	return ud.LastName
}
func (ud *userDomain) GetIsActive() bool {
	return ud.IsActive
}

func (ud *userDomain) EncryptPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	ud.Password = string(bytes)
}
