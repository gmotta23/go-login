package tests

import (
	"gmtc/login/database"
	"gmtc/login/models"
	"gmtc/login/services"
	"time"
)

func ResetDatabase() {
	database.Connect()

	database.Session.Exec("drop schema public cascade;")
	database.Session.Exec("create schema public;")

	database.Session.AutoMigrate(models.GetAllModels()...)
}

func RegisterUser() (*database.UserSubset, string) {
	password := "Test"
	hashedPassword, _ := services.HashPassword(password)
	userData := models.User{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  hashedPassword,
		BirthDate: time.Now(),
	}
	user, _ := database.CreateUser(userData)
	return user, password
}
