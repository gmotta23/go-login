package tests

import (
	"bytes"
	"encoding/json"
	"gmtc/login/controllers"
	"gmtc/login/setup"
	"net/http"
	"net/http/httptest"
	"testing"

	. "gopkg.in/check.v1"
)

type UserRouteSuite struct{}

var _ = Suite(&UserRouteSuite{})

func (s *UserRouteSuite) SetUpTest(c *C) {
	resetDatabase()
}

func makeRequest(method string, url string, body any) (*http.Request, error) {
	if body == nil {
		return http.NewRequest(method, url, nil)
	}

	encodedBody, _ := json.Marshal(body)
	buffer := bytes.NewBuffer(encodedBody)

	return http.NewRequest(method, url, buffer)
}

func TestUserRoute(t *testing.T) { TestingT(t) }

func (s *UserRouteSuite) TestPing(c *C) {
	router := setup.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := makeRequest("GET", "/api/ping", nil)

	router.ServeHTTP(w, req)

	c.Assert(200, Equals, w.Code)
	c.Assert(w.Body.String(), Equals, "{\"message\":\"pong!\"}")
}

func (s *UserRouteSuite) TestRegister(c *C) {
	router := setup.SetupRouter()
	w := httptest.NewRecorder()
	payload := controllers.RegisterBody{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: "06-07-1995",
	}

	req, _ := makeRequest("POST", "/api/auth/register", payload)

	router.ServeHTTP(w, req)

	c.Assert(201, Equals, w.Code)

	var responseBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &responseBody)

	user, _ := responseBody["user"].(map[string]interface{})

	_, idExists := user["ID"]
	c.Assert(idExists, Equals, true)
	c.Assert(user["Email"], Equals, payload.Email)
	_, passwordExists := user["Password"]
	c.Assert(passwordExists, Equals, false)
}
