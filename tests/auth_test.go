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

func (s *AuthSuite) TestBasic3(c *C) {
	fmt.Println("hello")
	fmt.Println("hello 2")
	fmt.Println("hello 2")
	c.Assert(42, Equals, "42") // fails
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

func (s *AuthSuite) TestBasic4(c *C) {
	// c.Assert(42, Equals, "42") // fails
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

func (s *AuthSuite) TestRegister(c *C) {
	userData := schemas.CreateUserData{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: time.Now(),
	}

	services.CreateUser(userData)
}

func (s *AuthSuite) TestGet(c *C) {
	result1 := services.GetUsers()

	fmt.Println("users found")
	fmt.Println(result1)
	fmt.Println(result1.RowsAffected)

	userData := schemas.CreateUserData{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: time.Now(),
	}

	services.CreateUser(userData)

	result2 := services.GetUsers()

	fmt.Println("users found")
	fmt.Println(result2)
	fmt.Println(result2.RowsAffected)
}
