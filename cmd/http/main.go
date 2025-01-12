package main

import (
	"log"
	"server-go/internal/config"
	"server-go/internal/config/logger"
	"server-go/internal/router"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logger.InitLogger(cfg.App)
	router.InitRouter(cfg.HTTP)
}
