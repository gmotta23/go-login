package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	BirthDate string `json:"birthDate" binding:"required"`
}

func RegisterAuthHTTPEndpoints(router *gin.RouterGroup) *gin.RouterGroup {
	authGroup := router.Group("/auth")

	authGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong from auth!"})
	})

	authGroup.POST("/register", func(c *gin.Context) {
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

	return router
}
