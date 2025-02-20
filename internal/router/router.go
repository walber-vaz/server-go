package router

import (
	"fmt"
	"net/http"
	"os"
	"server-go/internal/config"
	"server-go/internal/config/logger"
	"server-go/internal/config/rest_errors"
	"server-go/internal/handlers/users_handlers"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter(cfg *config.HTTP, uh users_handlers.UserHandlerInterface) {
	logger.Info("[INFO] - Sistema Fly - Iniciando o servidor")
	switch cfg.Env {
	case "test":
		gin.SetMode(gin.TestMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "development":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	corsConfig := cors.Config{
		AllowOrigins:     strings.Split(cfg.AllowOrigins, ","),
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	initRoutes(&router.RouterGroup, uh)

	router.NoRoute(func(c *gin.Context) {
		err := rest_errors.NewNotFoundError("Ops! Rota não encontrada.")
		logger.Error("[ERROR] - Sistema Fly - Rota não encontrada: ", err)
		c.JSON(err.Status, err)
	})

	listAddress := fmt.Sprintf("%s:%s", cfg.URL, cfg.Port)
	if err := router.Run(listAddress); err != nil {
		logger.Error("[ERROR] - Sistema Fly - Erro ao iniciar o servidor: ", err)
		os.Exit(1)
	}
}
