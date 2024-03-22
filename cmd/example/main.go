package main

import (
	"flag"
	"fmt"
	lab2 "lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	outputFile = flag.String("o", "", "File to output result")
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()
	expression := flag.Arg(0)
	file := flag.Arg(2)

	res, _ := lab2.CalculatePostfix(expression)
	fmt.Println(res)
	fmt.Println(file)
}
