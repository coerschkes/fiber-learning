package database

import "gorm.io/gorm"

type DatabaseConnector interface {
	Connect() *gorm.DB
}
