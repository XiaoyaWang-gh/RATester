package gin

import (
	"fmt"
	"testing"
)

func TestGetFloat64_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		value    any
		expected float64
	}

	testCases := []testCase{
		{
			name:     "Test with float64 value",
			key:      "key1",
			value:    float64(123.456),
			expected: 123.456,
		},
		{
			name:     "Test with int value",
			key:      "key2",
			value:    int(123),
			expected: 123,
		},
		{
			name:     "Test with string value",
			key:      "key3",
			value:    "123.456",
			expected: 123.456,
		},
		{
			name:     "Test with nil value",
			key:      "key4",
			value:    nil,
			expected: 0,
		},
		{
			name:     "Test with unsupported type",
			key:      "key5",
			value:    make(chan int),
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				Keys: make(map[string]any),
			}
			c.Keys[tc.key] = tc.value

			result := c.GetFloat64(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
