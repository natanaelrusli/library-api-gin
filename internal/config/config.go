package config

import (
	"log"
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

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	return &Config{
		DBConfig: initDbConfig(),
	}
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
