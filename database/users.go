package database

import (
	"gmtc/login/models"
)

func CreateUser(userData models.User) (*models.User, error) {
	result := Session.Create(&userData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userData, nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := Session.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
