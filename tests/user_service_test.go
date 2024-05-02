package tests

import (
	"gmotta/login/models"
	"gmotta/login/services"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

type UserServiceSuite struct{}

var _ = Suite(&UserServiceSuite{})

func (s *UserServiceSuite) SetUpTest(c *C) {
	resetDatabase()
}

func TestUserService(t *testing.T) { TestingT(t) }

func (s *UserServiceSuite) TestCreateUser(c *C) {
	userData := models.User{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: time.Now(),
	}

	result, _ := services.CreateUser(userData)
	c.Assert(result.RowsAffected, Equals, int64(1))
}

func (s *UserServiceSuite) TestGet(c *C) {
	resultBefore, _ := services.GetUsers()

	c.Assert(len(resultBefore), Equals, 0)

	userData := models.User{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: time.Now(),
	}
	services.CreateUser(userData)

	resultAfter, _ := services.GetUsers()

	c.Assert(len(resultAfter), Equals, 1)
}
