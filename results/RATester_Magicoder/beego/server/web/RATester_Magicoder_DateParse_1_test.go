package web

import (
	"fmt"
	"testing"
	"time"
)

func TestDateParse_1(t *testing.T) {
	tests := []struct {
		name        string
		dateString  string
		format      string
		expected    time.Time
		expectedErr error
	}{
		{
			name:       "Valid Date",
			dateString: "2022-01-01",
			format:     "2006-01-02",
			expected:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:       "Invalid Date",
			dateString: "2022-13-01",
			format:     "2006-01-02",
			expectedErr: &time.ParseError{
				Layout:     "2006-01-02",
				Value:      "2022-13-01",
				LayoutElem: "01",
				ValueElem:  "13",
				Message:    "month out of range",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := DateParse(test.dateString, test.format)
			if err != nil {
				if test.expectedErr == nil {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != test.expectedErr.Error() {
					t.Errorf("Expected error: %v, but got: %v", test.expectedErr, err)
				}
			} else {
				if !result.Equal(test.expected) {
					t.Errorf("Expected: %v, but got: %v", test.expected, result)
				}
			}
		})
	}
}
