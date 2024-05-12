package tests

import (
	"bytes"
	"encoding/json"
	"gmtc/login/controllers"
	"gmtc/login/setup"
	. "gmtc/login/tests"
	"net/http"
	"net/http/httptest"
	"testing"

	. "gopkg.in/check.v1"
)

type UserRouteSuite struct{}

var _ = Suite(&UserRouteSuite{})

func (s *UserRouteSuite) SetUpTest(c *C) {
	ResetDatabase()
}

func makeRequest(method string, url string, body any) (*http.Request, error) {
	if body == nil {
		return http.NewRequest(method, url, nil)
	}

	encodedBody, _ := json.Marshal(body)
	buffer := bytes.NewBuffer(encodedBody)

	return http.NewRequest(method, url, buffer)
}

func loadResponseBody(w *httptest.ResponseRecorder) map[string]interface{} {
	var responseBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &responseBody)

	return responseBody
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

	responseBody := loadResponseBody(w)

	user := responseBody["user"].(map[string]interface{})
	token := responseBody["token"].(string)

	c.Assert(user["ID"], NotNil)
	c.Assert(user["Password"], IsNil)
	c.Assert(user["Email"], Equals, payload.Email)
	c.Assert(token, NotNil)
}

func (s *UserRouteSuite) TestLogin(c *C) {
	userDB, password := RegisterUser()

	router := setup.SetupRouter()
	w := httptest.NewRecorder()

	payload := controllers.LoginBody{
		Email:    userDB.Email,
		Password: password,
	}

	req, _ := makeRequest("POST", "/api/auth/login", payload)

	router.ServeHTTP(w, req)

	c.Assert(200, Equals, w.Code)

	responseBody := loadResponseBody(w)

	user := responseBody["user"].(map[string]interface{})
	token := responseBody["token"].(string)

	c.Assert(user["ID"], Equals, userDB.ID)
	c.Assert(user["Password"], IsNil)
	c.Assert(user["Email"], Equals, payload.Email)
	c.Assert(token, NotNil)
}
