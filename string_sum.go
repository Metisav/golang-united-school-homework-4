package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	// Use when the input is not valid ( contains characters )
	errorNotValidInput = errors.New("not valid input string")
)

func StringSum(input string) (output string, err error) {

	// First step check string is valid
	//
	isNotValidRune := func(c rune) bool {
		if c == '-' || c == '+' || c == ' ' {
			return false
		}

		if c > '0' && c < '9' {
			return false
		}

		return true
	}
	b := strings.IndexFunc(input, isNotValidRune) == -1

	if !b {
		return "", fmt.Errorf("%w: \"%s\"", errorNotValidInput, input)
	}

	if len(strings.TrimSpace(input)) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	input = strings.TrimSpace(input)
	ops := strings.Split(input, "+")

	if len(ops) != 2 {
		return "", fmt.Errorf("%w ", errorNotTwoOperands)
	}

	summ := 0
	for _, op := range ops {
		op, _ := strconv.Atoi(op)
		summ += op
	}

	return strconv.Itoa(summ), nil
}
