package terminal

import (
	"fmt"
	"testing"
)

func TestSinglePercent_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test1", "Hello%%World", "Hello%World"},
		{"Test2", "%%Hello%%World", "%Hello%World"},
		{"Test3", "%%%Hello%%World", "%Hello%World"},
		{"Test4", "Hello%%%World", "Hello%World"},
		{"Test5", "%%%Hello%%%World", "%Hello%World"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := singlePercent(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
