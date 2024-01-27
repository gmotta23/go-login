package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserData struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	BirthDate string `json:"birthDate" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var createUserData CreateUserData

	if err := c.ShouldBindJSON(&createUserData); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	fmt.Printf("Received registration request: %+v\n", createUserData)

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
