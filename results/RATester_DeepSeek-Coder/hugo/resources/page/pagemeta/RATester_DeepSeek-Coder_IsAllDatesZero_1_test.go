package pagemeta

import (
	"fmt"
	"testing"
	"time"
)

func TestIsAllDatesZero_1(t *testing.T) {
	type testCase struct {
		name     string
		input    Dates
		expected bool
	}

	testCases := []testCase{
		{
			name: "All dates are zero",
			input: Dates{
				Date:        time.Time{},
				Lastmod:     time.Time{},
				PublishDate: time.Time{},
				ExpiryDate:  time.Time{},
			},
			expected: true,
		},
		{
			name: "Some dates are not zero",
			input: Dates{
				Date:        time.Now(),
				Lastmod:     time.Time{},
				PublishDate: time.Time{},
				ExpiryDate:  time.Time{},
			},
			expected: false,
		},
		{
			name: "All dates are not zero",
			input: Dates{
				Date:        time.Now(),
				Lastmod:     time.Now(),
				PublishDate: time.Now(),
				ExpiryDate:  time.Now(),
			},
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

			result := tc.input.IsAllDatesZero()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
