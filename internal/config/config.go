package config

import (
	"github.com/joho/godotenv"
	"os"
)

type (
	Container struct {
		App  *App
		DB   *DB
		HTTP *HTTP
	}

	App struct {
		Name string
		Env  string
	}

	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	HTTP struct {
		Port         string
		Env          string
		URL          string
		AllowOrigins string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	http := &HTTP{
		Port:         os.Getenv("HTTP_PORT"),
		Env:          os.Getenv("APP_ENV"),
		URL:          os.Getenv("HTTP_URL"),
		AllowOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	return &Container{
		App:  app,
		DB:   db,
		HTTP: http,
	}, nil
}
