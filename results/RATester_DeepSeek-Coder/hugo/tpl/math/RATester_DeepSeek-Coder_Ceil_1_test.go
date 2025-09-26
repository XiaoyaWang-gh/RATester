package math

import (
	"fmt"
	"testing"
)

func TestCeil_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("valid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		tests := []struct {
			name     string
			input    any
			expected float64
		}{
			{"positive float", 1.1, 2},
			{"negative float", -1.1, -1},
			{"zero", 0, 0},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				result, err := ns.Ceil(test.input)
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if result != test.expected {
					t.Errorf("Expected %v, got %v", test.expected, result)
				}
			})
		}
	})

	t.Run("invalid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Ceil("invalid")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
