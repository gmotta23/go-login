package routes

import (
	"gmtc/login/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthHTTPEndpoints(router *gin.RouterGroup) *gin.RouterGroup {
	authGroup := router.Group("/auth")

	authGroup.POST("/register", controllers.Register)
	authGroup.POST("/login", controllers.Login)

	return router
}
