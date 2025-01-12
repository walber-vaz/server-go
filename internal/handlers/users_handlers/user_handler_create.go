package users_handlers

import (
	"net/http"
	"server-go/internal/config/validation"

	"github.com/gin-gonic/gin"
	"server-go/internal/schemas/request"
)

func UsersHandlerCreate(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Status, restErr)
		return
	}

	ctx.JSON(http.StatusOK, userRequest)
}
