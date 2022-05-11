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
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()
	fmt.Println(*eFlag, *oFlag, *fFlag)
	if *eFlag != "" && *fFlag != "" {
		panic("Only one input sorce allowed!")
	}
	var input io.Reader
	var output io.Writer

	if *eFlag != "" {
		input := strings.NewReader(*eFlag)
	}
	if *fFlag != "" {
		input, err := os.Open(*fFlag)
		errorHandler(err)
	}
	if *oFlag != "" {
		output, err := os.Create(*oFlag)
		errorHandler(err)
	} else {
		output = os.Stdout
	}
	handler := lab2.ComputeHandler{Input: input, Output: output}
	err := handler.Compute()
	errorHandler(err)
	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
