package pagination

import (
	"fmt"
	"testing"
)

func TestSetNums_1(t *testing.T) {
	paginator := &Paginator{}

	testCases := []struct {
		name     string
		input    interface{}
		expected int64
	}{
		{"valid int", 10, 10},
		{"valid int64", int64(20), 20},
		{"valid string", "30", 30},
		{"invalid string", "abc", 0},
		{"invalid type", true, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			paginator.SetNums(tc.input)
			if paginator.nums != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, paginator.nums)
			}
		})
	}
}
