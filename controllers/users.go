package controllers

import (
	"net/http"

	"gmotta/login/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	result, err := services.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})

}
