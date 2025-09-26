package admin

import (
	"fmt"
	"testing"
)

func TestIsSuccess_1(t *testing.T) {
	t.Run("TestIsSuccess", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			status   int
			expected bool
		}{
			{"200 is success", 200, true},
			{"201 is success", 201, true},
			{"299 is success", 299, true},
			{"300 is not success", 300, false},
			{"301 is not success", 301, false},
			{"400 is not success", 400, false},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				r := &Result{Status: tc.status}
				result := r.IsSuccess()
				if result != tc.expected {
					t.Errorf("Expected %v, got %v", tc.expected, result)
				}
			})
		}
	})
}
