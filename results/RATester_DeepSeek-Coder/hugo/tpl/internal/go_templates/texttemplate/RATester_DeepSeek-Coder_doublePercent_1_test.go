package template

import (
	"fmt"
	"testing"
)

func TestDoublePercent_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test case 1", "10%% off", "10%% off"},
		{"Test case 2", "20% off", "20%% off"},
		{"Test case 3", "No discount", "No discount"},
		{"Test case 4", "%50 off", "%%50 off"},
		{"Test case 5", "100%% off", "100%% off"},
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
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
