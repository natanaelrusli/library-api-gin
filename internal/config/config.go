package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig DatabaseConfig
}

type DatabaseConfig struct {
	Port     int64
	Host     string
	DbName   string
	Username string
	Password string
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	return &Config{
		DBConfig: initDbConfig(),
	}, nil
}

func initDbConfig() DatabaseConfig {
	host := os.Getenv("DATABASE_HOST")
	port, _ := strconv.ParseInt(os.Getenv("DATABASE_PORT"), 10, 32)
	dbName := os.Getenv("DATABASE_NAME")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	return DatabaseConfig{
		Host:     host,
		Port:     port,
		DbName:   dbName,
		Username: dbUsername,
		Password: dbPassword,
	}
}
