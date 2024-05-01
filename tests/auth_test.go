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
	c.Assert(42, Equals, "42") // fails
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

// func (s *AuthSuite) TestRegister(c *C) {
// 	fmt.Println("testing register")
// 	userData := schemas.CreateUserData{
// 		Email:     "a@b.com",
// 		Name:      "Gustavo",
// 		Password:  "Test",
// 		BirthDate: time.Now(),
// 	}

// 	result, _ := services.CreateUser(userData)
// 	fmt.Println("hello")
// 	fmt.Println(result.Statement.RowsAffected)

// 	fmt.Println(result.Statement.Vars...)
// }

func (s *AuthSuite) TestGet(c *C) {
	result1, _ := services.GetUsers()

	// fmt.Println("users found")
	fmt.Println(result1)
	// fmt.Println(result1.RowsAffected)

	userData := schemas.CreateUserData{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: time.Now(),
	}

	resultCreate, _ := services.CreateUser(userData)

	fmt.Println(resultCreate.Statement.RowsAffected)

	result2, _ := services.GetUsers()

	// fmt.Println("users found")
	fmt.Println(result2)
	// fmt.Println(result2.RowsAffected)
}
