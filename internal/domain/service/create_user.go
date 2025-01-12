package service

import (
	"fmt"
	"go.uber.org/zap"
	"server-go/internal/config/logger"
	"server-go/internal/config/rest_errors"
	"server-go/internal/domain"
)

func (ud *userDomainService) CreateUser(userDomain domain.UserDomainInterface) *rest_errors.RestErr {
	logger.Info("[INFO] - Sistema Fly - User domain - CreateUser", zap.String("journey", "CreateUser"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
