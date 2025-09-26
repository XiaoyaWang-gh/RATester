package validation

import (
	"fmt"
	"testing"
)

func TestIP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := Validation{}

	type testCase struct {
		input    interface{}
		expected bool
	}

	testCases := []testCase{
		{"192.168.1.1", true},
		{"256.256.256.256", false},
		{"255.255.255.255", true},
		{"127.0.0.1", true},
		{"localhost", false},
		{"", false},
	}

	for _, tc := range testCases {
		result := v.IP(tc.input, "")
		if result.Ok != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result.Ok)
		}
	}
}
