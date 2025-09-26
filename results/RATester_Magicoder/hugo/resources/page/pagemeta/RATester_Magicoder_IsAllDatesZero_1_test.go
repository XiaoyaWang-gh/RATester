package pagemeta

import (
	"fmt"
	"testing"
	"time"
)

func TestIsAllDatesZero_1(t *testing.T) {
	type test struct {
		name     string
		dates    Dates
		expected bool
	}

	tests := []test{
		{
			name: "All dates are zero",
			dates: Dates{
				Date:        time.Time{},
				Lastmod:     time.Time{},
				PublishDate: time.Time{},
				ExpiryDate:  time.Time{},
			},
			expected: true,
		},
		{
			name: "Some dates are not zero",
			dates: Dates{
				Date:        time.Now(),
				Lastmod:     time.Time{},
				PublishDate: time.Time{},
				ExpiryDate:  time.Time{},
			},
			expected: false,
		},
		{
			name: "All dates are not zero",
			dates: Dates{
				Date:        time.Now(),
				Lastmod:     time.Now(),
				PublishDate: time.Now(),
				ExpiryDate:  time.Now(),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.dates.IsAllDatesZero()
			if got != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, got)
			}
		})
	}
}
