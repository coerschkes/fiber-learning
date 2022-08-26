package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/coerschkes/fiber-learning/config"
	"github.com/coerschkes/fiber-learning/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnector struct {
	DB *gorm.DB
}

func NewPostgresConnector() *PostgresConnector {
	return &PostgresConnector{}
}

// connect to the database
func (c PostgresConnector) Connect() *gorm.DB {
	DB, err := gorm.Open(postgres.Open(c.buildConnectionString()))
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.Note{}, &model.User{})

	log.Println("Connection to database " + config.LoadProperty(config.DB_HOST) + "/" + config.LoadProperty(config.DB_NAME) + " established.")
	return DB
}

// build the postgres connection url
func (c PostgresConnector) buildConnectionString() string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.LoadProperty(config.DB_HOST),
		c.loadPort(),
		config.LoadProperty(config.DB_USER),
		config.LoadProperty(config.DB_PASSWORD),
		config.LoadProperty(config.DB_NAME))
	return dsn
}

func (c PostgresConnector) loadPort() uint64 {
	port, err := strconv.ParseUint(config.LoadProperty(config.DB_PORT), 10, 32)
	if err != nil {
		log.Println("Unable to parse port")
		port = 5432
	}
	return port
}
