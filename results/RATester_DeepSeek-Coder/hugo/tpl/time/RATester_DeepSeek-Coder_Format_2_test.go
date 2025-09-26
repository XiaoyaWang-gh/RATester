package time

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/gohugoio/hugo/common/htime"
)

func TestFormat_2(t *testing.T) {
	ns := &Namespace{
		timeFormatter: htime.TimeFormatter{},
		location:      time.UTC,
	}

	testCases := []struct {
		name     string
		layout   string
		v        any
		expected string
		err      error
	}{
		{
			name:     "valid time",
			layout:   time.RFC3339,
			v:        "2022-01-02T15:04:05Z",
			expected: "2022-01-02T15:04:05Z",
			err:      nil,
		},
		{
			name:     "invalid time",
			layout:   time.RFC3339,
			v:        "2022-01-02T15:04:05Z",
			expected: "",
			err:      errors.New("invalid time"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Format(tc.layout, tc.v)
			if err != tc.err {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("Expected result %s, got %s", tc.expected, result)
			}
		})
	}
}
