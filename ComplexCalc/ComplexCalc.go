// This is a simple calculator in GoLang.  It does not do order of operations, or supports parens.
// TODO:  Add parens support.  Add OoO.  Add more functions.

package ComplexCalc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

func parseArgs(c string) ([]string, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))

	var tok rune
	var result = make([]string, 0)
	for tok != scanner.EOF {
		tok = s.Scan()
		value := strings.TrimSpace(s.TokenText())
		if len(value) > 0 {
			result = append(result, s.TokenText())
		}
	}
	return result, nil
}

/*func processStack(e []string) (float64, error) {
	result := 0.0
	for _, v := range e {
		c := strings.Split(v, " ")
		if len(c)-1 < 2 {
			return 0.0, errors.New("error: some arguments are not supplied")
		}
		num1, num2, err := parseArgs(c)
		if err != nil {
			return 0.0, err
		}
		switch c[1] {
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0.0 {
				return 0.0, errors.New("error: you tried to divide by zero.")
			}
			result = num1 / num2
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "^":
			result = int(math.Pow(float64(num1), float64(num2)))
		case default:
			result = 0, errors.New("Error, unrecognized operator")
		}
	}
	return result, nil
}*/

func main() {
	expressions := make([]string, 1)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("gocalc>")
		for scanner.Scan() {
			expressions = append(expressions, scanner.Text())

			//Here we look for parens, then split the string, and process the substring.

			//Here we perform **

			//Here we perform *,/

			//Here wer process the remaining stack
			res, err := processStack(expressions)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
			fmt.Print("gocalc>")
		}
	}

	//The new calc
	// calculate 12 + 4 * 3 / 5 - 2
	//input := "12 + 4 * 3 / 5 - 2"   //12.4
	//input := "(12 + 4 * 3) / 5 - 2" //2.8

	// parse input expression to infix notation
	infixTokens, err := ComplexCalc.Scan(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("Infix Tokens:")
	fmt.Println(infixTokens)

	// convert infix notation to postfix notation(RPN) on whitespace
	// postfixTokens, err := ComplexCalc.Parse(infixTokens)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Postfix(RPN) Tokens:")
	// for _, t := range postfixTokens {
	// 	fmt.Printf("%v ", t.GetDescription())
	// }
	// fmt.Println()

	// evaluate tokens
	result, err := ComplexCalc.Evaluate(postfixTokens)
	if err != nil {
		panic(err)
	}

	// outputs the result as an INT
	fmt.Printf("Result: %v", result)

}
