package tests

import (
	"gmtc/login/utils"
	"testing"

	. "gopkg.in/check.v1"
)

type ConvertUtilsSuite struct{}

var _ = Suite(&ConvertUtilsSuite{})

func TestConvertUtils(t *testing.T) { TestingT(t) }

func (s *ConvertUtilsSuite) TestUintToString(c *C) {
	uintNumber := uint(8)
	result := utils.UintToString(uintNumber)
	c.Assert(result, Equals, "8")
}

func (s *ConvertUtilsSuite) TestFloat64ToString(c *C) {
	float64Number := float64(32.123)
	result := utils.Float64ToString(float64Number)
	c.Assert(result, Equals, "32")
}
