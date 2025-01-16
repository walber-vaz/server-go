package main

import (
	"log"
	"server-go/internal/config"
	"server-go/internal/config/database/mongodb"
	"server-go/internal/config/logger"
	"server-go/internal/domain/service"
	"server-go/internal/handlers/users_handlers"
	"server-go/internal/router"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	logger.InitLogger(cfg.App)
	mongodb.InitConnection(cfg.DB)

	userService := service.NewUserDomainService()
	userHandler := users_handlers.NewUserHandlerInterface(userService)

	router.InitRouter(cfg.HTTP, userHandler)
}
