package tests

import (
	"gmtc/login/util"
	"testing"

	. "gopkg.in/check.v1"
)

type ConvertUtilSuite struct{}

var _ = Suite(&ConvertUtilSuite{})

func TestConvertUtil(t *testing.T) { TestingT(t) }

func (s *ConvertUtilSuite) TestUintToString(c *C) {
	uintNumber := uint(8)
	result := util.UintToString(uintNumber)
	c.Assert(result, Equals, "8")
}

func (s *ConvertUtilSuite) TestFloat64ToString(c *C) {
	float64Number := float64(32.123)
	result := util.Float64ToString(float64Number)
	c.Assert(result, Equals, "32")
}
