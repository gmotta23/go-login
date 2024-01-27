package main

import (
	"gmotta/login/database"
	"gmotta/login/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r := gin.Default()

	api := r.Group("/api")

	routes.SetupRouter(api)

	r.Run()

}
