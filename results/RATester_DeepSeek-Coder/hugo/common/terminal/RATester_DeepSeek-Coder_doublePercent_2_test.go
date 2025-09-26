package terminal

import (
	"fmt"
	"testing"
)

func TestDoublePercent_2(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test1", "10%%", "10%%"},
		{"Test2", "20% off", "20%% off"},
		{"Test3", "No discount", "No discount"},
		{"Test4", "%50 off", "%%50 off"},
		{"Test5", "100%% off", "100%% off"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := doublePercent(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
