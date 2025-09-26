package gin

import (
	"fmt"
	"testing"
)

func TestMustGet_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected any
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			context: &Context{
				Keys: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
			key:      "key1",
			expected: "value1",
		},
		{
			name: "Test case 2",
			context: &Context{
				Keys: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
			key:      "key3",
			expected: "Key \"key3\" does not exist",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); r != nil {
					if r != tc.expected {
						t.Errorf("Expected panic message to be \"%v\", but got \"%v\"", tc.expected, r)
					}
				}
			}()

			result := tc.context.MustGet(tc.key)
			if result != tc.expected {
				t.Errorf("Expected result to be \"%v\", but got \"%v\"", tc.expected, result)
			}
		})
	}
}
