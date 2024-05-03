package database

import (
	"gmtc/login/models"
	"time"
)

type UserSubset struct {
	ID        uint
	Name      string
	Email     string
	BirthDate time.Time
}

func CreateUser(userData models.User) (*UserSubset, error) {
	result := Session.Create(&userData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &UserSubset{
			ID:        userData.ID,
			Name:      userData.Name,
			Email:     userData.Email,
			BirthDate: userData.BirthDate},
		nil
}

func GetUsers() ([]UserSubset, error) {
	var users []UserSubset
	result := Session.Model(&models.User{}).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
