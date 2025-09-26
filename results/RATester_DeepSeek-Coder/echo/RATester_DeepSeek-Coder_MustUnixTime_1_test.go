package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestMustUnixTime_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		source   string
		expected time.Time
		hasError bool
	}

	testCases := []testCase{
		{
			name:     "valid unix time",
			source:   "1633673600",
			expected: time.Unix(1633673600, 0),
			hasError: false,
		},
		{
			name:     "invalid unix time",
			source:   "invalid",
			expected: time.Time{},
			hasError: true,
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
			binder := &ValueBinder{}
			binder = binder.MustUnixTime(tc.source, &dest)

			if tc.hasError {
				if binder.BindError() == nil {
					t.Errorf("Expected error but got nil")
				}
			} else {
				if binder.BindError() != nil {
					t.Errorf("Unexpected error: %v", binder.BindError())
				}
				if !dest.Equal(tc.expected) {
					t.Errorf("Expected %v but got %v", tc.expected, dest)
				}
			}
		})
	}
}
