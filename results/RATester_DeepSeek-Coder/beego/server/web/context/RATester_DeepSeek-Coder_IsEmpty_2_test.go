package context

import (
	"fmt"
	"testing"
)

func TestIsEmpty_2(t *testing.T) {
	testCases := []struct {
		name   string
		status int
		want   bool
	}{
		{"Status 201", 201, true},
		{"Status 204", 204, true},
		{"Status 304", 304, true},
		{"Status 200", 200, false},
		{"Status 300", 300, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			output := &BeegoOutput{Status: tc.status}
			got := output.IsEmpty()
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}
