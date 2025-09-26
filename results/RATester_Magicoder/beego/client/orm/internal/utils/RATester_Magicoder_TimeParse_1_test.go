package utils

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestTimeParse_1(t *testing.T) {
	tests := []struct {
		name        string
		dateString  string
		format      string
		expectedTP  time.Time
		expectedErr error
	}{
		{
			name:        "Valid Date",
			dateString:  "2022-01-01",
			format:      "2006-01-02",
			expectedTP:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: nil,
		},
		{
			name:        "Invalid Date",
			dateString:  "2022-13-01",
			format:      "2006-01-02",
			expectedTP:  time.Time{},
			expectedErr: errors.New("parsing time \"2022-13-01\" as \"2006-01-02\": month out of range"),
		},
		{
			name:        "Invalid Format",
			dateString:  "2022-01-01",
			format:      "2006-02-01",
			expectedTP:  time.Time{},
			expectedErr: errors.New("parsing time \"2022-01-01\" as \"2006-02-01\": month out of range"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tp, err := TimeParse(test.dateString, test.format)
			if err != nil && err.Error() != test.expectedErr.Error() {
				t.Errorf("Expected error: %v, got: %v", test.expectedErr, err)
			}
			if tp != test.expectedTP {
				t.Errorf("Expected time: %v, got: %v", test.expectedTP, tp)
			}
		})
	}
}
