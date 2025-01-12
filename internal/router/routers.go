package router

import (
	"github.com/gin-gonic/gin"
	"server-go/internal/handlers/users_handlers"
)

func initRoutes(r *gin.RouterGroup) {
	baseUrl := r.Group("/api/v1")

	users := baseUrl.Group("/users")
	{
		users.POST("", users_handlers.UsersHandlerCreate)
	}
}
