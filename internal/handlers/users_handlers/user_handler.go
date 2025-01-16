package users_handlers

import (
	"server-go/internal/domain/service"

	"github.com/gin-gonic/gin"
)

func NewUserHandlerInterface(service service.UserDomainService) UserHandlerInterface {
	return &userHandlerInterface{
		service,
	}
}

type UserHandlerInterface interface {
	UsersHandlerCreate(ctx *gin.Context)
}

type userHandlerInterface struct {
	service service.UserDomainService
}
