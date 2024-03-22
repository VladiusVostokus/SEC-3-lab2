package lab2

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type CalculatePostfixSuite struct{}

var _ = Suite(&CalculatePostfixSuite{})

func (s *CalculatePostfixSuite) TestCalculatePostfixSimpleExpression(c *C) {
	result, err := CalculatePostfix("4 2 -")
	c.Assert(err, IsNil)
	c.Assert(result, Equals, 2)
}
