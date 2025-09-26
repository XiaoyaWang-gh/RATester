package paths

import (
	"fmt"
	"testing"
)

func TestIshex_2(t *testing.T) {
	testCases := []struct {
		name  string
		input byte
		want  bool
	}{
		{"Test Case 1", '0', true},
		{"Test Case 2", '9', true},
		{"Test Case 3", 'a', true},
		{"Test Case 4", 'f', true},
		{"Test Case 5", 'A', true},
		{"Test Case 6", 'F', true},
		{"Test Case 7", 'g', false},
		{"Test Case 8", 'z', false},
		{"Test Case 9", ' ', false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ishex(tc.input)
			if got != tc.want {
				t.Errorf("Expected %t, but got %t", tc.want, got)
			}
		})
	}
}
