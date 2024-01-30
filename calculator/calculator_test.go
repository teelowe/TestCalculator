package calculator

import (
	"errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		operand1 float64
		operator string
		operand2 float64
		expected float64
		err      error
	}{
		{1, "+", 1, 2, nil},
		{5, "-", 3, 2, nil},
		{2, "*", 3, 6, nil},
		{6, "/", 2, 3, nil},
		{6, "/", 0, 0, errors.New("division by zero is not allowed")},
		{1, "%", 1, 0, errors.New("invalid operator")},
	}

	for _, test := range tests {
		result, err := Calculate(test.operand1, test.operand2, test.operator)

		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if result != test.expected {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}

func TestParseOperand(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      error
	}{
		{"1", 1, nil},
		{"3.14", 3.14, nil},
		{"abc", 0, errors.New("invalid operand")},
	}

	for _, test := range tests {
		result, err := ParseOperand(test.input)

		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if result != test.expected {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}
