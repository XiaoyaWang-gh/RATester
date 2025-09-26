package types

import (
	"fmt"
	"testing"
	"time"
)

func TestToDuration_1(t *testing.T) {
	t.Run("Test with valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			input    any
			expected time.Duration
		}{
			{"Test with int", 10, 10 * time.Nanosecond},
			{"Test with float64", 10.5, 1050000000 * time.Nanosecond},
			{"Test with string", "10", 10 * time.Nanosecond},
			{"Test with string of float", "10.5", 1050000000 * time.Nanosecond},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				result := ToDuration(tc.input)
				if result != tc.expected {
					t.Errorf("Expected %v, got %v", tc.expected, result)
				}
			})
		}
	})

	t.Run("Test with invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name  string
			input any
		}{
			{"Test with nil", nil},
			{"Test with unsupported type", make(chan int)},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected a panic, but didn't get one")
					}
				}()

				ToDuration(tc.input)
			})
		}
	})
}
