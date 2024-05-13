package tests

import (
	"gmtc/login/controllers"
	"gmtc/login/setup"
	. "gmtc/login/tests"
	"gmtc/login/util"
	"net/http/httptest"
	"testing"

	. "gopkg.in/check.v1"
)

type UserRouteSuite struct{}

var _ = Suite(&UserRouteSuite{})

func (s *UserRouteSuite) SetUpTest(c *C) {
	ResetDatabase()
}

func TestUserRoute(t *testing.T) { TestingT(t) }

func (s *UserRouteSuite) TestPing(c *C) {
	router := setup.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := MakeRequest("GET", "/api/ping", nil)

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

	req, _ := MakeRequest("POST", "/api/auth/register", payload)

	router.ServeHTTP(w, req)

	c.Assert(201, Equals, w.Code)

	responseBody := LoadResponseBody(w)

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

	req, _ := MakeRequest("POST", "/api/auth/login", payload)

	router.ServeHTTP(w, req)

	c.Assert(200, Equals, w.Code)

	responseBody := LoadResponseBody(w)

	user := responseBody["user"].(map[string]interface{})
	token := responseBody["token"].(string)

	c.Assert(util.Float64ToString(float64(user["ID"].(float64))), Equals, util.UintToString(userDB.ID))
	c.Assert(user["Password"], IsNil)
	c.Assert(user["Email"], Equals, payload.Email)
	c.Assert(token, NotNil)
}
