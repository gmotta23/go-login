package tests

import (
	"gmtc/login/database"
	"gmtc/login/models"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

type UserDatabaseSuite struct{}

var _ = Suite(&UserDatabaseSuite{})

func (s *UserDatabaseSuite) SetUpTest(c *C) {
	resetDatabase()
}

func TestUserService(t *testing.T) { TestingT(t) }

func (s *UserDatabaseSuite) TestCreateUser(c *C) {
	userData := models.User{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: time.Now(),
	}

	result, _ := database.CreateUser(userData)
	c.Assert(result.RowsAffected, Equals, int64(1))
}

func (s *UserDatabaseSuite) TestGet(c *C) {
	resultBefore, _ := database.GetUsers()

	c.Assert(len(resultBefore), Equals, 0)

	userData := models.User{
		Email:     "a@b.com",
		Name:      "Gustavo",
		Password:  "Test",
		BirthDate: time.Now(),
	}
	database.CreateUser(userData)

	resultAfter, _ := database.GetUsers()

	c.Assert(len(resultAfter), Equals, 1)
}
