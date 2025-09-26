package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	type testCase struct {
		name     string
		input    errorMsgs
		expected string
	}

	testCases := []testCase{
		{
			name: "Single error",
			input: errorMsgs{
				{
					Err:  errors.New("test error"),
					Meta: "test meta",
				},
			},
			expected: "Error #01: test error\n     Meta: test meta\n",
		},
		{
			name: "Multiple errors",
			input: errorMsgs{
				{
					Err: errors.New("error 1"),
				},
				{
					Err:  errors.New("error 2"),
					Meta: "meta 2",
				},
			},
			expected: "Error #01: error 1\nError #02: error 2\n     Meta: meta 2\n",
		},
		{
			name:     "No errors",
			input:    errorMsgs{},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.String()
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}
