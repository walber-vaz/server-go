package repository

import (
	"context"
	"server-go/internal/config/rest_errors"
	"server-go/internal/domain"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(database *mongo.Database) UseRepository {
	return &userRepository{database}
}

type userRepository struct {
	database *mongo.Database
}

type UseRepository interface {
	CreateUser(ctx context.Context, userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *rest_errors.RestErr)
}
