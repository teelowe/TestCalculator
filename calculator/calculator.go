package calculator

import (
	"errors"
	"strconv"
)

// Calculate the result of two operands and an associated operator
func Calculate(operand1, operand2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, errors.New("division by zero is not allowed")
		}
		return operand1 / operand2, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

// ParseOperand provided as a string and return a float64, error
func ParseOperand(operand string) (float64, error) {
	value, err := strconv.ParseFloat(operand, 64)
	if err != nil {
		return 0, errors.New("invalid operand")
	}
	return value, nil
}
