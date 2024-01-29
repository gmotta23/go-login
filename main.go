package main

import (
	"fmt"
	"gmotta/login/database"
	"gmotta/login/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	database.Connect()

	r := gin.Default()

	api := r.Group("/api")

	routes.SetupRouter(api)

	r.Run()

}
