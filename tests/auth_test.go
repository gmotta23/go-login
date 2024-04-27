package tests

import (
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestBasic(c *C) {
	// c.Assert(42, Equals, "42") // fails
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}
