package setup

import (
	"fmt"
	"gmtc/login/database"
	"gmtc/login/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	routes.SetupRouter(api)

	return r
}

func SetupDatabase() {
	database.Connect()
}

func StartServer() {
	LoadEnv()

	SetupDatabase()

	r := SetupRouter()

	r.Run()
}
