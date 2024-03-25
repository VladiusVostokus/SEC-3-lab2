package main

import (
	"flag"
	lab2 "lab2"
	"os"
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

	handler := &lab2.ComputeHandler {
		    Input: bytes.NewBufferString(inputExpression),
		    Output: os.Stdout,
		}
	
	if inputExpression == "" {
		if inputFile == "" {
			os.Stderr.WriteString("No expression to compute")
			return
		} else {
			f, err := os.Open(inputFile)
			if err != nil {
				os.Stderr.WriteString("Incorrect file to open")
				return
			}

			handler.Input = f
		
			if outputFile != "" {
				f2, err := os.Create(outputFile)
				if err != nil {
					outErr := string(err.Error())
					os.Stderr.WriteString(outErr)
					return
				}
				handler.Output = f2
			}
			err = handler.Compute()
			if err != nil {
				outErr := string(err.Error())
				os.Stderr.WriteString(outErr)
				return
			}
		}
	} else {

		if inputFile != "" {
			os.Stderr.WriteString("Can not be 2 source of expression")
			return
		}
		if outputFile != "" {
			f, err := os.Create(outputFile)
				if err != nil {
					outErr := string(err.Error())
					os.Stderr.WriteString(outErr)
					return
				}
			handler.Output = f
			err = handler.Compute()

			if err != nil {
				outErr := string(err.Error())
				os.Stderr.WriteString(outErr)
				return
			}

		} else {
			err := handler.Compute()

			if err != nil {
				outErr := string(err.Error())
				os.Stderr.WriteString(outErr)
				return
			}
		}
	}	
}

