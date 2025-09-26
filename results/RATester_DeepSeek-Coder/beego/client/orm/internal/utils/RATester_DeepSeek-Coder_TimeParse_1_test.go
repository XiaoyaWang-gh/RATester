package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeParse_1(t *testing.T) {
	testCases := []struct {
		name        string
		dateString  string
		format      string
		expectedTP  time.Time
		expectedErr error
	}{
		{
			name:        "Valid date and format",
			dateString:  "2022-01-02 15:04:05",
			format:      "2006-01-02 15:04:05",
			expectedTP:  time.Date(2022, 1, 2, 15, 4, 5, 0, time.UTC),
			expectedErr: nil,
		},
		{
			name:        "Invalid date",
			dateString:  "2022-13-02 15:04:05",
			format:      "2006-01-02 15:04:05",
			expectedTP:  time.Time{},
			expectedErr: &time.ParseError{},
		},
		{
			name:        "Invalid format",
			dateString:  "2022-01-02 15:04:05",
			format:      "2006-02-01 15:04:05",
			expectedTP:  time.Time{},
			expectedErr: &time.ParseError{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tp, err := TimeParse(tc.dateString, tc.format)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if !tp.Equal(tc.expectedTP) {
				t.Errorf("Expected time %v, got %v", tc.expectedTP, tp)
			}
		})
	}
}
