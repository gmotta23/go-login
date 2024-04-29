package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB

func Connect() (*gorm.DB, error) {
	var err error

	dsn := os.Getenv("DATABASE_DSN")

	Session, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return Session, err

}
