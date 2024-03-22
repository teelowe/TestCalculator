//go:build integration

package calculator

import (
	"errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		operand1 float64
		operator string
		operand2 float64
		expected float64
		err      error
	}{
		{"addition", 1, "+", 1, 2, nil},
		{"subtraction", 5, "-", 3, 2, nil},
		{"multiplication", 2, "*", 3, 6, nil},
		{"division", 6, "/", 2, 3, nil},
		{"divide by zero", 6, "/", 0, 0, errors.New("division by zero is not allowed")},
		{"modulo", 1, "%", 1, 0, errors.New("invalid operator")},
		{"divide zero by 6", 0, "/", 6, 0, nil},
		{"negative result", 1, "-", 3, -2, nil},
		{"multiply by 0", 3, "*", 0, 0, nil},
		{"negative addition", 6, "+", -7, -1, nil},
		{"true floats", 1.0, "*", 5.12, 5.12, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// uncomment these to use sub test specific setup/teardown
			// SubTestSetup(t)
			// defer SubTestTearDown(t)

			result, err := Calculate(test.operand1, test.operand2, test.operator)

			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error: %v, got: %v", test.err, err)
			}

			if result != test.expected {
				t.Errorf("Expected result: %v, got: %v", test.expected, result)
			}
		})
	}
}
