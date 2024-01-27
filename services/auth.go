package services

import (
	"gmotta/login/database"
	"gmotta/login/models"
	"gmotta/login/schemas"
)

func NewUserFromCreateUserData(userData schemas.CreateUserData) models.User {
	return models.User{
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		BirthDate: userData.BirthDate,
	}
}

func CreateUser(userData schemas.CreateUserData) {
	user := NewUserFromCreateUserData(userData)
	database.Session.Create(&user)
}
