package main

import (
	"io"
	"io/ioutil"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buf, err := ioutil.ReadAll(ch.Input)
	if err != nil {
		return err
	}
	str := strings.Trim(string(buf), "\n")
	output, err := PrefixToInfix(str)
	if err != nil {
		return err
	}
	res := []byte(output)
	_, err = ch.Output.Write(res)
	return nil
}
