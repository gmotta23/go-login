package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB

func Connect() (*gorm.DB, error) {
	var err error

	dsn := "host=db-dev user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	Session, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return Session, err

}
