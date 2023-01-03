package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConfigurationProperty string

const (
	DB_HOST       ConfigurationProperty = "DB_HOST"
	DB_NAME       ConfigurationProperty = "DB_NAME"
	DB_USER       ConfigurationProperty = "DB_USER"
	DB_PORT       ConfigurationProperty = "DB_PORT"
	DB_PASSWORD   ConfigurationProperty = "DB_PASSWORD"
	SERVER_PORT   ConfigurationProperty = "SERVER_PORT"
	PROPERTY_FILE string                = ".env"
)

func LoadProperty(key ConfigurationProperty) string {
	err := godotenv.Load(PROPERTY_FILE)
	if err != nil {
		fmt.Print("Error loading property file")
	}
	return os.Getenv(string(key))
}
