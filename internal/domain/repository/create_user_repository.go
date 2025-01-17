package repository

import (
	"context"
	"fmt"
	"server-go/internal/config/logger"
	"server-go/internal/config/rest_errors"
	"server-go/internal/domain"

	"go.uber.org/zap"
)

func (userRepository *userRepository) CreateUser(ctx context.Context, userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("[INFO] - Sistema Fly - User repository - CreateUser", zap.String("journey", "CreateUser"))

	collection := userRepository.database.Collection("users")

	value, err := userDomain.GetJsonValue()
	if err != nil {
		return nil, rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get json value from userDomain: %s", err.Error()))
	}
	result, err := collection.InsertOne(ctx, value)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to insert user: %s", err.Error()))
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
