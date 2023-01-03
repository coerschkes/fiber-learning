package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type configurationProperty string

const (
	dbHost       configurationProperty = "DB_HOST"
	dbName       configurationProperty = "DB_NAME"
	dbUser       configurationProperty = "DB_USER"
	dbPort       configurationProperty = "DB_PORT"
	dbPassword   configurationProperty = "DB_PASSWORD"
	serverPort   configurationProperty = "SERVER_PORT"
	propertyFile configurationProperty = ".env"
)

func LoadProperty(key configurationProperty) string {
	err := godotenv.Load(string(propertyFile))
	if err != nil {
		fmt.Print("Error loading property file")
	}
	return os.Getenv(string(key))
}

func LoadDBHost() string {
	return LoadProperty(dbHost)
}

func LoadDBName() string {
	return LoadProperty(dbName)
}

func LoadDBUser() string {
	return LoadProperty(dbUser)
}

func LoadDBPort() string {
	return LoadProperty(dbPort)
}

func LoadDBPassword() string {
	return LoadProperty(dbPassword)
}

func LoadServerPort() string {
	return LoadProperty(serverPort)
}
