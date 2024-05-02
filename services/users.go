package services

import (
	"gmtc/login/database"
	"gmtc/login/models"

	"gorm.io/gorm"
)

func CreateUser(userData models.User) (*gorm.DB, error) {
	result := database.Session.Create(&userData)

	if result.Error != nil {
		return nil, result.Error
	}

	return result, nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := database.Session.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
