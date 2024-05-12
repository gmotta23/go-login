package tests

import (
	"fmt"
	"gmtc/login/services"
	"testing"

	. "gopkg.in/check.v1"
)

type JWTServiceSuite struct{}

var _ = Suite(&JWTServiceSuite{})

func TestJWTService(t *testing.T) { TestingT(t) }

func (s *JWTServiceSuite) TestCreateJWTToken(c *C) {
	id := uint(10)
	result, _ := services.CreateJWTToken(id)

	fmt.Println("hello!!!")
	fmt.Println(result)
	fmt.Println("Got here!")

	c.Assert(1, Equals, 2)
}
