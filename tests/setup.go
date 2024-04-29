package tests

import (
	"gmotta/login/database"
	"gmotta/login/models"
)

func resetDatabase() {
	database.Connect()

	database.Session.Exec("drop schema public cascade;")
	database.Session.Exec("create schema public;")

	database.Session.AutoMigrate(&models.User{})
}
