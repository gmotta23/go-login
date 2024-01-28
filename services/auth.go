package services

import (
	"fmt"
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

func CreateUser(userData schemas.CreateUserData) *gorm.DB {
	user := NewUserFromCreateUserData(userData)

	fmt.Println("registering user")

	result := database.Session.Create(&user)

	if result.Error != nil {
		// Handle the error internally
		fmt.Println(result.Error)
	}

	return result
}
