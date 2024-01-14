package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rsc.io/quote"
)

type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	BirthDate string `json:"birthDate" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong!"})
	})
	// create route for register
	// create route for login POST
	// create protected route

	r.GET("/quote", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": quote.Glass()})
	})

	r.POST("/auth/register", func(c *gin.Context) {
		fmt.Println("hit register")

		var request RegisterRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
			return
		}

		// body, err := c.GetRawData() // no validation
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		// 	return
		// }

		fmt.Printf("Received registration request: %+v\n", request)

		// get info from request body
		// c.bod

		c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	})

	r.Run()
}
