package lab2


import (
	"io"
	"strconv"
	"bufio"
)


// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	
	reader := bufio.NewReader(ch.Input)
	writer := ch.Output
	expr, _ := reader.ReadString('\n')
	res, err := CalculatePostfix(string(expr))

	if err != nil {
		return err
	}
	output := strconv.Itoa(res)
	writer.Write([]byte(output))
	
	return nil
}
