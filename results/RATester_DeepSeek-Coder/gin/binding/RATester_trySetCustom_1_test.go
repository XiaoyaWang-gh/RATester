package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrySetCustom_1(t *testing.T) {
	type testStruct struct {
		value string
	}

	testCases := []struct {
		name     string
		val      string
		value    reflect.Value
		expected bool
	}{
		{
			name:     "Test case 1",
			val:      "test",
			value:    reflect.ValueOf(testStruct{value: "test"}),
			expected: true,
		},
		{
			name:     "Test case 2",
			val:      "test",
			value:    reflect.ValueOf(testStruct{value: "test"}),
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

			isSet, err := trySetCustom(tc.val, tc.value)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if isSet != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, isSet)
			}
		})
	}
}
