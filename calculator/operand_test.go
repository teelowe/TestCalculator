//go:build unit

package calculator

import (
	"errors"
	"testing"
)

func TestParseOperand(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		err      error
	}{
		{"int", "1", 1, nil},
		{"float", "3.14", 3.14, nil},
		{"NaN", "abc", 0, errors.New("invalid operand")},
		// added by Thom
		{"also NaN", "a13b4", 0, errors.New("invalid operand")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// uncomment these to use sub test specific setup/teardown
			// SubTestSetup(t)
			// defer SubTestTearDown(t)

			result, err := ParseOperand(test.input)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error: %v, got: %v", test.err, err)
			}
			if result != test.expected {
				t.Errorf("Expected result: %v, got: %v", test.expected, result)
			}
		})
	}
}
