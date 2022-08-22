package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConfigurationProperty string

const (
	DB_HOST     ConfigurationProperty = "DB_HOST"
	DB_NAME     ConfigurationProperty = "DB_NAME"
	DB_USER     ConfigurationProperty = "DB_USER"
	DB_PORT     ConfigurationProperty = "DB_PORT"
	DB_PASSWORD ConfigurationProperty = "DB_PASSWORD"
)

// LoadProperty func to get env value
func LoadProperty(key ConfigurationProperty) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	// Return the value of the variable
	return os.Getenv(string(key))
}
