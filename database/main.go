package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB

func Connect() (*gorm.DB, error) {
	var err error

	dsn := os.Getenv("DATABASE_DSN")

	fmt.Println("Connecting to database")
	fmt.Println(dsn)

	Session, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return Session, err

}
