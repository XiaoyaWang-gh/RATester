package strings

import (
	"fmt"
	"testing"
)

func TestRuneCount_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			input    any
			expected int
		}{
			{"hello", 5},
			{"🌍", 2},
			{"🌍🌍", 4},
			{"🌍🌍🌍", 6},
			{"🌍🌍🌍🌍", 8},
		}

		for _, tc := range testCases {
			actual, err := ns.RuneCount(tc.input)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.RuneCount(make(chan int))
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
