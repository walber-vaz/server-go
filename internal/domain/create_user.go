package domain

import (
	"fmt"
	"go.uber.org/zap"
	"server-go/internal/config/logger"
	"server-go/internal/config/rest_errors"
)

func (ud *userDomain) CreateUser() *rest_errors.RestErr {
	logger.Info("[INFO] - Sistema Fly - User domain - CreateUser", zap.String("journey", "CreateUser"))
	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
