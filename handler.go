package lab2

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	if ch.Input == nil {
		return fmt.Errorf("input is not specified")
	}
	if ch.Output == nil {
		return fmt.Errorf("output is not specified")
	}
	buf, err := ioutil.ReadAll(ch.Input)
	if string(buf) == "" {
		return fmt.Errorf("invalid input expression")
	}
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
