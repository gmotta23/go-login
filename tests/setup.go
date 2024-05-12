package tests

import (
	"bytes"
	"encoding/json"
	"gmtc/login/database"
	"gmtc/login/models"
	"gmtc/login/services"
	"net/http"
	"net/http/httptest"
	"time"
)

func ResetDatabase() {
	database.Connect()

	database.Session.Exec("drop schema public cascade;")
	database.Session.Exec("create schema public;")

	database.Session.AutoMigrate(models.GetAllModels()...)
}

func MakeRequest(method string, url string, body any) (*http.Request, error) {
	if body == nil {
		return http.NewRequest(method, url, nil)
	}

	encodedBody, _ := json.Marshal(body)
	buffer := bytes.NewBuffer(encodedBody)

	return http.NewRequest(method, url, buffer)
}

func LoadResponseBody(w *httptest.ResponseRecorder) map[string]interface{} {
	var responseBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &responseBody)

	return responseBody
}

func RegisterUser() (*database.UserSubset, string) {
	password := "Test"
	hashedPassword, _ := services.HashPassword(password)
	userData := models.User{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  hashedPassword,
		BirthDate: time.Now(),
	}
	user, _ := database.CreateUser(userData)
	return user, password
}
