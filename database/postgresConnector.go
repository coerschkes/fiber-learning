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

type PostgresConnector struct{}

func NewPostgresConnector() *PostgresConnector {
	return &PostgresConnector{}
}

func (c PostgresConnector) Connect() *gorm.DB {
	DB, err := gorm.Open(postgres.Open(c.buildConnectionString()))
	if err != nil {
		panic("failed to connect database. shutting down.")
	}
	DB.AutoMigrate(&model.Note{})
	log.Println("Connection to database " + config.LoadDBHost() + "/" + config.LoadDBName() + " established.")
	return DB
}

func (c PostgresConnector) buildConnectionString() string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.LoadDBHost(),
		c.loadPort(),
		config.LoadDBUser(),
		config.LoadDBPassword(),
		config.LoadDBName())
	return dsn
}

func (c PostgresConnector) loadPort() uint64 {
	port, err := strconv.ParseUint(config.LoadDBPort(), 10, 32)
	if err != nil {
		log.Println("Unable to parse port")
		port = 5432
	}
	return port
}
