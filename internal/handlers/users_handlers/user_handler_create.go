package users_handlers

import (
	"net/http"
	"server-go/internal/config/logger"
	"server-go/internal/config/validation"
	"server-go/internal/domain"
	"server-go/internal/schemas/request"
	"server-go/internal/view"

	"github.com/gin-gonic/gin"
)

func (uh *userHandlerInterface) UsersHandlerCreate(ctx *gin.Context) {
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

	if err := uh.sv.CreateUser(domainUser); err != nil {
		logger.Error("[ERROR] - Sistema Fly - Create User: ", err)
		ctx.JSON(err.Status, err)
		return
	}

	logger.Info("[INFO] - Sistema Fly - Create User: User created successfully")
	ctx.JSON(http.StatusCreated, view.DomainToResponse(domainUser))
}
