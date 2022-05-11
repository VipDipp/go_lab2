package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: document this function.
// PrefixToPostfix converts
func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		errorHandler(err)
		input := strings.TrimSuffix(str, "\r\n")
		input = strings.ReplaceAll(input, " ", "")
		PrefixToInfix(input)
	}
}

func PrefixToInfix(input string) (string, error) {
	fmt.Println(input)
	var symbols []rune
	var symbol rune
	count := true
	done := false
	var buf rune
	var output []string
	str_buf := ""
	end := ""
	for i, el := range input {
		if el == '+' || el == '-' || el == '*' || el == '/' || el == '^' {

			if i == len(input)-1 {
				fmt.Println("you can't end with symbol!")
				end = "error"
				break
			}
			if count == false && len(symbols) != 0 {
				symbol, symbols = symbols[len(symbols)-1], symbols[:len(symbols)-1]
				count = true
				output = append(output, string(buf))
				if _, err := strconv.Atoi(output[0]); err == nil && done == false {
					done = true
					output = append(output, string(symbol))
					symbols = append(symbols, rune(el))
					continue
				}
				output = append(output, ")")
				output = append(output, string(symbol))
			}
			symbols = append(symbols, rune(el))
		} else if el >= '0' && el <= '9' {
			if count {
				buf = el
				count = false
				if len(symbols) != 0 {
					output = append(output, "(")
				}
				continue
			}
			if len(symbols) != 0 {
				symbol, symbols = symbols[len(symbols)-1], symbols[:len(symbols)-1]
				output = append(output, string(buf))
				output = append(output, (" " + string(symbol) + " "))
			} else {
				if len(output) != 0 {
					str_buf, output = output[len(output)-1], output[:len(output)-1]
				}

				output = append(output, string(buf))
				output = append(output, str_buf)
			}
			buf = el

			if i == len(input)-1 {
				output = append(output, (string(buf) + ")"))
			}
		}
	}
	fmt.Println(output)
	for _, el := range output {
		end = end + el
	}
	fmt.Println(end)
	return "TODO", fmt.Errorf("TODO")
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
