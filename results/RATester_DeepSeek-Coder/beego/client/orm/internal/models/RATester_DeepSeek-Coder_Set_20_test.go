package models

import (
	"fmt"
	"testing"
	"time"
)

func TestSet_20(t *testing.T) {
	testCases := []struct {
		name     string
		input    time.Time
		expected DateTimeField
	}{
		{
			name:     "Test with a valid time",
			input:    time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: DateTimeField(time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name:     "Test with a valid time in different timezone",
			input:    time.Date(2022, time.January, 1, 0, 0, 0, 0, time.FixedZone("UTC+3", 3*60*60)),
			expected: DateTimeField(time.Date(2022, time.January, 1, 0, 0, 0, 0, time.FixedZone("UTC+3", 3*60*60))),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var e DateTimeField
			e.Set(tc.input)
			if e != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, e)
			}
		})
	}
}
