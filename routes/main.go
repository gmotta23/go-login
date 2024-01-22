package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong!"})
	})

	RegisterAuthHTTPEndpoints(router)

	return router
}
