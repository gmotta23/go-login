package tests

import (
	"fmt"
	"gmotta/login/schemas"
	"gmotta/login/services"
	"io"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

type AuthSuite struct{}

var _ = Suite(&AuthSuite{})

func (s *AuthSuite) SetUpTest(c *C) {
	resetDatabase()
}

func TestAuth(t *testing.T) { TestingT(t) }

func (s *AuthSuite) TestBasic4(c *C) {
	// c.Assert(42, Equals, "42") // fails
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

func (s *AuthSuite) TestRegister(c *C) {
	fmt.Println("testing register")
	userData := schemas.CreateUserData{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: time.Now(),
	}

	result, _ := services.CreateUser(userData)
	c.Assert(result.RowsAffected, Equals, int64(1))
}

func (s *AuthSuite) TestGet(c *C) {
	resultBefore, _ := services.GetUsers()

	c.Assert(len(resultBefore), Equals, 0)

	userData := schemas.CreateUserData{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: time.Now(),
	}
	services.CreateUser(userData)

	resultAfter, _ := services.GetUsers()

	c.Assert(len(resultAfter), Equals, 1)
}
