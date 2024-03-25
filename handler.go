package lab2

import "strconv"


// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := handler.Input
	writer := handler.Output
	res, err := CalculatePostfix()
	if err != nil {
		outErr := string(err.Error())
		writer.WriteString(outErr)
		return
	}
	write.WriteString(strconv.Itoa(res))
	return nil
}
