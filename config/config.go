package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Konfig function untuk get key dari env
func Config(key string) string {
	// memuat file .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
