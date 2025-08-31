package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config armazena as configurações do banco de dados
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() *Config {
	// Carrega as variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}
}