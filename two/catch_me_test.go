package two

import (
	"testing"
)

func TestCatchMe(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test case LLRR=",
			input:    "LLRR=",
			expected: "210122",
		},
		{
			name:     "Test case ==RLL",
			input:    "==RLL",
			expected: "000210",
		},
		{
			name:     "Test case =LLRR",
			input:    "=LLRR",
			expected: "221012",
		},
		{
			name:     "Test case RRL=R",
			input:    "RRL=R",
			expected: "012001",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CatchMe(tc.input)
			if result != tc.expected {
				t.Errorf("input(%q) = actual %q; expected %q",
					tc.input, result, tc.expected)
			}
		})
	}
}
