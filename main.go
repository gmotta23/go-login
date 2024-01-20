package main

import (
	"fmt"
	"gmotta/login/models"
	"net/http"

	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rsc.io/quote"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	BirthDate string `json:"birthDate" binding:"required"`
}

func main() {
	stmts, err := gormschema.New("postgres").Load(&models.User{}, &models.Pet{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)

	fmt.Println(models.Pet{})

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database")

	db.AutoMigrate(&Product{})
	// CHECK MIGRATIONS IN https://atlasgo.io/guides/orms/gorm

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
