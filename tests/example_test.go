package tests

import (
	"fmt"
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

type ExampleSuite struct{}

var _ = Suite(&ExampleSuite{})

func (s *ExampleSuite) SetUpTest(c *C) {
	fmt.Println("Hello from setu up test! (Example)!!!")
	fmt.Println("Hello from setu up test! (Example)!!!")
}

func TestExample(t *testing.T) { TestingT(t) }

func (s *ExampleSuite) TestBasic1(c *C) {
	// c.Assert(42, Equals, "42") // fails
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

func (s *ExampleSuite) TestBasic2(c *C) {
	// c.Assert(42, Equals, "42") // fails
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}
