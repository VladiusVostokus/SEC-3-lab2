package main

import (
	"flag"
	lab2 "lab2"
	"os"
	"strconv"
	"bytes"
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

	handler := &lab2.ComputeHandler {
		    Input: bytes.NewBufferString(inputExpression),
		    Output: os.Stdout,
		}
	
	if inputExpression == "" {

		if inputFile == "" {
			os.Stderr.WriteString("No expression to compute")
			return
		} else {
			f, _ := os.Open(inputFile)
			handler.Input = f
			
			handler.Compute()
			/*

			if outputFile != "" {
				f, err := os.Create(outputFile)

				if err != nil {
					outErr := string(err.Error())
					os.Stderr.WriteString(outErr)
				}
				_, err2 := f.WriteString(strconv.Itoa(res))
				defer f.Close()
		
				if err2 != nil {
					outErr := string(err2.Error())
					os.Stderr.WriteString(outErr)
	    		}
			} else {
				output := strconv.Itoa(res)
				os.Stdout.WriteString(output)
			}*/
		}

	} else {

		if inputFile != "" {
			os.Stderr.WriteString("Can not be 2 source of expression")
			return
		}

		res, _ := lab2.CalculatePostfix(inputExpression)

		if outputFile != "" {
			f, err := os.Create(outputFile)
			if err != nil {
				outErr := string(err.Error())
				os.Stderr.WriteString(outErr)
			}

			_, err2 := f.WriteString(strconv.Itoa(res))

			defer f.Close()
		
			if err2 != nil {
				outErr := string(err2.Error())
				os.Stderr.WriteString(outErr)
	    	}

		} else {
			output := strconv.Itoa(res)
			os.Stdout.WriteString(output)
		}
	}	
}
