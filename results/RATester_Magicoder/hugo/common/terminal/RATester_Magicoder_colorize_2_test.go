package terminal

import (
	"fmt"
	"testing"
)

func Testcolorize_2(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		color    string
		expected string
	}{
		{
			name:     "Test case 1",
			s:        "Hello",
			color:    "red",
			expected: "redHello",
		},
		{
			name:     "Test case 2",
			s:        "World",
			color:    "blue",
			expected: "blueWorld",
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := colorize(test.s, test.color)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
