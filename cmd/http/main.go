package main

import (
	"log"
	"server-go/internal/config"
	"server-go/internal/router"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	router.InitRouter(cfg.HTTP)
}
