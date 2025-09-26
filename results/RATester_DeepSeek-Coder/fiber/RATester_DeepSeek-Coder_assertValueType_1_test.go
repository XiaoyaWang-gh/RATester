package fiber

import (
	"fmt"
	"testing"
)

func TestAssertValueType_1(t *testing.T) {
	type testStruct struct {
		value int
	}

	testCases := []struct {
		name     string
		input    testStruct
		expected int
	}{
		{
			name:     "success",
			input:    testStruct{value: 1},
			expected: 1,
		},
		{
			name:     "failure",
			input:    testStruct{value: 2},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := assertValueType[int](tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
