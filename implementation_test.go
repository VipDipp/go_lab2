package lab2

import (
	"fmt"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestPrefixToInfix(c *C) {
	samples := map[string]string{
		"++555":                                 "(5 + 5 + 5)",
		"+ 5 * - 4 2 3":                         "(5)+(4 - 2 * 3)",
		"++33-+212-+21":                         "(3 + 3)+(2 + 1 - 2)-(2 + 1)",
		"555-+-+1245":                           "555-(1 + 2 - 4 + 5)",
		"/*123/*798":                            "(1 * 2 / 3)/(79 * 8)",
		"^123/*-+123124":                        "(12 ^ 3)/(1 + 2 - 312 * 4)",
		"5748/*-+1241245":                       "5748/(1 + 2 - 4124 * 5)",
		"5748/*-+1241245/*-+1321313/*-+1321654": "5748/(1 + 2 - 4124 * 5)/(1 + 3 - 2131 * 3)/(1 + 3 - 2165 * 4)",
		"brokenwatch":                           "What da dis",
		"/*-+^11":                               "TOO MANY SYMBOLS!!!",
	}
	for prefix, expected := range samples {
		res, err := PrefixToInfix(prefix)
		if err != nil {
			c.Assert(err, ErrorMatches, expected)
		} else {
			c.Assert(res, Equals, expected)
		}
	}
}

func ExamplePrefixToInfix() {
	res, err := PrefixToInfix("++555")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
