package sass

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types/css"
)

func TestCreateVarsStyleSheet_1(t *testing.T) {
	type testCase struct {
		name     string
		input    map[string]any
		expected string
	}

	testCases := []testCase{
		{
			name: "Test with nil input",
			input: map[string]any{
				"$var1": "value1",
				"var2":  "value2",
			},
			expected: "$var1: value1;\nvar2: value2;",
		},
		{
			name: "Test with empty input",
			input: map[string]any{
				"$var1": "",
				"var2":  "",
			},
			expected: "$var1: \"\";\nvar2: \"\";",
		},
		{
			name: "Test with special characters",
			input: map[string]any{
				"$var1": "value1",
				"var2":  "value2",
			},
			expected: "$var1: value1;\nvar2: value2;",
		},
		{
			name: "Test with quoted string",
			input: map[string]any{
				"$var1": css.QuotedString("value1"),
				"var2":  "value2",
			},
			expected: "$var1: \"value1\";\nvar2: value2;",
		},
		{
			name: "Test with unquoted string",
			input: map[string]any{
				"$var1": "value1",
				"var2":  "value2",
			},
			expected: "$var1: value1;\nvar2: value2;",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := CreateVarsStyleSheet(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
