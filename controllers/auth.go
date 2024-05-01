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

	result, err := services.CreateUser(createUserData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
