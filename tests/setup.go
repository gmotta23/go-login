package tests

import (
	"gmtc/login/database"
	"gmtc/login/models"
)

func ResetDatabase() {
	database.Connect()

	database.Session.Exec("drop schema public cascade;")
	database.Session.Exec("create schema public;")

	database.Session.AutoMigrate(models.GetAllModels()...)
}
