package tests

import (
	"fmt"
	"io"
	"testing"

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