package lab2


import (
	"io"
	"strconv"
)


// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	
	reader := ch.Input
	writer := ch.Output
	barr := make([]byte, 0, 1)
	oneElemArr := make([]byte, 1)

	for {
		_, err := reader.Read(oneElemArr)
		if err != nil {
			break
		}	
		barr = append(barr, oneElemArr[0])
	}
	res, err := CalculatePostfix(string(barr))

	if err != nil {
		return err
	}
	output := strconv.Itoa(res)
	writer.Write([]byte(output))
	
	return nil
}
