package validation

import (
	"fmt"
	"testing"
)

func TestRequired_1(t *testing.T) {
	v := Validation{}

	type TestStruct struct {
		Field string
	}

	testCases := []struct {
		name     string
		input    TestStruct
		expected bool
	}{
		{
			name:     "valid case",
			input:    TestStruct{Field: "test"},
			expected: true,
		},
		{
			name:     "invalid case",
			input:    TestStruct{Field: ""},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := v.Required(tc.input, "Field")
			if result.Ok != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result.Ok)
			}
		})
	}
}
