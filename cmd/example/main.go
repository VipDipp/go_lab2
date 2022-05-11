package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/VipDipp/go_lab2"
)

var (
	eFlag = flag.String("e", "", "Expression to compute")
	oFlag = flag.String("o", "", "Output file")
	fFlag = flag.String("f", "", "Input file")
)

func main() {
	flag.Parse()
	if *eFlag != "" && *fFlag != "" {
		panic("Only one input sorce allowed!")
	}
	var input io.Reader
	var output io.Writer
	var err error

	if *eFlag != "" {
		input = strings.NewReader(*eFlag)
	}
	if *fFlag != "" {
		input, err = os.Open(*fFlag)
		errorHandler(err)
	}
	if *oFlag != "" {
		output, err = os.Create(*oFlag)
		errorHandler(err)
	} else {
		output = os.Stdout
	}
	handler := lab2.ComputeHandler{Input: input, Output: output}
	err = handler.Compute()
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
