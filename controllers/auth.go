package controllers

import (
	"fmt"
	"net/http"

	"gmotta/login/schemas"
	"gmotta/login/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var createUserData schemas.CreateUserData

	if err := c.ShouldBindJSON(&createUserData); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	services.CreateUser(createUserData)

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
