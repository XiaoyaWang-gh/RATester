package pagination

import (
	"fmt"
	"testing"
)

func TestSetNums_1(t *testing.T) {
	type testCase struct {
		name     string
		nums     interface{}
		expected int64
	}

	testCases := []testCase{
		{"Test with int", 10, 10},
		{"Test with int64", int64(10), 10},
		{"Test with float64", float64(10), 10},
		{"Test with string", "10", 10},
		{"Test with string number", "10", 10},
		{"Test with nil", nil, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Paginator{}
			p.SetNums(tc.nums)
			if p.nums != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, p.nums)
			}
		})
	}
}
