package routes

import (
	"gmotta/login/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthHTTPEndpoints(router *gin.RouterGroup) *gin.RouterGroup {
	authGroup := router.Group("/auth")

	authGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong from auth!"})
	})

	authGroup.POST("/register", controllers.CreateUser)

	return router
}
