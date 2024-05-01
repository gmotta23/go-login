package services

import (
	"gmotta/login/database"
	"gmotta/login/models"
	"gmotta/login/schemas"

	"gorm.io/gorm"
)

func NewUserFromCreateUserData(userData schemas.CreateUserData) models.User {
	return models.User{
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		BirthDate: userData.BirthDate,
	}
}

func CreateUser(userData schemas.CreateUserData) (*gorm.DB, error) {
	user := NewUserFromCreateUserData(userData)

	result := database.Session.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return result, nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := database.Session.Find(&models.User{})

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
