package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"server-go/internal/config"
	"server-go/internal/config/rest_errors"
	"strings"
)

func InitRouter(cfg *config.HTTP) {
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

	initRoutes(&router.RouterGroup)

	router.NoRoute(func(c *gin.Context) {
		err := rest_errors.NewNotFoundError("Ops! Rota não encontrada.")
		c.JSON(err.Status, err)
	})

	listAddress := fmt.Sprintf("%s:%s", cfg.URL, cfg.Port)
	if err := router.Run(listAddress); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
		os.Exit(1)
	}
}
