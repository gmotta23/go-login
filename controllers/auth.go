package controllers

import (
	"fmt"
	"net/http"
	"time"

	"gmtc/login/models"
	"gmtc/login/services"

	"github.com/gin-gonic/gin"
)

type RegisterBody struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	BirthDate string `json:"birthDate" binding:"required"`
}

func SerializeUserData(userData RegisterBody) (models.User, error) {
	loc, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		return models.User{}, err
	}

	birthDate, err := time.ParseInLocation("02-01-2006", userData.BirthDate, loc)

	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		BirthDate: birthDate,
	}, nil
}

func CreateUser(c *gin.Context) {
	var createUserData RegisterBody

	if err := c.ShouldBindJSON(&createUserData); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	userData, err := SerializeUserData(createUserData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	result, err := services.CreateUser(userData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
