package main

import (
	"flag"
	lab2 "lab2"
	"os"
	"strconv"
)

var (
	inputExpression string
	inputFile string
	outputFile string
)

func main() {
	flag.StringVar(&inputExpression, "e", "", "Expression to compute")
	flag.StringVar(&inputFile, "f", "","File with expression")
	flag.StringVar(&outputFile, "o", "", "File to output result")

	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()
	
	if outputFile != "" {
		res, _ := lab2.CalculatePostfix(inputExpression)
		f, err := os.Create(outputFile)
		if err != nil {
			os.Stderr.WriteString(string(err.Error()))
		}
		_, err2 := f.WriteString(strconv.Itoa(res))
		defer f.Close()
		
		if err2 != nil {
			os.Stderr.WriteString(string(err2.Error()))
		}
	} else {
		f, _ := os.Open(inputFile)
		barr := make([]byte, 0, 1)
		oneElemArr := make([]byte, 1)

		for {
			_, err := f.Read(oneElemArr)
			if err != nil {
				break
			}
			barr = append(barr, oneElemArr[0])
		}
		res, _ := lab2.CalculatePostfix(string(barr))
		output := strconv.Itoa(res)
		os.Stdout.Write([]byte(output))
		defer f.Close()
	}
}
