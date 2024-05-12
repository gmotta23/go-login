package controllers

import (
	"fmt"
	"net/http"
	"time"

	"gmtc/login/database"
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

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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

func Register(c *gin.Context) {
	var createUserData RegisterBody

	if err := c.ShouldBindJSON(&createUserData); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	userData, err := SerializeUserData(createUserData)
	userData.Password, _ = services.HashPassword(userData.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	user, err := database.CreateUser(userData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": err.Error()})
		return
	}

	token, err := services.CreateJWTToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user, "token": token})
}

func Login(c *gin.Context) {
	var loginData LoginBody

	if err := c.ShouldBindJSON(&loginData); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	userInDB, err := database.FindUserByEmail(loginData.Email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found", "error": err.Error()})
	}

	passwordMatch := services.CheckPasswordHash(loginData.Password, userInDB.Password)

	fmt.Println(loginData.Password)
	fmt.Println(userInDB.Password)

	if passwordMatch {
		userSubset := database.ToUserSubset(*userInDB)
		token, _ := services.CreateJWTToken(userSubset.ID)
		c.JSON(http.StatusOK, gin.H{"user": userSubset, "token": token})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect user or password"})

}
