package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"os"
	"strings"
)

type ComputeHandlerSuite struct{}

var _ = Suite(&ComputeHandlerSuite{})

func (s *ComputeHandlerSuite) TestCorrectExpression(c *C) {
	expected := "6"
	output := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  bytes.NewBufferString("4 2 +"),
		Output: output,
	}
	err := handler.Compute()
	c.Assert(err, IsNil)
	c.Assert(output.String(), Equals, expected)
}

func (s *ComputeHandlerSuite) TestIncorrectExpression(c *C) {
	handler := ComputeHandler{
		Input:  strings.NewReader("4 2 ;"),
		Output: os.Stdout,
	}
	err := handler.Compute()
	c.Assert(err, NotNil)
}
