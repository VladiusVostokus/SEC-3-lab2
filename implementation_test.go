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

func (s *CalculatePostfixSuite) TestCalculatePostfixComplexExpression(c *C) {
	result, err := CalculatePostfix("4 2 - 3 * 5 +")
	c.Assert(err, IsNil)
	c.Assert(result, Equals, 11)
}
func (s *CalculatePostfixSuite) TestCalculatePostfixComplexExpression2(c *C) {
	result, err := CalculatePostfix("2 2 + 3 ^ 4 - 10 /")
	c.Assert(err, IsNil)
	c.Assert(result, Equals, 6)
}
func (s *CalculatePostfixSuite) TestCalculatePostfixEmptyExpression(c *C) {
	_, err := CalculatePostfix("")
	c.Assert(err, NotNil)
}
func (s *CalculatePostfixSuite) TestCalculatePostfixInvalidSymbol(c *C) {
	_, err := CalculatePostfix("4 2 $")
	c.Assert(err, NotNil)
}
func (s *CalculatePostfixSuite) TestCalculatePostfixInsufficientOperands(c *C) {
	_, err := CalculatePostfix("+")
	c.Assert(err, NotNil)
}
func (s *CalculatePostfixSuite) TestCalculatePostfixSpaceExpression(c *C) {
	_, err := CalculatePostfix(" ")
	c.Assert(err, NotNil)
}
