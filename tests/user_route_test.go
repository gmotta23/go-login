package tests

import (
	"gmotta/login/setup"
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

func TestUserRoute(t *testing.T) { TestingT(t) }

func (s *UserRouteSuite) TestPing(c *C) {
	router := setup.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)

	router.ServeHTTP(w, req)

	c.Assert(200, Equals, w.Code)
	c.Assert(w.Body.String(), Equals, "{\"message\":\"pong!\"}")
}
