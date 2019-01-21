package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	PostgresURL string
	JwtSecret   string
	JwtIssuer   string
}

func GetConfig() Config {
	var config Config
	var err error
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "prod" {
		if config.Port, err = strconv.Atoi(os.Getenv("PROD_PORT")); err != nil {
			log.Fatal(err)
		}
		config.PostgresURL = os.Getenv("PROD_POSTGRES_URL")
		config.JwtSecret = os.Getenv("PROD_JWT_SECRET")
		config.JwtIssuer = os.Getenv("PROD_JWT_ISSUER")
		return config
	}
	// "dev"
	if config.Port, err = strconv.Atoi(os.Getenv("DEV_PORT")); err != nil {
		log.Fatal(err)
	}
	config.PostgresURL = os.Getenv("DEV_POSTGRES_URL")
	config.JwtSecret = os.Getenv("DEV_JWT_SECRET")
	config.JwtIssuer = os.Getenv("DEV_JWT_ISSUER")
	return config
}
