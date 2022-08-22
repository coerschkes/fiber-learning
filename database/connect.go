package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/coerschkes/fiber-learning/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	port, err := strconv.ParseUint(config.LoadProperty(config.DB_PORT), 10, 32)
	if err != nil {
		log.Println("Unable to parse port")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.LoadProperty(config.DB_HOST),
		port,
		config.LoadProperty(config.DB_USER),
		config.LoadProperty(config.DB_PASSWORD),
		config.LoadProperty(config.DB_NAME))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection to database " + config.LoadProperty(config.DB_HOST) + "/" + config.LoadProperty(config.DB_NAME) + " established.")
}
