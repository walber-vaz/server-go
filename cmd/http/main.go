package main

import (
	"log"
	"server-go/internal/config"
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

	userService := service.NewUserDomainService()
	userHandler := users_handlers.NewUserHandlerInterface(userService)

	logger.InitLogger(cfg.App)
	router.InitRouter(cfg.HTTP, userHandler)
}
