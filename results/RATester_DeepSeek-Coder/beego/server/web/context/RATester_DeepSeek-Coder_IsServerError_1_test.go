package context

import (
	"fmt"
	"testing"
)

func TestIsServerError_1(t *testing.T) {
	tests := []struct {
		name   string
		status int
		want   bool
	}{
		{"Test 1", 500, true},
		{"Test 2", 501, true},
		{"Test 3", 599, true},
		{"Test 4", 400, false},
		{"Test 5", 499, false},
		{"Test 6", 600, false},
		{"Test 7", 0, false},
		{"Test 8", 1000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			output := &BeegoOutput{
				Status: tt.status,
			}
			if got := output.IsServerError(); got != tt.want {
				t.Errorf("BeegoOutput.IsServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}
