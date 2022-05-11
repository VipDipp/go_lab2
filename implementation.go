package main

import (
	"bufio"
	"fmt"
	"os"
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
	var buf rune
	var output []string
	end := ""
	for i, el := range input {
		if el == '+' || el == '-' || el == '*' || el == '/' || el == '^' {

			if i == len(input)-1 {
				fmt.Println("you can't end with symbol!")
				end = "error"
				break
			}
			if count == false {
				symbol, symbols = symbols[len(symbols)-1], symbols[:len(symbols)-1]
				count = true
				output = append(output, (string(buf) + ")" + string(symbol)))
			}
			symbols = append(symbols, rune(el))
		} else if el >= '0' && el <= '9' {
			if count {
				buf = el
				count = false
				output = append(output, "(")
				continue
			}
			symbol, symbols = symbols[len(symbols)-1], symbols[:len(symbols)-1]
			output = append(output, (string(buf) + " " + string(symbol) + " "))
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
