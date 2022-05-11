package lab2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	errorHandler(err)
	input := strings.TrimSuffix(str, "\r\n")
	input = strings.ReplaceAll(input, " ", "")
	PrefixToInfix(input)
}

func PrefixToInfix(input string) (string, error) {
	var symbols []rune
	var symbol, buf rune
	count, done := true, false
	var output []string
	str_buf, end := "", ""
	for i, el := range input {
		if el == '+' || el == '-' || el == '*' || el == '/' || el == '^' {

			if i == len(input)-1 {
				err := errors.New("you can't end with symbol!")
				end = ""
				return end, err
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
				output = append(output, "(")
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
		} else if i == len(input)-1 {
			err := errors.New("What da dis")
			end = ""
			return end, err
		}
	}
	if len(symbols) != 0 {
		err := errors.New("TOO MANY SYMBOLS!!!")
		end = ""
		return end, err
	}
	for _, el := range output {
		end = end + el
	}
	return end, nil
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
