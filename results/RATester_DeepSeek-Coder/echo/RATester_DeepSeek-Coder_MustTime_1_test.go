package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestMustTime_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		sourceParam string
		layout      string
		expected    time.Time
		shouldError bool
	}

	testCases := []testCase{
		{
			name:        "valid time",
			sourceParam: "2022-01-02T15:04:05Z",
			layout:      time.RFC3339,
			expected:    time.Date(2022, 1, 2, 15, 4, 5, 0, time.UTC),
			shouldError: false,
		},
		{
			name:        "invalid time",
			sourceParam: "not a time",
			layout:      time.RFC3339,
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var dest time.Time
			b := &ValueBinder{
				ValueFunc: func(sourceParam string) string {
					if sourceParam == tc.sourceParam {
						return tc.sourceParam
					}
					return ""
				},
			}
			b.MustTime(tc.sourceParam, &dest, tc.layout)
			if (b.errors != nil) != tc.shouldError {
				t.Errorf("Expected error %t, but got %t", tc.shouldError, b.errors != nil)
			}
			if !tc.shouldError && dest != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, dest)
			}
		})
	}
}
