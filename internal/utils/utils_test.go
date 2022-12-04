package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyPassword(t *testing.T) {
	testCase := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "verify",
			input:    "Tested8*",
			expected: true,
		},
		{
			name:     "verify first",
			input:    "T8*ested",
			expected: true,
		},
		{
			name:     "length not valid",
			input:    "T",
			expected: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			output := VerifyPassword(tc.input)
			assert.Equal(t, output, tc.expected)
		})
	}
}
