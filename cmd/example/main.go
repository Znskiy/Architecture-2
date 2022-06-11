package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "https://github.com/Da6hkin/Architecture-2"
)

var (
	inputExpression     = flag.String("e", "", "Expression to convert")
	inputFileExpression = flag.String("fe", "", "File with an expression to convert")
	outputFile          = flag.String("o", "", "File for the converted expression")
)
var (
	reader io.Reader
	writer io.Writer
	err    error
)

func main() {
	flag.Parse()
	if *inputExpression != "" && *inputFileExpression != "" {
		fmt.Fprintln(os.Stderr, "There should be 1 input")
		os.Exit(1)
	} else if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else if *inputFileExpression != "" {
		reader = strings.NewReader(*inputFileExpression)
		reader, err = os.Open(*inputFileExpression)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Filepath is incorrect")
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "There should be 1 input")
		os.Exit(1)
	}
	if *outputFile != "" {
		writer, err = os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}
	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	res, _ := lab2.prefixToInfix("+ 2 2")
	fmt.Println(res)
}
