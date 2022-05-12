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
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	result := strings.ReplaceAll(input, " ", "")
	if len(result) < 1 {
		return "", fmt.Errorf("empty input: %w", errorEmptyInput)
	}

	splited := strings.Split(result, "+")
	if len(splited) == 2 {
		return makeMath(splited[0], splited[1], plus)
	}

	splited = strings.Split(result, "-")
	if len(splited) == 2 {
		return makeMath(splited[0], splited[1], minus)
	} else if len(splited) == 3 && len(splited[0]) == 0 {
		return makeMath("-"+splited[1], splited[2], minus)
	}

	return result, fmt.Errorf("number of arguments is wrong: %w", errorNotTwoOperands)
}

func plus(i1, i2 int) int {
	return i1 + i2
}

func minus(i1, i2 int) int {
	return i1 - i2
}

func makeMath(s1, s2 string, operation func(int, int) int) (string, error) {
	i1, err := strconv.Atoi(s1)
	if err != nil {
		return "", fmt.Errorf("invalid syntax: %w", errorNotTwoOperands)
	}

	i2, err := strconv.Atoi(s2)
	if err != nil {
		return "", fmt.Errorf("invalid syntax: %w", errorNotTwoOperands)
	}

	return strconv.Itoa(operation(i1, i2)), nil
}
