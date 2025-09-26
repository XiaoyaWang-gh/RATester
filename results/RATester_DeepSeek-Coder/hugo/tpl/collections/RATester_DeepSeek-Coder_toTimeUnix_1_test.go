package collections

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestToTimeUnix_1(t *testing.T) {
	ns := &Namespace{
		loc: time.UTC,
	}

	testCases := []struct {
		name     string
		input    reflect.Value
		expected int64
	}{
		{
			name:     "valid time.Time",
			input:    reflect.ValueOf(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)),
			expected: 1609459200,
		},
		{
			name:     "invalid type",
			input:    reflect.ValueOf("not a time.Time"),
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ns.toTimeUnix(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
