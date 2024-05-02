package tests

import (
	"gmtc/login/services"
	"testing"

	. "gopkg.in/check.v1"
)

type HashServiceSuite struct{}

var _ = Suite(&HashServiceSuite{})

func TestHashService(t *testing.T) { TestingT(t) }

func (s *HashServiceSuite) TestHashPassword(c *C) {
	password := "test password"
	hash, _ := services.HashPassword(password)
	hash2, _ := services.HashPassword(password)

	c.Assert(password, Not(Equals), hash)
	c.Assert(password, Not(Equals), hash2)
	c.Assert(hash, Not(Equals), hash2)
}

func (s *HashServiceSuite) TestCheckPasswordHash(c *C) {
	password := "test password"

	hash, _ := services.HashPassword(password)

	match := services.CheckPasswordHash(password, hash)

	c.Assert(match, Equals, true)
}
