package router

import (
	"server-go/internal/handlers/users_handlers"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.RouterGroup, uh users_handlers.UserHandlerInterface) {
	baseUrl := r.Group("/api/v1")

	users := baseUrl.Group("/users")
	{
		users.POST("", uh.UsersHandlerCreate)
	}
}
