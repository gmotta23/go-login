package routes

import (
	"gmotta/login/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUsersHTTPEndpoints(router *gin.RouterGroup) *gin.RouterGroup {
	usersGroup := router.Group("/users")

	usersGroup.GET("/list", controllers.GetUsers)

	return router
}
