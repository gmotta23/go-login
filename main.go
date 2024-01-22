package main

import (
	"gmotta/login/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	api := r.Group("/api")

	routes.SetupRouter(api)

	r.Run()

}
