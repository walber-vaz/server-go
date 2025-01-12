package users_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-go/internal/config/logger"
	"server-go/internal/config/validation"
	"server-go/internal/domain"
	"server-go/internal/domain/service"
	"server-go/internal/schemas/request"
)

func UsersHandlerCreate(ctx *gin.Context) {
	logger.Info("[INFO] - Sistema Fly - Create User")
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("[ERROR] - Sistema Fly - Create User: ", err)
		restErr := validation.ValidateUserError(err)

		ctx.JSON(restErr.Status, restErr)
		return
	}

	domainUser := domain.NewUserDomain(
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Email,
		userRequest.Password,
		userRequest.Phone,
		userRequest.Role,
		userRequest.IsActive,
	)
	serviceUser := service.NewUserDomainService()

	if err := serviceUser.CreateUser(domainUser); err != nil {
		logger.Error("[ERROR] - Sistema Fly - Create User: ", err)
		ctx.JSON(err.Status, err)
		return
	}

	logger.Info("[INFO] - Sistema Fly - Create User: User created successfully")
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
