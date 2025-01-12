package users_handlers

import (
	"fmt"
	"net/http"
	"server-go/internal/config/logger"
	"server-go/internal/config/validation"

	"github.com/gin-gonic/gin"
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

	logger.Info(fmt.Sprintf("[INFO] - Sistema Fly - Create User: %v", userRequest))
	ctx.JSON(http.StatusOK, userRequest)
}
