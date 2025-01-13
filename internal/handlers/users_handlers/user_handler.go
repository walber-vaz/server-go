package users_handlers

import (
	"server-go/internal/domain/service"

	"github.com/gin-gonic/gin"
)

func NewUserHandlerInterface(sv service.UserDomainService) UserHandlerInterface {
	return &userHandlerInterface{
		sv: sv,
	}
}

type UserHandlerInterface interface {
	UsersHandlerCreate(ctx *gin.Context)
}

type userHandlerInterface struct {
	sv service.UserDomainService
}
