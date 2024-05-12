package tests

import (
	"gmtc/login/services"
	"gmtc/login/utils"
	"testing"

	. "gopkg.in/check.v1"
)

type JWTServiceSuite struct{}

var _ = Suite(&JWTServiceSuite{})

func TestJWTService(t *testing.T) { TestingT(t) }

func (s *JWTServiceSuite) TestJWTFlow(c *C) {
	id := uint(10)
	signedString, _ := services.CreateJWTToken(id)

	jwtContent, _ := services.ValidateAndParseJWTToken(signedString)

	c.Assert(jwtContent.ID, Equals, utils.UintToString(id))
}
