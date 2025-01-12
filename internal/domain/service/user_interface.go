package service

import (
	"server-go/internal/config/rest_errors"
	"server-go/internal/domain"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct{}

type UserDomainService interface {
	CreateUser(domain.UserDomainInterface) *rest_errors.RestErr
	//UpdateUser(string) *rest_errors.RestErr
	//FindUser(string) (domain.UserDomainInterface, *rest_errors.RestErr)
	//DeleteUser(string) *rest_errors.RestErr
}
